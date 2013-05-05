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

func NewInitCmd(opts []Option, destpath []string) HgCmd {
	cmd, _ := NewHgCmd("init", opts, destpath, new(initOpts))
	return *cmd
}

// TODO	Implement the flags for hg init.

// Be aware of the fact that Init() cannot be used to initialize the repo you
// want the (current) Hg CS to work on, as the Hg CS requires an existing repo
// before you can connect it. But Init() can be used to create any new repo
// besides the one the Hg CS is running for.
func (hgcl *HgClient) Init(opts []Option, destpath string) error {
	dest := string(destpath[0])
	fa, err := filepath.Abs(dest)
	if err != nil {
		return fmt.Errorf("Init() -> filepath.Abs(): %s", err)
	}
	if dest == "" || dest == "." || fa == hgcl.RepoRoot() /*&& cmdOpts.Mq == false*/ {
		return errors.New("HgClient.Init: path for new repo must be different" +
			" from the Command Server repo path")
	}

	cmd := NewInitCmd(opts, []string{destpath})
	_, err = cmd.Exec(hgcl)
	return err
}
