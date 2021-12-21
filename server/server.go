package server

import (
	"net/http"
)

type Server struct{}

func (s *Server) GetAssets(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Cache-Control", "no-cache")
	http.ServeFile(rw, req, "."+req.URL.Path)
}
