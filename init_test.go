// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

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

	path, err := filepath.Abs(testdir + "/gohg-init-success/")
	if err != nil {
		t.Error(err)
	}
	err = os.RemoveAll(path)
	if err != nil {
		t.Fatal(err)
	}

	err = hct.Init(Destpath(path))
	if err != nil {
		t.Error(err)
	}
}

func TestHgClient_Init_Existing_Should_Fail(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	err := hct.Init(Destpath(hct.RepoRoot()))
	if err == nil {
		t.Error(errors.New("HgClient.Init() did not fail in an existing Hg working copy"))
	}
}

func TestHgClient_Init_New_Mq_Should_Succeed(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	path, err := filepath.Abs(testdir + "/gohg-init-mq-success/")
	if err != nil {
		t.Error(err)
	}
	err = os.RemoveAll(path)
	if err != nil {
		t.Fatal(err)
	}

	err = hct.Init(Destpath(path))
	if err != nil {
		t.Error(err)
	}

	err = hct.Init(Mq(true), Cwd(path))
	if err != nil {
		t.Error(err)
	}
}
