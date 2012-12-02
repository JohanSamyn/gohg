// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg_test

import (
	"testing"
)

func TestHgClient_Summary(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	var expected string = "parent: -1:000000000000 tip (empty repository)\n" +
		"branch: default\n" +
		"commit: (clean)\n" +
		"update: (current)\n"
	got, err := hct.Summary(nil)
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Fatalf("Test Summary: expected:\n%s\n but got:\n%s\n", expected, got)
	}
}
