package requesthead

import (
	"net/http"
	"encoding/json"
)

var (
	Header http.Header
)

func AddHeader(originHeader, addHeader *http.Header) {
	for key, headerItem := range *addHeader {
		for _, headerItemItem := range headerItem {
			originHeader.Set(key, headerItemItem)
		}
	}
}

func RequestHeaderForGet(inputHeadJson string) *http.Header {
	getHeader := requestHeaderBase()
	return addInputHead(inputHeadJson, getHeader)
}

func RequestHeaderForPost(inputHeadJson string) *http.Header {
	postHeader := requestHeaderBase()
	return addInputHead(inputHeadJson, postHeader)
}

func requestHeaderBase() *http.Header {
	baseHeader := make(http.Header)
	//baseHeader.Set(Header_Key_App_Key, cons.AppKey)
	//baseHeader.Set(Header_Key_User_Id, cons.UserId)
	//baseHeader.Set(Header_Key_Timestamp, date.GetTimestamp())
	//baseHeader.Set(Header_Key_Nonce, util.GetNonce())
	return &baseHeader
}

func addInputHead(inputHeadJson string, header *http.Header) *http.Header {
	if inputHeadJson == "" {
		return header
	}
	var inputHeaderMap map[string]string
	err := json.Unmarshal([]byte(inputHeadJson), &inputHeaderMap)
	if err != nil {
		return header
	}
	for key, value := range inputHeaderMap {
		header.Set(key, value)
	}
	return header
}
