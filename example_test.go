// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

// This program is an example of how to use the gohg library.

package gohg_test

import (
	"bitbucket.org/gohg/gohg"
	"log"
)

func ExampleHgClient_Connect() {
	hc := gohg.NewHgClient()
	if err := hc.Connect("hg", ".", []string{""}); err != nil {
		log.Fatal(err)
	}
	defer hc.Close()
	// Call some useful methods on hc here.
}
