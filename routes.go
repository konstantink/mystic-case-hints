package main

import (
	"fmt"
	"net/http"
)

func initSystemHandlers() {
	http.HandleFunc("/favicon.ico", fileHandlerFunc("./images/favicon.ico"))
	http.HandleFunc("/sitemap.xml", fileHandlerFunc("./sitemap.xml"))
	http.HandleFunc("/robots.txt", fileHandlerFunc("./robots.txt"))
	http.Handle("/static/", http.StripPrefix("/static/", folderHandler("./static")))
	http.Handle("/images/", http.StripPrefix("/images/", folderHandler("./images")))
}

func initHintsHandlers(prefix string) {
	http.HandleFunc(fmt.Sprintf("%s/haunted-castle", prefix), boxHints("Haunted Castle"))
	http.HandleFunc(fmt.Sprintf("%s/time-machine", prefix), boxHints("Time Machine"))
	http.HandleFunc(fmt.Sprintf("%s/vampires-tale", prefix), boxHints("Vampire's Tale"))
	http.HandleFunc(fmt.Sprintf("%s/ufo-crash", prefix), boxHints("UFO Crash"))
	http.HandleFunc(fmt.Sprintf("%s/school-of-magic", prefix), boxHints("School of Magic"))
	http.HandleFunc(fmt.Sprintf("%s/lost-island", prefix), boxHints("Lost Island"))
	http.HandleFunc(fmt.Sprintf("%s/simulation-theory", prefix), boxHints("Simulation Theory"))
	http.HandleFunc(fmt.Sprintf("%s/national-treasure", prefix), boxHints("National Treasure"))
	http.HandleFunc(fmt.Sprintf("%s/unfinished-case-of-holmes", prefix), boxHints("Unfinished case of Holmes"))
	http.HandleFunc(fmt.Sprintf("%s/christmas-adventure", prefix), boxHints("Christmas Adventure"))
}

func initRedirectHandlers() {
	var paths = []string{
		"/haunted-castle-hints",
		"/time-machine-hints",
		"/vampires-tale-hints",
		"/ufo-crash-hints",
		"/school-of-magic-hints",
		"/lost-island-hints",
		"/simulation-theory-hints",
		"/national-treasure-hints",
		"/unfinished-case-of-holmes-hints",
		"/christmas-adventure-hints",
	}
	for i := range paths {
		http.Handle(paths[i], redirectToHints(paths[i]))
	}
}

func initRoutes() {
	initSystemHandlers()
	initRedirectHandlers()
	initHintsHandlers("/hints")
}
