package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the home page of Go Bank!")
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}