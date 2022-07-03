package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// The static/ dir is hosted at /static/
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.FS(staticContent))))

	// The public/ dir is hosted at /
	r.PathPrefix("/").Handler(http.FileServer(http.FS(publicContent)))

	return r
}
