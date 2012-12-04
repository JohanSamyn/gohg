// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg_test

import (
	"bitbucket.org/gohg/gohg"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

var testdir string

func setup(t *testing.T) (hct *gohg.HgClient) {
	// Set var hgexe to whatever is appropriate for your situation.
	// You can also change it to test with different versions of Mercurial.
	hgexe := "hg"

	var err error
	testdir, err = ioutil.TempDir("", "gohg_test_")
	if err != nil {
		t.Fatal(err)
	}

	var cmd *exec.Cmd
	cmd = exec.Command(hgexe, "--cwd", testdir, "init")
	if err = cmd.Run(); err != nil {
		t.Fatal(err)
	}

	repo := testdir

	hct = gohg.NewHgClient()
	cfg := make([]string, 0)
	err = hct.Connect(hgexe, repo, cfg)
	if err != nil {
		t.Fatal(err)
	}
	return hct
}

func teardown(t *testing.T, hct *gohg.HgClient) {
	err := hct.Close()
	if err != nil {
		t.Errorf("from Close(): %s", string(err.Error()))
	}
	err = os.RemoveAll(testdir)
	if err != nil {
		t.Errorf("teardown(): %s", string(err.Error()))
	}
}
