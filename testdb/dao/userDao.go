package dao

import (
	"database/sql"
	"fmt"
	"testdb/models"
)

// 用户表结构体
type User struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	Age      int    `db:"age"`
	RealName string `db:"realName"`
}

// 查询数据，指定字段名
func StructQueryField(accountId int) {

	user := new(User)
	row := MysqlDb.QueryRow("select accountId, userName, age from tb_users where accountId = ? ", accountId)

	if err := row.Scan(&user.Id, &user.Name, &user.Age); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return
	}

	fmt.Println(user.Id, user.Name, user.Age)
}

// UpdateFooBar 更新
func InsertTxTest(user *User, tx *sql.Tx) (err error) {

	//_, err = tx.Exec("update `foo` set `x` = ? where `id` = ?", x, id)
	//if err != nil {
	//	return
	//}
	//_, err = tx.Exec("update `bar` set `y` = ? where `id` = ?", y, id)
	//if err != nil {
	//	return
	//}

	//res, err := pool.Exec("insert into `users` (`name`) values (?)", name)

	_, err = tx.Exec("INSERT INTO `tb_users` (`accountId`,`userName`,`realName`) values (?,?,?) ", user.Id, user.Name, user.RealName)
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO `tb_users` (`accountId`,`userName`,`realName`) values (?,?,?) ", user.Id, user.Name, user.RealName)
	if err != nil {
		return err
	}

	return
}

// UpdateFooBar 更新
func InsertWithOutTxTest(user *User) (err error) {

	_, err = MysqlDb.Exec("INSERT INTO `tb_users` (`accountId`,`userName`,`realName`) values (?,?,?) ", user.Id, user.Name, user.RealName)
	if err != nil {
		return err
	}

	_, err = MysqlDb.Exec("INSERT INTO `tb_users` (`accountId`,`userName`,`realName`) values (?,?,?) ", user.Id, user.Name, user.RealName)
	if err != nil {
		return err
	}

	return
}

// 使用user, password进行查询
func RetrieveUserByUserNameAndPassword(userInfo *models.LoginBody) (user *User) {

	user1 := new(User)
	row := MysqlDb.QueryRow("select accountId, userName, age from tb_users where userName = ? AND password = ? ", userInfo.UserName, userInfo.Password)

	if err := row.Scan(&user1.Id, &user1.Name, &user1.Age); err != nil {
		//fmt.Printf("scan failed, err:%v", err)
		//fmt.Println("")
		//return user1
	}

	//fmt.Println(user1.Id, user1.Name, user1.Age)

	return user1
}
