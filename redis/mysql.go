package main

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "michong:michong@tcp(192.168.1.34:3306)/haochang_api?charset=utf8mb4")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer db.Close()

	queryData(db)

}

func queryData(db *sql.DB) {
	rows, _ := db.Query("select * from user_back limit 2")

	columns, _ := rows.Columns()              //读出查询出的列字段名
	fmt.Println(columns)
	fmt.Println(len(columns))
	scans := make([]interface{}, len(columns))
	values := make([][]byte, len(columns))

	for i := range values {
		scans[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scans...)
		if err != nil {
			fmt.Println(err)
			continue
		}
		row := make(map[string]string)
		for i, v := range values {
			row[columns[i]] = string(v)
		}
		fmt.Println(row)
	}

	rows.Close()

}