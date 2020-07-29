package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testdb/dao"
	"testdb/models"
)

func JsonLogin(c *gin.Context) {
	var json models.LoginBody

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := dao.RetrieveUserByUserNameAndPassword(&json)

	if result.Id > 0 {
		//登录成功
		c.JSON(200, gin.H{
			"code":    20000,
			"status":  "login Successed",
			"message": "login Successed",
			"data":    "",
		})
	} else {
		//用户名或密码错误
		c.JSON(200, gin.H{
			"code":    20001,
			"status":  "login failed",
			"message": "login failed",
			"data":    "",
		})
	}
}

func SimpleLogin(c *gin.Context) {

	c.JSON(200, gin.H{
		"code":    20000,
		"status":  "posted",
		"message": "test message",
		"data":    "hello world",
	})
}
