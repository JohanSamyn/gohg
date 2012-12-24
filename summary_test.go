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
	scmd := NewSummary()
	// scmd := NewSummary(0)

	// scmd = scmd.SetRepo(hct.RepoRoot())
	// scmd = scmd.SetRepo("C:/DEV/go/src/bitbucket.org/gohg/gohg")
	// scmd = scmd.SetRemote(true)

	// scmd, err := NewSummary(int(0))

	// scmd, err := NewSummary(O_remote(true), O_mq(true), O_test(true))

	got, err := hct.Summary(scmd)
	// got, err := hct.Summary(NewSummary())
	if err != nil {
		t.Error(err)
	}
	if string(got) != expected {
		t.Fatalf("Test Summary: expected:\n%s\n but got:\n%s\n", expected, got)
	}
}
