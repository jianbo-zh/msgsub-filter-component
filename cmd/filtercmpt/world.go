package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func world() {
	db, err := sql.Open("mysql", "root:WYG0pTooU6@tcp(192.168.190.150:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// stmtIns, err := db.Prepare("INSERT INTO mytable(name) VALUES( ? )")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer stmtIns.Close()

	stmtOut, err := db.Prepare("SELECT * FROM mytable WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	// for i := 0; i < 25; i++ {
	// 	_, err = stmtIns.Exec("Hello" + strconv.Itoa(i)) // Insert tuples (i, i^2)
	// 	if err != nil {
	// 		panic(err.Error()) // proper error handling instead of panic in your app
	// 	}
	// }

	var id int
	var name string

	err = stmtOut.QueryRow(13).Scan(&id, &name)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("id: %d, name: %s", id, name)

}
