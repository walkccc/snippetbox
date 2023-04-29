package main

import "github.com/walkccc/snippetbox/internal/models"

// templateData holds structure for any dynamic data that we want to pass to
// our HTML templates.
type templateData struct {
	Snippet *models.Snippet
}
