// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"testing"
)

// TODO	Add tests for the --remote and --mq options, and for setting the repo explicitely.

func TestHgClient_Summary_NoOptions(t *testing.T) {

	// GIVEN a new repo
	hct := setup(t)
	defer teardown(t, hct)

	// AND the known result for 'hg summary' for a new repo
	var expected string = "parent: -1:000000000000 tip (empty repository)\n" +
		"branch: default\n" +
		"commit: (clean)\n" +
		"update: (current)\n"

	// WHEN I call the 'hg summary' command on it
	// example 1
	got, err := hct.Summary()
	// got, err := hct.Summary(Limit(2))

	// example 2
	// got, err := hct.Summary(Profile(true))

	// example 3
	// p := Profile(true)
	// got, err := hct.Summary(p)

	// example 4
	// var p Profile
	// p = true
	// got, err := hct.Summary(p)

	if err != nil {
		t.Error(err)
	}

	// THEN the resulting info should be as expected
	if string(got) != expected {
		t.Fatalf("Test Summary: expected:\n%s\n but got:\n%s\n", expected, got)
	}
}
