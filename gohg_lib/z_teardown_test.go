// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package gohg_lib_test

import (
	. "gohg/gohg_lib"
	"os"
	"testing"
)

// TestTearDown closes the connection to the Hg CS after all tests.

func TestTearDown(t *testing.T) {
	err := Hgclient.Close()
	if err != nil {
		t.Error("from Close(): " + string(err.Error()))
	}
	err = os.RemoveAll(Tempdir)
	if err != nil {
		t.Error("TearDown(): " + string(err.Error()))
	}
}
