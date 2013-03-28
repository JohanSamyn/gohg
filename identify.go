// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type identifyOpts struct {
	Insecure
	Mq
	RemoteCmd
	Rev
	ShowBookmarks
	ShowBranch
	ShowId
	ShowNum
	ShowTags
	Ssh
	Debug
	Profile
	Traceback
}

func (cmd *identifyOpts) String() string {
	return fmt.Sprintf(
		"identifyOpts = {\n    bookmarks: (%T) %t\n    branch: (%T) %t\n    id: (%T) %t\n"+
			"    mq: (%T) %t\n    num: (%T) %t\n    rev: (%T) %t\n    tags: (%T) %t\n"+
			"    debug: (%T) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.ShowBookmarks, cmd.ShowBookmarks, cmd.ShowBranch, cmd.ShowBranch,
		cmd.ShowId, cmd.ShowId, cmd.Mq, cmd.Mq, cmd.ShowNum, cmd.ShowNum,
		cmd.Rev, cmd.Rev, cmd.ShowTags, cmd.ShowTags,
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Identify provides the 'hg identify' command.
func (hgcl *HgClient) Identify(opts ...optionAdder) ([]byte, error) {
	cmdOpts := new(identifyOpts)
	// apply gohg defaults (that differ from type default)
	cmdOpts.ShowBookmarks = true
	cmdOpts.ShowBranch = true
	cmdOpts.ShowId = true
	cmdOpts.ShowNum = true
	cmdOpts.ShowTags = true
	hgcmd, err := hgcl.buildCommand("identify", cmdOpts, opts)
	if err != nil {
		return nil, err
	}
	data, err := hgcl.runcommand(hgcmd)
	return data, err
}
