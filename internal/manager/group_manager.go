package manager

import (
	"MiniIM/internal/dao/pool"
	"MiniIM/internal/models"
	"MiniIM/pkg/common/request"

	"github.com/google/uuid"
)

type groupManager struct {
}

var GroupManager = new(groupManager)

func (*groupManager) CreateGroup(req *request.GroupRequest) error {
	user, err := UserManager.GetUserByUuid(req.UserUuid)
	if err != nil {
		return err
	}

	req.GroupUuid = uuid.New().String()
	newGroup := &models.Group{
		Uuid:   req.GroupUuid,
		UserID: user.ID,
		Name:   req.GroupName,
	}

	pool.GetDB().Create(&newGroup)
	if err := GroupMemberManager.AddMember(newGroup.ID, user.ID); err != nil {
		return err
	}
	return nil
}

// func (*groupManager) JoinGroup(req *request.GroupRequest) error {
// 	user, err := UserManager.GetUserByUuid(req.UserUuid)
// 	if err != nil {
// 		return err
// 	}

// }
