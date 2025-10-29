package view

import (
	"bytes"
	"github.com/yskim308/fake-news/data"
	"html/template"
	"log"
)

type EntryGetter interface {
	GetEntry(id int) (data.Post, error)
}

func GeneratePage(id int, repo EntryGetter) (string, error) {
	const templateFilePath = "./main.html"
	tmpl := template.Must(template.ParseFiles(templateFilePath))

	post, err := repo.GetEntry(id)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, post)
	generatedHTML := buffer.String()

	return generatedHTML, nil
}
