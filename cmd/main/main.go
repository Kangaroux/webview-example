package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/kangaroux/webview-example/api"
	"github.com/webview/webview"
)

var htmlTemplate = ""

const (
	htmlTemplatePath = "static/index.template.html"
)

func init() {
	data, err := ioutil.ReadFile(htmlTemplatePath)

	if err != nil {
		log.Fatal(err)
	}

	htmlTemplate = string(data)
}

func buildFromTemplate(templateString string, data interface{}) string {
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

	html := buildFromTemplate(htmlTemplate, struct{ Host string }{addr})

	w.Navigate("data:text/html," + url.QueryEscape(html))
	w.Run()
}
