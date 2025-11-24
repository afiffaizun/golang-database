package main

import "database/sql"

// insertCustomer menyisipkan satu baris baru ke tabel customers.
func insertCustomer(db *sql.DB, name, email string) error {
	const query = `
		INSERT INTO customers (name, email)
		VALUES (?, ?)
	`
	_, err := db.Exec(query, name, email)
	return err
}

// deleteCustomer menghapus baris customers berdasarkan email.
func deleteCustomer(db *sql.DB, email string) error {
	const query = `
		DELETE FROM customers
		WHERE email = ?
	`
	_, err := db.Exec(query, email)
	return err
}
