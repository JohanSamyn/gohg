// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

import (
	"testing"
)

func TestHgClient_Update(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	var expected string = "0 files updated, 0 files merged, 0 files removed, 0 files unresolved\n"

	got, err := hct.Update(nil, nil)
	if err != nil {
		t.Error(err)
	}

	if string(got) != expected {
		t.Fatalf("Test Update: expected:\n%s\n but got:\n%s\n", expected, got)
	}
}
