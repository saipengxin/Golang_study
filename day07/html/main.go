package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// 加载模板文件，* 代表加载所有
	//r.LoadHTMLGlob("tmpl/index.html")
	//r.LoadHTMLGlob("tmpl/*")
	r.LoadHTMLFiles("tmpl/index.html") // LoadHTMLFiles只能指定到文件名字，不能指定文件路径

	r.GET("index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{"title": "我的标题"})
	})

	r.Run(":8800")
}
