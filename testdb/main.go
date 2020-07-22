package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testdb/dao"
	"testdb/services"
)

//文档
//https://blog.csdn.net/embinux/article/details/84031620

func main() {
	dao.Init()
	fmt.Print("test code")
	//StructQueryField()
	services.RetriveUserInfo(1)
	services.RetriveUserInfo(2)
	services.RetriveUserInfo(3)
}
