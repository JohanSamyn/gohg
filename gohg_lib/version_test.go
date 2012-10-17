// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package gohg_lib_test

import (
	"testing"
)

func TestHgClient_Version_Minimal(t *testing.T) {
	ver, err := hct.Version()
	if err != nil {
		t.Fatal(err)
	}
	minimalversion := "1.9"
	if ver < minimalversion {
		t.Error("HgClient.Version(): expected value >= " + minimalversion + " but got " + ver)
	}
}

func TestHgClient_Version_AsConnected(t *testing.T) {
	ver, err := hct.Version()
	if err != nil {
		t.Fatal(err)
	}
	if ver != hct.HgVersion {
		t.Error("HgClient.Version(): expected value " + hct.HgVersion + " but got " + ver)
	}
}
