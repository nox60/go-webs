package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"testdb/models"
)

func ListFunctionsData(c *gin.Context) {

	resultMsg := new(models.HttpResult)
	resultMsg.Code = 20000
	resultMsg.Msg = "获取数据成功"

	parentIdStr := c.Param("parentId")
	parentId, _ := strconv.Atoi(parentIdStr)

	functions := make([]models.FunctionNode, 0)

	if parentId == 0 {
		//build mock data for test
		function1 := models.FunctionNode{1, "test1", "/a/11", 0, false}
		function2 := models.FunctionNode{2, "test2", "/a/22", 0, false}
		function3 := models.FunctionNode{3, "test3", "/a/33", 0, false}
		function4 := models.FunctionNode{4, "test4", "/a/44", 0, true}
		function5 := models.FunctionNode{5, "test5", "/a/55", 0, false}

		functions = append(functions, function1)
		functions = append(functions, function2)
		functions = append(functions, function3)
		functions = append(functions, function4)
		functions = append(functions, function5)
	}

	resultMsg.Data = functions
	c.JSON(200, resultMsg)

	//
	////var dataLists models.PageListDataResult
	//functions := make([]models.FunctionNode, 0)
	//
	//functions1 := make([]models.FunctionNode, 0)
	//
	//functions2 := make([]models.FunctionNode, 0)
	//
	////build mock data for test
	//function1 := models.FunctionNode{1, "test1", "/a/11", ""}
	//function2 := models.FunctionNode{2, "test2", "/a/22", ""}
	//function3 := models.FunctionNode{3, "test3", "/a/33", ""}
	//function4 := models.FunctionNode{4, "test4", "/a/44", ""}
	//function5 := models.FunctionNode{5, "test5", "/a/55", ""}
	//
	//function6 := models.FunctionNode{6, "test6", "/a/66", ""}
	//function7 := models.FunctionNode{7, "test7", "/a/77", ""}
	//
	//function8 := models.FunctionNode{8, "test8", "/a/88", ""}
	//
	//functions2 = append(functions2, function8)
	//
	//function7.FunctionChild = functions2
	//
	//functions1 = append(functions1, function6)
	//functions1 = append(functions1, function7)
	//
	//function5.FunctionChild = functions1
	//
	//functions = append(functions, function1)
	//functions = append(functions, function2)
	//functions = append(functions, function3)
	//functions = append(functions, function4)
	//functions = append(functions, function5)
	//
	////results, totalCount, err := dao.RetrieveSampleData(&fetchDataRequestBody)
	//
	//resultMsg.Data = functions
	//c.JSON(200, resultMsg)
}
