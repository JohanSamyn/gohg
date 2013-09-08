// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

import (
	"errors"
)

type commitOpts struct {
	globalOpts

	AddRemove
	Amend
	CloseBranch
	Date
	Exclude
	Include
	Logfile
	Message
	SubRepos
	User

	debugOpts
}

func (cmdOpts *commitOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewCommitCmd(opts []HgOption, files []string) HgCmd {
	cmd, _ := NewHgCmd("commit", opts, files, new(commitOpts))
	return *cmd
}

func (hgcl *HgClient) Commit(opts []HgOption, files []string) error {
	cmd := NewCommitCmd(opts, files)

	// Either make sure there is an editor configured for firing up in case
	// there is no commit message provided, or catch the lack of that message.
	// For now we catch it.
	var err error
	// We have to build the command to have any values in cmd.cmdOpts.
	cmd.cmd, err = cmd.buildCommand()
	if err != nil {
		return err
	}
	// Maybe this would be easier if commitOpts was a map ?
	if cmd.cmdOpts.(*commitOpts).Message == "" {
		return errors.New("Commit(): please provide a non-empty commit message.")
	}

	_, err = cmd.Exec(hgcl)
	return err
}
