package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Nickname string `sql:"nickname"`
	UserId   int64  `sql:"user_id"`
}

func main() {

	conn, err := sql.Open("mysql", "michong:michong@tcp(192.168.1.34:3306)/ck_new?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	sql := "select user_id, nickname from users limit 1"

	var user User
	err = conn.QueryRow(sql).Scan(&user.UserId, &user.Nickname)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
