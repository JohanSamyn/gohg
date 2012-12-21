// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"strings"
	"testing"
)

func TestPrependStringToSlice(t *testing.T) {
	expected := []string{"cmd", "option1", "option2"}
	got := PrependStringToSlice(expected[0], expected[1:])
	if strings.Join(got, ", ") != strings.Join(expected, ", ") {
		t.Fatalf("Expected %s but got %s", expected, got)
	}
}
