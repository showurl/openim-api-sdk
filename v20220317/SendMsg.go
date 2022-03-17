package v20220317

import (
	"github.com/showurl/openim-api-sdk/v20220317/apiStruct"
	"github.com/showurl/openim-api-sdk/v20220317/conf"
)

func SendMsg(
	req *apiStruct.ReqSendMsg,
	options ...Option,
) (resp apiStruct.RespSendMsg, err error) {
	response := &struct {
		apiStruct.ResponseBase
		Data apiStruct.RespSendMsg `json:"data"`
	}{}
	err = adminApiProxy(conf.ApiRouterSendMsg, req, response, options...)
	if err != nil {
		return resp, err
	}
	return response.Data, nil
}
