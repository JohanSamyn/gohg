// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

// This program is an example of how you can use the gohg library.
// Run it in the working dir of any Mercurial repo, or pass the
// path of the repo as the second argument to hc.Connect().
// It is my Go implementation of the same Python program that serves as an
// example for the python-hglib (see http://mercurial.selenic.com/wiki/PythonHglib).

package main

import (
	hg "bitbucket.org/gohg/gohg"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	hc := hg.NewHgClient()
	if err := hc.Connect("", "", nil, false); err != nil {
		log.Fatal(err)
	}
	defer hc.Disconnect()

	t, err := hc.Tip([]hg.HgOption{hg.Template("{rev}")}, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	revs, err := strconv.Atoi(string(t))

	m, err := hc.Manifest(nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// don't count empty value after last \n
	files := len(strings.Split(string(m), "\n")) - 1

	h, err := hc.Heads([]hg.HgOption{hg.Template("{rev}\n")}, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// don't count empty value after last \n
	heads := len(strings.Split(string(h), "\n")) - 1

	b, err := hc.Branches([]hg.HgOption{hg.Quiet(true)}, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// don't count empty value after last \n
	branches := len(strings.Split(string(b), "\n")) - 1

	tg, err := hc.Tags([]hg.HgOption{hg.Quiet(true)}, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// don't count empty value after last \n
	// don't count tip
	tags := len(strings.Split(string(tg), "\n")) - 1 - 1

	au, err := hc.Log([]hg.HgOption{hg.Template("{author}\n")}, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	var am = make(map[string]bool)
	var sn string
	for _, n := range strings.Split(strings.TrimRight(string(au), "\n"), "\n") {
		sn = string(n)
		_, ok := am[sn]
		if !ok {
			am[sn] = true
		}
	}
	authors := len(am)

	mg, err := hc.Log([]hg.HgOption{hg.Rev("merge()"), hg.Template("{rev}\n")}, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// don't count empty value after last \n
	merges := len(strings.Split(string(mg), "\n")) - 1

	fmt.Printf("some stats for repo %s:\n"+
		"%8d revisions\n"+
		"%8d merges\n"+
		"%8d files\n"+
		"%8d heads\n"+
		"%8d branches\n"+
		"%8d tags\n"+
		"%8d authors\n",
		hc.RepoRoot(), revs, merges, files, heads, branches, tags, authors)
}
