package services

import (
	"fmt"
	"testdb/dao"
	"testdb/models"
)

func GetFunctionsByParentId(fetchDataBody *models.FunctionNode) (dataResBody []models.FunctionNode, err error) {
	return dao.GetFunctionsByParentId(fetchDataBody)
}

func AddFunction(function *models.FunctionNode) {
	tx, err := dao.MysqlDb.Begin()

	if err != nil {
		return
	}
	defer func() {
		switch {
		case err != nil:
			fmt.Println(err)
			fmt.Println("rollback error")
		default:
			fmt.Println("commit ")
			err = tx.Commit()
		}
	}()

	err = dao.AddFunction(function, tx)
}

func GetFunctionById(fetchDataBody *models.FunctionNode) (dataResBody models.FunctionNode, err error) {
	dataRes, err := dao.GetFunctionById(fetchDataBody)
	tempParentId := dataRes.ParentFunctionId
	parents := make([]int, 0)

	//递归获取所有的父节点
	for {
		if tempParentId == 0 {
			parents = append(parents, tempParentId)
			break
		}
		parents = append(parents, tempParentId)

		//查询父节点
		tempFetchDataBody := new(models.FunctionNode)
		tempFetchDataBody.FunctionId = dataRes.ParentFunctionId
		tempData, _ := dao.GetFunctionById(tempFetchDataBody)
		tempParentId = tempData.ParentFunctionId
	}

	dataRes.Parents = parents
	return dataRes, err
}

func UpdateFunctionById(function *models.FunctionNode) {
	tx, err := dao.MysqlDb.Begin()

	if err != nil {
		return
	}
	defer func() {
		switch {
		case err != nil:
			fmt.Println(err)
			fmt.Println("rollback error")
		default:
			fmt.Println("commit ")
			err = tx.Commit()
		}
	}()

	err = dao.UpdateFunctionById(function, tx)
}
