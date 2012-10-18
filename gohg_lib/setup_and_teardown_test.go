// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package gohg_lib_test

import (
	. "gohg/gohg_lib"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"testing"
)

var Tempdir string

func setup(t *testing.T) (hct *HgClient) {
	var err error
	Tempdir, err = ioutil.TempDir("", "gohg_test_")
	if err != nil {
		log.Fatal(err.Error())
	}

	// set this var to whatever suits you (default: "hg")
	hgexe := "M:\\DEV\\hg-stable\\hg"

	var cmd *exec.Cmd
	cmd = exec.Command(hgexe, "--cwd", Tempdir, "init")
	if err = cmd.Run(); err != nil {
		log.Fatal(err.Error())
	}
	var repo = Tempdir

	hct = NewHgClient()
	cfg := make([]string, 0)
	err = hct.Connect(hgexe, repo, cfg)
	if err != nil {
		log.Fatal(err)
	}
	return hct
}

func teardown(t *testing.T, hct *HgClient) {
	err := hct.Close()
	if err != nil {
		t.Error("from Close(): " + string(err.Error()))
	}
	err = os.RemoveAll(Tempdir)
	if err != nil {
		t.Error("teardown(): " + string(err.Error()))
	}
}
