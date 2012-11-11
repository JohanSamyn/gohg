// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

// This program is an example of how to use the gohg library.

package main

import (
	"bitbucket.org/gohg/gohg/gohg_lib"
	"fmt"
	"log"
)

var hgexe string
var repo string

func init() {
	// Read the gohg.ini file to determine hgexe and repo
	// or set them here manually.
	// hgexe = "M:/DEV/hg-stable/hg"
	// hgexe := "hg"
	// repo := "C:/DEV/go/src/bitbucket.org/gohg/gohg"
	// repo = "C:/DEV/thg-qt-work"
}

func main() {
	fmt.Println("========== Begin of example1 ==========")

	var err error
	fmt.Printf("Using Mercurial repo at: %s\n", repo)
	fmt.Println("--------------------")

	hc := gohg_lib.NewHgClient()
	var cfg []string
	if err = hc.Connect(hgexe, repo, cfg); err != nil {
		log.Fatal(err)
	}
	defer hc.Close()
	defer func() { fmt.Println("========== End of example1 ==========") }()

	var v string
	if v, err = hc.Version(); err != nil {
		// fmt.Println(err)
		// return
		log.Println(err)
	}
	fmt.Printf("[[version]]: %s\n", v)

	fmt.Println("--------------------")

	var s string
	if s, err = hc.Summary(); err != nil {
		fmt.Println(err)
		return
	}
	// use 'go run example1.go | less' (or more) to view big results (such as a full log)
	fmt.Printf("[[Summary]]:\n%s", s)

	// give time to see the Hg CS session live and die from Process Explorer
	// fmt.Print("waiting...")
	// time.Sleep(3 * time.Second)
}
