// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type commitOpts struct {
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
func (hgcl *HgClient) Commit(files []string, opts ...optionAdder) ([]byte, error) {
	hgcmd, err := hgcl.buildCommand("commit", new(commitOpts), opts, files)
	if err != nil {
		return nil, err
	}
	return hgcl.runcommand(hgcmd)
}
