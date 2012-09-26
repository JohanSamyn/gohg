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

func TestHgClient_Init_New_Should_Succeed(t *testing.T) {
	tempdir := os.TempDir()
	path := tempdir + "\\gohg-init-success\\"
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
	tempdir := os.TempDir()
	path := tempdir + "\\gohg-init-failure\\"
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
