package controllers

import (
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Login page of Go Bank!")
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to your account dashboard!")
}
