package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // import implicito
)

func Conectar() (*sql.DB, error) {
	strConn := "root:@/golang?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", strConn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
