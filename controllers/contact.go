package controllers

import (
	"fmt"
	"net/http"
)

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Contact page of Go Bank!")
}
