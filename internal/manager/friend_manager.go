package manager

import (
	"MiniIM/internal/dao/pool"
	"MiniIM/internal/models"
	"MiniIM/pkg/common/request"
	"fmt"
)

type friendManager struct {
}

var FriendManager = new(friendManager)

func (f *friendManager) GetFriends(uuid string) ([]string, error) {
	if !UserManager.CheckUserExistByUuid(uuid) {
		return nil, fmt.Errorf("user(uuid = %s) not exists", uuid)
	}

	var uuidList []string
	pool.GetDB().Raw("SELECT u.uuid FROM friends AS uf JOIN users AS u ON uf.friend_id = u.id WHERE uf.user_id = ?", uuid).Scan(&uuidList)
	return uuidList, nil
}

func (f *friendManager) AddFriends(req request.FriendRequest) error {
	user, err := UserManager.GetUserByUuid(req.UserUuid)
	if err != nil {
		return err
	}
	friend, err := UserManager.GetUserByUuid(req.FriendUuid)
	if err != nil {
		return err
	}

	friends := &models.Friend{
		UserID:   user.ID,
		FriendID: friend.ID,
	}
	pool.GetDB().Create(&friends)
	return nil
}

func (f *friendManager) DelFriends(req request.FriendRequest) error {
	user, err := UserManager.GetUserByUuid(req.UserUuid)
	if err != nil {
		return err
	}
	friend, err := UserManager.GetUserByUuid(req.FriendUuid)
	if err != nil {
		return err
	}

	pool.GetDB().Where("user_id = ? AND friend_id = ?", user.ID, friend.ID).Delete(&models.Friend{})
	return nil
}
