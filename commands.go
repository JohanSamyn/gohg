// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
	"strings"
)

func command(hgcl *HgClient, cmd []string) (data []byte, err error) {
	// boilerplate code for all commands

	// fmt.Printf("cmd = %s\nopts = %v\n", cmd[0], cmd[1:])

	data, hgerr, ret, err := hgcl.run(cmd)
	if err != nil {
		return nil, fmt.Errorf("from hgcl.run(): %s", err)
	}
	// Maybe make this 2 checks, to differentiate between ret and hgerr?
	if ret != 0 || hgerr != nil {
		return nil, fmt.Errorf("%s(): returncode=%d\nhgerr:\n%s\n",
			strings.Title(cmd[0]), ret, string(hgerr))
	}
	return data, nil
}

func commandOld(hgcl *HgClient, cmd string, opts []string) (data []byte, err error) {
	// boilerplate code for all commands

	cmdline := PrependStringToSlice(cmd, opts)
	data, hgerr, ret, err := hgcl.run(cmdline)
	if err != nil {
		return nil, fmt.Errorf("from hgcl.run(): %s", err)
	}
	// Maybe make this 2 checks, to differentiate between ret and hgerr?
	if ret != 0 || hgerr != nil {
		return nil, fmt.Errorf("%s(): returncode=%d\nhgerr:\n%s\n",
			strings.Title(cmd), ret, string(hgerr))
	}
	return data, nil
}

// Add provides the 'hg add' command.
func (hgcl *HgClient) Add(opts []string) ([]byte, error) {
	data, err := commandOld(hgcl, "add", opts)
	return data, err
}

// Identify provides the 'hg identify' command.
func (hgcl *HgClient) Identify(opts []string) ([]byte, error) {
	data, err := commandOld(hgcl, "identify", opts)
	return data, err
}

// Add provides the 'hg log' command.
func (hgcl *HgClient) Log(opts []string) ([]byte, error) {
	data, err := commandOld(hgcl, "log", opts)
	return data, err
}

// Status provides the 'hg status' command.
func (hgcl *HgClient) Status(opts []string) ([]byte, error) {
	data, err := commandOld(hgcl, "status", opts)
	return data, err
}
