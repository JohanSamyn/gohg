// Copyright (C) 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package gohg_lib_test

import (
	. "gohg/gohg_lib"
	"io/ioutil"
	"log"
	"os/exec"
	"testing"

	"fmt"
)

var Tempdir string

// var Hgclient hgclient

// TestSetup makes a connection to the Hg CS once, for all tests to use.
func TestSetup(t *testing.T) {

	var err error
	Tempdir, err = ioutil.TempDir("", "gohg_test_")
	fmt.Println(Tempdir)

	// now create an empty Hg repo inthere
	var cmd *exec.Cmd
	cmd = exec.Command("M:\\DEV\\hg-stable\\hg", "--cwd", Tempdir, "init")
	if err = cmd.Run(); err != nil {
		log.Fatal(err)
	}

	// Hgclient = newHgClient()
	if Hgclient.Connected != true {
		var repo = Tempdir
		cfg := make([]string, 0)
		err = Hgclient.Connect("M:\\DEV\\hg-stable\\hg", repo, cfg)
		if err != nil {
			log.Fatal(err)
		}
	}
}
