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

type groupAPI struct {
}

var GroupAPI = new(groupAPI)

// method: POST
//
// body
// - request.GroupRequest
//
// response
// - response: groupUuid
func (*groupAPI) CreateGroup(c *gin.Context) {
	var req request.GroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Logger.Debug("josn error", zap.Any("err", err))
		c.JSON(http.StatusBadRequest, response.Fail("json error"))
		return
	}

	if err := manager.GroupManager.CreateGroup(&req); err != nil {
		log.Logger.Debug("CreateGroup error", zap.Any("err", err))
		c.JSON(http.StatusOK, response.Fail(err.Error()))
		return
	}

	log.Logger.Sugar().Debugf("return json: %v", response.Success(req.GroupUuid))
	c.JSON(http.StatusOK, response.Success(req.GroupUuid))
}

// method: POST
//
// body
// - request.GroupRequest
//
// response
// - response: groupUuid
func (*groupAPI) JoinGroup(c *gin.Context) {
	var req request.GroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Logger.Debug("josn error", zap.Any("err", err))
		c.JSON(http.StatusBadRequest, response.Fail("json error"))
		return
	}

	if err := manager.GroupManager.JoinGroup(&req); err != nil {
		log.Logger.Debug("JoinGroup error", zap.Any("err", err))
		c.JSON(http.StatusOK, response.Fail(err.Error()))
		return
	}

	log.Logger.Sugar().Debugf("return json: %v", response.Success(req.GroupUuid))
	c.JSON(http.StatusOK, response.Success(req.GroupUuid))
}
