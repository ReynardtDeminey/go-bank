package router

import (
	"net/http"

	controllers "github.com/ReynardtDeminey/go-bank/controllers/static"
	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Landing)

	http.ListenAndServe(":8080", r)
}
