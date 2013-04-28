// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

import (
	"fmt"
	"testing"
)

func TestHgClient_Clone_To_New_Should_Succeed(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	// the source repo was created in setup(), and then used as the repo for
	// connecting the Hg CS

	dest, err := createTempdir(t)
	if err != nil {
		t.Error(err)
	}
	defer destroyTempdir(dest)

	err1 := hct.Clone(hct.RepoRoot(), dest)
	if err1 != nil {
		t.Fatalf("Test Clone failed: %s\n but got:\n%s\n", err1)
	}

	hc2 := NewHgClient()
	var cfg []string
	err2 := hc2.Connect(hct.HgExe(), dest, cfg)
	_ = hc2.Disconnect()
	if err1 != nil || err2 != nil {
		t.Fatalf("Test Clone failed: %s\n but got:\n%s\n", err2)
	}
}

func TestHgClient_Clone_To_Existing_Should_Fail(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	// the source repo was created in setup(), and then used as the repo for
	// connecting the Hg CS

	expected := "runcommand: Clone(): returncode=-1\n" +
		"cmd: clone " + hct.RepoRoot() + " " + hct.RepoRoot() + "\n" +
		"hgerr:\n" +
		"abort: destination '" + hct.RepoRoot() + "' is not empty\n\n"

	dest := hct.RepoRoot()

	g := hct.Clone(hct.RepoRoot(), dest)
	got := fmt.Sprint(g)
	if got != expected {
		t.Fatalf("Test Clone: expected:\n%s\n but got:\n%s\n", expected, got)
	}
}
