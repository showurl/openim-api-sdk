package v20220317

import (
	"github.com/showurl/openim-api-sdk/v20220317/apiStruct"
)

var tokenMgr = make(map[string]string)

func GetToken(userid, secret string) (string, error) {
	if token, ok := tokenMgr[userid]; ok {
		return token, nil
	}
	// 获取token
	resp, err := UserToken(&apiStruct.ReqUserToken{
		Secret:      secret,
		Platform:    apiStruct.AdminPlatformID,
		UserID:      userid,
		OperationID: NewOperationID(),
	})
	if err != nil {
		return "", err
	}
	tokenMgr[userid] = resp.Token
	return resp.Token, nil
}
