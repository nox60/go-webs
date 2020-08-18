package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testdb/models"
)

func ListFunctionsData(c *gin.Context) {

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	var fetchDataRequestBody models.FetchDataRequestBody
	var dataLists models.PageListDataResult
	functions := make([]models.PermissionNode, 0)

	if err := c.ShouldBindJSON(&fetchDataRequestBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fetchDataRequestBody.ItemId = -1

	//results, totalCount, err := dao.RetrieveSampleData(&fetchDataRequestBody)

	resultMsg.Data = dataLists
	c.JSON(200, resultMsg)
}
