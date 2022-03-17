package apiStruct

//注册新用户
//用户注册新账号时，AppServer 调用此在OpenIM系统中创建新用户。

/*
	参数名	类型	必选	说明
	secret	string	是	OpenIM秘钥
	platform	int	是	用户注册的平台类型
	userID	string	是	用户 ID
	nickname	string	是	用户昵称
	faceURL	string	否	用户头像URL
	gender	int	否	用户性别
	phoneNumber	string	否	用户手机号码
	birth	int	否	用户生日
	email	string	否	邮箱地址
	ex	string	否	扩展字段
	operationID	string	是	操作ID
*/
type ReqUserRegister struct {
	Secret      string   `json:"secret"`
	Platform    Platform `json:"platform"`
	UserID      string   `json:"userID"`
	Nickname    string   `json:"nickname"`
	FaceURL     string   `json:"faceURL"`
	Gender      Gender   `json:"gender"`
	PhoneNumber string   `json:"phoneNumber"`
	Birth       int64    `json:"birth"`
	Email       string   `json:"email"`
	Ex          string   `json:"ex"`
	OperationID string   `json:"operationID"`
}

/*
	参数名	类型	说明
	errCode	int	0成功，非0失败
	errMsg	string	错误信息
	userID	string	用户ID
	token	string	用户token
	expiredTime	int	token过期时间戳（秒）
*/
type RespUserRegister struct {
	UserID      string `json:"userID"`
	Token       string `json:"token"`
	ExpiredTime int    `json:"expiredTime"`
}
