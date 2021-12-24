package main

import (
	"log"
	"net/http"

	"github.com/kangaroux/webview/server"
	"github.com/webview/webview"
)

func main() {
	go func() {
		server := server.Server{}

		http.HandleFunc("/static/", server.StaticAssets)

		if err := http.ListenAndServe(":8000", nil); err != nil {
			log.Fatal(err)
		}
	}()

	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate(`data:text/html,
<html>
	<head>
		<link type="text/css" rel="stylesheet" href="http://localhost:8000/public/style.css">
		<script src="http://localhost:8000/public/app.js"></script>
	</head>
</html>
	`)
	w.Run()
}
