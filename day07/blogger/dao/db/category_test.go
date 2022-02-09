package db

import (
	"github.com/saipengxin/study/day07/blogger/model"
	"testing"
)

func init() {
	// parseTime=true,将MySQL中的时间类型，自动解析为go结构体中的时间类型,不加报错
	dns := "root:root@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestInsertCategory(t *testing.T) {
	var category = &model.Category{}
	category.CategoryName = "测试分类1"
	category.CategoryNo = 1

	categoryId, err := InsertCategory(category)
	if err != nil {
		panic(err)
	}
	t.Logf("categoryId:%#v", categoryId)
}

func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err != nil {
		panic(err)
	}
	t.Logf("category:%#v", category)
}

func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64
	categoryIds = append(categoryIds, 1, 2)
	categoryList, err := GetCategoryList(categoryIds)
	if err != nil {
		panic(err)
	}
	for _, v := range categoryList {
		t.Logf("id:%d,categoryList:%#v", v.CategoryId, v)
	}

}

func TestGetAllCategoryList(t *testing.T) {
	categoryList, err := GetAllCategoryList()
	if err != nil {
		panic(err)
	}
	for _, v := range categoryList {
		t.Logf("id:%d,categoryList:%#v", v.CategoryId, v)
	}
}
