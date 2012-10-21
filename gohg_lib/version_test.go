// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package gohg_lib_test

import (
	"testing"
)

func TestHgClient_Version_Minimal(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	ver, err := hct.Version([]string{""})
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

	_, err := hct.Version([]string{"-q"})
	if err != nil {
		t.Error("HgClient.Version() failed with -q flag")
	}
}

func TestHgClient_Version_Verbose(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	_, err := hct.Version([]string{"-v"})
	if err != nil {
		t.Error("HgClient.Version() failed with -v flag")
	}
}

func TestHgClient_Version_WrongFlag_Should_Fail(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	_, err := hct.Version([]string{"-wrongflag"})
	if err == nil {
		t.Error("HgClient.Version() did not fail with wrong flag")
	}
}

func TestHgClient_Version_HelpFlag_Should_Fail(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	_, err := hct.Version([]string{"-h"})
	if err == nil {
		t.Error("HgClient.Version() did not fail with -h flag")
	}
}

func TestHgClient_Version_AsConnected(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	ver, err := hct.Version([]string{""})
	if err != nil {
		t.Fatal(err)
	}
	if ver != hct.GetHgVersion() {
		t.Error("HgClient.Version(): expected value " + hct.GetHgVersion() + " but got " + ver)
	}
}
