package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/saipengxin/study/day07/blogger/model"
)

// InsertCategory 添加分类
func InsertCategory(category *model.Category) (categoryId int64, err error) {
	sqlStr := "insert into category(category_name,category_no) value (?,?)"
	result, err := DB.Exec(sqlStr, category.CategoryName, category.CategoryNo)
	if err != nil {
		return
	}
	categoryId, err = result.LastInsertId()
	return
}

// GetCategoryById 获取单个分类
func GetCategoryById(id int64) (category *model.Category, err error) {
	category = &model.Category{}
	sqlStr := "select id,category_name,category_no from category where id = ?"
	err = DB.Get(category, sqlStr, id)
	return
}

// GetCategoryList 获取多个分类
func GetCategoryList(categoryIds []int64) (categoryList []*model.Category, err error) {
	sqlStr, args, err := sqlx.In("select id,category_name,category_no from category where id in (?)", categoryIds)
	if err != nil {
		return
	}
	// 查询
	err = DB.Select(&categoryList, sqlStr, args...)
	return
}

// GetAllCategoryList 获取所有分类
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	sqlStr := "select id,category_name,category_no from category order by category_no asc"
	err = DB.Select(&categoryList, sqlStr)
	return
}
