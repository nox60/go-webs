package dao

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"testdb/models"
)

// 使用accountId获取用户信息
func RetrievePermissionData(fetchDataBody *models.FetchDataRequestBody) (dataResBody []models.ItemDataBody, totalCount int, err error) {

	// 通过切片存储
	results := make([]models.ItemDataBody, 0)

	// 获取数据的临时对象
	var dataObj models.ItemDataBody

	// 查询条件
	var queryStm strings.Builder

	// 总记录条数查询条件
	var countQueryStm strings.Builder

	// 查询条件
	var fetchArgs = make([]interface{}, 0)

	queryStm.WriteString(" SELECT `itemId`, `createTime`,`itemContent`,`itemStar`,`itemType`,`itemTitle`,`itemStatus`,`itemDesc` FROM tb_items WHERE 1=1 ")
	countQueryStm.WriteString(" SELECT COUNT(*) AS totalCount FROM tb_items WHERE 1=1 ")
	// 查询条件.
	if fetchDataBody.ItemId > -1 {
		queryStm.WriteString(" AND itemId = ? ")
		countQueryStm.WriteString(" AND itemId = ? ")
		fetchArgs = append(fetchArgs, fetchDataBody.ItemId)
	}
	queryStm.WriteString("LIMIT ?,? ")

	// 分页查询记录
	stmt, _ := MysqlDb.Prepare(queryStm.String())
	stmtCount, _ := MysqlDb.Prepare(countQueryStm.String())
	defer stmt.Close()
	defer stmtCount.Close()

	// 先查询总条数count(*)
	countResult := stmtCount.QueryRow(fetchArgs...)

	if err := countResult.Scan(&totalCount); err != nil {
		fmt.Printf("scan failed, err:%v", err)
	}

	// 查询分页数据
	fetchArgs = append(fetchArgs, fetchDataBody.GetStartByPageAndLimit())
	fetchArgs = append(fetchArgs, fetchDataBody.Limit)
	queryResults, err := stmt.Query(fetchArgs...)

	if err != nil {
		fmt.Println(err)
		return results, 0, err
	}

	for queryResults.Next() {
		queryResults.Scan(&dataObj.ItemId,
			&dataObj.CreateTime,
			&dataObj.ItemContent,
			&dataObj.ItemStar,
			&dataObj.ItemType,
			&dataObj.ItemTitle,
			&dataObj.ItemStatus,
			&dataObj.ItemDesc)

		results = append(results, dataObj)
	}

	return results, totalCount, err
}

func GetFunctionsByParentId(fetchDataBody *models.FunctionNode) (dataResBody []models.FunctionNode, err error) {
	// 通过切片存储
	results := make([]models.FunctionNode, 0)

	// 获取数据的临时对象
	var dataObj models.FunctionNode

	// 查询条件
	var queryStm strings.Builder

	// 查询条件
	var fetchArgs = make([]interface{}, 0)

	queryStm.WriteString(" SELECT a.`function_id`,a.`number`,a.`order`,a.`name`,a.`path`,a.`parent_function_id`, ")
	queryStm.WriteString(" IF(b.function_id IS NULL,1,0) AS leaf, ")
	queryStm.WriteString(" IF(b.function_id IS NULL,0,1) AS hasChildren ")
	queryStm.WriteString(" FROM tb_functions AS a  ")
	queryStm.WriteString(" LEFT JOIN tb_functions AS b ON a.function_id = b.parent_function_id ")
	queryStm.WriteString(" WHERE 1=1 ")
	// 查询条件.
	queryStm.WriteString(" AND a.parent_function_id = ? ")
	fetchArgs = append(fetchArgs, fetchDataBody.ParentFunctionId)

	queryStm.WriteString(" GROUP BY a.`function_id` ")
	queryStm.WriteString(" ORDER BY a.`order` DESC ")

	// 查询记录
	stmt, _ := MysqlDb.Prepare(queryStm.String())
	defer stmt.Close()

	// 查询数据
	queryResults, err := stmt.Query(fetchArgs...)

	if err != nil {
		fmt.Println(err)
		return results, err
	}

	for queryResults.Next() {
		queryResults.Scan(&dataObj.FunctionId,
			&dataObj.Number,
			&dataObj.Order,
			&dataObj.Name,
			&dataObj.Path,
			&dataObj.ParentFunctionId,
			&dataObj.Leaf,
			&dataObj.HasChildren,
		)

		var items = make([]models.FunctionItem, 0)

		var tempItem models.FunctionItem
		tempItem.ItemName = "itemA" + strconv.Itoa(dataObj.FunctionId)

		var tempItem2 models.FunctionItem
		tempItem2.ItemName = "itemB" + strconv.Itoa(dataObj.FunctionId)

		items = append(items, tempItem)
		items = append(items, tempItem2)

		dataObj.Items = &items
		results = append(results, dataObj)
	}

	return results, err
}

