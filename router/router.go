package router

import (
	"net/http"

	controllers "github.com/ReynardtDeminey/go-bank/controllers"

	"github.com/gorilla/mux"
)

// Router defines the routes for Go Bank!
func Router() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Landing)
	r.HandleFunc("/contact", controllers.Contact)
	r.HandleFunc("/login", controllers.Login)
	r.HandleFunc("/signup", controllers.SignUp)
	r.HandleFunc("/dashboard", controllers.Dashboard)

	http.ListenAndServe(":8080", r)
}
