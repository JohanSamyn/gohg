// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"testing"
)

func TestHgClient_Tip(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	var expected string = "changeset:   -1:000000000000\n" +
		"tag:         tip\n" +
		"user:        \n" +
		"date:        Thu Jan 01 00:00:00 1970 +0000\n" +
		"\n"

	got, err := hct.Tip()
	if err != nil {
		t.Error(err)
	}

	if string(got) != expected {
		t.Fatalf("Test Tip: expected:\n%s\n but got:\n%s\n", expected, got)
	}
}

func TestHgClient_Tip_Rev(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	var expected string = "\"-1\n\""

	got, err := hct.Tip(Template("{rev}\n"))
	if err != nil {
		t.Error(err)
	}

	if string(got) != expected {
		t.Fatalf("Test Tip: expected:\n%s\n but got:\n%s\n", expected, got)
	}
}
