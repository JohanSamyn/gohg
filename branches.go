// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type branchesOpts struct {
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

	Active
	Closed
	// Mq

	Debug
	Profile
	Time
	Traceback
}

func (cmd *branchesOpts) String() string {
	return fmt.Sprintf(
		"branchesOpts = {\n    "+
			// "Active: (%T) %t\n    Closed: (%T) %t\n    Mq: (%T) %t\n"+
			"Active: (%T) %t\n    Closed: (%T) %t\n"+
			"debug: (%T) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Active, cmd.Active, cmd.Closed, cmd.Closed,
		// cmd.Mq, cmd.Mq,
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Branches provides the 'hg branches' command.
func (hgcl *HgClient) Branches(opts ...optionAdder) ([]byte, error) {
	hgcmd, err := hgcl.buildCommand("branches", new(branchesOpts), opts, nil)
	if err != nil {
		return nil, err
	}
	data, err := hgcl.runcommand(hgcmd)
	return data, err
}
