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

	cmd := make([]string, 4)
	cmd[0] = "log"
	cmd[1] = "--limit"
	cmd[2] = "2"
	cmd[3] = "-v"
	res, _ := hgcl.ExecCmd(cmd)
	fmt.Printf("%s\n", hgcl.ShowLastCmd())
	fmt.Printf("result:\n%s", string(res))
}
