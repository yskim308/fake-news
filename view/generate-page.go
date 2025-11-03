package view

import (
	"bytes"
	"context"
	"html/template"
	"log"

	"github.com/yskim308/fake-news/data"
)

type EntryGetter interface {
	GetEntry(ctx context.Context, id string) (data.Post, error)
}

func GeneratePage(ctx context.Context, id string, repo EntryGetter) (string, error) {
	const templateFilePath = "./view/main.html"
	tmpl := template.Must(template.ParseFiles(templateFilePath))

	post, err := repo.GetEntry(ctx, id)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, post)
	generatedHTML := buffer.String()

	return generatedHTML, nil
}
