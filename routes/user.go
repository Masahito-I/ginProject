// routes/user.go

package routes

import (
	"github.com/gin-gonic/gin"
	"ginProject/handlers"
)

func InitUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/:id", handlers.GetUser)
		userGroup.POST("/", handlers.CreateUser)
	}
}
