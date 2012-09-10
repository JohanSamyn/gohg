// Copyright (C) 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package main

import (
	"fmt"
	"gohg"
	"log"
	"os"

// "time"
)

func main() {
	var err error
	var repo = "."

	if len(os.Args) > 1 {
		repo = os.Args[1]
	}
	cfg := make([]string, 0)
	err = gohg.Connect("M:\\DEV\\hg-stable\\hg", repo, cfg)
	if err != nil {
		log.Fatal(err)
	}

	// do whatever you want to do via the Hg CS
	var encoding string
	encoding, err = gohg.GetEncoding()
	if err != nil {
		log.Fatal("from GetEncoding : " + string(err.Error()))
	}
	fmt.Printf("--------------------\ncommand -> getencoding\ndata -> %s\n", encoding)

	var data []byte
	var ret int32
	// hgcmd := []string{"log", "-l", "2"}
	hgcmd := []string{"summary"}
	data, ret, err = gohg.RunCommand(hgcmd)
	if err != nil {
		log.Fatal("from RunCommand : " + string(err.Error()))
	}
	fmt.Printf("--------------------\ncommand -> runcommand\nhgcmd -> %s\ndata ->\n%s\nreturncode -> %d\n",
		hgcmd, data, ret)

	// give time to see the Hg CS session live and die from Process Explorer
	// fmt.Print("waiting...")
	// time.Sleep(3 * time.Second)

	err = gohg.Close()
	if err != nil {
		log.Fatal("from Close(): " + string(err.Error()))
	}

	os.Exit(0)
}
