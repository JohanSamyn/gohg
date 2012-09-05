package main

import (
	"fmt"
	"gohg"
	"log"
	"os"
)

func main() {
	var repo = "."
	if len(os.Args) > 1 {
		repo = os.Args[1]
	}
	cfg := make([]string, 0)
	err := gohg.Connect("M:\\DEV\\hg-stable\\hg", repo, cfg)
	if err != nil {
		log.Fatal(err)
	}

	// do whatever you want to do via the Hg CS

	err = gohg.Close()
	if err != nil {
		fmt.Println("from Close():", err)
		os.Exit(1)
	}
	os.Exit(0)
}
