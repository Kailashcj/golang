package db

import (
	"database/sql"
	"fmt"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//CreateDatabase Mysql db connections
func CreateDatabase() (*sql.DB, error) {
	serverName := "localhost:3306"
	user := "sammy"
	password := "password"
	dbName := "demo"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", user, password, serverName, dbName)
	fmt.Println(connectionString)
	/*
		Perhaps counter-intuitively, sql.Open() does not establish any connections to the database,
		nor does it validate driver connection parameters.
		Instead, it simply prepares the database abstraction for later use.
	*/
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	fmt.Println(db)
	fmt.Println("checking db connectivity")
	return db, nil
}
