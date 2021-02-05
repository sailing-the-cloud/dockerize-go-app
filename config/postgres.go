package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"time"
)

/**
 * Created by Manggala Pramuditya Wiryawan on 07/11/19 Nov, 2019
 * email : manggala.wiryawan@gmail.com
 */
const (
	MaxIdleConnection = 10
	MaxOpenConnection = 10
)

func ConnectDB() *sql.DB {
	return CreateDBConnection(fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME")))

}

// CreateDBConnection function for creating database connection
func CreateDBConnection(descriptor string) *sql.DB {
	db, err := sql.Open("postgres", descriptor)
	if err != nil {
		defer db.Close()
		return db
	}

	db.SetMaxIdleConns(MaxIdleConnection)
	db.SetMaxOpenConns(MaxOpenConnection)
	db.SetConnMaxLifetime(time.Second)

	return db
}

// CloseDb function for closing database connection
func CloseDb(db *sql.DB) {
	if db != nil {
		db.Close()
		db = nil
	}
}
