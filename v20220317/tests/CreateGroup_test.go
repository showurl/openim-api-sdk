package tests

import (
	v20220317 "github.com/showurl/openim-api-sdk/v20220317"
	"github.com/showurl/openim-api-sdk/v20220317/apiStruct"
	"github.com/showurl/openim-api-sdk/v20220317/conf"
	"testing"
)
/*
=== RUN   TestCreateGroup
    CreateGroup_test.go:31: apiStruct.RespCreateGroup{CreatorUserID:"openIM123456", GroupID:"cf8e2fc76b4b316097dab9eb517032c2", GroupName:"管理员群1", MemberCount:2, OwnerUserID:"openIM123456"}
--- PASS: TestCreateGroup (0.30s)
PASS
 */
func TestCreateGroup(t *testing.T)  {
	t.Setenv("UrlPrefix", "http://api.openim:10000")
	t.Setenv("DefaultUserID", `openIM123456`)
	t.Setenv("DefaultUserSecret", `tuoyun`)
	resp, err := v20220317.CreateGroup(&apiStruct.ReqCreateGroup{
		OwnerUserID:  conf.DefaultUserID,
		GroupType:    1,
		GroupName:    "管理员群1",
		Notification: "",
		Introduction: "",
		FaceURL:      "https://go.dev/images/gophers/pilot-bust.svg",
		Ex:           `{"k","v"}`,
		OperationID:  v20220317.NewOperationID(),
		MemberList: []apiStruct.Member{{
			UserID:    "15666355529",
			RoleLevel: 1,
		}},
	}, v20220317.WithUser(`openIM333`, conf.DefaultUserSecret))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}
