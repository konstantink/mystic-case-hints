package main

import "net/http"

func check(err error, w http.ResponseWriter) bool {
	if err != nil {
		http.Error(w, err.Error(), 500)
		return false
	}
	return true
}
