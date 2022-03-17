package tests

import (
	v20220317 "github.com/showurl/openim-api-sdk/v20220317"
	"github.com/showurl/openim-api-sdk/v20220317/apiStruct"
	"testing"
)

/*
=== RUN   TestInviteUserToGroup
    InviteUserToGroup_test.go:24: apiStruct.RespInviteUserToGroups{{UserID:"openIM654321", Result:0}}
--- PASS: TestInviteUserToGroup (0.12s)
PASS
*/
func TestInviteUserToGroup(t *testing.T)  {
	t.Setenv("UrlPrefix", "http://api.openim:10000")
	t.Setenv("DefaultUserID", `openIM123456`)
	t.Setenv("DefaultUserSecret", `tuoyun`)
	resp, err := v20220317.InviteUserToGroup(&apiStruct.ReqInviteUserToGroup{
		GroupID:           "cf8e2fc76b4b316097dab9eb517032c2",
		OperationID:       v20220317.NewOperationID(),
		InvitedUserIDList: []string{
			"openIMAdmin",
		},
		Reason:            "用户请求加入",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}
