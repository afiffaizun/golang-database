package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	// Koneksi ke MySQL di Docker
	// Format: username:password@tcp(host:port)/database_name
	// Untuk Docker, host tetap localhost karena port sudah di-mapping
	db, err := sql.Open("mysql", "root:rootpassword@tcp(localhost:3306)/belajar_golang")
	if err != nil {
		log.Fatal("Error koneksi database:", err)
	}
	defer db.Close()

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test koneksi dengan retry (karena Docker butuh waktu untuk start)
	fmt.Println("Menunggu MySQL siap...")
	for i := 0; i < 10; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		fmt.Printf("Mencoba koneksi... (attempt %d/10)\n", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Error ping database:", err)
	}
	fmt.Println("✓ Berhasil terhubung ke database MySQL di Docker!")

	// Buat tabel jika belum ada
	createTable(db)

	// Insert data
	fmt.Println("\n--- Insert Data ---")
	insertUser(db, "John Doe", "john@example.com")
	insertUser(db, "Jane Smith", "jane@example.com")

	// Query semua data
	fmt.Println("\n--- Query Semua Data ---")
	users := getAllUsers(db)
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}

	// Query data berdasarkan ID
	fmt.Println("\n--- Query Data by ID ---")
	user := getUserByID(db, 1)
	if user != nil {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}

	// Update data
	fmt.Println("\n--- Update Data ---")
	updateUser(db, 1, "John Updated", "john.updated@example.com")

	// Delete data
	fmt.Println("\n--- Delete Data ---")
	deleteUser(db, 2)

	// Query lagi setelah update dan delete
	fmt.Println("\n--- Data Setelah Update & Delete ---")
	users = getAllUsers(db)
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}
}

// Fungsi untuk membuat tabel
func createTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Error membuat tabel:", err)
	}
	fmt.Println("✓ Tabel users berhasil dibuat/tersedia")
}

// Fungsi untuk insert data
func insertUser(db *sql.DB, name, email string) {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	result, err := db.Exec(query, name, email)
	if err != nil {
		log.Printf("Error insert data: %v\n", err)
		return
	}

	id, _ := result.LastInsertId()
	fmt.Printf("✓ Data berhasil diinsert dengan ID: %d\n", id)
}

// Fungsi untuk mendapatkan semua user
func getAllUsers(db *sql.DB) []User {
	query := "SELECT id, name, email FROM users"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error query data:", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Fatal("Error scan data:", err)
		}
		users = append(users, user)
	}

	return users
}

// Fungsi untuk mendapatkan user berdasarkan ID
func getUserByID(db *sql.DB, id int) *User {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := db.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("User dengan ID %d tidak ditemukan\n", id)
			return nil
		}
		log.Fatal("Error scan data:", err)
	}

	return &user
}

// Fungsi untuk update data
func updateUser(db *sql.DB, id int, name, email string) {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	result, err := db.Exec(query, name, email, id)
	if err != nil {
		log.Printf("Error update data: %v\n", err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("✓ Data berhasil diupdate, rows affected: %d\n", rowsAffected)
}

// Fungsi untuk delete data
func deleteUser(db *sql.DB, id int) {
	query := "DELETE FROM users WHERE id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error delete data: %v\n", err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("✓ Data berhasil didelete, rows affected: %d\n", rowsAffected)
}