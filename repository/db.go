package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Repository struct {
	db *sql.DB
}

func (r *Repository) Connect() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL not sent in env")
	}

	var err error
	DBinstance, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("Open error:", err)
	}

	err = DBinstance.Ping()
	if err != nil {
		log.Fatal("Ping error:", err)
	}

	r.db = DBinstance
	fmt.Println("connected to database")
}
