package main

import (
	"fmt"
	"github.com/yskim308/fake-news/repository"
	"io"
	"log"
	"net/http"
)

func main() {
	repo := repository.Repository{}
	repo.Connect()

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	})

	port := 4000
	fmt.Printf("listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
