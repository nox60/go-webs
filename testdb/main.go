package main

import (
	_ "github.com/go-sql-driver/mysql"
	"testdb/dao"
	"testdb/services"
)

//文档
//https://blog.csdn.net/embinux/article/details/84031620

func main() {
	dao.Init()
	//StructQueryField()
	//services.RetriveUserInfo(1)
	//services.RetriveUserInfo(2)
	//services.RetriveUserInfo(3)
	services.InsertTest()

}
