package controllers

import (
	"fmt"
	"net/http"
)

func Landing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the landing page of Go Bank!")
}
