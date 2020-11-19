package router

import (
	"github.com/gin-gonic/gin"
	"go-manager/service"
)

func UserRouter(group *gin.RouterGroup) {
	group.POST("user", service.SaveUser)
	group.GET("user/:id", service.GetUserById)
}
