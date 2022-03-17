package v20220317

import (
	"encoding/json"
	"github.com/showurl/openim-api-sdk/v20220317/apiStruct"
	"github.com/showurl/openim-api-sdk/v20220317/conf"
	"github.com/showurl/openim-api-sdk/v20220317/utils"
)

func UserToken(req *apiStruct.ReqUserToken) (resp apiStruct.RespUserToken, err error) {
	buf, err := utils.HttpClient.Post2Api(conf.ApiRouterUserToken, req, "")
	if err != nil {
		return resp, err
	}
	response := &struct {
		apiStruct.ResponseBase
		Data apiStruct.RespUserToken `json:"data"`
	}{}
	err = json.Unmarshal(buf, response)
	if err != nil || response == nil {
		return resp, err
	}
	return response.Data, nil
}
