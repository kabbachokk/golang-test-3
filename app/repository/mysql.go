package repository

import (
	"database/sql"
)

// mysqlRepo
type mysqlRepo struct {
	conn *sql.DB
}

// NewMysqlRepo
func NewMysqlRepo(conn *sql.DB) *mysqlRepo {
	return &mysqlRepo{conn}
}
