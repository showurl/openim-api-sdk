package v20220317

import (
	"github.com/showurl/openim-api-sdk/v20220317/apiStruct"
	"github.com/showurl/openim-api-sdk/v20220317/conf"
)

func UpdateUserInfo(
	req *apiStruct.ReqUpdateUserInfo,
	options ...Option,
) (resp apiStruct.RespUpdateUserInfo, err error) {
	response := &struct {
		apiStruct.ResponseBase
		Data apiStruct.RespUpdateUserInfo `json:"data"`
	}{}
	err = adminApiProxy(conf.ApiRouterUpdateUserInfo, req, response, options...)
	if err != nil {
		return resp, err
	}
	return response.Data, nil
}
