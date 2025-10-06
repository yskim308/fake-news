package repository

import (
	"fmt"
	"log"
)

func CreateEntry(title string, thumbnailURL string, imageURL string) {
	db := DBinstance
	if db == nil {
		fmt.Println("DB NOT INITIALIZED")
		return
	}

	_, err := db.Exec(`
		INSERT INTO POSTS ( title, thumbnail_url, image_url)
		VALUES ($1, $2, $3)
	`, title, thumbnailURL, imageURL)

	if err != nil {
		log.Printf("failed to create entry: %v", err)
	}
}

func GetEntry(id int) {
	db := DBinstance
	if db == nil {
		fmt.Println("DB NOT INITIALIZED")
		return
	}

	var title, thumbnailURL, imageURL string

	err := db.QueryRow(`
		SELECT title, thumbnail_url, image_url
		FROM posts 
		WHERE id=$1
		`, id).Scan(&title, &thumbnailURL, &imageURL)
	if err != nil {
		log.Printf("failed to fetch entry: %v", err)
	}
}
