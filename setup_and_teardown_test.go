// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

var testdir string

func setup(t *testing.T) (hct *HgClient) {
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

	hct = NewHgClient()
	cfg := make([]string, 0)
	err = hct.Connect(hgexe, repo, cfg)
	if err != nil {
		t.Fatal(err)
	}
	return hct
}

func teardown(t *testing.T, hct *HgClient) {
	err := hct.Disconnect()
	if err != nil {
		t.Errorf("from Disconnect(): %s", err.Error())
	}
	err = os.RemoveAll(testdir)
	if err != nil {
		t.Errorf("teardown(): %s", err.Error())
	}
}

func createFile(file string, data string, hct *HgClient) error {
	f, err := os.Create(hct.RepoRoot() + "/" + file)
	if err != nil {
		return err
	}
	_, err = f.Write([]byte(data))
	f.Sync()
	f.Close()
	return err
}

func createAndCommitFile(t *testing.T, hct *HgClient) error {

	err := createFile("/a", "aaa", hct)
	if err != nil {
		t.Fatal(err)
	}
	// add all there is to add to the repo,
	_, err = hct.Add(nil)
	// commit it
	var cmd *exec.Cmd
	cmd = exec.Command(hct.HgExe(), "--cwd", hct.RepoRoot(), "commit", "-Am\"first commit\"")
	if err = cmd.Run(); err != nil {
		t.Fatal(err)
	}

	return nil
}
