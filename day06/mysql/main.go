package main

import (
	"database/sql" // 使用数据库方法我们还是使用golang原生的，所以这个包也要导入
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 三方包我们只要他的init方法就可以了
)

var db *sql.DB // 这是一个连接池对象，连接都是从这里取出来的，定义成公共的变量

type user struct {
	id   int
	age  int
	name string
}

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
		fmt.Println("连接失败：", err)
		return
	}

	// 1.写查询的sql
	sqlStr1 := "delete from user where id = ?;"

	// 2.执行sql语句
	ret, err := db.Exec(sqlStr1, 3)
	if err != nil {
		fmt.Println("删除失败", err)
		return
	}

	// 获取受影响的行数，删除不返回LastInsertID,所以LastInsertID一直为0
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("获取受影响的行数失败", err)
		return
	}
	fmt.Printf("删除成功，%d行受影响\n", rows)
}
