// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package gohg_lib_test

import (
	"os"
	"testing"
)

var pathSummary = "\\gohg-summary-success\\"

func TestHgClient_Summary(t *testing.T) {
	defer cleanupSummary(t)
	path := Tempdir + pathSummary
	err := os.RemoveAll(path)
	if err != nil {
		t.Fatal(err)
	}

	var data string
	var summary string = "parent: -1:000000000000 tip (empty repository)\n" +
		"branch: default\n" +
		"commit: (clean)\n" +
		"update: (current)\n"
	data, err = hct.Summary()
	if err != nil {
		t.Error(err)
	}
	if data != summary {
		t.Fatalf("Test Summary: expected:\n%s and got:\n%s", summary, data)
	}
}

func cleanupSummary(t *testing.T) {
	err := os.RemoveAll(Tempdir + pathSummary)
	if err != nil {
		t.Log(err)
	}
}
