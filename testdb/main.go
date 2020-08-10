package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"testdb/controller"
	"testdb/dao"
	"testdb/utils"
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

	//以下接口不需要鉴权
	api.POST("/checkLogin", controller.JsonLogin)

	api.Use(Authorize())
	// 以下接口都需要鉴权，验证token的正确性
	api.GET("/userInfo", controller.UserInfo)
	api.POST("/listSampleData", controller.ListSampleData)
	api.POST("/addItem", controller.AddItem)
	api.DELETE("/deleteItem/:itemId", controller.DeleteItem)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 首先判断token解析是否合法，如果不合法则提示访问未授权
		xToken := c.Request.Header.Get("X-Token")

		parsedToken, err := utils.JwtParse(xToken)

		if err != nil {
			fmt.Println("拦截，不让通过")
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			return
		} else {
			fmt.Println(parsedToken)
			//每次请求只有要刷新token
			refreshedToken := utils.RefreshToken(parsedToken)
			fmt.Println(refreshedToken)

			fmt.Println("允许通过")
			//刷新token
			c.Writer.Header().Set("x-token-rep", refreshedToken)
			c.Next()

		}
	}
}
