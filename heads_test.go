// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

import (
	"os"
	"os/exec"
	"testing"
)

func TestHgClient_Heads(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	var expected string = "0\n"

	f, err := os.Create(hct.RepoRoot() + "/a")
	_, _ = f.Write([]byte{'a', 'a', 'a'})
	f.Sync()
	f.Close()

	cmd := exec.Command(hct.HgExe(), "--cwd", hct.RepoRoot(), "ci", "-Am\"first commit\"")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	got, err := hct.Heads([]HgOption{Template("{rev}\n")}, nil)
	if err != nil {
		t.Error(err)
	}

	if string(got) != expected {
		t.Fatalf("Test Heads: expected:\n%s\n but got:\n%s\n", expected, got)
	}
}
