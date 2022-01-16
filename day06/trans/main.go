package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// 初始化连接
func initDb() (err error) {
	// 用户名:密码@tcp(ip:端口)/数据库名?参数键值对
	dsn := "root:root@tcp(127.0.0.1:3306)/gostudy?charset=utf8"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	db.SetMaxOpenConns(5) // 最大连接数量 5
	db.SetMaxIdleConns(3) // 最大空闲连接 3
	return nil
}

func main() {
	err := initDb()
	if err != nil {
		fmt.Println("连接数据库失败,err", err)
		return
	}

	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("开启事务失败, err:%v\n", err)
		return
	}

	sqlStr1 := "update user set money = money-500 where id = 1"
	res1, err := tx.Exec(sqlStr1)
	if err != nil {
		fmt.Println("sql语句执行失败，err", err)
		tx.Rollback()
		return
	}
	row1, err := res1.RowsAffected()
	if err != nil {
		fmt.Println("获取数据失败，err", err)
		tx.Rollback()
		return
	}

	sqlStr2 := "update user set money = money+500 where id = 2"
	res2, err := tx.Exec(sqlStr2)
	if err != nil {
		fmt.Println("sql语句执行失败，err", err)
		tx.Rollback()
		return
	}
	row2, err := res2.RowsAffected()
	if err != nil {
		fmt.Println("获取数据失败，err", err)
		tx.Rollback()
		return
	}

	if row1 == 1 && row2 == 1 {
		tx.Commit()
	} else {
		tx.Rollback()
	}
}
