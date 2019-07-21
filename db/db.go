package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:12345@tcp(127.0.0.1)/dbname")

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
