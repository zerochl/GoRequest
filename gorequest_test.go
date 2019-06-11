package GoRequest

import (
	"testing"
	"encoding/json"
	"GoRequest/request/apirequest"
	"log"
)

func initRequest() {
	initRequestEntity := &InitRequestEntity{
		BaseUrl: "http://172.16.11.229:8086/",
		HeaderJson: "{\"BaseHeader1\": \"1233\",\"BaseHeader2\":\"value2\"}",
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
	// https://wuliao.epro.sogou.com/ask?id=1027267&cb=_sg6432acd6358e759f&ssi0=2053&wsg=w-0_dx-99&_v=f3efed63
	GetRequest("api/login/callback", "{\"test\": 3,\"test2\":\"value2\"}", "{\"head1\": \"hhhh\",\"head2\":\"value2\"}", apiRequestCallBack)

	PostRequest("api/login/callback", "{\"test\": 3,\"test2\":\"value2\"}", "{\"head1\": \"hhhh\",\"head2\":\"value2\"}", apiRequestCallBack)

	quit := make(chan int)
	<- quit
}