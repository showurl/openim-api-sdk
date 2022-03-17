package apiStruct

//创建群组
//APP管理员创建群组

/*
	参数名	必选	说明
	memberList	是	指定初始群成员
	ownerUserID	是	群主UserID
	operationID	是	操作ID
*/
type ReqCreateGroup struct {
	OwnerUserID  string   `json:"ownerUserID"`
	GroupType    int      `json:"groupType"`
	GroupName    string   `json:"groupName"`
	Notification string   `json:"notification"`
	Introduction string   `json:"introduction"`
	FaceURL      string   `json:"faceURL"`
	Ex           string   `json:"ex"`
	OperationID  string   `json:"operationID"`
	MemberList   []Member `json:"memberList"`
}
type Member struct {
	UserID    string `json:"userID"`
	RoleLevel int    `json:"roleLevel"`
}

/*
	参数名	类型	说明
	errCode	int	0成功，非0失败
	errMsg	string	错误信息
	creatorUserID	json对象	创建者userID，
	groupID	string	群ID
	groupName	string	群名称
	memberCount	int	群成员数量
	ownerUserID	string	群主UserID
*/
type RespCreateGroup struct {
	CreatorUserID string `json:"creatorUserID"`
	GroupID       string `json:"groupID"`
	GroupName     string `json:"groupName"`
	MemberCount   int    `json:"memberCount"`
	OwnerUserID   string `json:"ownerUserID"`
}
