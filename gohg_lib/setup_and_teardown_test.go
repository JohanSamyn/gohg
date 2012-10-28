// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package gohg_lib_test

import (
	. "gohg/gohg_lib"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

var testdir string

func setup(t *testing.T) (hct *HgClient) {
	var err error
	testdir, err = ioutil.TempDir("", "gohg_test_")
	if err != nil {
		t.Fatal(err)
	}

	// Set var hgexe to whatever is appropriate for your situation.
	// You can also change it to test with different versions of Mercurial.
	// hgexe := "hg"
	hgexe := "M:/DEV/hg-stable/hg"

	var cmd *exec.Cmd
	cmd = exec.Command(hgexe, "--cwd", testdir, "init")
	if err = cmd.Run(); err != nil {
		t.Fatal(err)
	}

	repo := testdir

	hct = NewHgClient()
	cfg := make([]string, 0)
	err = hct.Connect(hgexe, repo, cfg)
	if err != nil {
		t.Fatal(err)
	}
	return hct
}

func teardown(t *testing.T, hct *HgClient) {
	err := hct.Close()
	if err != nil {
		t.Error("from Close(): " + string(err.Error()))
	}
	err = os.RemoveAll(testdir)
	if err != nil {
		t.Error("teardown(): " + string(err.Error()))
	}
}
