// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
// "fmt"
// "testing"
)

// func TestConnect(*testing.T) {
// 	fmt.Printf("Hgclient.Connected = %v\n", Hgclient.Connected)
// 	var err error
// 	var repo = "."

// 	if Hgclient.Connected {
// 		Close()
// 	}

// 	cfg := make([]string, 0)
// 	err = Hgclient.Connect("M:/DEV/hg-stable/hg" /*"hg"*/, repo, cfg)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func TestRun(t *testing.T) {
// 	hct := setup(t)
// 	defer teardown(t, hct)

// 	var err error
// 	// var data []byte
// 	// var ret int32
// 	// hgcmd := []string{"log", "-l", "2"}
// 	hgcmd := []string{"summary"}
// 	// data, ret, err = run(hgcmd)
// 	_, _, err = hct.run(hgcmd)
// 	if err != nil {
// 		t.Errorf("from run : %s", string(err.Error()))
// 	}
// }

// func TestClose(t *testing.T) {
// 	err := Hgclient.Close()
// 	if err != nil {
// 		t.Errorf("from Close(): %s", string(err.Error()))
// 	}
// }