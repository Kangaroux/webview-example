# webview-example

This repo contains an example/starter for building a [webview](https://github.com/webview/webview) app in Go.

Webview is a cross-platform library for creating desktop web applications (similar to Electron).

This repo compiles to a 6.7MB executable on Pop!OS 22.04.

## Quickstart

### Prerequisites

Check the [webview](https://github.com/webview/webview#prerequisites) docs for what dependencies you'll need.

### How it Works

This example uses a HTTP server for serving the content. The server listens on a random available port
so there is no concern for conflicts.

Static files and HTML can be found in `server/`. These files are embedded into the app using the
[embed](https://pkg.go.dev/embed) package.

It's not included in this example, but you could implement a REST API or use websockets to handle the
communication between the UI and the server. The `webview` library also provides `Eval()` and `Bind()`
functions for calling JS and Go code respectively. I have not done any benchmarking, but I imagine that
any difference between these choices is negligible.

### Running the App

```bash
go run .
```

The window first appears as a blank white page while the browser engine is initializing. There isn't
a way to avoid this with the current webview version, however there is an [open feature request](https://github.com/webview/webview/issues/495)
to support this.

### Developer Tools

Developer tools can be enabled/disabled by passing a bool to `webview.New()`. You can access the devtools
via right click > `Inspect Element`.

You can disable the devtools when building (or running) the app by passing `-ldflags "-X main.debug=false"`.
For example:

```bash
# Build to ./app and disable devtools
go build -ldflags "-X main.debug=false" -o app .
```

There doesn't seem to be any performance hit to leaving devtools on, but probably worth disabling if you
don't want users snooping around your app.