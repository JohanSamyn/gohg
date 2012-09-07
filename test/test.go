// Copyright (C) 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package main

import (
	// "bytes"
	"fmt"
	"gohg"
	"log"
	"os"
	// "time"
)

func main() {
	var repo = "."
	if len(os.Args) > 1 {
		repo = os.Args[1]
	}
	cfg := make([]string, 0)
	err := gohg.Connect("M:\\DEV\\hg-stable\\hg", repo, cfg)
	if err != nil {
		log.Fatal(err)
	}

	// do whatever you want to do via the Hg CS
	// hgcmd := []string{"log", "-l", "2"}
	hgcmd := []string{"branches"}
	gohg.RunCommand(hgcmd)

	// give time to see the Hg CS session live and die from Process Explorer
	// time.Sleep(1 * time.Second)

	err = gohg.Close()
	if err != nil {
		log.Fatal("from Close(): " + string(err.Error()))
	}

	os.Exit(0)
}
