package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func fileHandlerFunc(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	}
}

func folderHandler(path string) http.Handler {
	return http.FileServer(http.Dir(path))
}

func redirectToHints(path string) http.Handler {
	url := fmt.Sprintf("/hints/%s", strings.Replace(path, "-hints", "", 1))
	return http.RedirectHandler(url, http.StatusMovedPermanently)
}

func hauntedCastleHints(w http.ResponseWriter, r *http.Request) {
	context := struct {
		BoxName string
		GtmID   string
		Hints   []Level
	}{
		BoxName: "Haunted Castle",
		GtmID:   "",
		Hints: []Level{
			{
				Name: "Envelop #1",
				Steps: []Sublevel{
					{
						Name: "Double sided photo of the hall",
						Tips: []Hint{
							{
								Name: "Tip 1",
								Text: "Find the differences between the two photos and make a word from the letters in squares with differences.",
							},
							{
								Name: "Answer",
								Text: "Twelve. One and two.",
							},
						},
					},
					{
						Name: "Laminated photo",
						Tips: []Hint{
							{
								Name: "Tip 1",
								Text: "Enlighten the photo with a flashlight.",
							},
							{
								Name: "Answer",
								Text: "Ghost has a hat and a cane.",
							},
						},
					},
				},
			},
		},
	}
	files := []string{
		"./templates/base.html",
		"./templates/views/header.html",
		"./templates/views/hints.html",
		// "./templates/views/main.html",
		// "./templates/views/feedback.html",
	}

	tpl, err := template.ParseFiles(files...)
	if !check(err, w) {
		return
	}

	tpl.ExecuteTemplate(w, "base", context)
}
