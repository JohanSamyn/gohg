// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

// This program is the same as in the README.md file.
// It is meant for testing that code.

package main

import (
	. "bitbucket.org/gohg/gohg"
	"fmt"
	"log"
)

func main() {
	var err error
	hc := NewHgClient()
	if err = hc.Connect("", "", nil); err != nil {
		log.Fatal(err)
	}
	defer hc.Disconnect()

	var summ []byte
	if summ, err = hc.Summary(nil, nil); err != nil {
		log.Println(err)
	}
	fmt.Printf("\"summary\" for repo %s:\n%s\n", hc.RepoRoot(), summ)

	var l []byte
	files := []string{}
	if l, err = hc.Log([]Option{Limit(2)}, files); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\"log -l 2\" for repo %s:\n%s\n", hc.RepoRoot(), l)
}
