package main

import (
	"database/sql" // 使用数据库方法我们还是使用golang原生的，所以这个包也要导入
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 三方包我们只要他的init方法就可以了
)

func main() {
	// 用户名:密码@tcp(ip:端口)/数据库名?参数键值对
	dsn := "root:root@tcp(127.0.0.1:3306)/gostudy?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("open mysql fail:", err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败：", err)
		return
	} else {
		fmt.Println("连接成功")
	}

}
