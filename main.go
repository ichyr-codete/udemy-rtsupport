package main

import (
	"net/http"
)

func main() {
	router := NewRouter()
	http.Handle("/", router)
	http.ListenAndServe(":4000", nil)
}
