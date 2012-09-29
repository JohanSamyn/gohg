// Copyright (C) 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package gohg_lib_test

import (
	. "gohg/gohg_lib"
	// "fmt"
	"testing"
)

// func TestConnect(*testing.T) {
// 	fmt.Printf("HgClient.Connected = %v\n", HgClient.Connected)
// 	var err error
// 	var repo = "."

// 	if HgClient.Connected {
// 		Close()
// 	}

// 	cfg := make([]string, 0)
// 	err = HgClient.Connect("M:\\DEV\\hg-stable\\hg", repo, cfg)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func TestGetEncoding(t *testing.T) {
	var err error
	var encoding string
	encoding, err = HgClient.GetEncoding()
	if err != nil {
		t.Error("from GetEncoding : " + string(err.Error()))
	}
	if encoding == "" {
		t.Error("GetEncoding did not return a valid encoding")
	}
}

func TestRunCommand(t *testing.T) {
	var err error
	// var data []byte
	// var ret int32
	// hgcmd := []string{"log", "-l", "2"}
	hgcmd := []string{"summary"}
	// data, ret, err = RunCommand(hgcmd)
	_, _, err = HgClient.RunCommand(hgcmd)
	if err != nil {
		t.Error("from RunCommand : " + string(err.Error()))
	}
}

// func TestClose(t *testing.T) {
// 	err := HgClient.Close()
// 	if err != nil {
// 		t.Error("from Close(): " + string(err.Error()))
// 	}
// }
