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

func RetriveUserByUserNameAndPassword(loginBody *models.LoginBody) (user *dao.User) {
	return dao.RetrieveUserByUserNameAndPassword(loginBody)
}

func AddUser(user *models.User) {
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

	err = dao.AddUser(user, tx)

	// 将用户加入对应的组

}

func UpdateUser(user *models.User) {
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

	err = dao.UpdateUserByAccountId(user, tx)

	// 先删除用户对应的组信息
	err = dao.DeleteUserRoleByAccountId(user.AccountId, tx)

	// 将用户加入对应的组

}
