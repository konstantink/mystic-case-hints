package main

import (
	"context"
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

func initHandlers(prefix string, handler HandlerWrapper) {
	http.HandleFunc(fmt.Sprintf("%s/haunted-castle", prefix), handler("Haunted Castle"))
	http.HandleFunc(fmt.Sprintf("%s/time-machine", prefix), handler("Time Machine"))
	http.HandleFunc(fmt.Sprintf("%s/vampires-tale", prefix), handler("Vampire's Tale"))
	http.HandleFunc(fmt.Sprintf("%s/ufo-crash", prefix), handler("UFO Crash"))
	http.HandleFunc(fmt.Sprintf("%s/school-of-magic", prefix), handler("School of Magic"))
	http.HandleFunc(fmt.Sprintf("%s/lost-island", prefix), handler("Lost Island"))
	http.HandleFunc(fmt.Sprintf("%s/simulation-theory", prefix), handler("Simulation Theory"))
	http.HandleFunc(fmt.Sprintf("%s/national-treasure", prefix), handler("National Treasure"))
	http.HandleFunc(fmt.Sprintf("%s/unfinished-case-of-holmes", prefix), handler("Unfinished case of Holmes"))
	http.HandleFunc(fmt.Sprintf("%s/christmas-adventure", prefix), handler("Christmas Adventure"))
}

func initRedirectHandlers() {
	var paths = []string{
		"/haunted-castle",
		"/time-machine",
		"/vampires-tale",
		"/ufo-crash",
		"/school-of-magic",
		"/lost-island",
		"/simulation-theory",
		"/national-treasure",
		"/unfinished-case-of-holmes",
		"/christmas-adventure",
	}
	for i := range paths {
		http.Handle(fmt.Sprintf("%s-hints", paths[i]), redirectTo("hints", paths[i]))
		http.Handle(fmt.Sprintf("%s-additional-task", paths[i]), redirectTo("additional-task", paths[i]))
	}
}

func initRoutes(ctx context.Context) {
	initSystemHandlers()
	initRedirectHandlers()
	initHandlers("/hints", boxHints)
	initHandlers("/additional-task", boxAdditionalTask)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/v1/htmx/feedback", feedbackHandler(ctx))
}
