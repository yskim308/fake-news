package repository

import (
	"fmt"
	"github.com/yskim308/fake-news/data"
	"log"
)

func (r *Repository) CreateEntry(submision data.Submission) (string, error) {
	db := r.db
	if db == nil {
		fmt.Println("DB NOT INITIALIZED")
		return "", fmt.Errorf("database connection not initialized")
	}

	var insertedID string

	err := db.QueryRow(`
        INSERT INTO POSTS (title, thumbnail_url, image_url)
        VALUES ($1, $2, $3)
        RETURNING id
    `, submision.Title, submision.ThumbnailURL, submision.ImageURL).Scan(&insertedID) // 4. Scan the returned ID into your variable

	if err != nil {
		log.Printf("failed to create entry: %v", err)
		return "", err
	}

	// 5. Return the new ID
	return insertedID, nil
}

func (r *Repository) GetEntry(id string) (data.Post, error) {
	db := r.db
	if db == nil {
		fmt.Println("DB NOT INITIALIZED")
		return data.Post{}, fmt.Errorf("database connection not initialized")
	}

	var post data.Post

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
