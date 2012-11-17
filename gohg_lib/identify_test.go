// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg_lib_test

import (
	// "errors"
	// "fmt"
	"testing"
	// "os/exec"
)

func TestHgClient_Identify_EmptyRepo(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	var expected string = "000000000000 tip\n"
	got, err := hct.Identify(nil)
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Fatalf("Test Identify: expected:\n%s\n but got:\n%s\n", expected, got)
	}
}

// One test should be enough. If that one gets thru and returns a reasonable result,
// then the others are supposed to do that as well. After all, we're not testing the
// hg identify command as such here; only if we can get data from the Hg CS using that command.

// func TestHgClient_Identify_OneCset(t *testing.T) {
// 	hct := setup(t)
// 	// defer teardown(t, hct)

// 	var err error
// 	var cmd *exec.Cmd
// 	cmd = exec.Command("echo aaaa > " + hct.RepoRoot() + "\a")
// 	if err = cmd.Run(); err != nil {
// 		t.Fatal(err)
// 	}
// 	cmd = exec.Command(hct.HgExe(), "--cwd", hct.RepoRoot(), "ci", "-Am\"first commit\"")
// 	if err = cmd.Run(); err != nil {
// 		t.Fatal(err)
// 	}

// 	// use a regex to match the (any) hash
// 	// var expected string = "000000000000 tip\n"
// 	// got, err := hct.Identify(nil)
// 	// if err != nil {
// 	// 	t.Error(err)
// 	// }
// 	// if got != expected {
// 	// 	t.Fatalf("Test Identify: expected:\n%s but got:\n%s\n", expected, got)
// 	// }
// }

// func TestHgClient_Identify_Dirty(t *testing.T) {
// 	// should show a "+" at the end of the hash
// }

// func TestHgClient_Identify_TwoParents(t *testing.T) {
// 	// should show 2 hashes, and no "+"
// }

// func TestHgClient_Identify_TwoParents_Dirty(t *testing.T) {
// 	// should show 2 hashes, and a "+"
// }
