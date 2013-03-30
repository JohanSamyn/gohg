// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type logOpts struct {
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

	Branch
	Copies
	Date
	Follow
	Git
	Graph
	Keyword
	Limit
	Mq
	NoMerges
	Patch
	Prune
	Removed
	Rev
	Stat
	Style
	Template
	User

	Debug
	Profile
	Time
	Traceback
}

func (cmd *logOpts) String() string {
	return fmt.Sprintf(
		"logOpts = {\n    limit: (%T) %t\n    rev: (%T) %t\n    mq: (%T) %t\n"+
			"    template: (%T) %t\n"+
			"    debug: (%T) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Limit, cmd.Limit, cmd.Rev, cmd.Rev, cmd.Mq, cmd.Mq,
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Log provides the 'hg log' command.
func (hgcl *HgClient) Log(files []string, opts ...optionAdder) ([]byte, error) {
	hgcmd, err := hgcl.buildCommand("log", new(logOpts), opts, files)
	if err != nil {
		return nil, err
	}
	data, err := hgcl.runcommand(hgcmd)
	return data, err
}
