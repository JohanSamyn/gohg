// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"errors"
	"fmt"
	"path/filepath"
)

type initOpts struct {
	Config
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

	Insecure
	// Mq
	Remote
	Ssh

	Debug
	Profile
	Time
	Traceback
}

func (cmdOpts *initOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

// TODO	Implement the flags for hg init.

// Init provides the 'hg init' command.
//
// Be aware of the fact that it cannot be used to initialize the repo you want
// the (current) Hg CS to work on, as the Hg CS requires an existing repo.
// But Init() can be used to create any new repo besides the one the Hg CS is
// running for.
func (hgcl *HgClient) Init(destpath string, opts ...optionAdder) error {
	cmdOpts := new(initOpts)
	params := []string{destpath}
	hgcmd, err := hgcl.buildCommand("init", cmdOpts, opts, params)
	if err != nil {
		return err
	}

	fa, err := filepath.Abs(destpath)
	if err != nil {
		return fmt.Errorf("Init() -> filepath.Abs(): %s", err)
	}
	if destpath == "" || destpath == "." || fa == hgcl.RepoRoot() /*&& cmdOpts.Mq == false*/ {
		return errors.New("HgClient.Init: path for new repo must be different" +
			" from the Command Server repo path")
	}

	_, err = hgcl.runcommand(hgcmd)
	return err
}
