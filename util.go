// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

func prependStringToSlice(cmd string, opts []string) []string {
	// adds a string as the first element of an existing slice-of-strings

	c := []string{cmd}
	if opts == nil || len(opts) == 0 {
		return c
	}
	fullcmd := make([]string, len(c)+len(opts))
	copy(fullcmd, c)
	copy(fullcmd[len(c):], opts)
	return fullcmd
}
