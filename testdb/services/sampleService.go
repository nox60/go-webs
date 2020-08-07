package services

import (
	"testdb/dao"
	"testdb/models"
)

func RetrieveSampleData(fetchDataBody *models.FetchDataRequestBody) (dataResBody []models.DataResponseBody) {
	return dao.RetrieveSampleData(fetchDataBody)
}
