package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "log"

	// 简化操作
	"github.com/jmoiron/sqlx"
)

type member struct {
	id   int64
	name string
}

func main() {
	// InsertData()
	SelectIds()
	SelectNames()
	SelectData()


}


// 全局数据库变量 DB
var Db *sqlx.DB

// 初始化 go初始化调用方法
func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(119.3.222.192:3306)/go_db?charset=utf8")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

// 插入数据
func InsertData(){
	// id自增长
	r, err := Db.Exec("insert into test_table(id, name)values(?, ?)", "0", "mario3")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}



func SelectIds(){
	// 获取id集合没问题
	ids  := []int{}
	err := Db.Select(&ids, "select id from test_table")
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}

	fmt.Println("select succ:", ids)
}


func SelectNames(){
	// 获取id集合没问题
	names  := []string{}
	err := Db.Select(&names, "select name from test_table")
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}

	fmt.Println("select succ:", names)
}

func SelectData(){
	var datas []member
	err := Db.Select(&datas, "select * from test_table")
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}

	fmt.Println("select succ:", datas)
}

