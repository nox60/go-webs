package dao

import (
	"database/sql"
	"fmt"
	"testdb/models"
)

// 使用accountId获取用户信息
func RetrieveSampleData(fetchDataBody *models.FetchDataRequestBody) (dataResBody []models.ItemDataBody, totalCount int, err error) {

	// 通过切片存储
	results := make([]models.ItemDataBody, 0)

	var dataObj models.ItemDataBody

	//分页查询记录
	//queryResults, err := MysqlDb.Query("SELECT `itemId`, `createTime`,`itemContent`,`itemStar`,`itemType`,`itemTitle`,`itemStatus`,`itemDesc` FROM tb_items LIMIT ?,? ", params)
	stmt, _ := MysqlDb.Prepare("SELECT `itemId`, `createTime`,`itemContent`,`itemStar`,`itemType`,`itemTitle`,`itemStatus`,`itemDesc` FROM tb_items LIMIT ?,?")
	defer stmt.Close()

	queryResults, err := stmt.Query(fetchDataBody.GetStartByPageAndLimit(), fetchDataBody.Limit)

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

	//查询总条数
	countResult := MysqlDb.QueryRow("SELECT count(*) AS totalCount FROM tb_items ")

	totalCount = 0

	if err := countResult.Scan(&totalCount); err != nil {
		fmt.Printf("scan failed, err:%v", err)
	}

	fmt.Println(totalCount)

	return results, totalCount, err
}

func AddItem(itemData *models.ItemDataBody, tx *sql.Tx) (err error) {

	_, err = tx.Exec("INSERT INTO `tb_items` (`createTime`,`itemContent`,`itemStar`,`itemType`,`itemTitle`,`itemStatus`,`itemDesc`) "+
		"values (now(),?,?,?,?,?,?) ",
		itemData.ItemContent,
		itemData.ItemStar,
		itemData.ItemType,
		itemData.ItemTitle,
		itemData.ItemStatus,
		itemData.ItemDesc)
	if err != nil {
		return err
	}
	return
}

func DeleteItem(itemId int, tx *sql.Tx) (err error) {
	_, err = tx.Exec("DELETE FROM `tb_items` WHERE itemId = ? ",
		itemId)
	if err != nil {
		return err
	}
	return
}
