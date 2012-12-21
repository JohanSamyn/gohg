// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"os"
	"os/exec"
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
	got, err := hct.Verify(nil)
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

	// have to make the working dir dirty
	f, err := os.Create(hct.RepoRoot() + "/a")
	_, _ = f.Write([]byte{'a', 'a', 'a'})
	f.Sync()
	f.Close()
	// add all there is to add to the repo,
	_, err = hct.Add(nil)
	// commit it
	var cmd *exec.Cmd
	cmd = exec.Command(hct.HgExe(), "--cwd", hct.RepoRoot(), "commit", "-Am\"first commit\"")
	if err = cmd.Run(); err != nil {
		t.Fatal(err)
	}
	// cause some integrity problem
	err = os.Rename(hct.RepoRoot()+"\\.hg\\store\\00manifest.i",
		hct.RepoRoot()+"\\.hg\\store\\00manifest.i_BAK")
	if err != nil {
		t.Fatal(err)
	}

	// now we can perform the real test

	_, err = hct.Verify(nil)
	if err == nil {
		t.Fatalf("Test Verify: did not get expected error")
	}
}
