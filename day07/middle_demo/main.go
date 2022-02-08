package main

import "github.com/gin-gonic/gin"

func Demo1(c *gin.Context) {
	c.String(200, "中间件开始执行了\n")
	c.Next() // 手动调用路由处理函数执行
	c.String(200, "中间件执行结束了\n")
}

func Demo2(c *gin.Context) {
	c.String(200, "中间件开始执行了\n")
	c.String(200, "中间件执行结束了\n")
	// 等待中间件执行完成后自动执行。
}

func main() {
	r := gin.Default()

	r.GET("demo1", Demo1, func(c *gin.Context) {
		c.String(200, "路由处理函数执行\n")
	})
	r.GET("demo2", Demo2, func(c *gin.Context) {
		c.String(200, "路由处理函数执行\n")
	})

	r.Run(":8800")
}
