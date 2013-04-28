// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"errors"
)

type commitOpts struct {
	Config
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

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

	Debug
	Profile
	Time
	Traceback
}

func (cmdOpts *commitOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func (hgcl *HgClient) Commit(opts []Option, files []string) error {
	cmdOpts := new(commitOpts)
	hgcmd, err := hgcl.buildCommand("commit", cmdOpts, opts, files)
	if err != nil {
		return err
	}

	// Either make sure there is an editor configured for firing up in case
	// there is no commit message provided, or catch the lack of that message.
	// For now we catch it.
	if cmdOpts.Message == "" {
		return errors.New("Commit(): please provide a non-empty commit message.")
	}

	_, err = hgcl.runcommand(hgcmd)
	return err
}
