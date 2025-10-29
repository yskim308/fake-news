package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/yskim308/fake-news/data"
	"github.com/yskim308/fake-news/repository"
	"github.com/yskim308/fake-news/view"
)

func main() {
	repo := &repository.Repository{}
	repo.Connect()

	http.HandleFunc("/news/articles/", func(w http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		pathPrefix := "/news/articles/"
		if path == pathPrefix {
			http.Error(w, "article ID missing", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(strings.TrimPrefix(path, pathPrefix))
		if err != nil {
			http.Error(w, "id is not an integer, invalid path", http.StatusBadRequest)
			return
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
			return
		}
	})

	http.HandleFunc("/submit", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
			return
		}
		defer req.Body.Close()

		var submission data.Submission

		err := json.NewDecoder(req.Body).Decode(&submission)
		if err != nil {
			log.Printf("Error decoding JSON: %v", err)
			http.Error(w, "invalid request body format", http.StatusBadRequest)
			return
		}

		var id int64
		id, err = repo.CreateEntry(submission)
		if err != nil {
			log.Printf("Error creating entry in database: %v", err)
			http.Error(w, "error creating entry in database: %v", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/JSON")
		w.WriteHeader(http.StatusOK)
	})

	port := 4000
	fmt.Printf("listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
