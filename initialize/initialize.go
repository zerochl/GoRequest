package initialize

import (
	"GoRequest/request/apirequest"
	"GoRequest/request/requesthead"
	"GoRequest/common/cons"
	"net/http"
)

type InitService struct {
	apiRequest *apirequest.ApiRequest
}

func Init(baseUrl string) *InitService {
	initService := &InitService{}
	requesthead.Header = make(http.Header)
	initService.apiRequest = apirequest.NewApiRequest(cons.RequestPoolMaxIdle, cons.RequestPoolKeepAlive, cons.RequestTimeOut, cons.RequestKeepAlive, baseUrl)
	return initService
}

func (initService *InitService) GetApiRequest() *apirequest.ApiRequest {
	return initService.apiRequest
}
