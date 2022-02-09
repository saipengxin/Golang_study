package model

// Category 定义分类结构体
type Category struct {
	CategoryId   int64  `db:"id"`
	CategoryName string `db:"category_name"`
	CategoryNo   int64  `db:"category_no"`
}
