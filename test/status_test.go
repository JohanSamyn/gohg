// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg_lib_test

import (
	"os"
	"testing"
)

func TestHgClient_Status_Clean(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	// status should be empty for clean working dir
	expected := []byte{}
	got, err := hct.Status([]string{"-mardcui"})
	if err != nil {
		t.Error(err)
	}
	if string(got) != string(expected) {
		t.Fatalf("Test Status (clean): expected:\n%s\n but got:\n%s\n", expected, got)
	}
}

func TestHgClient_Status_Dirty(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	// status should not be empty for dirty working dir

	// have to make the working dir dirty !
	f, err := os.Create(hct.RepoRoot() + "/a")
	_, _ = f.Write([]byte{'a', 'a', 'a'})
	f.Sync()
	f.Close()
	// add all there is to add to the repo
	_, err = hct.Add(nil)

	// now we can perform the real test
	expected := []byte{}
	got, err := hct.Status([]string{"-mardcui"})
	if err != nil {
		t.Error(err)
	}
	if string(got) == string(expected) {
		t.Fatalf("Test Status (dirty): expected:\n%s\n but got:\n%s\n", expected, got)
	}
}
