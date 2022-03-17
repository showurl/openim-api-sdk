package tests

import (
	v20220317 "github.com/showurl/openim-api-sdk/v20220317"
	"github.com/showurl/openim-api-sdk/v20220317/apiStruct"
	"testing"
)

/*
=== RUN   TestKickGroup
    KickGroup_test.go:24: apiStruct.RespKickGroups{apiStruct.RespKickGroup{UserID:"openIM123456", Result:-1}}
--- PASS: TestKickGroup (0.10s)
PASS
 */
func TestKickGroup(t *testing.T)  {
	t.Setenv("UrlPrefix", "http://api.openim:10000")
	t.Setenv("DefaultUserID", `openIM123456`)
	t.Setenv("DefaultUserSecret", `tuoyun`)
	resp, err := v20220317.KickGroup(&apiStruct.ReqKickGroup{
		GroupID:          "cf8e2fc76b4b316097dab9eb517032c2",
		OperationID:      v20220317.NewOperationID(),
		Reason:           "用户退出群组",
		KickedUserIDList: []string{
			`openIM123456`,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}
