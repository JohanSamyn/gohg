// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

// This program is an example of how to use the gohg library.

package main

import (
	. "gohg/gohg_lib"
	"fmt"
	"log"
)

func main() {
	fmt.Println("========== Begin of example1 ==========")

	var err error
	hgcmd := "M:\\DEV\\hg-stable\\hg"
	// hgcmd := "hg"
	repo := "C:\\DEV\\go\\src\\gohg"
	fmt.Printf("Using Mercurial repo at: %s\n", repo)
	fmt.Println("--------------------")

	hc := NewHgClient()
	var cfg []string
	if err = hc.Connect(hgcmd, repo, cfg); err != nil {
		log.Fatal(err)
	}
	defer hc.Close()
	defer func() { fmt.Println("========== End of example1 ==========") }()

	var v, fv string
	if v, fv, err = hc.Version(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("[[version]]: %s\n[[fullversion]]:\n%s", v, fv)

	fmt.Println("--------------------")

	var s string
	if s, err = hc.Summary(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("[[Summary]]:\n%s", s)
}