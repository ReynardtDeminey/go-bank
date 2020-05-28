package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the home page of Go Bank!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	http.ListenAndServe(":8080", r)
}
