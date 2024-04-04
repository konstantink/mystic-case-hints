package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
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

func boxHints(boxName string) func(w http.ResponseWriter, r *http.Request) {
	var box2Hints = map[string]string{
		"Haunted Castle":            "haunted-castle.json",
		"Time Machine":              "time-machine.json",
		"Vampire's Tale":            "vampires-tale.json",
		"UFO Crash":                 "ufo-crash.json",
		"School of Magic":           "school-of-magic.json",
		"National Treasure":         "national-treasure.json",
		"Unfinished case of Holmes": "unfinished-case-of-holmes.json",
		"Lost Island":               "lost-island.json",
		"Simulation Theory":         "simulation-theory.json",
		"Christmas Adventure":       "christmas-adventure.json",
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var hints []Level
		hintsStream, _ := os.ReadFile(fmt.Sprintf("./hints/%s", box2Hints[boxName]))
		err := json.Unmarshal(hintsStream, &hints)
		if err != nil {
			log.Print(err.Error())
		}

		context := struct {
			BoxName string
			GtmID   string
			Hints   []Level
		}{
			BoxName: boxName,
			GtmID:   "",
			Hints:   hints,
		}
		files := []string{
			"./templates/base.html",
			"./templates/views/header.html",
			"./templates/views/hints.html",
			"./templates/views/footer.html",
		}

		funcs := template.FuncMap{
			"split": func(val string) []string {
				return strings.Split(val, "|")
			},
		}

		tpl, err := template.New("hints").Funcs(funcs).ParseFiles(files...)
		if !check(err, w) {
			return
		}

		tpl.ExecuteTemplate(w, "base", context)
	}
}
