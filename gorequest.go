package GoRequest

import (
	"log"
	"GoRequest/initialize"
	"GoRequest/common/entity/response"
	"GoRequest/common/cons"
	"GoRequest/common/util"
	"GoRequest/common/util/maplist"
	"encoding/json"
	"GoRequest/request/requesthead"
)

var (
	initService *initialize.InitService
)

func InitRequest(initRequestEntity *InitRequestEntity) (result string) {
	log.Println("in InitRequest")
	defer catchError(&result)
	initService = initialize.Init(initRequestEntity.BaseUrl)
	return response.NewBaseResponse(cons.ResponseCodeSuccess, "", nil).ToJson()
}

func GetRequest(route string, paramJson string, headerJson string, requestCallBack ApiRequestCallBack) (result string) {
	log.Println("in GetRequest")
	defer catchError(&result)
	paramMapList := getGetParam(paramJson)
	realGetParam := getGetRealParam(paramMapList)
	header := requesthead.RequestHeaderForGet(headerJson)
	initService.GetApiRequest().Get(route + realGetParam, header, requestCallBack)
	return response.NewBaseResponse(cons.ResponseCodeSuccess, "", nil).ToJson()
}

func PostRequest(route string, paramJson string, headerJson string, requestCallBack ApiRequestCallBack) (result string) {
	log.Println("in PostRequest")
	defer catchError(&result)
	realPostParam := getPostRealParam(paramJson)
	header := requesthead.RequestHeaderForPost(headerJson)
	initService.GetApiRequest().Post(route, realPostParam, header, requestCallBack)
	return response.NewBaseResponse(cons.ResponseCodeSuccess, "", nil).ToJson()
}

func getParam(paramJson string) map[string] interface{} {
	if paramJson == "" {
		return nil
	}
	var paramOriginMap  map[string] interface{}
	err := json.Unmarshal([]byte(paramJson), &paramOriginMap)
	if err != nil {
		return nil
	}
	return paramOriginMap
}

func getGetParam(paramJson string) *util.MapList {
	paramOriginMap := getParam(paramJson)
	if paramOriginMap == nil {
		return nil
	}
	paramMapList := util.NewMapList()
	for key, value :=range paramOriginMap {
		paramMapList.Push(&maplist.Elements{Key: key, Value: value})
	}
	return paramMapList
}

func getGetRealParam(paramMapList *util.MapList) string {
	realGetParam := ""
	if paramMapList != nil {
		realGetParam = paramMapList.GetRequestParam()
	}
	return realGetParam
}

func getPostRealParam(paramJson string) map[string] interface{} {
	if paramJson == "" {
		return nil
	}
	var realPostParam  map[string] interface{}
	err := json.Unmarshal([]byte(paramJson), &realPostParam)
	if err != nil {
		return nil
	}
	return realPostParam
}