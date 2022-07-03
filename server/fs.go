package server

import (
	"embed"
	"io/fs"
)

var (
	//go:embed public/*
	public           embed.FS
	publicContent, _ = fs.Sub(public, "public")

	//go:embed static/*
	static           embed.FS
	staticContent, _ = fs.Sub(static, "static")
)