func GetPermissionData(fetchDataBody *models.FetchDataRequestBody) (dataResBody models.ItemDataBody, err error) {

	// 获取数据的临时对象
	var dataObj models.ItemDataBody

	// 查询条件
	var queryStm strings.Builder

	// 查询条件
	var fetchArgs = make([]interface{}, 0)

	queryStm.WriteString(" SELECT `itemId`, `createTime`,`itemContent`,`itemStar`,`itemType`,`itemTitle`,`itemStatus`,`itemDesc` FROM tb_items WHERE 1=1 ")

	// 查询条件.
	if fetchDataBody.ItemId > -1 {
		queryStm.WriteString(" AND itemId = ? ")
		fetchArgs = append(fetchArgs, fetchDataBody.ItemId)
	}

	// 分页查询记录
	stmt, _ := MysqlDb.Prepare(queryStm.String())
	defer stmt.Close()

	// 查询
	queryResults := stmt.QueryRow(fetchArgs...)

	if err != nil {
		fmt.Println(err)
		return dataObj, err
	}

	queryResults.Scan(&dataObj.ItemId,
		&dataObj.CreateTime,
		&dataObj.ItemContent,
		&dataObj.ItemStar,
		&dataObj.ItemType,
		&dataObj.ItemTitle,
		&dataObj.ItemStatus,
		&dataObj.ItemDesc)

	return dataObj, err
}

func AddFunction(function *models.FunctionNode, tx *sql.Tx) (err error) {

	_, err = tx.Exec("INSERT INTO `tb_functions` (`number`,`order`,`name`,`path`,`parent_function_id`) "+
		"values (?,?,?,?,?) ",
		function.Number,
		function.Order,
		function.Name,
		function.Path,
		function.ParentFunctionId)
	if err != nil {
		return err
	}
	return
}

func AddFunctionItem(item *models.FunctionItem, tx *sql.Tx) (err error) {
	_, err = tx.Exec(" INSERT INTO `tb_functions_items` (`item_name`,`function_id`) "+
		" values (?,?) ",
		item.ItemName,
		item.FunctionId)
	if err != nil {
		return err
	}
	return
}

func AddRole(role *models.Role, tx *sql.Tx) (roleId int64, err error) {
	ret, err := tx.Exec("INSERT INTO `tb_roles` (`code`,`name`,`status`) "+
		"values (?,?,?) ",
		role.Code,
		role.Name,
		role.Status)
	if err != nil {
		return 0, err
	}

	if roleId, err = ret.LastInsertId(); nil == err {
		fmt.Println("LastInsertId:", roleId)
	}

	return roleId, err
}

func AddRoleFunction(role *models.Role, tx *sql.Tx) (err error) {
	for index, value := range role.Functions {
		fmt.Println("index:", index, "value:", value)
		_, err := tx.Exec("INSERT INTO `tb_roles_functions` (`role_id`,`function_id`) "+
			"values (?,?) ",
			role.RoleId,
			value)
		//新增
		if err != nil {
			return err
		}
	}

	return
}

