package dao

import (
	"testdb/models"
)

// 使用accountId获取用户信息
func RetrieveSampleData(accountId int) (dataResBody []models.DataResponseBody) {

	// 通过切片存储
	results := make([]models.DataResponseBody, 0)

	var user1 models.DataResponseBody

	queryResults, _ := MysqlDb.Query("select accountId, userName, age from tb_users where accountId = ? ", accountId)

	for queryResults.Next() {
		queryResults.Scan(&user1.Id, &user1.Author, &user1.Importance)
		results = append(results, user1)
	}

	return results
}
