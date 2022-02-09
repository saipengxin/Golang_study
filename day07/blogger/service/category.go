package service

import (
	"fmt"
	"github.com/saipengxin/study/day07/blogger/dao/db"
	"github.com/saipengxin/study/day07/blogger/model"
)

func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n", err)
		return
	}
	return
}
