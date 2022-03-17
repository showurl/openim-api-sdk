package v20220317

import (
	"github.com/showurl/openim-api-sdk/v20220317/apiStruct"
	"github.com/showurl/openim-api-sdk/v20220317/conf"
)

func UserRegister(
	req *apiStruct.ReqUserRegister,
	options ...Option,
) (resp apiStruct.RespUserRegister, err error) {
	response := &struct {
		apiStruct.ResponseBase
		Data apiStruct.RespUserRegister `json:"data"`
	}{}
	err = adminApiProxy(conf.ApiRouterUserRegister, req, response, options...)
	if err != nil {
		return resp, err
	}
	return response.Data, nil
}
