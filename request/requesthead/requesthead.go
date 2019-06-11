package requesthead

import (
	"net/http"
	"encoding/json"
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
	return AddInputHead(inputHeadJson, getHeader)
}

func RequestHeaderForPost(inputHeadJson string) *http.Header {
	postHeader := requestHeaderBase()
	return AddInputHead(inputHeadJson, postHeader)
}

func requestHeaderBase() *http.Header {
	baseHeader := make(http.Header)
	return &baseHeader
}

func AddInputHead(inputHeadJson string, header *http.Header) *http.Header {
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
