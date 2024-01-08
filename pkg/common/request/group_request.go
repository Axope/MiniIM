package request

type GroupRequest struct {
	UserUuid  string `json:"userUuid"`
	GroupUuid string `json:"groupUuid"`
	GroupName string `json:"groupName"`
}
