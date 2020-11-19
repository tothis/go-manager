package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("token") == "token" {
			println("--- 校验token成功 ---")
			// 验证通过继续执行下一个中间件
			c.Next()
		} else {
			// 验证不通过时终止请求
			c.Abort()
			println("--- 校验token ---")
			c.JSON(http.StatusUnauthorized, "验证失败")
		}
	}
}
