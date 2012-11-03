// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg_lib

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
)

// TODO	Implement the flags for hg init.

// Init provides the 'hg init' command.
//
// Be aware of the fact that it cannot be used to initialize the repo you want
// the (current) Hg CS to work on, as the Hg CS requires an existing repo.
// But Init() can be used to create any new repo outside the one the Hg CS is
// running for.
// func (hgcl *HgClient) Init(path string, args []string) error {
func (hgcl *HgClient) Init(path string) error {
	var err error
	var fa string
	fa, err = filepath.Abs(path)

	if path == "" || path == "." || fa == hgcl.Repo() {
		return errors.New("HgClient.Init: path for new repo must be different" +
			" from the Command Server repo path")
	}

	var data []byte
	var ret int32

	hgcmd := []string{"init", fa}
	data, ret, err = hgcl.run(hgcmd)
	if err != nil {
		return fmt.Errorf("from run(): %s", string(err.Error()))
	}
	// Will have to capture the "e" channel to be able to return a useful
	// error message in case of failure.
	if ret != 0 {
		return fmt.Errorf("HgClient.Init():\npath -> %s\ndata ->\n%s\nret -> %s",
			fa, string(data), strconv.Itoa(int(ret)))
	}

	return nil
}
