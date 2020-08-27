package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"testdb/models"
	"testdb/services"
)

func ListFunctionsData(c *gin.Context) {

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	parentIdStr := c.Param("parentId")
	parentId, _ := strconv.Atoi(parentIdStr)

	functions := make([]models.FunctionNode, 0)

	fetchBody := new(models.FunctionNode)

	fetchBody.ParentFunctionId = parentId

	functions, _ = services.GetFunctionsByParentId(fetchBody)

	resultMsg.Data = functions
	c.JSON(200, resultMsg)
}

func ListNodesData(c *gin.Context) {

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	functions := make([]models.FunctionNode, 0)

	node1 := models.FunctionNode{1, 1, 1, "test1", "adasdf", 0, false, true}
	node2 := models.FunctionNode{2, 3, 2, "test2", "asdfasf", 0, true, false}
	node3 := models.FunctionNode{3, 4, 3, "test3", "asdf", 0, false, true}

	functions = append(functions, node1)
	functions = append(functions, node2)
	functions = append(functions, node3)

	resultMsg.Data = functions
	c.JSON(200, resultMsg)

}

func AddOrUpdateFunction(c *gin.Context) {

	var funcionReq models.FunctionNode

	if err := c.ShouldBindJSON(&funcionReq); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if funcionReq.FunctionId == 0 {
		// 新增
		services.AddFunction(&funcionReq)
	} else {
		// 更新
		// services.UpdateItemById(&funcionReq)
		services.UpdateFunctionById(&funcionReq)
	}

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "新增功能点成功"
	c.JSON(200, resultMsg)
}

func GetFunctionById(c *gin.Context) {

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	idStr := c.Param("id")

	id, _ := strconv.Atoi(idStr)

	fetchBody := new(models.FunctionNode)

	fetchBody.FunctionId = id

	resultBody, _ := services.GetFunctionById(fetchBody)

	resultMsg.Data = resultBody

	c.JSON(200, resultMsg)
}
