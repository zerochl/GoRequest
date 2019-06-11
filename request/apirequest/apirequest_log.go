package apirequest

import (
	"net/http"
	"log"
	"bytes"
	"encoding/json"
	"io"
)

func (apiRequest *ApiRequest) outputRequestLog(req *http.Request)  {
	header := req.Header
	var buffer bytes.Buffer
	for key, headerItem := range header {
		buffer.WriteString(key + ":")
		for _, headerItemItem := range headerItem {
			buffer.WriteString(headerItemItem + " ")
		}
		buffer.WriteString("\n")
	}
	log.Println("method:", req.Method, ";url:", req.URL.String() + ";\nheader:", buffer.String())
	if req.Form != nil {
		log.Println(req.Form.Encode())
	}
	//log.Println("req:", req.URL.Query())
	//// 未解决body读取之后导致不能再次读取的问题
	//if nil != req.Body {
	//	buff := new(bytes.Buffer)
	//	io.Copy(buff, req.Body)
	//	//body := req.Body.(io.Reader)
	//	//buf := body.(*bytes.Buffer)
	//	//body, _ := ioutil.ReadAll(req.Body)
	//	log.Println("param:", string(buff.Bytes()))
	//}
}

func (apiRequest *ApiRequest) outputResponseLog(bodyBytes []byte, httpCode int)  {
	if bodyBytes == nil {
		log.Println("status->", httpCode)
	} else {
		log.Println("status->", httpCode, ";body:", string(bodyBytes))
	}
}

func paramToBody(param interface{}) (body io.Reader) {
	buff := new(bytes.Buffer)
	err := json.NewEncoder(buff).Encode(param)
	if err != nil {
		log.Println("failed to marshal user payload:", err.Error())
		return nil
	}
	return buff
}