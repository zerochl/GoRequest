package apirequest

import (
	"net/http"
	"GoRequest/request/apirequest/threadpool"
	"crypto/tls"
	"net"
	"time"
	"GoRequest/request/apirequest/entity"
	"GoRequest/request/apirequest/cons"
	"encoding/json"
	"io"
	"log"
	"GoRequest/request/requesthead"
)

type ApiRequest struct {
	threadPool        *threadpool.ThreadPool
	requestClient     *http.Client
	baseUrl           string
	header            *http.Header
	retryCount        int
	retryIntervalTime int
}

type RequestCallback interface {
	OnSuccess(response string)
	OnError(errorCode int, errorMsg string)
}

func NewApiRequest(poolMaxIdle, poolKeepAlive, requestTimeOut, requestKeepAlive int, baseUrl string, header *http.Header) *ApiRequest {
	dialer := &net.Dialer{
		Timeout:   time.Duration(requestTimeOut) * time.Second,
		KeepAlive: time.Duration(requestKeepAlive) * time.Second,
	}
	transport := &http.Transport{
		DialContext:     dialer.DialContext,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		//Dial:            TimeoutDialer(time.Duration(connectTimeOut), time.Duration(readTimeOut)),
	}
	requestClient := &http.Client{
		Transport: transport,
	}

	return &ApiRequest{
		requestClient:     requestClient,
		retryCount:        3,
		retryIntervalTime: 5,
		baseUrl:           baseUrl,
		header:            header,
		threadPool:        threadpool.NewThreadPool(poolMaxIdle, poolKeepAlive),
	}
}

func (apiRequest *ApiRequest) Post(route string, param interface{}, header *http.Header, requestCallBack RequestCallback) {
	apiRequest.prepareRequest(http.MethodPost, apiRequest.baseUrl+route, paramToBody(param), header, requestCallBack)
}

func (apiRequest *ApiRequest) PostWithRes(route string, param interface{}, header *http.Header, requestCallBack RequestCallback, res interface{}) {
	apiRequest.prepareRequestWithRes(http.MethodPost, apiRequest.baseUrl+route, paramToBody(param), header, requestCallBack, res)
}

func (apiRequest *ApiRequest) Get(route string, header *http.Header, requestCallBack RequestCallback) {
	log.Println("apiRequest.baseUrl+route:", apiRequest.baseUrl+route)
	apiRequest.prepareRequest(http.MethodGet, apiRequest.baseUrl+route, nil, header, requestCallBack)
}

func (apiRequest *ApiRequest) prepareRequest(method, url string, body io.Reader, header *http.Header, requestCallBack RequestCallback)  {
	apiRequest.threadPool.AddTask(func() error {
		log.Println("in thread")
		newRequest, err := http.NewRequest(method, url, body)
		if err != nil {
			requestCallBack.OnError(cons.ErrorCodeRequestNew, err.Error())
			return err
		}
		// 设置请求头
		newRequest.Header = make(http.Header)
		requesthead.AddHeader(&newRequest.Header, apiRequest.header)
		requesthead.AddHeader(&newRequest.Header, header)
		// 创建response接收对象
		res := &entity.RootResponse{}
		// 执行网络请求
		errorCode, err, _ := apiRequest.tryRequest(newRequest, res)
		if err != nil {
			// 请求异常
			requestCallBack.OnError(errorCode, err.Error())
			return err
		}
		// 请求成功
		if res.Code != cons.ErrorCodeRequestOK {
			// 服务端返回异常
			requestCallBack.OnError(res.Code, res.Message)
			return nil
		}
		dataJsonByte, err := json.Marshal(res)
		if err != nil {
			requestCallBack.OnError(cons.ErrorCodeRequestReadException, err.Error())
			return nil
		}
		requestCallBack.OnSuccess(string(dataJsonByte))
		return nil
	})
}

func (apiRequest *ApiRequest) prepareRequestWithRes(method, url string, body io.Reader, header *http.Header, requestCallBack RequestCallback, res interface{})  {
	apiRequest.threadPool.AddTask(func() error {
		log.Println("in thread")
		newRequest, err := http.NewRequest(method, url, body)
		if err != nil {
			requestCallBack.OnError(cons.ErrorCodeRequestNew, err.Error())
			return err
		}
		// 设置请求头
		newRequest.Header = make(http.Header)
		requesthead.AddHeader(&newRequest.Header, apiRequest.header)
		requesthead.AddHeader(&newRequest.Header, header)
		// 执行网络请求
		errorCode, err, response := apiRequest.tryRequest(newRequest, res)
		if err != nil && response == "" {
			// 请求异常
			requestCallBack.OnError(errorCode, err.Error())
			return err
		}
		if err != nil && response != "" {
			requestCallBack.OnSuccess(response)
			return nil
		}
		dataJsonByte, err := json.Marshal(res)
		if err != nil {
			requestCallBack.OnError(cons.ErrorCodeRequestReadException, err.Error())
			return nil
		}
		requestCallBack.OnSuccess(string(dataJsonByte))
		return nil
	})
}