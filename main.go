package main

import (
	"fmt"
	"net/http"

	"github.com/ReynardtDeminey/go-bank/router"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the home page of Go Bank!")
}

func main() {
	router.Router()

}
