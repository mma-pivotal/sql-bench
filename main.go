package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func insertRecord(db *sql.DB, count int) {
	for i := 0; i < count; i++ {
		_, err := db.Exec("insert into test_table (name,description) values ('2nd','2nd record');")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	db, err := sql.Open("mysql", "root:1VPnSsXuX7vXpEYm86uzK7X5mAvLTs@tcp(10.193.28.29:3306)/mysql")
	if err != nil {
		fmt.Println(err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	insertRecord(db, 10000)

}
