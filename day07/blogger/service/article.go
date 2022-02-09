package service

import (
	"fmt"
	"github.com/saipengxin/study/day07/blogger/dao/db"
	"github.com/saipengxin/study/day07/blogger/model"
)

// 获取文章和分类的列表
func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	// 获取文章列表
	articleInfoList, err := db.GetAricleList(pageNum, pageSize)
	if err != nil {
		fmt.Printf("get article list failed, err:%v\n", err)
		return
	}
	if len(articleInfoList) == 0 {
		return
	}
	// 将文章列表中的文章分类获取到
	categoryIds := getCategoryIds(articleInfoList)
	// 获取文章分类的信息
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n", err)
		return
	}

	// 将文章信息和文字对应的分类信息整合到一起
	for _, article := range articleInfoList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.CategoryId
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {
LABEL:
	for _, article := range articleInfoList {
		categoryId := article.CategoryId
		for _, id := range ids {
			if id == categoryId {
				continue LABEL
			}
		}
		ids = append(ids, categoryId)
	}
	return
}
