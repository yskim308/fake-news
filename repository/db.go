package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DBinstance *sql.DB

func Connect() {
	connStr := "postgres://young@localhost:5432/postgres?sslmode=disable"

	var err error
	DBinstance, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("Open error:", err)
	}

	err = DBinstance.Ping()
	if err != nil {
		log.Fatal("Ping error:", err)
	}

	fmt.Println("connected to database")
}
