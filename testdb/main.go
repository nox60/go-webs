package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testdb/dao"
	"testing"
	"time"
)

//文档
//https://blog.csdn.net/embinux/article/details/84031620

const (
	USER_NAME = "root"
	PASS_WORD = "kkkk!@#$.com"
	HOST      = "www.scbridge.cn"
	PORT      = "18999"
	DATABASE  = "infomana"
	CHARSET   = "utf8"
)

var MysqlDbErr error

func main() {
	testing.Init()
	fmt.Print("test code")
	//StructQueryField()
	dao.StructQueryField()
}

// 初始化链接
func init() {

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)

	// 打开连接失败
	dao.MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)
	//defer MysqlDb.Close();
	if MysqlDbErr != nil {
		log.Println("dbDSN: " + dbDSN)
		panic("数据源配置不正确: " + MysqlDbErr.Error())
	}

	// 最大连接数
	dao.MysqlDb.SetMaxOpenConns(100)
	// 闲置连接数
	dao.MysqlDb.SetMaxIdleConns(20)
	// 最大连接周期
	dao.MysqlDb.SetConnMaxLifetime(100 * time.Second)

	if MysqlDbErr = dao.MysqlDb.Ping(); nil != MysqlDbErr {
		panic("数据库链接失败: " + MysqlDbErr.Error())
	}

}