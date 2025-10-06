package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// connect to database
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL not sent in env")
	}

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("Open error:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Ping error:", err)
	}

	fmt.Println("connected to database")

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	})

	port := 4000
	fmt.Printf("listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
