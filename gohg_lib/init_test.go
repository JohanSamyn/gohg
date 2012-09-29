// Copyright (C) 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package gohg_lib_test

import (
	. "gohg/gohg_lib"
	"errors"
	"os"
	"testing"
)

var tempdir = os.TempDir()
var pathSuccess = "\\gohg-init-success\\"
var pathFailure = "\\gohg-init-failure\\"

// I have a feeling all this is too much testing Mercurial
// instead of testing the HgClient.Init() method.
//
// I should test if it returns the right error when I pass it an empty path
// or "." as the path, or if the path is the same as the repo of the Hg CS.
// Or if RunCommand() produced an error.
// Or if the "e" channel returned the right error message.

func TestHgClient_Init_New_Should_Succeed(t *testing.T) {
	defer cleanupSuccess(t)
	path := tempdir + pathSuccess
	err := os.RemoveAll(path)
	if err != nil {
		t.Fatal(err)
	}

	err = HgClient.Init(path)
	if err != nil {
		t.Error(err)
	}
}

func TestHgClient_Init_Existing_Should_Fail(t *testing.T) {
	defer cleanupFailure(t)

	path := tempdir + pathFailure
	err := os.RemoveAll(path)
	if err != nil {
		t.Fatal(err)
	}
	// finding a subfolder '.hg' is enough for Mercurial to think
	// there is already a repo there.
	err = os.MkdirAll(path+".hg", 0777)
	if err != nil {
		t.Fatal(err)
	}

	err = HgClient.Init(path)
	if err == nil {
		t.Error(errors.New("HgClient.Init() did not fail in an existing Hg working copy"))
	}
}

func cleanupSuccess(t *testing.T) {
	err := os.RemoveAll(tempdir + pathSuccess)
	if err != nil {
		t.Log(err)
	}
}

func cleanupFailure(t *testing.T) {
	err := os.RemoveAll(tempdir + pathFailure)
	if err != nil {
		t.Log(err)
	}
}
