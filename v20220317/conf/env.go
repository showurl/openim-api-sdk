package conf

import "os"

var (
	UrlPrefix         = `http://127.0.0.1:10000`
	DefaultUserID     = `openIM123456`
	DefaultUserSecret = `tuoyun`
)

func init() {
	if v := os.Getenv("UrlPrefix"); v != "" {
		UrlPrefix = v
	}
	if v := os.Getenv("DefaultUserID"); v != "" {
		DefaultUserID = v
	}
	if v := os.Getenv("DefaultUserSecret"); v != "" {
		DefaultUserSecret = v
	}
}
