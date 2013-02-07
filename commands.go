// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
	"strings"
)

func commandOld(hgcl *HgClient, cmd string, opts []string) (data []byte, err error) {
	// boilerplate code for all commands

	cmdline := PrependStringToSlice(cmd, opts)
	var hgerr []byte
	var ret int32
	data, hgerr, ret, err = runInHg(hgcl, "runcommand", cmdline)
	if err != nil {
		return nil, fmt.Errorf("from runInHg(): %s", err)
	}
	// Maybe make this 2 checks, to differentiate between ret and hgerr?
	if ret != 0 || hgerr != nil {
		return nil, fmt.Errorf("%s(): returncode=%d\nhgerr:\n%s\n",
			strings.Title(cmd), ret, string(hgerr))
	}
	return data, nil
}

// Log provides the 'hg log' command.
func (hgcl *HgClient) Log(opts []string) ([]byte, error) {
	data, err := commandOld(hgcl, "log", opts)
	return data, err
}

// Status provides the 'hg status' command.
func (hgcl *HgClient) Status(opts []string) ([]byte, error) {
	data, err := commandOld(hgcl, "status", opts)
	return data, err
}
