// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

// This program is an example of how to use the gohg library.

package main

import (
	"bitbucket.org/gohg/gohg"
	"fmt"
	"log"
)

func main() {
	// Set var hgexe to whatever is appropriate for your situation.
	// You can also change it to test with different versions of Mercurial.
	hgexe := "hg"
	repo := "."

	fmt.Println("========== Begin of example1 ==========")

	var err error
	fmt.Printf("Using Mercurial repo at: %s\n", repo)
	fmt.Println("--------------------")

	hc := gohg.NewHgClient()
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

	var s []byte
	if s, err = hc.Summary([]string{}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("[[summary]]:\n%s", s)

	fmt.Println("--------------------")

	var l []byte
	if l, err = hc.Log([]string{"-l", "2"}); err != nil {
		fmt.Println(err)
		return
	}
	// use 'go run example1.go | less' (or more) to view big results (such as a full log)
	fmt.Printf("[[log -l 2]]:\n%s", l)

	// give time to see the Hg CS session live and die from Process Explorer
	// fmt.Print("waiting...")
	// time.Sleep(3 * time.Second)
}
