package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
	"time"
)

const (
	USER_NAME = "root"
	PASS_WORD = "kkkk!@#$.com"
	HOST      = "www.scbridge.cn"
	PORT      = "18999"
	DATABASE  = "infomana"
	CHARSET   = "utf8"
)

var MysqlDb *sql.DB
var MysqlDbErr error

func main() {
	testing.Init()
	fmt.Print("test code")
	StructQueryField()
}

// 初始化链接
func init() {

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)

	// 打开连接失败
	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)
	//defer MysqlDb.Close();
	if MysqlDbErr != nil {
		log.Println("dbDSN: " + dbDSN)
		panic("数据源配置不正确: " + MysqlDbErr.Error())
	}

	// 最大连接数
	MysqlDb.SetMaxOpenConns(100)
	// 闲置连接数
	MysqlDb.SetMaxIdleConns(20)
	// 最大连接周期
	MysqlDb.SetConnMaxLifetime(100 * time.Second)

	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		panic("数据库链接失败: " + MysqlDbErr.Error())
	}

}

// 用户表结构体
type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// 查询数据，指定字段名
func StructQueryField() {

	user := new(User)
	row := MysqlDb.QueryRow("select accountId, userName, age from tb_users where accountId = ? ", 3)
	if err := row.Scan(&user.Id, &user.Name, &user.Age); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return
	}
	fmt.Println(user.Id, user.Name, user.Age)
}






