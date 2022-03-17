package apiStruct

//发送单聊群聊消息
//管理员通过后台接口发送单聊群聊消息，可以以管理员身份发消息，也可以以其他用户的身份发消息，通过sendID区分。

/*
	参数名	类型	必选	说明
	operationID	string	是	操作ID，保持唯一，建议用当前时间微秒+随机数
	sendID	string	是	发送者ID
	recvID	string	否	接收者ID，单聊为用户ID，如果是群聊，则不填
	groupID	string	否	群聊ID，如果为单聊，则不填
	senderNickname	string	否	发送者昵称
	senderFaceURL	string	否	发送者头像
	senderPlatformID	int	否	发送者平台号，模拟用户发送时填写， 1->IOS,2->Android,3->Windows,4->OSX,5->Web,5->MiniWeb,7->Linux
	forceList	json数组	否	当聊天类型为群聊时，使用@指定强推用户userID列表
	content	json对象	是	消息的具体内容，内部是json 对象，其他消息的详细字段请参考消息类型格式描述文档
	contentType	int	是	消息类型，101表示文本，102表示图片..详细参考消息类型格式描述文档
	sessionType	int	是	发送的消息是单聊还是群聊,单聊为1，群聊为2
	isOnlineOnly	bool	否	改字段设置为true时候，发送的消息服务器不会存储，接收者在线才会收到并存储到本地，不在线该消息丢失，当消息的类型为113->typing时候，接收者本地也不会做存储
	offlinePushInfo	json对象	否	离线推送的具体内容，如果不填写，使用服务器默认推送标题
	title	string	否	推送的标题
	desc	string	否	推送的具体描述
	ex	string	否	扩展字段
	iOSPushSound	string	否	IOS的推送声音
	iOSBadgeCount	bool	否	IOS推送消息是否计入桌面图标未读数
*/
type ReqSendMsg struct {
	OperationID      string          `json:"operationID"`
	SendID           string          `json:"sendID"`
	RecvID           string          `json:"recvID"`
	GroupID          string          `json:"groupID"`
	SenderNickname   string          `json:"senderNickname"`
	SenderFaceURL    string          `json:"senderFaceURL"`
	SenderPlatformID Platform        `json:"senderPlatformID"`
	ForceList        []string        `json:"forceList"`
	Content          MsgContext      `json:"content"`
	ContentType      MsgContentType  `json:"contentType"`
	SessionType      MsgSessionType  `json:"sessionType"`
	IsOnlineOnly     bool            `json:"isOnlineOnly"`
	OfflinePushInfo  OfflinePushInfo `json:"offlinePushInfo"`
}
type OfflinePushInfo struct {
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Ex            string `json:"ex"`
	IOSPushSound  string `json:"iOSPushSound"`
	IOSBadgeCount bool   `json:"iOSBadgeCount"`
}
type MsgContext struct {
	Text string `json:"text"`
}
type MsgContentType int
type MsgSessionType int

const (
	MsgSessionTypeSingle MsgSessionType = iota + 1 // 单聊
	MsgSessionTypeGroup                     // 群聊
)

const (
	MsgContentTypeText MsgContentType = iota + 101 // 101 文本
	Picture                                        // 图片
	Voice                                          // 语音
	Video                                          // 视频
	File                                           // 文件
	AtText                                         // @文字消息
	Merger                                         // 合并消息
	Card                                           // 卡片消息
	Location                                       // 定位信息
	Custom                                         // 自定义
	Revoke                                         // 撤回
	HasReadReceipt                                 // 已读
	Typing                                         // 输入中
	Quote                                          // 引用消息
)

/*
	参数名	类型	说明
	errCode	int	0成功，非0失败
	errMsg	string	错误信息
	sendTime	int	消息发送的具体时间，具体为毫秒的时间戳
	serverMsgID	string	服务器生成的消息的唯一ID
	clientMsgID	string	客户端生成的消息唯一ID，默认情况使用这个为主键
*/
type RespSendMsg struct {
	ServerMsgID string `json:"serverMsgID"`
	ClientMsgID string `json:"clientMsgID"`
	SendTime    int64  `json:"sendTime"`
}
