// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"testing"
)

func TestHgClient_Version_Minimal(t *testing.T) {

	// This a bit silly, as the setup function will already have failed
	// if a Hg version prior to 1.9 is used.

	hct := setup(t)
	defer teardown(t, hct)

	ver, err := hct.Version()
	if err != nil {
		t.Fatal(err)
	}
	minimalversion := "1.9"
	if ver < minimalversion {
		t.Error("HgClient.Version(): expected value >= " + minimalversion + " but got " + ver)
	}
}

func TestHgClient_Version_Short(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	_, err := hct.Version()
	if err != nil {
		t.Error("HgClient.Version() failed")
	}
}

func TestHgClient_Version_AsConnected(t *testing.T) {

	// This is also a bit silly, as both version strings are obtained
	// with the same code.

	hct := setup(t)
	defer teardown(t, hct)

	ver, err := hct.Version()
	if err != nil {
		t.Fatal(err)
	}
	if ver != hct.hgVersion() {
		t.Error("HgClient.Version(): expected value " + hct.hgVersion() + " but got " + ver)
	}
}
