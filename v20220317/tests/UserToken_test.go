package tests

import (
	v20220317 "github.com/showurl/openim-api-sdk/v20220317"
	"github.com/showurl/openim-api-sdk/v20220317/apiStruct"
	"testing"
)
/*
=== RUN   TestUserToken
    UserToken_test.go:27: apiStruct.RespUserToken{UserID:"abcdefghi1234567890", Token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiJhYmNkZWZnaGkxMjM0NTY3ODkwIiwiUGxhdGZvcm0iOiJJT1MiLCJleHAiOjE5NjI4NzM1MzgsIm5iZiI6MTY0NzUxMzUzOCwiaWF0IjoxNjQ3NTEzNTM4fQ.IZtfh0SzuJXymmUJmlyi4FAPz5IhqrlSyL8xWT7MFMs", ExpiredTime:1962873538}
--- PASS: TestUserToken (0.02s)
PASS
 */
func TestUserToken(t *testing.T)  {
	t.Setenv("UrlPrefix", "http://api.openim:10000")
	t.Setenv("DefaultUserID", `openIM123456`)
	t.Setenv("DefaultUserSecret", `tuoyun`)
	resp, err := v20220317.UserToken(&apiStruct.ReqUserToken{
		Secret:      `tuoyun`,
		Platform:    apiStruct.IOSPlatformID,
		UserID:      "abcdefghi1234567890",
		OperationID: v20220317.NewOperationID(),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}
