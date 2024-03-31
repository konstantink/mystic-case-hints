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
	http.HandleFunc(fmt.Sprintf("%s/haunted-castle", prefix), hauntedCastleHints)
}

func initRedirectHandlers() {
	var paths = []string{
		"/haunted-castle-hints",
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
