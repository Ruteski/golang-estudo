package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // import implicito
)

func main() {
	strConn := "root:@/golang?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", strConn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão aberta!")

	rows, err := db.Query("select * from usuarios")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println(rows)
}
