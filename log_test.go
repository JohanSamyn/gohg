// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

import (
	"testing"
)

func TestHgClient_Log_NewRepo(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	data, err := hct.Log([]HgOption{Limit(99)}, nil)
	if err != nil {
		t.Error(err)
	}
	if data != nil {
		t.Fatal("Newly created repo should have empty log")
	}
}

// func TestHgClient_Log_Empty(t *testing.T) {
// 	hct := setup(t)
// 	defer teardown(t, hct)

// 	data, err := hct.Log([]HgOption{Rev("tip:0")}, nil)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if data != nil {
// 		t.Fatal("Empty repo should have empty log")
// 	}
// }

func TestHgClient_Log_NotEmpty(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	err := createAndCommitFile(t, hct, "/a", "aaa")
	if err != nil {
		t.Fatal(err)
	}

	data, err := hct.Log([]HgOption{Rev("tip:0")}, nil)
	if err != nil {
		t.Error(err)
	}
	if data == nil {
		t.Fatal("Non-empty repo should have non-empty log")
	}
}
