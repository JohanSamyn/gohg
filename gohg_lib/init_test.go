// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package gohg_lib_test

import (
	"errors"
	"os"
	"testing"
)

var pathSuccess = "\\gohg-init-success\\"
var pathFailure = "\\gohg-init-failure\\"

// I have a feeling all this is too much testing Mercurial
// instead of testing the HgClient.Init() method.
//
// I should test if it returns the right error when I pass it an empty path
// or "." as the path, or if the path is the same as the repo of the Hg CS.
// Or if run() produced an error.
// Or if the "e" channel returned the right error message.

func TestHgClient_Init_New_Should_Succeed(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	defer cleanupInitSuccess(t)
	path := testdir + pathSuccess
	err := os.RemoveAll(path)
	if err != nil {
		t.Fatal(err)
	}

	err = hct.Init(path)
	if err != nil {
		t.Error(err)
	}
}

func TestHgClient_Init_Existing_Should_Fail(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	defer cleanupInitFailure(t)

	path := testdir + pathFailure
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

	err = hct.Init(path)
	if err == nil {
		t.Error(errors.New("HgClient.Init() did not fail in an existing Hg working copy"))
	}
}

func cleanupInitSuccess(t *testing.T) {
	err := os.RemoveAll(testdir + pathSuccess)
	if err != nil {
		t.Log(err)
	}
}

func cleanupInitFailure(t *testing.T) {
	err := os.RemoveAll(testdir + pathFailure)
	if err != nil {
		t.Log(err)
	}
}
