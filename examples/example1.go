// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

// This program is an example of how you can use the gohg library.

package main

import (
	hg "bitbucket.org/gohg/gohg"
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

	hc := hg.NewHgClient()
	var cfg []string
	if err = hc.Connect(hgexe, repo, cfg); err != nil {
		log.Fatal(err)
	}
	defer hc.Disconnect()
	defer func() { fmt.Println("========== End of example1 ==========") }()

	var v string
	if v, err = hc.Version(); err != nil {
		log.Println(err)
	}
	fmt.Printf("[[version]]: %s\n", v)

	fmt.Println("--------------------")

	var i []byte
	if i, err = hc.Identify(nil /*[]hg.Option{hg.Verbose(true)}*/, []string{""}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("[[identify]]:\n%s", i)

	// fmt.Printf("%s", hg.SprintfOpts(*(new(hg.IdentifyOpts))))

	fmt.Println("--------------------")

	var s []byte
	if s, err = hc.Summary(nil, nil); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("[[summary]]:\n%s", s)

	// fmt.Printf("%s", SprintfOpts(*(new(SummaryOpts))))

	fmt.Println("--------------------")

	var limit hg.Limit = 2
	var verbose hg.Verbose = true
	opts := []hg.Option{limit, verbose}
	params := []string{}
	var l []byte
	// if l, err = hc.Log2([]hg.Option{hg.Limit(2), hg.Verbose(true)}, []hg.Param{}); err != nil {
	if l, err = hc.Log(opts, params); err != nil {
		fmt.Println(err)
		return
	}
	// use 'go run example1.go | less' (or more) to view big results (such as a full log)
	fmt.Printf("[[log -v -l 2 branches_test.go]]:\n%s", l)

	// give time to see the Hg CS session live and die from Process Explorer
	// fmt.Print("waiting...")
	// time.Sleep(3 * time.Second)
}
