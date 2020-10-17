package dao

import (
	"database/sql"
	"fmt"
	"strconv"
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

	// 查询条件
	var queryStm strings.Builder

	// 查询条件
	var fetchArgs = make([]interface{}, 0)

	fetchArgs = append(fetchArgs, userInfo.UserName)
	fetchArgs = append(fetchArgs, userInfo.Password)

	queryStm.WriteString(" select account_id, user_name, age from tb_users where user_name = ? AND password = ? ")

	user1 := new(User)

	//row := MysqlDb.QueryRow("select account_id, user_name, age from tb_users where user_name = ? AND password = ? ", userInfo.UserName, userInfo.Password)

	// 分页查询记录
	stmt, err := MysqlDb.Prepare(queryStm.String())
	if err != nil {
		fmt.Printf("SQL Prepare error, err:%v", err)
	}
	defer stmt.Close()

	// 查询
	queryResults, err := stmt.Query(fetchArgs...)

	if err != nil {
		fmt.Println(err)
		return user
	}

	for queryResults.Next() {
		queryResults.Scan(
			&user1.Id,
			&user1.Name,
			&user1.Age)
	}

	// if err := row.Scan(&user1.Id, &user1.Name, &user1.Age); err != nil {
	//fmt.Printf("scan failed, err:%v", err)
	//fmt.Println("")
	//return user1

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

	queryStm.WriteString(" SELECT a.`account_id`, a.`user_name`, a.`real_name`, ")
	queryStm.WriteString(" GROUP_CONCAT(DISTINCT(CONCAT_WS('|!|', c.role_id, c.name))) AS roleStr ")
	queryStm.WriteString(" FROM tb_users a ")
	queryStm.WriteString(" LEFT JOIN tb_users_roles b ON a.account_id = b.account_id ")
	queryStm.WriteString(" LEFT JOIN tb_roles c ON b.role_id = c.role_id ")
	queryStm.WriteString(" WHERE 1=1 ")

	countQueryStm.WriteString(" SELECT COUNT(*) AS totalCount ")
	countQueryStm.WriteString(" FROM tb_users a ")
	countQueryStm.WriteString(" LEFT JOIN tb_users_roles b ON a.account_id = b.account_id ")
	countQueryStm.WriteString(" LEFT JOIN tb_roles c ON b.role_id = c.role_id ")
	countQueryStm.WriteString(" WHERE 1=1 ")

	// 查询条件.
	if fetchDataBody.AccountId > -1 {
		queryStm.WriteString(" AND a.account_id = ? ")
		countQueryStm.WriteString(" AND a.account_id = ? ")
		fetchArgs = append(fetchArgs, fetchDataBody.AccountId)
	}

	queryStm.WriteString(" GROUP BY a.`account_id` ")
	queryStm.WriteString(" ORDER BY a.`account_id` DESC ")
	queryStm.WriteString(" LIMIT ?,? ")

	countQueryStm.WriteString(" GROUP BY a.`account_id` ")
	countQueryStm.WriteString(" ORDER BY a.`account_id` DESC ")

	// 分页查询记录
	stmt, err := MysqlDb.Prepare(queryStm.String())
	if err != nil {
		fmt.Printf("SQL Prepare error, err:%v", err)
	}

	stmtCount, err := MysqlDb.Prepare(countQueryStm.String())
	if err != nil {
		fmt.Printf("COUNT SQL Prepare error, err:%v", err)
	}

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
		fmt.Println(err)
		return results, 0, err
	}

	for queryResults.Next() {

		dataObj.RoleStr = ""
		dataObj.RoleIds = []int{}

		queryResults.Scan(&dataObj.AccountId,
			&dataObj.UserName,
			&dataObj.RealName,
			&dataObj.RoleStr)

		if strings.Index(dataObj.RoleStr, "|!|") > 0 {
			var roles = make([]string, 0)
			roles = strings.Split(dataObj.RoleStr, ",")

			if len(roles) > 0 {
				for _, roleTemp := range roles {
					roleTempArray := strings.Split(roleTemp, "|!|")
					var roleTemp models.Role
					roleIdInt, _ := strconv.Atoi(roleTempArray[0])
					roleTemp.RoleId = int64(roleIdInt)
					roleTemp.Name = roleTempArray[1]
					dataObj.Roles = append(dataObj.Roles, roleTemp)
					dataObj.RoleIds = append(dataObj.RoleIds, roleIdInt)
				}
			}
		}

		results = append(results, dataObj)
	}

	// 处理权限点信息

	return results, totalCount, err
}

func AddUser(user *models.User, tx *sql.Tx) (accountId int, err error) {
	ret, err := tx.Exec(" INSERT INTO `tb_users` (`user_name`,`real_name`) "+
		" values (?,?) ",
		user.UserName,
		user.RealName)
	var accountId64 int64
	if accountId64, err = ret.LastInsertId(); nil == err {
		fmt.Println("LastInsertId:", accountId)
	}
	accountId = int(accountId64)
	return accountId, err
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

func AddUserRole(accountId int, roleId int, tx *sql.Tx) (err error) {
	_, err = tx.Exec(" INSERT INTO `tb_users_roles` (`account_id`,`role_id`) "+
		" values (?,?) ",
		accountId,
		roleId)
	if err != nil {
		return err
	}
	return
}

func DeleteUserByAccountId(accountId int, tx *sql.Tx) (err error) {
	_, err = tx.Exec("DELETE FROM `tb_users` WHERE account_id = ? ",
		accountId)
	if err != nil {
		return err
	}
	return
}

func DeleteUserRoleByAccountId(accountId int, tx *sql.Tx) (err error) {
	_, err = tx.Exec("DELETE FROM `tb_users_roles` WHERE account_id = ? ",
		accountId)
	if err != nil {
		return err
	}
	return
}
