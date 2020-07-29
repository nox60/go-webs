package services

import (
	"fmt"
	"testdb/dao"
	"testdb/models"
)

func RetriveUserInfo(accountId int) {
	dao.StructQueryField(accountId)
}

func InsertTest() {
	tx, err := dao.MysqlDb.Begin()

	if err != nil {
		return
	}
	defer func() {
		switch {
		case err != nil:
			fmt.Println("rollback error")
			// tx.Rollback()
		default:
			fmt.Println("commit ")
			err = tx.Commit()
		}
	}()

	user := dao.User{}

	user.Id = 888
	user.Name = "testUserName"
	user.RealName = "testUssssss"

	err = dao.InsertTxTest(&user, tx)

}

func InsertTestWithOutTx() {

	user := dao.User{}

	user.Id = 888
	user.Name = "testUserName"
	user.RealName = "testUssssss"

	err := dao.InsertWithOutTxTest(&user)

	fmt.Println(err)
}

func RetriveUserByUserNameAndPassword(loginBody *models.LoginBody) (user dao.User) {
	return dao.RetrieveUserByUserNameAndPassword(loginBody)
}
