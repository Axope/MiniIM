package api

import (
	"MiniIM/internal/manager"
	"MiniIM/internal/models"
	"MiniIM/pkg/common/response"
	"MiniIM/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type userAPI struct {
}

var UserAPI = new(userAPI)

// 注册或登录成功，response 中的 data 仅有uuid
// curl -H "Content-Type: application/json" -X POST --data '{"username": "test123","password": "qqqq"}' http://localhost:9876/user/register
func (u *userAPI) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Logger.Debug("josn error", zap.Any("err", err))
		c.JSON(http.StatusBadRequest, response.Fail("json error"))
		return
	}

	log.Logger.Debug("register", zap.Any("username", user.Username), zap.Any("password", user.Password))
	if err := manager.UserManager.Register(&user); err != nil {
		log.Logger.Debug("register error", zap.Any("err", err))
		c.JSON(http.StatusOK, response.Fail(err.Error()))
		return
	}

	log.Logger.Sugar().Infof("user[%s] register completed", user.Username)
	log.Logger.Sugar().Debugf("return json: %v", response.Success(user.Uuid))
	c.JSON(http.StatusOK, response.Success(user.Uuid))
}

func (u *userAPI) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Logger.Debug("josn error", zap.Any("err", err))
		c.JSON(http.StatusBadRequest, response.Fail("json error"))
		return
	}
	log.Logger.Debug("login", zap.Any("username", user.Username), zap.Any("password", user.Password))

	if err := manager.UserManager.Login(&user); err != nil {
		log.Logger.Debug("login error", zap.Any("err", err))
		c.JSON(http.StatusOK, response.Fail(err.Error()))
		return
	}

	log.Logger.Sugar().Infof("user[%s] login completed", user.Username)
	log.Logger.Sugar().Debugf("return json: %v", response.Success(user.Uuid))
	c.JSON(http.StatusOK, response.Success(user.Uuid))
}
