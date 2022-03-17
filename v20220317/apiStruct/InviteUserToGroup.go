package apiStruct

//邀请进群
//APP管理员邀请用户直接进群

/*
	参数名	类型	必选	说明
	groupID	string	是	群ID
	invitedUserIDList	json数组	是	被邀请进群的userID列表
	reason	string	否	进群原因
	operationID	string	是	操作ID
*/
type ReqInviteUserToGroup struct {
	GroupID           string   `json:"groupID"`
	OperationID       string   `json:"operationID"`
	InvitedUserIDList []string `json:"invitedUserIDList"`
	Reason            string   `json:"reason"`
}

/*
	参数名	类型	说明
	errCode	int	0成功，非0失败
	errMsg	string	错误信息
	userID	string	被邀请进群的UserID
	result	int	操作结果 0表示成功，其他表示失败
*/
type RespInviteUserToGroups []RespInviteUserToGroup
type RespInviteUserToGroup struct {
	UserID string `json:"userID"`
	Result int    `json:"result"`
}

func (r RespInviteUserToGroup) Success() bool {
	return r.Result == 0
}
