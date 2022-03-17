package v20220317

import (
	"github.com/showurl/openim-api-sdk/v20220317/apiStruct"
	"github.com/showurl/openim-api-sdk/v20220317/conf"
)

func KickGroup(
	req *apiStruct.ReqKickGroup,
	options ...Option,
) (resp apiStruct.RespKickGroups, err error) {
	response := &struct {
		apiStruct.ResponseBase
		Data apiStruct.RespKickGroups `json:"data"`
	}{}
	err = adminApiProxy(conf.ApiRouterKickGroup, req, response, options...)
	if err != nil {
		return resp, err
	}
	return response.Data, nil
}
