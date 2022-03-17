package v20220317

import (
	"github.com/showurl/openim-api-sdk/v20220317/apiStruct"
	"github.com/showurl/openim-api-sdk/v20220317/conf"
)

func CreateGroup(
	req *apiStruct.ReqCreateGroup,
	options ...Option,
) (resp apiStruct.RespCreateGroup, err error) {
	response := &struct {
		apiStruct.ResponseBase
		Data apiStruct.RespCreateGroup `json:"data"`
	}{}
	err = adminApiProxy(conf.ApiRouterCreateGroup, req, response, options...)
	if err != nil {
		return resp, err
	}
	return response.Data, nil
}
