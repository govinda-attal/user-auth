package usrstore

import (
	"database/sql"
	// ...
	_ "github.com/lib/pq"
)

func InitStore(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	return db, err
}
