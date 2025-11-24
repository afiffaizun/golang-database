package main

import (
	"fmt"
	"log"
)

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}
	defer db.Close()

	fmt.Println("✅ Berhasil terhubung ke database MySQL!")

	if err := insertCustomer(db, "John Doe", "john@example.com"); err != nil {
		log.Fatal("Insert customer gagal:", err)
	}
	fmt.Println("✅ Berhasil insert customer!")

	if err := deleteCustomer(db, "john@example.com"); err != nil {
		log.Fatal("Delete customer gagal:", err)
	}
	fmt.Println("🗑️ Customer terhapus!")
}
