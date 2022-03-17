package tests

import (
	v20220317 "github.com/showurl/openim-api-sdk/v20220317"
	"github.com/showurl/openim-api-sdk/v20220317/apiStruct"
	"github.com/showurl/openim-api-sdk/v20220317/conf"
	"testing"
	"time"
)

/*
=== RUN   TestUserRegister
    UserRegister_test.go:32: apiStruct.RespUserRegister{UserID:"abcdefghi1234567890", Token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiJhYmNkZWZnaGkxMjM0NTY3ODkwIiwiUGxhdGZvcm0iOiJMaW51eCIsImV4cCI6MTk2Mjg3MzE3OCwibmJmIjoxNjQ3NTEzMTc4LCJpYXQiOjE2NDc1MTMxNzh9.sg_pmomXqg4g7wsoavr3zQucktBvXFT68ADvhFr_8OQ", ExpiredTime:1962873178}
--- PASS: TestUserRegister (9.68s)
PASS
 */
func TestUserRegister(t *testing.T)  {
	t.Setenv("UrlPrefix", "http://api.openim:10000")
	t.Setenv("DefaultUserID", `openIM123456`)
	t.Setenv("DefaultUserSecret", `tuoyun`)
	resp, err := v20220317.UserRegister(&apiStruct.ReqUserRegister{
		Secret:      conf.DefaultUserSecret,
		Platform:    apiStruct.LinuxPlatformID,
		UserID:      "abcdefghi1234567890",
		Nickname:    "管理员创建的用户",
		FaceURL:     "https://go.dev/images/gophers/pilot-bust.svg",
		Gender:      apiStruct.GenderMale,
		PhoneNumber: "15666355530",
		Birth:       time.Now().Unix(),
		Email:       "showurl1@163.com",
		Ex:          `{"k","v"}`,
		OperationID: v20220317.NewOperationID(),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}
