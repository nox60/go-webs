package dao

import (
	"testdb/models"
)

// 使用accountId获取用户信息
func RetrieveSampleData(fetchDataBody *models.FetchDataRequestBody) (dataResBody []models.DataResponseBody) {

	// 通过切片存储
	results := make([]models.DataResponseBody, 0)

	var user1 models.DataResponseBody

	queryResults, _ := MysqlDb.Query("SELECT itemId, itemTilte, itemPrice FROM tb_items LIMIT ?,? ", fetchDataBody.Page, fetchDataBody.Limit)

	for queryResults.Next() {
		queryResults.Scan(&user1.ItemId, &user1.ItemTilte, &user1.ItemPrice)
		results = append(results, user1)
	}

	return results
}
