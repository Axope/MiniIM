package manager

import (
	"errors"
	"fmt"
	"MiniIM/internal/dao/pool"
	"MiniIM/internal/models"
	"MiniIM/pkg/log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userManager struct {
}

var UserManager = new(userManager)

func (u *userManager) CreateUser(user *models.User) error {
	if u.CheckUserExistByName(user.Username) {
		return fmt.Errorf("user[%s] already exists", user.Username)
	}
	user.Uuid = uuid.New().String()
	pool.GetDB().Create(&user)
	return nil
}

func (u *userManager) DeleteUser(uuid string) error {
	queryUser, err := u.GetUserByUuid(uuid)
	if err != nil {
		return err
	}

	pool.GetDB().Delete(&queryUser)
	return nil
}

func (u *userManager) UpdateUser(user *models.User) error {
	queryUser, err := u.GetUserByUuid(user.Uuid)
	if err != nil {
		return err
	}

	queryUser.Password = user.Password
	pool.GetDB().Save(queryUser)
	return nil
}

func (u *userManager) GetUserByUuid(uuid string) (*models.User, error) {
	db := pool.GetDB()
	var queryUser models.User
	if err := db.First(&queryUser, "uuid = ?", uuid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 记录不存在
			return nil, fmt.Errorf("uuid(%s) not found", uuid)
		} else {
			// 其他错误
			return nil, err
		}
	}
	return &queryUser, nil
}

func (u *userManager) GetUserByName(name string) (*models.User, error) {
	db := pool.GetDB()
	var queryUser models.User
	if err := db.First(&queryUser, "username = ?", name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 记录不存在
			return nil, fmt.Errorf("username(%s) not found", name)
		} else {
			// 其他错误
			return nil, err
		}
	}
	return &queryUser, nil
}

func (u *userManager) CheckUserExistByName(name string) bool {
	db := pool.GetDB()
	var queryUser models.User
	if err := db.First(&queryUser, "username = ?", name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 记录不存在
			return false
		} else {
			// 其他错误
			log.Logger.Sugar().Errorf("CheckUserExistByName Other error:", err)
			return false
		}
	}
	return true
}

func (u *userManager) CheckUserExistByUuid(uuid string) bool {
	db := pool.GetDB()
	var queryUser models.User
	if err := db.First(&queryUser, "uuid = ?", uuid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 记录不存在
			return false
		} else {
			// 其他错误
			log.Logger.Sugar().Errorf("CheckUserExistByUuid Other error:", err)
			return false
		}
	}
	return true
}

// 成功还会修改user的uuid
func (u *userManager) Register(user *models.User) error {
	// check...
	return u.CreateUser(user)
}

// 成功还会修改user的uuid
func (u *userManager) Login(user *models.User) error {
	queryUser, err := u.GetUserByName(user.Username)
	if err != nil {
		return err
	}
	if queryUser.Password != user.Password {
		return fmt.Errorf("password error")
	}
	user.Uuid = queryUser.Uuid
	return nil
}
