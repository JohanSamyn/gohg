// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg_lib

import (
	"errors"
	"fmt"
	"path/filepath"
)

// TODO	Implement the flags for hg init.

// Init provides the 'hg init' command.
//
// Be aware of the fact that it cannot be used to initialize the repo you want
// the (current) Hg CS to work on, as the Hg CS requires an existing repo.
// But Init() can be used to create any new repo outside the one the Hg CS is
// running for.
func (hgcl *HgClient) Init(path string) error {
	// func (hgcl *HgClient) Init(path string, args []string) error {
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

	hgcmd := []string{"init", fa}
	data, hgerr, ret, err := hgcl.run(hgcmd)
	if err != nil {
		return fmt.Errorf("from hgcl.run(): %s", err)
	}
	if ret != 0 || hgerr != nil {
		return fmt.Errorf("Init():\npath=%s\ndata:\n%s\nreturncode=%d\nhgerr:\n%s\n",
			fa, string(data), ret, string(hgerr))
	}
	return nil
}
