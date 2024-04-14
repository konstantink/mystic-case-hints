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

type HandlerWrapper func(string) http.HandlerFunc

var Box2File = map[string]string{
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

func fileHandlerFunc(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	}
}

func folderHandler(path string) http.Handler {
	return http.FileServer(http.Dir(path))
}

func redirectTo(prefix, path string) http.Handler {
	url := fmt.Sprintf("/%s%s", prefix, path)
	log.Printf("Redirect to: %s", url)
	return http.RedirectHandler(url, http.StatusMovedPermanently)
}

func boxHints(boxName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var hints []Level
		hintsStream, _ := os.ReadFile(fmt.Sprintf("./hints/%s", Box2File[boxName]))
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
			"./templates/views/last_word.html",
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

func boxAdditionalTask(boxName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Processing request %s", boxName)
		var task Task
		taskStream, _ := os.ReadFile(fmt.Sprintf("./additional_tasks/%s", Box2File[boxName]))
		err := json.Unmarshal(taskStream, &task)
		if err != nil {
			log.Print(err.Error())
		}
		context := struct {
			BoxName string
			GtmID   string
			Task    Task
		}{
			BoxName: boxName,
			GtmID:   "",
			Task:    task,
		}
		files := []string{
			"./templates/base.html",
			"./templates/views/header.html",
			"./templates/views/additional_task.html",
			"./templates/views/feedback_form.html",
			"./templates/views/last_word.html",
			"./templates/views/footer.html",
		}

		funcs := template.FuncMap{
			"safeHTML": func(val any) template.HTML {
				return template.HTML(fmt.Sprint(val))
			},
			"split": func(val string) []string {
				return strings.Split(val, "|")
			},
			"length": func(val []string) int {
				return len(val)
			},
		}

		tpl, err := template.New("additional_task").Funcs(funcs).ParseFiles(files...)
		if !check(err, w) {
			log.Print(err.Error())
			return
		}

		// tpl.ExecuteTemplate(os.Stdout, "base", context)
		tpl.ExecuteTemplate(w, "base", context)
	}
}
