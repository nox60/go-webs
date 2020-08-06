package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testdb/models"
)

func ListSampleData(c *gin.Context) {

	var fetchDataRequestBody models.FetchDataRequestBody
	fmt.Println(fetchDataRequestBody)

	if err := c.ShouldBindJSON(&fetchDataRequestBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//resultMsg := new(models.HttpResult)
	//resultMsg.Code = 20000
	//resultMsg.Msg = "获取用户信息成功"
	////登录成功
	//
	//userInfo := new(models.UserInfo)
	//
	//userInfo.Introduction = "I am a super administrator"
	//userInfo.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	//userInfo.Name = "Super Admin"
	//userInfo.Code = 100
	//userInfo.Roles = "[admin]"
	//resultMsg.Data = userInfo

	c.JSON(200, "")
}
