package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"testdb/dao"
	"testdb/models"
	"testdb/services"
)

func ListSampleData(c *gin.Context) {

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	var fetchDataRequestBody models.FetchDataRequestBody
	var dataLists models.PageListDataResult

	if err := c.ShouldBindJSON(&fetchDataRequestBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fetchDataRequestBody.ItemId = -1

	results, totalCount, err := dao.RetrieveSampleData(&fetchDataRequestBody)

	if err != nil {
		fmt.Println(err)
	}

	dataLists.TotalCounts = totalCount
	dataLists.DataLists = results

	resultMsg.Data = dataLists
	c.JSON(200, resultMsg)
}

func AddItem(c *gin.Context) {

	var itemDataBody models.ItemDataBody

	if err := c.ShouldBindJSON(&itemDataBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(itemDataBody)

	services.AddItem(&itemDataBody)

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "新增数据成功"
	c.JSON(200, resultMsg)
}

func DeleteItem(c *gin.Context) {
	itemIdStr := c.Param("itemId")
	itemId, _ := strconv.Atoi(itemIdStr)
	services.DeleteItem(itemId)

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "删除数据成功"
	c.JSON(200, resultMsg)
}

func GetItem(c *gin.Context) {

	itemIdStr := c.Param("itemId")

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	var fetchDataRequestBody models.FetchDataRequestBody

	if err := c.ShouldBindJSON(&fetchDataRequestBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	itemId, err := strconv.Atoi(itemIdStr)

	fetchDataRequestBody.ItemId = itemId

	results, err := dao.GetData(&fetchDataRequestBody)

	if err != nil {
		fmt.Println(err)
	}

	resultMsg.Data = results

	c.JSON(200, resultMsg)
}
