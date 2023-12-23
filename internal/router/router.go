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
	// router.Use(Cors())

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

		// userGroup.DELETE("/deleteUser/:uuid", api.UserAPI.DeleteUser())
		// userGroup.PUT("/:uuid", server.UpdateUser())
		// userGroup.GET("/QueryUser", server.QueryUser())
	}

	router.GET("/socket", Socket)

	return router
}

// func Cors() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		method := c.Request.Method
// 		origin := c.Request.Header.Get("Origin") //请求头部
// 		if origin != "" {
// 			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
// 			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
// 			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
// 			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
// 			c.Header("Access-Control-Allow-Credentials", "true")
// 		}
// 		//允许类型校验
// 		if method == "OPTIONS" {
// 			c.JSON(http.StatusOK, "ok!")
// 		}

// 		defer func() {
// 			if err := recover(); err != nil {
// 				log.Logger.Error("HttpError", zap.Any("HttpError", err))
// 			}
// 		}()

// 		c.Next()
// 	}
// }
