package router

import (
	"github.com/gin-gonic/gin"
	"go-manager/middleware"
)

func Server() {
	var Router = gin.Default()
	// 路由前缀
	ApiGroup := Router.Group("")
	// 放置中间件
	ApiGroup.Use(middleware.CheckToken())
	// 注册路由
	UserRouter(ApiGroup)
	// 启动服务
	Router.Run(":3000")
}
