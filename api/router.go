package api

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/static/").HandlerFunc(ServeStatic)

	return r
}
