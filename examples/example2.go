// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

// This program is an example of how you can use the gohg library.

package main

import (
	. "bitbucket.org/gohg/gohg"
	"fmt"
	"log"
)

func main() {
	// Set var hgexe to whatever is appropriate for your situation.
	// You can also change it to test with different versions of Mercurial.
	hgexe := "hg"
	repo := "."
	var cfg []string

	hgcl := NewHgClient()
	if err := hgcl.Connect(hgexe, repo, cfg); err != nil {
		log.Fatal(err)
	}
	defer hgcl.Disconnect()

	hc := NewLogCmd(nil, nil)
	o := make([]Option, 2)
	var lim Limit = 2
	o[0] = lim
	var verb Verbose = true
	o[1] = verb
	hc.SetOptions(o)
	// hc.SetParams([]string{"\"my param\""})
	cl, _ := hc.CmdLine(hgcl)
	fmt.Printf("%s\n", cl)
	res, _ := hc.Exec(hgcl)
	fmt.Printf("result:\n%s", string(res))
}
