package main

import (
	"go-manager/router"
	_ "gorm.io/driver/mysql"
)

func main() {
	// 启动服务
	router.Server()
}
