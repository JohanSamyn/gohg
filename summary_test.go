// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"testing"
)

// TODO	Add tests for the --remote and --mq options, and for setting the repo explicitely.

func TestHgClient_Summary_NoOptions(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	var expected string = "parent: -1:000000000000 tip (empty repository)\n" +
		"branch: default\n" +
		"commit: (clean)\n" +
		"update: (current)\n"

	// example 1
	got, err := hct.Summary()

	// example 2
	// got, err := hct.Summary(O_profile(true))

	// example 3
	// p := O_profile(true)
	// got, err := hct.Summary(p)

	// example 4
	// var p O_profile
	// p = true
	// got, err := hct.Summary(p)

	if err != nil {
		t.Error(err)
	}

	if string(got) != expected {
		t.Fatalf("Test Summary: expected:\n%s\n but got:\n%s\n", expected, got)
	}
}
