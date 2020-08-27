package dao

import (
	"database/sql"
	"fmt"
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

	/*
		SELECT a.function_id, a.name, b.function_id, IF(b.function_id IS NULL,0,1) AS leaf FROM tb_functions a
		LEFT JOIN tb_functions b ON a.function_id = b.parent_function_id
		WHERE
		a.parent_function_id = 0;
	*/

	queryStm.WriteString(" SELECT a.`function_id`,a.`number`,a.`order`,a.`name`,a.`path`,a.`parent_function_id`, ")
	queryStm.WriteString(" IF(b.function_id IS NULL,1,0) AS leaf, ")
	queryStm.WriteString(" IF(b.function_id IS NULL,0,1) AS hasChildren ")
	queryStm.WriteString(" FROM tb_functions AS a  ")
	queryStm.WriteString(" LEFT JOIN tb_functions AS b ON a.function_id = b.parent_function_id ")
	queryStm.WriteString(" WHERE 1=1 ")
	// 查询条件.
	queryStm.WriteString(" AND a.parent_function_id = ? ")
	fetchArgs = append(fetchArgs, fetchDataBody.ParentFunctionId)

	queryStm.WriteString(" GROUP BY a.`function_id`")

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

func DeleteFunction(itemId int, tx *sql.Tx) (err error) {
	_, err = tx.Exec("DELETE FROM `tb_items` WHERE itemId = ? ",
		itemId)
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

	fmt.Println("---------------------------------------+++")
	fmt.Println(queryStm.String())
	fmt.Println(fetchDataBody.FunctionId)
	fmt.Println(queryResult)

	queryResult.Scan(&dataObj.FunctionId,
		&dataObj.Number,
		&dataObj.Order,
		&dataObj.Name,
		&dataObj.Path,
		&dataObj.ParentFunctionId,
	)

	fmt.Println(dataObj)
	fmt.Println("---------------------------------------!!!")

	return dataObj, err
}
