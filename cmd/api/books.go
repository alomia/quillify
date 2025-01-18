package main

import (
	"net/http"
	"time"

	"github.com/alomia/quillify/internal/data"
	"github.com/labstack/echo/v4"
)

func (app *application) createBookHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "book created successfully",
	})
}

func (app *application) showBookHandler(c echo.Context) error {
	id, err := app.readIDParam(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	book := data.Book{
		ID:            id,
		Title:         "Cien años de soledad",
		Authors:       []string{"Gabriel García Márquez"},
		Publisher:     "Editorial Sudamericana",
		PublishedDate: "1967-05-30",
		Genres:        []string{"Fiction", "Magical Realism", "Literature"},
		ISBN:          "978-84-376-0494-7",
		Language:      "Spanish",
		Pages:         471,
		Description:   "Un clásico de la literatura universal que narra la historia de la familia Buendía a lo largo de varias generaciones en el pueblo ficticio de Macondo.",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Version:       1,
	}

	return c.JSON(http.StatusOK, book)
}
