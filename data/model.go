package data

type Post struct {
	Id           string
	Title        string
	ThumbnailUrl string
	ImageUrl     string
}

type Submission struct {
	Title        string `json:"title"`
	ThumbnailURL string `json:"thumbnailURL"`
	ImageURL     string `json:"imageURL"`
}
