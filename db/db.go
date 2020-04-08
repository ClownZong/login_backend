package db

import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"

// 连接数据库
func OpenDB() {
	db, err := sql.Open("mysql", "root:root@/logindb")
	if err != nil {
		fmt.Printf("open mysql databases logindb err: %s\n", err)
	}
}

// 插入数据
func InsertDB() {

	stmt, err := db.Prepare("insert user_info set id=?,name=?,password=?")
	if err != nil {
		fmt.Printf("insert db err: %s\n", err)
	}
	res, err := stmt.Exec(1, "admin", "admin")

}

// 更新数据
func UpdateDB() {
	stmt, err := db.Prepare("update user_info set password=? where name=?")
	if err != nil {
		fmt.Printf("update db err: %s\n", err)
	}
	res, err := stmt.Exec("admin", "admin")
	fmt.Printf("result: %s\n", res)
}

// 查询数据
func GetDB(){
	res, err := db.Query("select * from user_info")
	if err != nil {
		fmt.Printf("get db err: %s\n", err)
	}
	fmt.Printf("result: %s\n", res)
}
// 删除数据
func DeleteDB() {
	stmt, err := db.Prepare("delete from user_info where name=?")
	if err != nil {
		fmt.Printf("delete db err: %s\n", err)
	}
	res, err := stmt.Exec("admin")
	fmt.Printf("result: %s\n", res)
}