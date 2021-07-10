package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@123456789@/PeladeirosFut(tcp:localhost:3606)")//user:password@/dbname
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	con = db
	return db
}
