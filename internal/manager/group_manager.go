package manager

import (
	"MiniIM/internal/dao/pool"
	"MiniIM/internal/models"
	"MiniIM/pkg/common/request"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type groupManager struct {
}

var GroupManager = new(groupManager)

func GetGroupByUuid(groupUuid string) (*models.Group, error) {
	db := pool.GetDB()
	var queryGroup models.Group
	if err := db.First(&queryGroup, "uuid = ?", groupUuid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 记录不存在
			return nil, fmt.Errorf("group uuid(%s) not found", groupUuid)
		} else {
			// 其他错误
			return nil, err
		}
	}
	return &queryGroup, nil
}

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

func (*groupManager) JoinGroup(req *request.GroupRequest) error {
	user, err := UserManager.GetUserByUuid(req.UserUuid)
	if err != nil {
		return err
	}

	group, err := GetGroupByUuid(req.GroupUuid)
	if err != nil {
		return err
	}

	if err := GroupMemberManager.AddMember(group.ID, user.ID); err != nil {
		return err
	}
	return nil
}
