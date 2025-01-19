package data

import "time"

type Book struct {
	ID            int64     `json:"id"`
	Title         string    `json:"title"`
	Authors       []string  `json:"authors"`
	Publisher     string    `json:"publisher,omitempty"`
	PublishedDate string    `json:"published_date,omitempty"`
	Genres        []string  `json:"genres,omitempty"`
	ISBN          string    `json:"isbn,omitempty"`
	Language      string    `json:"language,omitempty"`
	Pages         int       `json:"pages,omitempty"`
	Description   string    `json:"description,omitempty"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
	Version       int       `json:"version"`
}
