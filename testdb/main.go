package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"testdb/controller"
	"testdb/dao"
	"time"
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
	api.POST("/login", controller.Login)

	api.Use(Authorize())
	// 以下接口都需要鉴权，验证token的正确性
	api.POST("/checkLogin", controller.JsonLogin)
	api.GET("/info", controller.Info)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 首先判断token解析是否合法，如果不合法则提示访问未授权
		xToken := c.Request.Header.Get("X-Token")

		if xToken == "" {
			fmt.Println("need x_token")
		}

		// 判断用户是否有请求相关接口的权限

		fmt.Println("before login......")
		test := c.Query("test")
		if test == "1" {
			fmt.Println("拦截，不让通过")
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			return
		} else {
			fmt.Println("允许通过")
			c.Next()
		}
	}
}

func JwtSign() {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	fmt.Fprintln(w, "Error extracting the key")
	//	fatal(err)
	//}

	//tokenString, err := token.SignedString([]byte(SecretKey))
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	fmt.Fprintln(w, "Error while signing the token")
	//	fatal(err)
	//}
}
