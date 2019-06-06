package response

import "mykeysdk/common/util"

type BaseResponse struct {
	ResultCode    int         `json:"result_code"`
	ResultMessage string      `json:"result_message"`
	ResultData    interface{} `json:"result_data"`
}

func NewBaseResponse(resultCode int, resultMessage string, resultData interface{}) *BaseResponse {
	return &BaseResponse{ResultCode: resultCode, ResultMessage: resultMessage, ResultData: resultData}
}

func (baseResponse *BaseResponse) ToJson() string {
	return util.ToJson(baseResponse)
}