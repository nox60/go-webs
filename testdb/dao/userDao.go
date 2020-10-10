package dao

import (
	"database/sql"
	"fmt"
	"strings"
	"testdb/models"
)

// 用户表结构体
type User struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Age      int    `db:"age"`
	RealName string `db:"realName"`
}

// 查询数据，指定字段名
func StructQueryField(accountId int) {

	user := new(User)
	row := MysqlDb.QueryRow("select account_id, user_name, age from tb_users where accountId = ? ", accountId)

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

	_, err = tx.Exec("INSERT INTO `tb_users` (`account_id`,`user_name`,`real_name`) values (?,?,?) ", user.Id, user.Name, user.RealName)
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO `tb_users` (`account_id`,`user_name`,`real_name`) values (?,?,?) ", user.Id, user.Name, user.RealName)
	if err != nil {
		return err
	}

	return
}

// UpdateFooBar 更新
func InsertWithOutTxTest(user *User) (err error) {

	_, err = MysqlDb.Exec("INSERT INTO `tb_users` (`account_id`,`user_name`,`real_name`) values (?,?,?) ", user.Id, user.Name, user.RealName)
	if err != nil {
		return err
	}

	_, err = MysqlDb.Exec("INSERT INTO `tb_users` (`account_id`,`user_name`,`real_name`) values (?,?,?) ", user.Id, user.Name, user.RealName)
	if err != nil {
		return err
	}

	return
}

// 使用user, password进行查询
func RetrieveUserByUserNameAndPassword(userInfo *models.LoginBody) (user *User) {

	user1 := new(User)
	row := MysqlDb.QueryRow("select account_id, user_name, age from tb_users where user_name = ? AND password = ? ", userInfo.UserName, userInfo.Password)

	if err := row.Scan(&user1.Id, &user1.Name, &user1.Age); err != nil {
		//fmt.Printf("scan failed, err:%v", err)
		//fmt.Println("")
		//return user1
	}

	//fmt.Println(user1.Id, user1.Name, user1.Age)

	return user1
}

// 使用accountId获取用户信息
func RetrieveUserByAccountId(accountId int) (user *User) {

	user1 := new(User)

	row := MysqlDb.QueryRow("select account_id, user_name, age from tb_users where accountId = ? ", accountId)

	if err := row.Scan(&user1.Id, &user1.Name, &user1.Age); err != nil {
	}

	return user1
}

// 分页获取用户信息
func RetrieveUsersData(fetchDataBody *models.User) (dataResBody []models.User, totalCount int, err error) {

	// 通过切片存储
	results := make([]models.User, 0)

	// 获取数据的临时对象
	var dataObj models.User

	// 查询条件
	var queryStm strings.Builder

	// 总记录条数查询条件
	var countQueryStm strings.Builder

	// 查询条件
	var fetchArgs = make([]interface{}, 0)

	queryStm.WriteString(" SELECT `account_id`,`user_name`,`real_name` FROM tb_users WHERE 1=1 ")
	countQueryStm.WriteString(" SELECT COUNT(*) AS totalCount FROM tb_users WHERE 1=1 ")
	// 查询条件.
	if fetchDataBody.AccountId > -1 {
		queryStm.WriteString(" AND account_id = ? ")
		countQueryStm.WriteString(" AND account_id = ? ")
		fetchArgs = append(fetchArgs, fetchDataBody.AccountId)
	}

	queryStm.WriteString("LIMIT ?,? ")

	// 分页查询记录
	stmt, _ := MysqlDb.Prepare(queryStm.String())
	stmtCount, _ := MysqlDb.Prepare(countQueryStm.String())
	defer stmt.Close()
	defer stmtCount.Close()

	// 先查询总条数count(*)
	countResult := stmtCount.QueryRow(fetchArgs...)

	if err := countResult.Scan(&totalCount); err != nil {
		fmt.Printf("scan failed, err:%v", err)
	}

	// 查询分页数据
	fetchArgs = append(fetchArgs, fetchDataBody.GetStartByPageAndLimit())
	fetchArgs = append(fetchArgs, fetchDataBody.Limit)
	queryResults, err := stmt.Query(fetchArgs...)

	if err != nil {
		return results, 0, err
	}

	for queryResults.Next() {
		queryResults.Scan(&dataObj.AccountId,
			&dataObj.UserName,
			&dataObj.RealName)
		results = append(results, dataObj)
	}

	return results, totalCount, err
}

func AddUser(user *models.User, tx *sql.Tx) (err error) {
	_, err = tx.Exec(" INSERT INTO `tb_users` (`user_name`,`real_name`) "+
		" values (?,?) ",
		user.UserName,
		user.RealName)
	if err != nil {
		return err
	}
	return
}

func UpdateUserByAccountId(user *models.User, tx *sql.Tx) (err error) {
	var queryStm strings.Builder
	queryStm.WriteString("UPDATE `tb_users` SET `user_name` = ?, `real_name` = ?  WHERE account_id = ? ")

	var args = make([]interface{}, 0)

	args = append(args, user.UserName)
	args = append(args, user.RealName)
	args = append(args, user.AccountId)

	_, err = tx.Exec(queryStm.String(), args...)

	if err != nil {
		return err
	}

	return
}
