package services

import (
	"testdb/dao"
	"testdb/models"
)

func RetrieveSampleData(fetchDataBody *models.FetchDataRequestBody) (dataResBody []models.DataResponseBody, totalCounts int, err error) {
	return dao.RetrieveSampleData(fetchDataBody)
}
