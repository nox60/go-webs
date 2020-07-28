package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"testdb/controller"
	"testdb/dao"
)

//文档
//https://blog.csdn.net/embinux/article/details/84031620

func main() {
	dao.Init()
	//StructQueryField()
	//services.RetriveUserInfo(1)
	//services.RetriveUserInfo(2)
	//services.RetriveUserInfo(3)
	//services.InsertTestWithOutTx()

	r := gin.Default()

	api := r.Group("/simple-api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api.GET("/someGet", controller.SimpleLogin)
	r.Run() // listen and serve on 0.0.0.0:8080

}
