package repository

import (
	"fmt"
	"log"
)

type Post struct {
	Id           int
	Iitle        string
	ThumbnailUrl string
	ImageUrl     string
}

func (r *Repository) CreateEntry(
	title string, 
	thumbnailURL string,
	imageURL string
) error {
	db := r.db
	if db == nil {
		fmt.Println("DB NOT INITIALIZED")
		return fmt.Errorf("database connection not initialized")
	}

	_, err := db.Exec(`
		INSERT INTO POSTS (title, thumbnail_url, image_url)
		VALUES ($1, $2, $3)
	`, title, thumbnailURL, imageURL)

	if err != nil {
		log.Printf("failed to create entry: %v", err)
	}
}

func (r *Repository) GetEntry(id int) (Post, error) {
	db := r.db
	if db == nil {
		fmt.Println("DB NOT INITIALIZED")
		return Post{}, fmt.Errorf("database connection not initialized")
	}

	var post Post

	err := db.QueryRow(`
		SELECT title, thumbnail_url, image_url
		FROM posts 
		WHERE id=$1
		`, id).Scan(&post.Title, &post.ThumbnailUrl, &post.ImageUrl)
	if err != nil {
		log.Printf("failed to fetch entry: %v", err)
	}
	return post, nil
}
