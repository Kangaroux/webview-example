package main

import (
	"log"
	"net"
	"net/http"

	"github.com/kangaroux/webview-example/server"
	"github.com/webview/webview"
)

// Having this as a string lets us configure it with the `-ldflags` option when building.
var debug = "true"

// Starts the API server and returns the host name it's listening on.
func startServer() string {
	addr := make(chan string)
	router := server.NewRouter()

	go func() {
		// Let the OS pick an open port
		listener, err := net.Listen("tcp", "127.0.0.1:0")

		if err != nil {
			log.Fatal(err)
		}

		log.Println("API listening on", listener.Addr().String())

		addr <- listener.Addr().String()

		if err := http.Serve(listener, router); err != nil {
			log.Fatal(err)
		}
	}()

	return <-addr
}

func main() {
	host := startServer()

	w := webview.New(debug == "true")
	defer w.Destroy()

	w.SetTitle("Webview Example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://" + host)
	w.Run()
}
