package services

import (
	"fmt"
	"testdb/dao"
	"testdb/models"
)

func RetrieveSampleData(fetchDataBody *models.FetchDataRequestBody) (dataResBody []models.ItemDataBody, totalCounts int, err error) {
	return dao.RetrieveSampleData(fetchDataBody)
}

func AddItem(itemData *models.ItemDataBody) {
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

	err = dao.AddItem(itemData, tx)
}
