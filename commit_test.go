// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

import (
	// "fmt"
	"strconv"
	"testing"
)

func TestHgClient_Commit(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	err1 := createFile("a.txt", "aaa\n", hct)
	if err1 != nil {
		t.Error(err1)
	}

	t1, err2 := hct.Tip(Template("{rev}"))
	if err2 != nil {
		t.Error(err2)
	}
	revsBefore, _ := strconv.Atoi(string(t1))
	revsBefore++
	// fmt.Printf("revsBefore: %d\n", revsBefore)

	var err error
	_, err = hct.Commit([]string{"a.txt"}, AddRemove(true), Message("first commit"), User("me"))
	if err != nil {
		t.Error(err)
	}
	// fmt.Printf("got: %v\n", got)

	t2, err4 := hct.Tip(Template("{rev}"))
	if err4 != nil {
		t.Error(err4)
	}
	revsAfter, _ := strconv.Atoi(string(t2))
	revsAfter += 1
	// fmt.Printf("revsAfter: %d\n", revsAfter)

	if !(revsAfter > revsBefore) {
		t.Fatalf("Test Commit: expected:\n%d\n but got:\n%d\n", revsBefore+1, revsAfter)
	}
	// t.Errorf("err: %s", "lkk")
}
