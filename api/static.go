package api

import "net/http"

func ServeStatic(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Cache-Control", "no-cache")
	http.ServeFile(w, req, "."+req.URL.Path)
}
