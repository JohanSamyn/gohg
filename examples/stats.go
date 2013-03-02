// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

// This program is an example of how to use the gohg library.

package main

import (
	. "bitbucket.org/gohg/gohg"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	hgexe := "M:/DEV/hg-stable/hg"
	repo := "M:/DEV/thg-qt"

	hc := NewHgClient()
	if err := hc.Connect(hgexe, repo, nil); err != nil {
		log.Fatal(err)
	}
	defer hc.Disconnect()

	t, err := hc.Tip(Template("{rev}"))
	if err != nil {
		fmt.Println(err)
		return
	}
	revs, err := strconv.Atoi(string(t))

	m, err := hc.Manifest()
	if err != nil {
		fmt.Println(err)
		return
	}
	files := len(strings.Split(string(m), "\n")) - 1 // don't count empty value after last \n

	h, err := hc.Heads(Template("{rev}\n"))
	if err != nil {
		fmt.Println(err)
		return
	}
	heads := len(strings.Split(string(h), "\n")) - 1 // don't count empty value after last \n

	fmt.Printf("some stats for repo %s:\n"+
		"%8d revisions\n"+
		"%8d files\n"+
		"%8d heads\n",
		hc.RepoRoot(), revs, files, heads)
}
