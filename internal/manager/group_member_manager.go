package manager

import (
	"MiniIM/internal/dao/pool"
	"MiniIM/internal/models"
	"MiniIM/pkg/log"
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

// 查询user的全部group的uuid
func (*groupMamberManager) GetGroupsByUuid(uuid string) ([]string, error) {
	user, err := UserManager.GetUserByUuid(uuid)
	if err != nil {
		return nil, err
	}

	var groupUuids []string
	log.Logger.Sugar().Debugf("SELECT g.uuid FROM group_members AS gm JOIN `groups` AS g ON gm.group_id = g.id WHERE gm.user_id = %s", user.ID)
	rows, err := pool.GetDB().Raw("SELECT g.uuid FROM group_members AS gm JOIN `groups` AS g ON gm.group_id = g.id WHERE gm.user_id = ?", user.ID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var uuid string
		if err := rows.Scan(&uuid); err != nil {
			return nil, err
		}
		groupUuids = append(groupUuids, uuid)
	}

	// 检查 rows.Next() 的错误
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return groupUuids, nil
}
