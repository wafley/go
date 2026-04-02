package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/school_db"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
        log.Fatal("Database connection failed:", err)
    }

	err = db.Ping()
    if err != nil {
        log.Fatal("Database not responding:", err)
    }

	return db
}