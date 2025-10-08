package view

import (
	"bytes"
	"html/template"
	"log"

	"github.com/yskim308/fake-news/repository"
)

func GeneratePage(id int, repo *repository.Repository) string {
	const templateFilePath = "./form.html"
	tmpl := template.Must(template.ParseFiles(templateFilePath))

	post, err := repo.GetEntry(id)
	if err != nil {
		log.Fatal(err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, post)
	generatedHTML := buffer.String()

	return generatedHTML
}
