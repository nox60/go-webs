package controller

import (
	"github.com/gin-gonic/gin"
	"testdb/models"
)

func ListFunctionsData(c *gin.Context) {

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	var dataLists models.PageListDataResult
	functions := make([]models.FunctionNode, 0)

	//build mock data for test
	function1 := models.FunctionNode{1, "test1", "/a/11", ""}
	function2 := models.FunctionNode{2, "test2", "/a/22", ""}
	function3 := models.FunctionNode{3, "test3", "/a/33", ""}
	function4 := models.FunctionNode{4, "test4", "/a/44", ""}
	function5 := models.FunctionNode{5, "test5", "/a/55", ""}

	functions = append(functions, function1)
	functions = append(functions, function2)
	functions = append(functions, function3)
	functions = append(functions, function4)
	functions = append(functions, function5)

	//results, totalCount, err := dao.RetrieveSampleData(&fetchDataRequestBody)

	resultMsg.Data = functions
	c.JSON(200, resultMsg)
}
