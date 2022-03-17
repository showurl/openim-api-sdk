package apiStruct

type ResponseBase struct {
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
	Data    struct {
		UserID      string `json:"userID"`
		Token       string `json:"token"`
		ExpiredTime int    `json:"expiredTime"`
	} `json:"data"`
}
