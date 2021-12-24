# webview-example

This repo contains an example/starter for building a [webview](https://github.com/webview/webview) app in Go.

Webview is a cross-platform library for creating desktop web applications (similar to Electron).

This example uses a HTTP server to serve static content and handle API requests. Because the server is running locally on the machine, you have full access to the network, file system, etc.

## Getting Started

The [webview](https://github.com/webview/webview) repo has some getting started info of its own. TL;DR: depending on your OS/distro, webview will use whatever web library it can find.

For Ubuntu, I had to install these:

```
sudo apt install -y libgtk-3-dev libwebkit2gtk-4.0-dev
```

Once installed, you can run the app as you would normally

```
go run cmd/main/main.go
```

## Quick Walkthrough

The first thing the app does is start a HTTP server. It opens a listener on the address `127.0.0.1:0` which tells the OS to pick a random open port.

For the webview side, a simple window is first created. The URL can be a data string so we just pass it the HTML directly. It doesn't have to be a data string, we could just as easily navigate to a URL on the server that displays the page.

In order for the frontend to know the host, the address of the HTTP server is injected into the HTML using the [html/template](https://pkg.go.dev/html/template) package.

At this point we have a fully functional web page being rendered. From here you could build out something like a React app and do any background work using the HTTP server.

## Caching

By default (at least for WebKit), caching seems to be enabled for the webview. The static file handler adds a `Cache-Control: no-cache` header which solves this issue for the static files. Caching the static files isn't necessary since we are serving them from the local machine.