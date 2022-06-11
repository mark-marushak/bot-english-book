package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var sqlxInstance *sqlx.DB

func Sqlx() *sqlx.DB {
	if sqlxInstance == nil {
		sqlxInstance = sqlx.NewDb(connect(), "postgres")
		//sqlxInstance = sqlx.MustConnect("postgres", GetPostgresConfig())
	}
	return sqlxInstance
}
