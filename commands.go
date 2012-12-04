// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

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

// Add provides the 'hg add' command.
func (hgcl *HgClient) Add(opts []string) ([]byte, error) {
	data, err := command(hgcl, "add", opts)
	return data, err
}

// Identify provides the 'hg identify' command.
func (hgcl *HgClient) Identify(opts []string) (string, error) {
	data, err := command(hgcl, "identify", opts)
	return string(data), err
}

// TODO	Implement the flags for hg init.

// Init provides the 'hg init' command.
//
// Be aware of the fact that it cannot be used to initialize the repo you want
// the (current) Hg CS to work on, as the Hg CS requires an existing repo.
// But Init() can be used to create any new repo outside the one the Hg CS is
// running for.
func (hgcl *HgClient) Init(path string, opts []string) error {
	var err1 error
	var fa string
	fa, err1 = filepath.Abs(path)
	if err1 != nil {
		return fmt.Errorf("Init() -> filepath.Abs(): %s", err1)
	}
	if path == "" || path == "." || fa == hgcl.RepoRoot() {
		return errors.New("HgClient.Init: path for new repo must be different" +
			" from the Command Server repo path")
	}

	allopts := prependStringToSlice(fa, []string{})
	_, err := command(hgcl, "init", allopts)
	return err
}

// Add provides the 'hg log' command.
func (hgcl *HgClient) Log(opts []string) ([]byte, error) {
	data, err := command(hgcl, "log", opts)
	return data, err
}

// Status provides the 'hg status' command.
func (hgcl *HgClient) Status(opts []string) ([]byte, error) {
	data, err := command(hgcl, "status", opts)
	return data, err
}

// Summary provides the 'hg summary' command.
func (hgcl *HgClient) Summary(opts []string) ([]byte, error) {
	data, err := command(hgcl, "summary", opts)
	return data, err
}

// Version implements the 'hg version -q' command,
// and only returns the version number.
func (hgcl *HgClient) Version() (string, error) {
	var err error
	if hgcl.hgVersion == "" {
		var data []byte
		data, err = command(hgcl, "version", []string{"-q"})
		if err == nil {
			ver := strings.Split(string(data), "\n")[0]
			ver = ver[strings.LastIndex(ver, " ")+1 : len(ver)-1]
			hgcl.hgVersion = ver
		}
	}
	return hgcl.hgVersion, err
}
