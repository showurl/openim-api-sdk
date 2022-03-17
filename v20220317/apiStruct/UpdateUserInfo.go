package apiStruct

//更新用户信息
//APP管理员更新用户信息

/*
	参数名	类型	必选	说明
	operationID	string	是	操作ID，保持唯一，建议用当前时间微秒+随机数
	nickname	string	否	用户昵称或者群昵称
	faceURL	string	否	用户头像或者群头像url，根据上下文理解
	gender	int	否	用户性别 1 表示男，2 表示女
	phoneNumber	string	否	用户手机号码，包括地区，(如香港：+852-xxxxxxxx)，
	birth	int	否	用户生日，Unix时间戳（秒）
	email	string	否	邮箱地址
	ex	string	否	用户扩展信息
	userID	string	是	用户 ID，必须保证IM内唯一
*/
type ReqUpdateUserInfo struct {
	UserID      string `json:"userID"`
	OperationID string `json:"operationID"`
	Nickname    string `json:"nickname"`
	FaceURL     string `json:"faceURL"`
	Gender      Gender    `json:"gender"`
	PhoneNumber string `json:"phoneNumber"`
	Birth       int64    `json:"birth"`
	Email       string `json:"email"`
	Ex          string `json:"ex"`
}
type Gender int

const (
	GenderUnknown Gender = iota
	GenderMale
	GenderFemale
)

/*
	参数名	类型	说明
	errCode	int	0成功，非0失败
	errCode	string	错误信息
*/
type RespUpdateUserInfo struct {
}
