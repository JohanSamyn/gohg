// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

// comment the one above, and use the one below to run this program
// package main

import (
	"bitbucket.org/gohg/gohg"
	"fmt"
	"log"
)

func main() {
	var err error
	hgexe := "M:/DEV/hg-stable/hg"
	repo := "C:/DEV/thg-qt-work/"
	var cfg []string
	hc := gohg.NewHgClient()
	if err = hc.Connect(hgexe, repo, cfg); err != nil {
		log.Fatal(err)
	}
	defer hc.Close()

	var s string
	if s, err = hc.Summary([]string{}); err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("[[Summary for repo %s]]:\n%s", repo, s)
}
