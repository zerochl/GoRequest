package GoRequest

import (
	"log"
	"GoRequest/common/entity/response"
	"GoRequest/common/cons"
)

type ApiRequestCallBack interface {
	OnSuccess(response string)
	OnError(errorCode int, errorMsg string)
}

func catchError(result *string) {
	if err := recover(); err != nil {
		log.Println("error:", err.(error).Error())
		*result = response.NewBaseResponse(cons.ResponseCodeErrorException, cons.MsgProgramError+err.(error).Error(), nil).ToJson()
	}
}