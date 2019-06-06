package apirequest

import (
	"testing"
	"log"
	"encoding/json"
	"net/http"
)

type GetRequestCallBack struct {
	RequestCallback
}

func (request *GetRequestCallBack) OnSuccess(response interface{})  {
	result, _ := json.Marshal(response)
	log.Println("response:", string(result))
}

func (request *GetRequestCallBack) OnError(errorCode int, errorMsg string)  {
	log.Println("error:", errorCode)
}

type AssetList struct {
	Chain string `json:"chain"`
	Currency string `json:"currency"`
	Codes [] struct{
		Code string `json:"code"`
		Symbol string `json:"symbol"`
	} `json:"codes"`
}

func TestApiRequest_Get(t *testing.T) {
	// https://app.mykey.tech/notification/client/config
	apiRequest := NewApiRequest(10, 5, 30, 5, "https://stg-app.mykey.tech/")
	getRequestCallBack := &GetRequestCallBack{}

	// https://app.mykey.tech/asset/list
	assetList := &AssetList{
		Chain: "eos",
		Currency: "USD",
		Codes: make([]struct{
			Code string `json:"code"`
			Symbol string `json:"symbol"`
		}, 1),
	}
	assetList.Codes[0].Code = "eosio.token"
	assetList.Codes[0].Symbol = "EOS"

	header := make(http.Header)
	header.Add("Test", "Test")

	apiRequest.Get("notification/client/config?test=test&test1=test1", &header, getRequestCallBack)

	//apiRequest.Post("asset/list", assetList, getRequestCallBack)
	//apiRequest.Get("notification/client/config?test=test&test1=test1", getRequestCallBack)
	//
	//apiRequest.Post("asset/list", assetList, getRequestCallBack)
	//apiRequest.Post("asset/list", assetList, getRequestCallBack)
	//apiRequest.Get("notification/client/config?test=test&test1=test1", getRequestCallBack)
	//apiRequest.Post("asset/list", assetList, getRequestCallBack)
	//apiRequest.Get("notification/client/config?test=test&test1=test1", getRequestCallBack)
	//apiRequest.Post("asset/list", assetList, getRequestCallBack)
	//apiRequest.Get("notification/client/config?test=test&test1=test1", getRequestCallBack)
	//apiRequest.Post("asset/list", assetList, getRequestCallBack)
	//
	//
	//apiRequest.Get("notification/client/config", getRequestCallBack)
	//apiRequest.Get("notification/client/config", getRequestCallBack)
	//apiRequest.Get("notification/client/config", getRequestCallBack)
	quit := make(chan int)
	<- quit
}
