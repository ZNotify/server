package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	switch os.Args[1] {
	case "ent":
		ent()
	case "download":
		download()
	case "schema":
		schema()
	default:
		log.Println("Unknown command " + os.Args[1])
	}
}
