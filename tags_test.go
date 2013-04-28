// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"testing"
)

func TestHgClient_Tags(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	var expected string = "tip                               -1:000000000000\n"

	got, err := hct.Tags(nil, nil)
	if err != nil {
		t.Error(err)
	}

	if string(got) != expected {
		t.Fatalf("Test Tags: expected:\n%s\n but got:\n%s\n", expected, got)
	}
}
