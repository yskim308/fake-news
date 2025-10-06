package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() {
	connStr := "postgres://young@localhost:5432/postgres?sslmode=disable"

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("Open error:", err)
	}
	defer db.Close()

	// CREATE TABLE
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL
		);
	`)
	if err != nil {
		log.Fatal("Create table error:", err)
	}

	// INSERT
	_, err = db.Exec(`INSERT INTO users (name, email) VALUES ($1, $2)`, "Alice", "alice@example.com")
	if err != nil {
		log.Fatal("Insert error:", err)
	}

	// READ
	var name, email string
	err = db.QueryRow(`SELECT name, email FROM users WHERE name=$1`, "Alice").Scan(&name, &email)
	if err != nil {
		log.Fatal("Query error:", err)
	}
	fmt.Println("Fetched:", name, email)
}
