// Copyright (C) 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package gohg_lib_test

import (
	. "gohg/gohg_lib"
	"testing"
)

func TestHgClient_Version_Minimal(t *testing.T) {
	ver, _, err := HgClient.Version()
	if err != nil {
		t.Fatal(err)
	}
	minimalversion := "1.9"
	if ver < minimalversion {
		t.Error("HgClient.Version(): expected value >= " + minimalversion + " but got " + ver)
	}
}

func TestHgClient_Version_AsConnected(t *testing.T) {
	ver, _, err := HgClient.Version()
	if err != nil {
		t.Fatal(err)
	}
	if ver != HgClient.HgVersion {
		t.Error("HgClient.Version(): expected value " + HgClient.HgVersion + " but got " + ver)
	}
}
