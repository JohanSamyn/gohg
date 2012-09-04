package main

import (
	"gohg"
	"fmt"
	"log"
	"os"
)

func main() {
	var repo = "."
	if len(os.Args) > 1 {
		repo = os.Args[1]
	}
	a := make([]string, 1)
	err := gohg.Connect("M:\\DEV\\hg-stable\\hg", repo, a)
	if err != nil {
		log.Fatal(err)
		// fmt.Println("from Connect():", err)
		// err = gohg.Close()
		// if err != nil {
		// 	fmt.Println("from Close():", err)
		// }
		// os.Exit(1)
	}
	// err = gohg.Connect("M:\\DEV\\hg-stable\\hg", repo, a)
	// if err != nil {
	// 	fmt.Println("from Connect():", err)
	// }

	err = gohg.Close()
	if err != nil {
		fmt.Println("from Close():", err)
		os.Exit(1)
	}
	os.Exit(0)
}
