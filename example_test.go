// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

// This program is an example of how to use the gohg library.

package gohg

import (
	"log"
)

func ExampleHgClient_Connect() {
	hc := NewHgClient()
	if err := hc.Connect("hg", ".", []string{""}, false); err != nil {
		log.Fatal(err)
	}
	defer hc.Disconnect()
	// Call some useful methods on hc here.
}
