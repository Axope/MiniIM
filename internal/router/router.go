package router

import (
	"net/http"

	"MiniIM/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Create() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/echo", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": ctx.PostForm("text"),
		})
	})

	router.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome!",
		})
	})

	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", api.UserAPI.Register)
		userGroup.POST("/login", api.UserAPI.Login)
	}

	friendGroup := router.Group("/friends")
	{
		friendGroup.GET("/getFriends", api.FriendAPI.GetFriends)
		friendGroup.POST("/addFriends", api.FriendAPI.AddFriends)
		friendGroup.DELETE("/delFriends", api.FriendAPI.DelFriends)
	}

	router.GET("/socket", Socket)

	return router
}
