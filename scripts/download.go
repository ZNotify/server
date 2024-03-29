package main

import (
	"archive/zip"
	"bytes"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const AssetUrl = "https://github.com/ZNotify/frontend/releases/download/bundle/build.zip"

func exist(filename string) bool {
	stat, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return !stat.IsDir()
}

func download() {
	isForce := flag.Bool("force", false, "force download")
	flag.Parse()

	if exist("app/api/web/static/index.html") && !*isForce {
		log.Println("web/static/index.html exists, skip download")
		log.Println("if you still want to download, please use -f")
		return
	}

	resp, err := http.Get(AssetUrl)
	if err != nil {
		panic(err)
	}

	log.Println("Downloading frontend assets...")

	buff := bytes.NewBuffer([]byte{})
	_, err = io.Copy(buff, resp.Body)
	if err != nil {
		panic(err)
	}

	func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	log.Println("Downloaded frontend assets, extracting...")

	reader := bytes.NewReader(buff.Bytes())

	zipFiles, err := zip.NewReader(reader, int64(buff.Len()))
	if err != nil {
		panic(err)
	}

	for _, file := range zipFiles.File {
		func(file *zip.File) {
			rc, err := file.Open()
			if err != nil {
				panic(err)
			}
			defer func(rc io.ReadCloser) {
				err := rc.Close()
				if err != nil {
					panic(err)
				}
			}(rc)

			// remove "build/" from file path
			target := filepath.Join("app/api/web/static", file.Name[6:])

			if file.FileInfo().IsDir() {
				err := os.MkdirAll(target, file.Mode())
				if err != nil {
					panic(err)
				}
				return
			} else {
				targetFile, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
				if err != nil {
					panic(err)
				}
				defer func(targetFile *os.File) {
					err := targetFile.Close()
					if err != nil {
						panic(err)
					}
				}(targetFile)

				_, err = io.Copy(targetFile, rc)
				if err != nil {
					panic(err)
				}

				log.Printf("Extracted %s", target)
			}
		}(file)
	}

	log.Println("Extracted frontend assets.")
}
