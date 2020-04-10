package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	sdb, err := sql.Open("mysql", "root:Y8NeLXer6u@tcp(192.168.190.150:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	db = sdb
}

func saveMsgSub(urlType string, url string, task string, tags string, content string) {
	stmt, err := db.Prepare(`INSERT INTO msgsub(type, url, task, tags, content) VALUES(?, ?, ?, ?, ?)`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	ret, err := stmt.Exec(urlType, url, task, tags, content)
	if err != nil {
		panic(err)
	}

	if LastInsertId, err := ret.LastInsertId(); err == nil {
		fmt.Println("LastInsertId:", LastInsertId)
	}
}