func UpdateRoleById(roleData *models.Role, tx *sql.Tx) (err error) {

	// 查询语句
	var queryStm strings.Builder
	queryStm.WriteString("UPDATE `tb_roles` SET `name` = ?, `code` = ?  WHERE role_id = ? ")

	// 查询条件
	var args = make([]interface{}, 0)

	args = append(args, roleData.Name)
	args = append(args, roleData.Code)
	args = append(args, roleData.RoleId)

	_, err = tx.Exec(queryStm.String(), args...)

	if err != nil {
		return err
	}

	return
}

func DeleteFunction(functionId int, tx *sql.Tx) (err error) {
	_, err = tx.Exec("DELETE FROM `tb_functions` WHERE function_id = ? ",
		functionId)
	if err != nil {
		return err
	}
	return
}

func UpdateFunctionById(functionData *models.FunctionNode, tx *sql.Tx) (err error) {

	//_, err = tx.Exec("UPDATE `tb_functions` SET `number` = ?, `order` = ?, `name` = ?, `path` = ?, `parent_function_id` = ?  WHERE itemId = ? ", itemData.ItemContent, itemData.ItemTitle, itemData.ItemType, itemData.ItemId)

	// 查询语句
	var queryStm strings.Builder
	queryStm.WriteString("UPDATE `tb_functions` SET `number` = ?, `order` = ?, `name` = ?, `path` = ?, `parent_function_id` = ?  WHERE function_id = ? ")

	// 查询条件
	var args = make([]interface{}, 0)

	args = append(args, functionData.Number)
	args = append(args, functionData.Order)
	args = append(args, functionData.Name)
	args = append(args, functionData.Path)
	args = append(args, functionData.ParentFunctionId)
	args = append(args, functionData.FunctionId)

	_, err = tx.Exec(queryStm.String(), args...)

	if err != nil {
		return err
	}

	return
}

func GetFunctionById(fetchDataBody *models.FunctionNode) (dataResBody models.FunctionNode, err error) {

	// 获取数据的临时对象
	var dataObj models.FunctionNode

	// 查询条件
	var queryStm strings.Builder

	// 查询条件
	var fetchArgs = make([]interface{}, 0)

	queryStm.WriteString(" SELECT a.`function_id`,a.`number`,a.`order`,a.`name`,a.`path`,a.`parent_function_id` ")
	queryStm.WriteString(" FROM tb_functions AS a  ")
	queryStm.WriteString(" WHERE 1=1 ")
	// 查询条件.
	queryStm.WriteString(" AND a.`function_id` = ? ")
	fetchArgs = append(fetchArgs, fetchDataBody.FunctionId)

	// 查询记录
	stmt, _ := MysqlDb.Prepare(queryStm.String())
	defer stmt.Close()

	// 查询数据
	queryResult := stmt.QueryRow(fetchArgs...)

	if err != nil {
		fmt.Println(err)
		return dataObj, err
	}

	queryResult.Scan(&dataObj.FunctionId,
		&dataObj.Number,
		&dataObj.Order,
		&dataObj.Name,
		&dataObj.Path,
		&dataObj.ParentFunctionId,
	)

	return dataObj, err
}

