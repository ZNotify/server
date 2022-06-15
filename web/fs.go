package web

import (
	"embed"
	"io/fs"
)

var UI fs.FS

func Init(f *embed.FS) {
	var err error
	UI, err = fs.Sub(f, "static")
	if err != nil {
		panic(err)
	}
}
