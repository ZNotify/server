package web

import (
	"embed"
	"io/fs"
)

var UI fs.FS

//go:embed "static/*"
var f embed.FS

func Init() {
	var err error
	UI, err = fs.Sub(f, "static")
	if err != nil {
		panic(err)
	}
}
