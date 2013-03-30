// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type identifyOpts struct {
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

	Insecure
	// Mq
	RemoteCmd
	Rev
	Bookmarks
	Branch
	Id
	Num
	Tags
	Ssh

	Debug
	Profile
	Time
	Traceback
}

func (cmd *identifyOpts) String() string {
	return fmt.Sprintf(
		"identifyOpts = {\n    bookmarks: (%T) %t\n    branch: (%T) %t\n    id: (%T) %t\n"+
			// "    mq: (%T) %t\n    num: (%T) %t\n    rev: (%T) %t\n    tags: (%T) %t\n"+
			"    num: (%T) %t\n    rev: (%T) %t\n    tags: (%T) %t\n"+
			"    debug: (%T) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Bookmarks, cmd.Bookmarks, cmd.Branch, cmd.Branch,
		cmd.Id, cmd.Id,
		// cmd.Mq, cmd.Mq,
		cmd.Num, cmd.Num,
		cmd.Rev, cmd.Rev, cmd.Tags, cmd.Tags,
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Identify provides the 'hg identify' command.
func (hgcl *HgClient) Identify(source string, opts ...optionAdder) ([]byte, error) {
	cmdOpts := new(identifyOpts)
	// apply gohg defaults (that differ from type default)
	cmdOpts.Bookmarks = true
	cmdOpts.Branch = true
	cmdOpts.Id = true
	cmdOpts.Num = true
	cmdOpts.Tags = true
	params := []string{source}
	hgcmd, err := hgcl.buildCommand("identify", cmdOpts, opts, params)
	if err != nil {
		return nil, err
	}
	data, err := hgcl.runcommand(hgcmd)
	return data, err
}
