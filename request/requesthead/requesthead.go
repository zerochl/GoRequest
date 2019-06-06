package requesthead

import (
	"net/http"
	"GoRequest/common/util"
)

var (
	Header http.Header
)

func AddHeader(originHeader, addHeader *http.Header)  {
	for key, headerItem := range *addHeader {
		for _, headerItemItem := range headerItem {
			originHeader.Set(key, headerItemItem)
		}
	}
}

func RequestHeaderForGet() *http.Header {
	getHeader := requestHeaderBase()
	return getHeader
}

func RequestHeaderForPost() *http.Header {
	postHeader := requestHeaderBase()
	return postHeader
}

func requestHeaderBase() *http.Header {
	baseHeader := make(http.Header)
	//baseHeader.Set(Header_Key_App_Key, cons.AppKey)
	//baseHeader.Set(Header_Key_User_Id, cons.UserId)
	//baseHeader.Set(Header_Key_Timestamp, date.GetTimestamp())
	//baseHeader.Set(Header_Key_Nonce, util.GetNonce())
	return &baseHeader
}