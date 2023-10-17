package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"ip.com/config"
)

func NewMysqlConn(c *config.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
