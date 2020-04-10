package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// database/sql，打开和关闭数据库；管理数据库连接池。

func initDB() *sql.DB{
	// 连接数据库
	// sql.Open()仅初始化一个sql.DB对象，进行数据库操作时，才会建立真正的网络连接。
	// 不要频繁Open,Close。一般开启一次，可以多次调用，很少掉关闭。
	var err error
	myDB :=  MyDB{}
	myDB.DB, err = sql.Open("mysql", "root:root@/logindb")
	if err != nil {
		fmt.Printf("open mysql databases logindb fail, err: %s\n", err)
	}
	// 设置数据库的最大连接数
	myDB.DB.SetConnMaxLifetime(100)
	// 设置数据库最大的闲置连接数
	myDB.DB.SetMaxIdleConns(10)
	// 验证连接，此时会尝试连接
	if err = myDB.DB.Ping(); err != nil {
		fmt.Printf("Open mysql database fail, err:%s\n", err)
	}
	return myDB.DB
}

type MyDB struct {
	*sql.DB
}

func CloseDB(myDB *MyDB) {
	// 如果不释放会占用连接池中的资源和系统内存。
	myDB.DB.Close()
}

// db.Prepare()预编译，灵活、高效、可以防止sql注入攻击。
// Prepare + Exec 可以实现插入、更新、删除操作。
// 查询，直接使用db.Query()，返回Rows。迭代查询使用Rows.Next()，读取每一行使用Rows.Scan。最后记得defer Rows.Close()
// 插入数据
func InsertDB(myDB *MyDB) {
	stmt, _ := myDB.DB.Prepare("insert user_info set id=?,name=?,password=?")
	_, err := stmt.Exec(1, "admin", "admin")
	if err != nil {
		fmt.Printf("insert db err: %s\n", err)
	}

}

// 更新数据
func  UpdateDB(myDB *MyDB) {
	stmt, _ := myDB.DB.Prepare("update user_info set password=? where name=?")

	res, err := stmt.Exec("admin", "admin")
	if err != nil {
		fmt.Printf("update db err: %s\n", err)
	}
	fmt.Printf("result: %s\n", res)
}

// 查询数据
func  GetDB(myDB *MyDB){
	res, err := myDB.DB.Query("select * from user_info")
	if err != nil {
		fmt.Printf("get db err: %s\n", err)
	}
	defer res.Close()
	fmt.Printf("result: %s\n", res)

}
// 删除数据
func  DeleteDB(myDB *MyDB) {
	stmt, _ := myDB.DB.Prepare("delete from user_info where name=?")
	res, err := stmt.Exec("admin")
	if err != nil {
		fmt.Printf("delete db err: %s\n", err)
	}
	fmt.Printf("result: %s\n", res)
}