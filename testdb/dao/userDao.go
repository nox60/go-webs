package dao

import (
	"fmt"
)

// 用户表结构体
type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// 查询数据，指定字段名
func StructQueryField() {

	user := new(User)
	row := MysqlDb.QueryRow("select accountId, userName, age from tb_users where accountId = ? ", 3)

	if err := row.Scan(&user.Id, &user.Name, &user.Age); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return
	}

	fmt.Println(user.Id, user.Name, user.Age)
}

// UpdateFooBar 更新
func UpdateFooBar(id int64, x, y string) (err error) {
	tx, err := MysqlDb.Begin()
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

	_, err = tx.Exec("update `foo` set `x` = ? where `id` = ?", x, id)
	if err != nil {
		return
	}
	_, err = tx.Exec("update `bar` set `y` = ? where `id` = ?", y, id)
	if err != nil {
		return
	}

	return
}
