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

func AddFunctionItem(functionItem *models.FunctionItem) {
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

	err = dao.AddFunctionItem(functionItem, tx)
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
		tempFetchDataBody.FunctionId = tempParentId
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

func DeleteFunction(functionId int) {
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

	err = dao.DeleteFunction(functionId, tx)
}

func AddRole(role *models.Role) {
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

	var roleId int64
	roleId, err = dao.AddRole(role, tx)
	role.RoleId = roleId
	err = dao.AddRoleFunction(role, tx)
}

func UpdateRole(role *models.Role) {
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

	err = dao.UpdateRoleById(role, tx)

	//先删除该角色对应的所有权限点
	err = dao.DeleteRolesAndFunctionsByRoleId(role.RoleId, tx)

	//重新添加角色对应的功能点
	err = dao.AddRoleFunction(role, tx)
}

func RetrieveRoleData(fetchDataBody *models.Role) (dataResBody []models.Role, totalCounts int, err error) {
	return dao.RetrieveRoleData(fetchDataBody)
}

func DeleteRole(roleId int64) {
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

	//删除角色本身
	err = dao.DeleteRole(roleId, tx)

	//删除角色对应的功能点
	err = dao.DeleteRolesAndFunctionsByRoleId(roleId, tx)
}

func GetRoleById(fetchDataBody *models.Role) (dataResBody models.Role, err error) {
	dataRes, err := dao.GetRoleById(fetchDataBody)
	return dataRes, err
}

func GetAllFunctions(node *models.FunctionNode) (err error, childIds []int) {
	var selfAndChildIds []int
	// var parentIds []int

	node.ParentIds = append(node.ParentIds, node.ParentFunctionId)
	if node.HasChildren {

		var parent models.FunctionNode
		parent.ParentFunctionId = node.FunctionId
		child, err := dao.GetFunctionsByParentId(&parent)

		if err != nil {
			fmt.Println(err)
		}

		node.Child = &child
		for i, _ := range child {
			// 将孩子节点和父亲节点进行关联
			//	child[i].ParentNode = node
			child[i].ParentIds = node.ParentIds
			_, childs := GetAllFunctions(&child[i])

			if len(childs) > 0 {
				for _, t := range childs {
					selfAndChildIds = append(selfAndChildIds, t)
				}
			}
		}
		node.ChildIds = selfAndChildIds

	}
	selfAndChildIds = append(selfAndChildIds, node.FunctionId)

	//处理父节点
	//tempNode := node
	//// var tempParentIds []int
	//
	//for {
	//	if tempNode.ParentFunctionId != 0 {
	//		node.ParentIds = append(node.ParentIds, tempNode.FunctionId)
	//		tempNode = tempNode.ParentNode
	//	} else {
	//		break
	//	}
	//}

	return err, selfAndChildIds
}

func DeleteFunctionItem(functionItemId int) {
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

	err = dao.DeleteFunctionItem(functionItemId, tx)
}

func GetFunctionItemById(fetchDataBody *models.FunctionItem) (dataResBody models.FunctionItem, err error) {
	dataRes, err := dao.GetFunctionItemById(fetchDataBody)
	return dataRes, err
}

func UpdateFunctionItemById(item *models.FunctionItem) {
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

	err = dao.UpdateFunctionItemById(item, tx)
}
