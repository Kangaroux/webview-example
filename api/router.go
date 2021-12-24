package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Don't serve a file index of the /static/ dir
	r.Path("/static/").Handler(http.NotFoundHandler())
	r.PathPrefix("/static/").HandlerFunc(ServeStatic)

	return r
}
