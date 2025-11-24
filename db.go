package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// connectDB membangun koneksi dan memastikan database siap digunakan.
func connectDB() (*sql.DB, error) {
	const dsn = "root:root@tcp(localhost:3306)/belajar_golang?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("gagal membuka koneksi: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("gagal melakukan ping: %w", err)
	}

	return db, nil
}
