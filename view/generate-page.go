package view

import (
	"bytes"
	"github.com/yskim308/fake-news/data"
	"html/template"
	"log"
)

type EntryGetter interface {
	GetEntry(id string) (data.Post, error)
}

func GeneratePage(id string, repo EntryGetter) (string, error) {
	const templateFilePath = "./view/main.html"
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
