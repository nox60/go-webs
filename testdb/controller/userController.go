package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testdb/dao"
	"testdb/models"
	"testdb/utils"
)

func JsonLogin(c *gin.Context) {

	var loginBody models.LoginBody
	fmt.Println(loginBody)

	if err := c.ShouldBindJSON(&loginBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := dao.RetrieveUserByUserNameAndPassword(&loginBody)

	resultMsg := new(models.HttpResult)

	//resultMap := utils.StructToMap(resultMsg)

	if result.Id > 0 {
		//登录成功
		resultMsg.Code = 20000
		resultMsg.Msg = "登录成功"

		//登录成功之后将用户能够使用的菜单权限信息，和其他信息一起编码放入token
		tokenPayload := new(models.TokenPayload)
		tokenPayload.AccountId = result.Id
		tokenPayload.MenuItems = "|001|002|003|004|"
		tokenJson, _ := json.Marshal(tokenPayload)
		jwtSignedToken := utils.JwtSign(string(tokenJson))
		resultMsg.Token = jwtSignedToken

		resultMsg.Data = ""
		//硬编码，先暂时未测试
		userInfo := new(models.UserInfo)
		userInfo.Introduction = "I am a super administrator"
		userInfo.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
		userInfo.Name = "Super Admin"
		userInfo.Code = 100

		//userInfoJson, _ := json.Marshal(userInfo)

		resultMsg.Data = userInfo

		//resultMsg.Data = string(userInfoJson)

		c.JSON(200, resultMsg)
	} else {
		//用户名或密码错误
		//resultMsg.ResultCode = 101
		resultMsg.Msg = "登录失败"
		c.JSON(200, gin.H{
			"code":    20000,
			"status":  "success",
			"message": "success",
			"data":    resultMsg,
		})
	}
}

func Info(c *gin.Context) {

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "登录成功"
	//登录成功

	userInfo := new(models.UserInfo)
	userInfo.Introduction = "I am a super administrator"
	userInfo.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	userInfo.Name = "Super Admin"
	userInfo.Code = 100

	resultMsg.Data = userInfo

	c.JSON(200, resultMsg)
}

func Login(c *gin.Context) {
	resultMsg := new(models.HttpResult)
	resultMsg.Token = "test"
	fmt.Println("login method")
	c.JSON(200, resultMsg)
}
