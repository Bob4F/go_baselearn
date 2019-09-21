package main
/*
   Go 学习 数据库操作
   database/sql 是Go 标准库提供的package 用于Sql和SQL-like 的通讯
	 1.sql.DB / DB
		代表一种数据库，能操作的数据库对象，具体的连接由内部连接池实现，对外不暴露
 */

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	// mysql的封装
)

// 第一次连接MySQL
// 并且查询
func FirstConnectionMySQL(){

	db,err := sql.Open("mysql","root:123456@tcp(119.3.222.192:3306)/go_db?charset=utf8")

	if err != nil {
		fmt.Println("firstConnectionMySQL Error INFO:",err)
		// panic(err)
	}
	var (
		id   int
		name string
	)
	// 查询用户id
	rows,sqlerr := db.Query("select id,name from test_table")
	if sqlerr != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id,&name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("id:",id,",name:",name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	Db.Exec("");
	fmt.Println("PING INFO:",db.Ping())
	// 关闭
	defer db.Close()
}

