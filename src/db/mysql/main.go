package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 定义一个全局 DB
var db *sql.DB

func insertData(username string, password string) (int64, error) {
	// 执行插入
	sqlStr := "insert into user_tb1(username, password) values (?, ?)"
	ret, err := db.Exec(sqlStr, username, password)
	if err != nil {
		fmt.Printf("insert faild, err:%v\n", err)
		return 0, err
	}

	// 获取数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get ID faild, err:%v\n", err)
		return 0, err
	}

	fmt.Printf("insert success, id:%d\n", id)
	return id, nil
}

type user struct {
	id       int64
	username string
	password string
}

// 单行查询
func queryRow(id int64) {
	sqlStr := "select id, username, password from user_tb1 where id = ?"
	var u user
	// 确保QueryRow之后调用Scan方法，否则持有的数据库连接不会被释放
	err := db.QueryRow(sqlStr, id).Scan(&u.id, &u.username, &u.password)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("user:%v\n", u)
}

// 单行查询
func query(id int) {
	sqlStr := "select id, username, password from user_tb1 where id > ?"
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	// 注意，关闭rows释放持有的数据库连接
	defer rows.Close()

	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("user:%v\n", u)
	}
}

func update(id int, username string) {

	sqlStr := "update user_tb1 set username = ? where id = ?"
	ret, err := db.Exec(sqlStr, username, id)
	if err != nil {
		fmt.Printf("更新失败, err:%v\n", err)
		return
	}

	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("更新行失败, err:%v\n", err)
		return
	}
	fmt.Printf("更新成功, 更新的行数 %d\n", rows)
}

func delete(id int64) {
	sqlStr := "delete from user_tb1 where id = ?"
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("删除失败, err:%v\n", err)
		return
	}

	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("删除行失败, err:%v\n", err)
		return
	}
	fmt.Printf("删除成功, 删除的行数 %d\n", rows)
}

func initDB() (err error) {

	// 这里不会校验账号密码是否正确
	// 用户名:密码@数据库名称
	db, err = sql.Open("mysql", "root:@/go_db?charset=utf8mb4")
	if err != nil {
		return err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {

	err := initDB()
	if err != nil {
		fmt.Printf("初始化失败 err:%v\n", err)
		return
	}
	fmt.Printf("初始化成功\n")

	//insertData("张三", "zs123")

	//queryRow(1)

	//query(1)

	//queryRow(1)
	//update(1, "chenken")
	//queryRow(1)

	id, _ := insertData("aken", "aken123")
	queryRow(id)
	delete(id)
	queryRow(id)
}
