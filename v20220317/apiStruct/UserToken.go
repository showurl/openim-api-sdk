package apiStruct

//获取用户token
//AppServer调用此接口获取token，IMSDK在login时需传入token
//换取管理员Token
//AppServer 通过调用（auth/user_token）获取或刷新管理员token以取得超级权限。在调用以下所有的api时，必须获取管理员token，并设置到请求header中。如果没有特别说明，请求方式一律为POST

/*
	参数名	类型	必选	说明
	secret	string	是	OpenIM秘钥 OpenIM秘钥
	platform	int	是	用户登录的平台类型 管理员填8
	userID	string	是	用户ID 管理员userID，此处的userID必须为配置文件config/config.yaml的appManagerUid的其中一个
	operationID	string	是	操作ID
*/
type Platform int

const (
	IOSPlatformID     Platform = 1
	AndroidPlatformID Platform = 2
	WindowsPlatformID Platform = 3
	OSXPlatformID     Platform = 4
	WebPlatformID     Platform = 5
	MiniWebPlatformID Platform = 6
	LinuxPlatformID   Platform = 7
	AdminPlatformID   Platform = 8
)

type ReqUserToken struct {
	Secret      string   `json:"secret"`
	Platform    Platform `json:"platform"`
	UserID      string   `json:"userID"`
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
type RespUserToken struct {
	UserID      string `json:"userID"`
	Token       string `json:"token"`
	ExpiredTime int    `json:"expiredTime"`
}
