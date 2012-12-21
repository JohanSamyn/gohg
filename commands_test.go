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

	data, err := Command(hct, "identify", []string{})
	if err != nil || data == nil || string(data) == "" {
		t.Fatalf("Did not perform correctly.")
	}
}

func TestCommand_runError(t *testing.T) {
	hct := setup(t)
	// do not defer, causing an error in the Command call
	// and so catching err in Command()
	teardown(t, hct)

	_, err := Command(hct, "", []string{})
	if err == nil || strings.Contains(err.Error(), "from hgcl.run():") == false {
		t.Fatalf("Did not get hgcl().run error.")
	}
}

func TestCommand_HgError(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	// catching hgerr in Command()
	_, err := Command(hct, "nonexistinghgcommand", []string{})
	if err == nil {
		t.Fatalf("Did not catch hgerr.")
	}
}
