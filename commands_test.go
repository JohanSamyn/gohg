// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"strings"
	"testing"
)

func TestCommand_Ok(t *testing.T) {
	// unnecessary in fact, as all the command tests going ok yields the same result

	hct := setup(t)
	defer teardown(t, hct)

	data, err := commandOld(hct, "identify", []string{})
	if err != nil || data == nil || string(data) == "" {
		t.Fatalf("Did not perform correctly.")
	}
}

func TestCommand_runError(t *testing.T) {
	hct := setup(t)
	// do not defer, causing an error in the Command call
	// and so catching err in Command()
	teardown(t, hct)

	_, err := commandOld(hct, "", []string{})
	if err == nil || strings.Contains(err.Error(), "from runInHg():") == false {
		t.Fatalf("Did not get runInHg() error.")
	}
}

func TestCommand_HgError(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	// catching hgerr in Command()
	_, err := commandOld(hct, "nonexistinghgcommand", []string{})
	if err == nil {
		t.Fatalf("Did not catch hgerr.")
	}
}
