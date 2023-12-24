package api

import (
	"MiniIM/internal/manager"
	"MiniIM/pkg/common/request"
	"MiniIM/pkg/common/response"
	"MiniIM/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type friendAPI struct {
}

var FriendAPI = new(friendAPI)

// method: GET
//
// param
// - uuid: string
//
// response
// - uuidList: array
func (f *friendAPI) GetFriends(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		log.Logger.Debug("uuid empty")
		c.JSON(http.StatusBadRequest, response.Fail("uuid empty"))
		return
	}
	uuidList, err := manager.FriendManager.GetFriends(uuid)
	if err != nil {
		log.Logger.Debug("GetFriends error", zap.Any("err", err))
		c.JSON(http.StatusOK, response.Fail(err.Error()))
		return
	}
	log.Logger.Sugar().Debugf("return json: %v", response.Success(uuidList))
	c.JSON(http.StatusOK, response.Success(uuidList))
}

// method: POST
//
// body
// - userUuid: string
// - friendUuid: string
//
// response
// - response: friendID
func (f *friendAPI) AddFriends(c *gin.Context) {
	var req request.FriendRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Logger.Debug("josn error", zap.Any("err", err))
		c.JSON(http.StatusBadRequest, response.Fail("json error"))
		return
	}

	if err := manager.FriendManager.AddFriends(req); err != nil {
		log.Logger.Debug("AddFriends error", zap.Any("err", err))
		c.JSON(http.StatusOK, response.Fail(err.Error()))
		return
	}

	log.Logger.Sugar().Debugf("return json: %v", response.Success(req))
	c.JSON(http.StatusOK, response.Success(req))
}

// method: POST
//
// body
// - userUuid: string
// - friendUuid: string
//
// response
// - response: friendID
func (f *friendAPI) DelFriends(c *gin.Context) {
	var req request.FriendRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Logger.Debug("josn error", zap.Any("err", err))
		c.JSON(http.StatusBadRequest, response.Fail("json error"))
		return
	}

	if err := manager.FriendManager.DelFriends(req); err != nil {
		log.Logger.Debug("DelFriends error", zap.Any("err", err))
		c.JSON(http.StatusOK, response.Fail(err.Error()))
		return
	}

	log.Logger.Sugar().Debugf("return json: %v", response.Success(req))
	c.JSON(http.StatusOK, response.Success(req))
}