// 获取角色信息
func RetrieveRoleData(fetchDataBody *models.Role) (dataResBody []models.Role, totalCount int, err error) {

	// 通过切片存储
	results := make([]models.Role, 0)

	// 获取数据的临时对象
	var dataObj models.Role

	// 查询条件
	var queryStm strings.Builder

	// 总记录条数查询条件
	var countQueryStm strings.Builder

	// 查询条件
	var fetchArgs = make([]interface{}, 0)

	queryStm.WriteString(" SELECT `role_id`,`name`,`code` FROM tb_roles WHERE 1=1 ")
	countQueryStm.WriteString(" SELECT COUNT(*) AS totalCount FROM tb_roles WHERE 1=1 ")
	// 查询条件.
	if fetchDataBody.RoleId > -1 {
		queryStm.WriteString(" AND role_id = ? ")
		countQueryStm.WriteString(" AND role_id = ? ")
		fetchArgs = append(fetchArgs, fetchDataBody.RoleId)
	}

	queryStm.WriteString("LIMIT ?,? ")

	// 分页查询记录
	stmt, _ := MysqlDb.Prepare(queryStm.String())
	stmtCount, _ := MysqlDb.Prepare(countQueryStm.String())
	defer stmt.Close()
	defer stmtCount.Close()

	// 先查询总条数count(*)
	countResult := stmtCount.QueryRow(fetchArgs...)

	if err := countResult.Scan(&totalCount); err != nil {
		fmt.Printf("scan failed, err:%v", err)
	}

	// 查询分页数据
	fetchArgs = append(fetchArgs, fetchDataBody.GetStartByPageAndLimit())
	fetchArgs = append(fetchArgs, fetchDataBody.Limit)
	queryResults, err := stmt.Query(fetchArgs...)

	if err != nil {
		return results, 0, err
	}

	for queryResults.Next() {
		queryResults.Scan(&dataObj.RoleId,
			&dataObj.Name,
			&dataObj.Code)
		results = append(results, dataObj)
	}

	return results, totalCount, err
}

func DeleteRole(roleId int64, tx *sql.Tx) (err error) {
	_, err = tx.Exec("DELETE FROM `tb_roles` WHERE role_id = ? ",
		roleId)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM `tb_roles_functions` WHERE role_id = ? ",
		roleId)
	if err != nil {
		return err
	}

	return
}

func DeleteRolesAndFunctionsByRoleId(roleId int64, tx *sql.Tx) (err error) {
	_, err = tx.Exec("DELETE FROM `tb_roles_functions` WHERE role_id = ? ",
		roleId)
	if err != nil {
		return err
	}
	return
}

func GetRoleById(fetchDataBody *models.Role) (dataResBody models.Role, err error) {

	// 获取数据的临时对象
	var dataObj models.Role

	// 查询条件
	var queryStm strings.Builder

	// 查询条件
	var fetchArgs = make([]interface{}, 0)

	queryStm.WriteString(" SELECT a.`role_id`,a.`name`,a.`code` ")
	queryStm.WriteString(" FROM tb_roles AS a  ")
	queryStm.WriteString(" WHERE 1=1 ")
	// 查询条件.
	queryStm.WriteString(" AND a.`role_id` = ? ")
	fetchArgs = append(fetchArgs, fetchDataBody.RoleId)

	// 查询记录
	stmt, _ := MysqlDb.Prepare(queryStm.String())
	defer stmt.Close()

	// 查询数据
	queryResult := stmt.QueryRow(fetchArgs...)

	if err != nil {
		fmt.Println(err)
		return dataObj, err
	}

	queryResult.Scan(&dataObj.RoleId,
		&dataObj.Name,
		&dataObj.Code,
	)

	//获取该角色所有的功能点
	// 获取数据的临时对象
	var tempFunction int64

	// 查询条件
	var queryStm2 strings.Builder

	// 通过切片存储
	results2 := make([]int64, 0)

	// 查询条件
	var fetchArgs2 = make([]interface{}, 0)

	queryStm2.WriteString(" SELECT a.`function_id` ")
	queryStm2.WriteString(" FROM tb_roles_functions AS a  ")
	queryStm2.WriteString(" WHERE 1=1 ")
	// 查询条件.
	queryStm2.WriteString(" AND a.`role_id` = ? ")
	fetchArgs2 = append(fetchArgs2, fetchDataBody.RoleId)

	// 查询记录
	stmt2, _ := MysqlDb.Prepare(queryStm2.String())
	defer stmt2.Close()

	// 查询数据
	queryResult2, err := stmt2.Query(fetchArgs2...)

	if err != nil {
		fmt.Println(err)
		return dataObj, err
	}

	for queryResult2.Next() {
		queryResult2.Scan(&tempFunction)

		results2 = append(results2, tempFunction)
	}

	dataObj.Functions = results2

	return dataObj, err
}
