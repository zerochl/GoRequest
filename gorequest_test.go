package GoRequest

import (
	"testing"
	"encoding/json"
	"GoRequest/request/apirequest"
	"log"
)

func initRequest() {
	initRequestEntity := &InitRequestEntity{
		BaseUrl: "http://172.16.11.168:8086/",
	}
	InitRequest(initRequestEntity)
}

type GetRequestCallBack struct {
	apirequest.RequestCallback
}

func (request *GetRequestCallBack) OnSuccess(response string)  {
	result, _ := json.Marshal(response)
	log.Println("response:", string(result))
}

func (request *GetRequestCallBack) OnError(errorCode int, errorMsg string)  {
	log.Println("error:", errorCode)
}

func TestGetRequest(t *testing.T) {
	initRequest()
	//GetRequest("", "{\"test\":\"value1\",\"test2\":\"value2\"}")
	apiRequestCallBack := &GetRequestCallBack{}
	GetRequest("api/login/callback", "{\"test\": 3,\"test2\":\"value2\"}", apiRequestCallBack)

	quit := make(chan int)
	<- quit
}