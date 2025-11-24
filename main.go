package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}
	defer db.Close()

	// Test koneksi
	err = db.Ping()
	if err != nil {
		log.Fatal("Gagal melakukan ping ke database:", err)
	}

	fmt.Println("✅ Berhasil terhubung ke database MySQL!")
}

func connectDB() (*sql.DB, error) {
	// DSN (Data Source Name) sesuai konfigurasi docker-compose.yml
	dsn := "root:root@tcp(localhost:3306)/belajar_golang?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("gagal membuka koneksi: %w", err)
	}

	// Test koneksi
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("gagal melakukan ping: %w", err)
	}

	return db, nil
}
