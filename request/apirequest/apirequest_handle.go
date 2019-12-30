package apirequest

import (
	"encoding/json"
	"net/http"
	"log"
	"time"
	"io"
	"bytes"
	"GoRequest/request/apirequest/cons"
	"GoRequest/common/util"
	"net"
)

// Do makes an HTTP request with the native `http.Do` interface and context passed in
func (apiRequest *ApiRequest) tryRequest(req *http.Request, res interface{}) (errorCode int, err error, response string) {
	for i := 0; i <= apiRequest.retryCount; i++ {
		errorCode, err, response = apiRequest.request(req, res)
		if err != nil {
			if response != "" {
				return errorCode, err, response
			}
			log.Println("http request failed error url:", req.URL.String(), ";error:", err)
			time.Sleep(time.Duration(apiRequest.retryIntervalTime) * time.Second)
			continue
		}
		return errorCode, nil, response
	}
	return errorCode, err, response
}

func (apiRequest *ApiRequest) request(req *http.Request, res interface{}) (errorCode int, err error, result string) {
	var (
		response  *http.Response
		bodyBytes []byte
	)
	apiRequest.outputRequestLog(req)
	response, err = apiRequest.requestClient.Do(req)
	if err != nil {
		log.Println("in request Do error:", err.Error())
		return cons.ErrorCodeRequestResponse, err, ""
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Println("in request Do response error url:", req.URL.String(), ";errorCode:", response.StatusCode)
		return response.StatusCode, cons.HTTP_CODE_ERROR, ""
	}

	bodyBytes, err = readAll(response.Body)
	if err != nil {
		log.Println("readAll error:", err.Error())
		return cons.ErrorCodeRequestReadException, err, ""
	}
	apiRequest.outputResponseLog(bodyBytes, http.StatusOK)
	if res != nil {
		err = json.Unmarshal(bodyBytes, &res)
		if err != nil {
			log.Println("cannot json unmarshal response:", string(bodyBytes), ";err:", err.Error())
			return cons.ErrorCodeRequestJsonException, err, string(bodyBytes)
		}
	}
	return cons.ErrorCodeRequestOK, nil, ""
}

func readAll(reader io.Reader) (b []byte, err error) {
	defer util.CatchError()
	buf := bytes.NewBuffer(make([]byte, 0, 64*1024))
	_, err = buf.ReadFrom(reader)
	return buf.Bytes(), err
}

func TimeoutDialer(connectTimeout time.Duration, readWriteTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, connectTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(readWriteTimeout))
		return conn, nil
	}
}