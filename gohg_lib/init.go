// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package gohg_lib

import (
	"errors"
	"path/filepath"
	"strconv"
)

// TODO	Implement the flags for hg init.

// Init provides the 'hg init' command.
func (hgcl *HgClient) Init(path string) error {
	var err error
	var fa string
	fa, err = filepath.Abs(path)

	if path == "" || path == "." || fa == hgcl.Repo {
		return errors.New("HgClient.Init: path for new repo must be different" +
			" from the Command Server repo path")
	}

	var data []byte
	var ret int32

	hgcmd := []string{"init", fa}
	data, ret, err = hgcl.run(hgcmd)
	if err != nil {
		return errors.New("from run(): " + string(err.Error()))
	}
	// Will have to capture the "e" channel to be able to return a useful
	// error message in case of failure.
	if ret != 0 {
		return errors.New("HgClient.Init():\npath -> " + fa + "\ndata ->\n" +
			string(data) + "\nret -> " + strconv.Itoa(int(ret)))
	}

	return nil
}
