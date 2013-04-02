// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"errors"
	"fmt"
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
	return fmt.Sprintf(
		"commitCmd = {\n    Debug: (%T) %t\n    Profile: (%T) %t\n"+
			"   Time: (%T) %t\n    Traceback: (%T) %t\n}\n",
		cmdOpts.Debug, cmdOpts.Debug, cmdOpts.Profile, cmdOpts.Profile,
		cmdOpts.Time, cmdOpts.Time, cmdOpts.Traceback, cmdOpts.Traceback)
}

// Commit provides the 'hg commit' command.
func (hgcl *HgClient) Commit(files []string, opts ...optionAdder) error {
	ciOpts := new(commitOpts)
	hgcmd, err := hgcl.buildCommand("commit", ciOpts, opts, files)
	if err != nil {
		return err
	}

	// Either make sure there is an editor configured for firing up in case
	// there is no commit message provided, or catch the lack of that message.
	// For now we catch it.
	if ciOpts.Message == "" {
		return errors.New("Commit(): please provide a non-empty commit message.")
	}

	_, err = hgcl.runcommand(hgcmd)
	return err
}
