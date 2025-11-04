package mysql

import (
	"database/sql"
	"fmt"
	"time"
)

type MYSQLDB struct {
	db *sql.DB
}

func New() *MYSQLDB {
	db, err := sql.Open("mysql", "user:password@(localhost:3308)/game-db")
	if err != nil {
		panic(fmt.Errorf("can not open the mysql db"))
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MYSQLDB{db}
}
