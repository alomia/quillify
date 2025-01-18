package data

import "time"

type Book struct {
	ID            int64     `json:"id"`
	Title         string    `json:"title"`
	Authors       []string  `json:"authors"`
	Publisher     string    `json:"publisher"`
	PublishedDate string    `json:"published_date"`
	Genres        []string  `json:"genres"`
	ISBN          string    `json:"isbn"`
	Language      string    `json:"language"`
	Pages         int       `json:"pages"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Version       int       `json:"version"`
}
