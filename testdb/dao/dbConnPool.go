package dao

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	USER_NAME = "root"
	PASS_WORD = "kkkk!@#$.com"
	HOST      = "www.scbridge.cn"
	PORT      = "18999"
	DATABASE  = "car1"
	CHARSET   = "utf8"
)

var MysqlDb *sql.DB
var MysqlDbErr error

// 初始化链接
func Init() {

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
