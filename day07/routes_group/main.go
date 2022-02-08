package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 路由组1 处理GET请求
	v1 := r.Group("/v1")
	v1.GET("/login", login)
	v1.GET("/submit", submit)

	// 路由组2 处理POST请求
	// {} 花括号不是必须的，写上是一种语法规范，可以看起来结构更加清晰
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	r.Run(":8800")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "saipx")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "saipx")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
