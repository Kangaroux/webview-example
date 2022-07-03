package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/kangaroux/webview-example/api"
	"github.com/webview/webview"
)

var htmlTemplate = ""

const (
	htmlTemplatePath = "static/index.template.html"
)

// Read the HTML template into memory
func init() {
	data, err := ioutil.ReadFile(htmlTemplatePath)

	if err != nil {
		log.Fatal(err)
	}

	htmlTemplate = string(data)
}

// A wrapper for rendering a template using the provided context
func renderTemplate(templateString string, data interface{}) string {
	tpl, err := template.New("index").Parse(templateString)

	if err != nil {
		log.Fatal(err)
	}

	out := &strings.Builder{}

	if err := tpl.Execute(out, data); err != nil {
		log.Fatal(err)
	}

	return out.String()
}

func main() {
	addrChan := make(chan string)

	// Start the HTTP server
	go func() {
		router := api.NewRouter()

		// Let the OS pick an open port
		listener, err := net.Listen("tcp", "127.0.0.1:0")

		if err != nil {
			log.Fatal(err)
		}

		log.Println("API listening on", listener.Addr().String())

		// http.Serve blocks, so send the address we're listening on back to
		// the main thread so we can use it for the webview
		addrChan <- listener.Addr().String()

		if err := http.Serve(listener, router); err != nil {
			log.Fatal(err)
		}
	}()

	debug := true
	addr := <-addrChan

	w := webview.New(debug)
	defer w.Destroy()

	w.SetTitle("Webview Example")
	w.SetSize(800, 600, webview.HintNone)

	// Inject the server address and render the full template
	html := renderTemplate(htmlTemplate, struct{ Host string }{addr})

	w.Navigate("data:text/html," + html)
	w.Run()
}
