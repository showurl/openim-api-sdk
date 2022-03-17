package tests

import (
	v20220317 "github.com/showurl/openim-api-sdk/v20220317"
	"github.com/showurl/openim-api-sdk/v20220317/apiStruct"
	"testing"
	"time"
)

/*
=== RUN   TestUpdateUserInfo
    UpdateUserInfo_test.go:28: apiStruct.RespUpdateUserInfo{}
--- PASS: TestUpdateUserInfo (0.14s)
PASS
*/
func TestUpdateUserInfo(t *testing.T) {
	t.Setenv("UrlPrefix", "http://api.openim:10000")
	t.Setenv("DefaultUserID", `openIM123456`)
	t.Setenv("DefaultUserSecret", `tuoyun`)
	resp, err := v20220317.UpdateUserInfo(&apiStruct.ReqUpdateUserInfo{
		UserID:      "15666355528",
		OperationID: v20220317.NewOperationID(),
		Nickname:    "改成28了",
		FaceURL:     "https://go.dev/images/gophers/pilot-bust.svg",
		Gender:      apiStruct.GenderMale,
		PhoneNumber: "15666355528",
		Birth:       time.Now().Unix(),
		Email:       "showurl@163.com",
		Ex:          `{"k","v"}`,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}
