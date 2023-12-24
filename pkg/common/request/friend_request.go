package request

type FriendRequest struct {
	UserUuid   string `json:"userUuid"`
	FriendUuid string `json:"friendUuid"`
}
