// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg_lib_test

import (
	"testing"
)

func TestHgClient_Log_EmptyRepo(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	// log should be empty for newly created repo
	data, err := hct.Log([]string{})
	if err != nil {
		t.Error(err)
	}
	if data != nil {
		t.Fatal("Empty repo should have empty log")
	}
}

func TestHgClient_Log_Tip(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	// log should be empty for newly created repo
	data, err := hct.Log([]string{"-r", "tip"})
	if err != nil {
		t.Error(err)
	}
	if data != nil {
		t.Fatal("Empty repo should have empty log")
	}
}
