package main

import (
	"github.com/gin-gonic/gin"
	"github.com/saipengxin/study/day07/blogger/controller"
	"github.com/saipengxin/study/day07/blogger/dao/db"
)

func main() {
	router := gin.Default()
	dns := "root:root@tcp(localhost:3306)/blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	// 加载静态资源文件（CSS JS）
	router.Static("/static/", "./static")
	// 加载模板文件
	router.LoadHTMLGlob("views/*")
	// 设置路由
	router.GET("/", controller.IndexHandle)

	router.Run(":8800")
}
