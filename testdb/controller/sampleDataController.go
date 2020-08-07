package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testdb/dao"
	"testdb/models"
)

func ListSampleData(c *gin.Context) {

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	var fetchDataRequestBody models.FetchDataRequestBody
	fmt.Println(fetchDataRequestBody)

	if err := c.ShouldBindJSON(&fetchDataRequestBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results, err := dao.RetrieveSampleData(&fetchDataRequestBody)

	if err != nil {
		fmt.Println(err)
	}

	resultMsg.Data = results
	c.JSON(200, resultMsg)
}
