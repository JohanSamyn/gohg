// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
	"log"
	"testing"
)

func TestClient_runcommand_IsConnected(t *testing.T) {
	hct := NewHgClient()
	// we deliberately do not call Connect()
	var err error
	_, err = hct.Identify("")
	if err == nil {
		fmt.Printf("error: %s\n", err)
		log.Fatal(fmt.Errorf("%s", "Did not detect disconnect!?"))
	}
}

// func TestConnect(*testing.T) {
// 	fmt.Printf("Hgclient.Connected = %v\n", Hgclient.Connected)
// 	var err error
// 	var repo = "."

// 	if Hgclient.Connected {
// 		Disconnect()
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

// func TestDisconnect(t *testing.T) {
// 	err := Hgclient.Disconnect()
// 	if err != nil {
// 		t.Errorf("from Disconnect(): %s", string(err.Error()))
// 	}
// }
