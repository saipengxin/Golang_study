package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 定义中间件
func myTime(c *gin.Context) {
	// 开始计时
	t := time.Now()

	// 执行函数
	c.Next()
	// 计算经过的时间
	t2 := time.Since(t)
	log.Printf("程序执行用时:%v\n", t2)
}

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 注册中间件
	r.Use(myTime)
	// {}为了代码规范
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}
	r.Run(":8800")
}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
}

func shopHomeHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
}
