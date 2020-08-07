package dao

import (
	"fmt"
	"testdb/models"
)

// 使用accountId获取用户信息
func RetrieveSampleData(fetchDataBody *models.FetchDataRequestBody) (dataResBody []models.DataResponseBody, err error) {

	// 通过切片存储
	results := make([]models.DataResponseBody, 0)

	var dataObj models.DataResponseBody

	startValue := fetchDataBody.GetStartByPageAndLimit()

	fmt.Println("))))))))))))))))))))))))))    " + string(startValue))

	queryResults, err := MysqlDb.Query("SELECT itemId, itemTitle, itemPrice FROM tb_items LIMIT ?,? ", fetchDataBody.GetStartByPageAndLimit(), fetchDataBody.Limit)

	if err != nil {
		fmt.Println(err)
		return results, err
	}

	for queryResults.Next() {
		queryResults.Scan(&dataObj.ItemId, &dataObj.ItemTitle, &dataObj.ItemPrice)
		results = append(results, dataObj)
	}

	return results, err
}
