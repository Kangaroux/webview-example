package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"

	"github.com/kangaroux/webview/api"
	"github.com/webview/webview"
)

func main() {
	addrChan := make(chan string)

	go func() {
		router := api.NewRouter()

		listener, err := net.Listen("tcp", "127.0.0.1:0")

		if err != nil {
			log.Fatal(err)
		}

		log.Println("API listening on", listener.Addr().String())
		addrChan <- listener.Addr().String()

		if err := http.Serve(listener, router); err != nil {
			log.Fatal(err)
		}
	}()

	debug := true
	addr := <-addrChan
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)
	data := fmt.Sprintf(`
	<html>
		<head>
			<link rel="stylesheet" href="http://%[1]s/static/style.css">
		</head>
		<body>
			<div id="root"></div>
			<script type="text/javascript">
				var API_HOST = "%[1]s";
				var API_URL = "http://" + API_HOST;
			</script>
			<script src="http://%[1]s/static/app.js"></script>
		</body>
	</html>
		`, addr)
	w.Navigate("data:text/html," + url.QueryEscape(data))
	w.Run()
}
