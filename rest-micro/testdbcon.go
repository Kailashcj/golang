package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var name string
	fmt.Println("A simple db connection")

	db, err := sql.Open("mysql", "sammy:password@tcp(127.0.0.1:3306)/demo")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO `test` (name) VALUES ('kailash')")

	if err != nil {
		panic(err.Error())
	}
	stmt, err := db.Prepare("SELECT * FROM `test`")
	if err != nil {
		fmt.Println("unable to read from test")
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(name)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	defer insert.Close()

}
