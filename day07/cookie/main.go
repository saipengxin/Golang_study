package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckLogin 定义中间件，检查是否存在cookie
func CheckLogin(c *gin.Context) {
	cookie, err := c.Cookie("login")
	if err == nil {
		if cookie == "123" {
			c.Next()
			return
		}
	}

	// 返回错误
	c.JSON(http.StatusUnauthorized, gin.H{"error": "请登录"})
	// 不调用后续的函数处理
	//c.Abort()
	return
}

func main() {
	r := gin.Default()

	r.GET("login", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("login", "123", 60, "/", "127.0.0.1", true, true)
	})
	r.GET("home", CheckLogin, func(c *gin.Context) {
		c.String(200, "首页")
	})

	r.Run(":8800")
}
