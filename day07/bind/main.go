package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	// binding:"required" 表示字段必填
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	// 创建路由
	r := gin.Default()

	r.GET("/loginURI/:user/:password", func(c *gin.Context) {
		// 定义接收数据的变量
		var login Login

		// 解析路由上的数据并绑定到结构体
		if err := c.ShouldBindUri(&login); err != nil {
			// 返回错误信息
			// gin.H 是一个自定义类型，我们可以借助这个类型自定义json格式的数据
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if login.User != "admin" || login.Password != "123456" {
			c.JSON(http.StatusBadRequest, gin.H{"state": 304})
			return
		}
		// c.JSON 返回json格式的数据，c.String返回字符串
		c.JSON(http.StatusOK, gin.H{"state": 200})
	})
	r.Run(":8800")
}
