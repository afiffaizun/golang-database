package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID int
	Name string
	Email string
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_golang")
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	fmt.Println("Meninggu MySQL siap...")
	for i := 0; i < 10; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		fmt.Printf("Mencoba Koneksi...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Error ping database", err)
	}
	fmt.Println("Berhasil terhubung MySQL")

	//Create Table
	createTable(db)

	//Inser Data
	fmt.Println("Insert Data")
	insertUser(db, "Apip", "apip@")
	insertUser(db, "Budi", "budi@gmail.com")
}

