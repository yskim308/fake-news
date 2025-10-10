package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/yskim308/fake-news/repository"
	"github.com/yskim308/fake-news/view"
)

func main() {
	repo := &repository.Repository{}
	repo.Connect()

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	})

	http.HandleFunc("/news/articles/", func(w http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		pathPrefix := "/news/articles/"
		if path == pathPrefix {
			http.Error(w, "article ID missing", http.StatusBadRequest)
		}

		id, err := strconv.Atoi(strings.TrimPrefix(path, pathPrefix))
		if err != nil {
			http.Error(w, "id is not an integer, invalid path", http.StatusBadRequest)
		}

		generatedHTML := view.GeneratePage(id, repo)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(generatedHTML))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		tmpl, err := template.ParseFiles("./view/form.html")
		if err != nil {
			log.Fatal("template not found: ", err)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal("error executing template for form: ", err)
		}
	})

	port := 4000
	fmt.Printf("listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
