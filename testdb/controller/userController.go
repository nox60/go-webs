package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"testdb/dao"
	"testdb/models"
)

func JsonLogin(c *gin.Context) {
	var json models.LoginBody

	fmt.Println(json)

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := dao.RetrieveUserByUserNameAndPassword(&json)

	resultMsg := new(models.SimpleCode)

	if result.Id > 0 {
		//登录成功
		resultMsg.ResultCode = 100
		resultMsg.Msg = "登录成功"
		resultMsg.Token = "4"
		//硬编码，先暂时未测试
		resultMsg.Roles = "['admin']"
		resultMsg.Introduction = "I am a super administrator"
		resultMsg.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
		resultMsg.Name = "Super Admin"
		resultMsg.AccountId = result.Id
		c.JSON(200, gin.H{
			"code":    20000,
			"status":  "success",
			"message": "success",
			"data":    resultMsg,
		})
	} else {
		//用户名或密码错误
		resultMsg.ResultCode = 101
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
	accountId, _ := strconv.Atoi(c.Query("token"))

	result := dao.RetrieveUserByAccountId(accountId)

	resultMsg := new(models.SimpleCode)

	if result.Id > 0 {
		//登录成功
		resultMsg.ResultCode = 100
		resultMsg.Msg = "登录成功"
		resultMsg.Token = "temp_token"
		//硬编码，先暂时未测试
		resultMsg.Roles = "['admin']"
		resultMsg.Introduction = "I am a super administrator"
		resultMsg.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
		resultMsg.Name = "Super Admin"
		resultMsg.AccountId = result.Id
		c.JSON(200, gin.H{
			"code":    20000,
			"status":  "success",
			"message": "success",
			"data":    resultMsg,
		})
	} else {
		//用户名或密码错误
		resultMsg.ResultCode = 101
		resultMsg.Msg = "登录失败"
		c.JSON(200, gin.H{
			"code":    20000,
			"status":  "success",
			"message": "success",
			"data":    resultMsg,
		})
	}
}

func Login(c *gin.Context) {

	c.JSON(200, gin.H{
		"code":    20000,
		"status":  "posted",
		"message": "test message",
		"data":    "hello world",
	})
}
