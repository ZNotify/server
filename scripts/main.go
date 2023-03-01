package main

import (
	"fmt"
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
	default:
		fmt.Println("Unknown command " + os.Args[1])
	}
}
