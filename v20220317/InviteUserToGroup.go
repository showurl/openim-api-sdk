package v20220317

import (
	"github.com/showurl/openim-api-sdk/v20220317/apiStruct"
	"github.com/showurl/openim-api-sdk/v20220317/conf"
)

func InviteUserToGroup(
	req *apiStruct.ReqInviteUserToGroup,
	options ...Option,
) (resp apiStruct.RespInviteUserToGroups, err error) {
	response := &struct {
		apiStruct.ResponseBase
		Data apiStruct.RespInviteUserToGroups `json:"data"`
	}{}
	err = adminApiProxy(conf.ApiRouterInviteUserToGroup, req, response, options...)
	if err != nil {
		return resp, err
	}
	return response.Data, nil
}
