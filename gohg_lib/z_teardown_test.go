// Copyright (C) 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package gohg_lib_test

import (
	. "gohg/gohg_lib"
	"log"
	"testing"
)

// TestTearDown closes the connection to the Hg CS after all tests.

func TestTearDown(*testing.T) {
	err := Close()
	if err != nil {
		log.Fatal("from Close(): " + string(err.Error()))
	}
}
