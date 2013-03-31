// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type statusOpts struct {
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

	Added
	All
	Change
	Clean
	Copies
	Deleted
	Ignored
	Modified
	NoStatus
	Print0
	Removed
	SubRepos
	Unknown

	Debug
	Profile
	Time
	Traceback
}

func (cmd *statusOpts) String() string {
	return fmt.Sprintf(
		"statusOpts = {\n    debug: (%T) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Status provides the 'hg status' command.
func (hgcl *HgClient) Status(files []string, opts ...optionAdder) ([]byte, error) {
	hgcmd, err := hgcl.buildCommand("status", new(statusOpts), opts, files)
	if err != nil {
		return nil, err
	}
	return hgcl.runcommand(hgcmd)
}
