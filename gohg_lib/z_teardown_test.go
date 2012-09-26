// Copyright (C) 2012, The gohg Authors. All rights reserved.
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
	err := Close()
	if err != nil {
		t.Error("from Close(): " + string(err.Error()))
	}

	// cleanup temporary folders/files
	tempdir := os.TempDir()
	path := tempdir + "\\gohg-init-success\\"
	err = os.RemoveAll(path)
	if err != nil {
		t.Log(err)
	}
	path = tempdir + "\\gohg-init-failure\\"
	err = os.RemoveAll(path)
	if err != nil {
		t.Log(err)
	}

}
