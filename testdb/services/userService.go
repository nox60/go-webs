package services

import "testdb/dao"

func RetriveUserInfo(accountId int) {
	dao.StructQueryField(accountId)
}

func UpdateTest() {
	tx, err := dao.MysqlDb.Begin()
	if err != nil {
		return
	}
	defer func() {
		switch {
		case err != nil:
			tx.Rollback()
		default:
			err = tx.Commit()
		}
	}()

	dao.UpdateTxTesting(10, "a", "b", tx)
}
