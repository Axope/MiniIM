package manager

import (
	"MiniIM/internal/dao/pool"
	"MiniIM/internal/models"
	"MiniIM/pkg/common/request"
	"MiniIM/pkg/log"
)

type friendManager struct {
}

var FriendManager = new(friendManager)

func (f *friendManager) GetFriends(uuid string) ([]string, error) {
	user, err := UserManager.GetUserByUuid(uuid)
	if err != nil {
		return nil, err
	}

	var uuidList []string
	log.Logger.Sugar().Debugf("SELECT u.uuid FROM friends AS uf JOIN users AS u ON uf.friend_id = u.id WHERE uf.user_id = %s", user.ID)
	rows, err := pool.GetDB().Raw("SELECT u.uuid FROM friends AS uf JOIN users AS u ON uf.friend_id = u.id WHERE uf.user_id = ?", user.ID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var uuid string
		if err := rows.Scan(&uuid); err != nil {
			return nil, err
		}
		uuidList = append(uuidList, uuid)
	}

	// 检查 rows.Next() 的错误
	if err := rows.Err(); err != nil {
		return nil, err
	}
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
