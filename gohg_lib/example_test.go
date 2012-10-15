// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

// This program is an example of how to use the gohg library.

package gohg_lib_test

import (
	. "gohg/gohg_lib"
	"log"
)

func ExampleHgClient_Connect() {
	hc := NewHgClient()
	if err := hc.Connect("hg", ".", []string{""}); err != nil {
		log.Fatal(err)
	}
	defer hc.Close()
}
