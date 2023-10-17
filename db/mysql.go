package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"ip.com/config"
)

func NewMysqlConn(c *config.Config) (*sql.DB, error) {
	// <username>:<password>@<addr>/<database>
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s", c.MySQL.User, c.MySQL.Pass, c.MySQL.Addr, c.MySQL.DB),
	)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
