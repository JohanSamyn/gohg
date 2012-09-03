package main

import (
	"gohg"
	"fmt"
	"os"
	// "time"
)

func main() {
	a := make([]string, 1)
	err := gohg.Connect("M:\\DEV\\hg-stable\\hg", "C:\\DEV\\go\\src\\golout", a)
	if err != nil {
		fmt.Print(err)
		err = gohg.Close()
		if err != nil {
			fmt.Print("from Close():", err)
		}
		os.Exit(1)
	}

	// // keeping the gohg lib 'open' for a while
	// fmt.Println("sleeping 5 seconds ...")
	// time.Sleep(5 * time.Second)

	// fmt.Println("before Close()")
	err = gohg.Close()
	if err != nil {
		fmt.Print("from Close():", err)
		os.Exit(1)
	}
	// fmt.Println("after Close()")
	os.Exit(0)
}
