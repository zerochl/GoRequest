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

func Init(baseUrl, headJson string) *InitService {
	initService := &InitService{}
	header := make(http.Header)
	requesthead.AddInputHead(headJson, &header)

	initService.apiRequest = apirequest.NewApiRequest(cons.RequestPoolMaxIdle, cons.RequestPoolKeepAlive, cons.RequestTimeOut, cons.RequestKeepAlive, baseUrl, &header)
	return initService
}

func (initService *InitService) GetApiRequest() *apirequest.ApiRequest {
	return initService.apiRequest
}
