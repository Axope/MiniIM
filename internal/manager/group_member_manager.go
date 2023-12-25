package manager

import (
	"MiniIM/internal/dao/pool"
	"MiniIM/internal/models"
)

type groupMamberManager struct {
}

var GroupMemberManager = new(groupMamberManager)

func (*groupMamberManager) AddMember(groupID uint, userID uint) error {
	newMember := &models.GroupMember{
		GroupID: groupID,
		UserID:  userID,
	}
	pool.GetDB().Create(&newMember)
	return nil
}
