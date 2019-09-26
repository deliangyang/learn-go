package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	logId   int
	logTime string
	path    string
	method  string
}

func main() {
	db, err := sql.Open("postgres", "host=192.168.1.16 port=5432 user=postgres password=postgres dbname=api_logs sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("select log_id, log_time, path, method from logs order by log_id desc limit 2")
	if err != nil {
		log.Fatal(err)
	}
	var users []User
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.logId, &user.logTime, &user.path, &user.method)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	fmt.Println(users)

	stmt, err := db.Prepare("select log_id, log_time, path, method from logs where method=$1 order by log_id desc limit $2")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	rows, err = stmt.Query("POST", 10)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.logId, &user.logTime, &user.path, &user.method)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	fmt.Println(users)

}
