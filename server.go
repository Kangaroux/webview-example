package main

import (
	"log"
	"net/http"
	"os"
)

type Server struct{}

var data []byte

func init() {
	f, err := os.Open("app.js")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if _, err := f.Read(data); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) GetAssets(rw http.ResponseWriter, req *http.Request) {
	rw.Write(data)
}
