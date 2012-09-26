// Copyright (C) 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package gohg_lib_test

import (
	. "gohg/gohg_lib"
	"log"
	"testing"
)

// TestSetup makes a connection to the Hg CS once, for all tests to use.

func TestSetup(*testing.T) {
	if HgClient.Connected != true {
		var err error
		var repo = "."
		cfg := make([]string, 0)
		err = Connect("M:\\DEV\\hg-stable\\hg", repo, cfg)
		if err != nil {
			log.Fatal(err)
		}
	}
}
