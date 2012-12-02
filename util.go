// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg_lib

import (
	"fmt"
)

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

func command(hgcl *HgClient, cmd string, opts []string) (data []byte, err error) {
	// boilerplate code for all commands

	cmdline := prependStringToSlice(cmd, opts)
	data, hgerr, ret, err := hgcl.run(cmdline)
	if err != nil {
		return nil, fmt.Errorf("from hgcl.run(): %s", err)
	}
	if ret != 0 || hgerr != nil {
		return nil, fmt.Errorf("Status(): returncode=%d\nhgerr:\n%s\n", data, string(hgerr))
	}
	return data, nil
}
