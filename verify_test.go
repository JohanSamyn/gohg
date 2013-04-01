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

	// GIVEN a new repo
	hct := setup(t)
	defer teardown(t, hct)

	// AND the known result for 'hg verify' on a new repo
	var expected string = "checking changesets\n" +
		"checking manifests\n" +
		"crosschecking files in changesets and manifests\n" +
		"checking files\n" +
		"0 files, 0 changesets, 0 total revisions\n"

	// WHEN I call 'hg verify' on it
	got, err := hct.Verify()
	if err != nil {
		t.Error(err)
	}

	// THEN it should not report any problem
	if string(got) != expected {
		t.Fatalf("Test Verify: expected:\n%s\n but got:\n%s\n", expected, got)
	}
} // TestHgClient_Verify_Healthy

func TestHgClient_Verify_Sick(t *testing.T) {

	// GIVEN a new repo
	hct := setup(t)
	defer teardown(t, hct)

	// AND at least one cset in the repo
	err := createAndCommitFile(t, hct)
	if err != nil {
		t.Fatal(err)
	}

	// AND we cause a defect to the repo
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

	// WHEN we call 'hg verify' on it
	_, err = hct.Verify()

	// THEN it should return an error
	if err == nil {
		t.Fatalf("Test Verify: did not get expected error")
	}
} // TestHgClient_Verify_Sick
