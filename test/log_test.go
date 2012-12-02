// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg_test

import (
	"os"
	"os/exec"
	"testing"
)

func TestHgClient_Log_NewRepo(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	// log should be empty for newly created repo
	data, err := hct.Log(nil)
	if err != nil {
		t.Error(err)
	}
	if data != nil {
		t.Fatal("Empty repo should have empty log")
	}
}

func TestHgClient_Log_Empty(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	// log should be empty for newly created repo
	data, err := hct.Log([]string{"-r", "tip"})
	if err != nil {
		t.Error(err)
	}
	if data != nil {
		t.Fatal("Empty repo should have empty log")
	}
}

func TestHgClient_Log_NotEmpty(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	// log should produce info for non-empty repo

	// have to make the working dir dirty !
	f, err := os.Create(hct.RepoRoot() + "/a")
	_, _ = f.Write([]byte{'a', 'a', 'a'})
	f.Sync()
	f.Close()
	// add all there is to add to the repo and commit
	var cmd *exec.Cmd
	cmd = exec.Command(hct.HgExe(), "--cwd", testdir, "commit", "-Am\"test commit\"")
	if err = cmd.Run(); err != nil {
		t.Fatal(err)
	}

	// now we can perform the real test
	data, err := hct.Log([]string{"-r", "tip"})
	if err != nil {
		t.Error(err)
	}
	if data == nil {
		t.Fatal("Non-empty repo should non-empty log")
	}
}
