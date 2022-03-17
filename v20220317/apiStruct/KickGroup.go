package apiStruct

//把用户踢出群
//APP管理员把用户从群里直接踢出

/*
	参数名	类型	必选	说明
	groupID	string	是	群ID
	kickedUserIDList	json数组	是	被踢UserID列表
	reason	string	否	被踢原因
	operationID	string	是	操作ID
*/
type ReqKickGroup struct {
	GroupID          string   `json:"groupID"`
	OperationID      string   `json:"operationID"`
	Reason           string   `json:"reason"`
	KickedUserIDList []string `json:"kickedUserIDList"`
}

/*
	参数名	类型	说明
	errCode	int	0成功，非0失败
	errMsg	string	错误信息
	userID	string	被踢UserID
	result	int	0成功，非0失败
*/
type RespKickGroups []RespKickGroup

type RespKickGroup struct {
	UserID string `json:"userID"`
	Result int    `json:"result"`
}

func (r RespKickGroup) Success() bool {
	return r.Result == 0
}
