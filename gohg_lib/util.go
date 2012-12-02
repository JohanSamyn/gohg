// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg_lib

import (
	"fmt"
)

func buildCmd(cmd []string, opts []string) []string {
	if opts == nil || len(opts) == 0 {
		return cmd
	}
	fullcmd := make([]string, len(cmd)+len(opts))
	copy(fullcmd, cmd)
	copy(fullcmd[len(cmd):], opts)
	return fullcmd
}

func command(hgcl *HgClient, cmd string, opts []string) (data []byte, err error) {
	// boilerplate code for all commands

	cmdline := buildCmd(cmd, opts)
	data, hgerr, ret, err := hgcl.run(cmdline)
	if err != nil {
		return nil, fmt.Errorf("from hgcl.run(): %s", err)
	}
	if ret != 0 || hgerr != nil {
		return nil, fmt.Errorf("Status(): returncode=%d\nhgerr:\n%s\n", data, string(hgerr))
	}
	return data, nil
}
