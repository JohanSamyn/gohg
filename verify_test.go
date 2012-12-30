// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"os"
	"path/filepath"
	"testing"
)

func TestHgClient_Verify_Healthy(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	var expected string = "checking changesets\n" +
		"checking manifests\n" +
		"crosschecking files in changesets and manifests\n" +
		"checking files\n" +
		"0 files, 0 changesets, 0 total revisions\n"
	got, err := hct.Verify()
	if err != nil {
		t.Error(err)
	}
	if string(got) != expected {
		t.Fatalf("Test Verify: expected:\n%s\n but got:\n%s\n", expected, got)
	}
}

func TestHgClient_Verify_Sick(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	err := addAndCommitFile(t, hct)
	if err != nil {
		t.Fatal(err)
	}
	// cause some integrity problem
	var f string
	f, err = filepath.Abs(hct.RepoRoot() + "/.hg/store/00manifest.i")
	if err != nil {
		t.Error(err)
	}
	err = os.Rename(f, f+"_BAK")
	if err != nil {
		t.Fatal(err)
	}

	_, err = hct.Verify()
	if err == nil {
		t.Fatalf("Test Verify: did not get expected error")
	}
}
