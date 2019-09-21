package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	// 简化操作
	"github.com/jmoiron/sqlx"
)

type member struct {
	Id   int `db:"id"`
	Name string `db:"name"`
}

func main() {
	// InsertData()
	//SelectIds()
	SelectNames()
	//SelectData()
	id := InsertData("Jack")
	SelectNames()
	update("alibaba",id)
	deleteResult := DeleteDataById(3)
	if deleteResult > 0{
		fmt.Println("删除成功")
	} else {
		fmt.Println("删除失败")
	}
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

// insert data And return id
func InsertData(name string)int{
	// id自增长
	r, err := Db.Exec("insert into test_table(id, name)values(?, ?)", "0", name)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return 0
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return 0
	}
	fmt.Println("insert succ:", id)
	return int(id)
}

// update if success return 1 else 0
func update(name string,id int){
	res,err := Db.Exec("update test_table set name=? where id=?",name,id)
	if err!=nil {
		fmt.Println("exec failed: ", err)
	}
	row,err:=res.RowsAffected()
	if err!=nil {
		fmt.Println("rows failed, ",err)
	}
	// 受影响行数
	fmt.Println("update succ:",row)
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

// select id and name ;struct property  first word must cap
func SelectData(){
	datas := []member{}
	err := Db.Select(&datas, "select id,name from test_table where id=?",1)
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	fmt.Println("select succ:", datas)
}

//Return the number of affected rows
func DeleteDataById(id int)int64{
	r,err := Db.Exec("DELETE FROM test_table WHERE id =?",id)
	if err != nil {
		fmt.Println("exec failed: ", err)
	}
	row,err:=r.RowsAffected()
	if err!=nil {
		fmt.Println("rows failed, ",err)
	}
	return row
}