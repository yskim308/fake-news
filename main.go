package main

import (
	"fmt"
	"github.om/yskim308/fake-news/db"
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	})

	port := 4000
	fmt.Printf("listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
