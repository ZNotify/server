package web

import (
	"embed"
	"io/fs"
	"net/http"
)

var StaticHttpFS http.FileSystem

//go:embed "static/*"
var f embed.FS

func init() {
	FS, err := fs.Sub(f, "static")
	if err != nil {
		panic(err)
	}
	StaticHttpFS = http.FS(FS)
}
