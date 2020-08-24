package controller

import (
	"github.com/gin-gonic/gin"
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
