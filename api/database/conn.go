package database

import (
	"database/sql"
	"time"
)

const (
	connString = "root:kkk123@tcp(localhost:33060)/home?allowNativePasswords=True"
)

func NewDbConn() (*sql.DB, error) {
	return sql.Open("mysql", connString)
}

func Init() error {
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return err
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}
