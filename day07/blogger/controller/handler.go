package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saipengxin/study/day07/blogger/service"
	"net/http"
)

func IndexHandle(c *gin.Context) {
	articleRecordList, err := service.GetArticleRecordList(1, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	allCategoryList, err := service.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n", err)
	}
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": allCategoryList,
	})
}
