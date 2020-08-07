package dao

import (
	"fmt"
	"testdb/models"
)

// 使用accountId获取用户信息
func RetrieveSampleData(fetchDataBody *models.FetchDataRequestBody) (dataResBody []models.DataResponseBody, err error) {

	// 通过切片存储
	results := make([]models.DataResponseBody, 0)

	var user1 models.DataResponseBody

	queryResults, err := MysqlDb.Query("SELECT itemId, itemTilte, itemPrice FROM tb_items LIMIT ?,? ", fetchDataBody.Page, fetchDataBody.Limit)

	if err != nil {
		fmt.Println(err)
		return results, err
	}

	for queryResults.Next() {
		queryResults.Scan(&user1.ItemId, &user1.ItemTilte, &user1.ItemPrice)
		results = append(results, user1)
	}

	return results, err
}
