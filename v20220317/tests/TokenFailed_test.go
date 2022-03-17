package tests

import (
	v20220317 "github.com/showurl/openim-api-sdk/v20220317"
	"github.com/showurl/openim-api-sdk/v20220317/apiStruct"
	"testing"
)

func TestTokenFailed(t *testing.T) {
	t.Setenv("UrlPrefix", "http://api.openim:10000")
	t.Setenv("DefaultUserID", `openIM123456`)
	t.Setenv("DefaultUserSecret", `tuoyun`)
	resp, err := v20220317.UserToken(&apiStruct.ReqUserToken{
		Secret:      `tuoyun`,
		Platform:    apiStruct.AdminPlatformID,
		UserID:      "openIM123456",
		OperationID: v20220317.NewOperationID(),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
	v20220317.SetToken(`openIM123456`, `asdfghjkl`)
	resp1, err := v20220317.SendMsg(&apiStruct.ReqSendMsg{
		OperationID:      v20220317.NewOperationID(),
		SendID:           "15666355528",
		RecvID:           "15666355529",
		GroupID:          "cf8e2fc76b4b316097dab9eb517032c2",
		SenderNickname:   "我是昵称",
		SenderFaceURL:    "https://go.dev/images/gophers/pilot-bust.svg",
		SenderPlatformID: apiStruct.AdminPlatformID,
		ForceList: []string{
			`15666355529`,
		},
		Content: apiStruct.MsgContext{
			Text: "内容",
		},
		ContentType:  apiStruct.MsgContentTypeText,
		SessionType:  apiStruct.MsgSessionTypeGroup,
		IsOnlineOnly: false,
		OfflinePushInfo: apiStruct.OfflinePushInfo{
			Title:         "您有新消息",
			Desc:          "",
			Ex:            "",
			IOSPushSound:  "",
			IOSBadgeCount: false,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp1)
}
