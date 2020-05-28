package controllers

import (
	"fmt"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Here you can open an account!")
}
