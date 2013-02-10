// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type identifyCmd struct {
	ShowBookmarks
	ShowBranch
	ShowId
	Mq
	ShowNum
	Rev
	ShowTags
	hgDebugOpts
}

func (cmd *identifyCmd) String() string {
	return fmt.Sprintf(
		"identifyCmd = {\n    bookmarks: (%T) %t\n    branch: (%T) %t\n    id: (%T) %t\n"+
			"    mq: (%T) %t\n    num: (%T) %t\n    rev: (%T) %t\n    tags: (%T) %t\n"+
			"    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.ShowBookmarks, cmd.ShowBookmarks, cmd.ShowBranch, cmd.ShowBranch,
		cmd.ShowId, cmd.ShowId, cmd.Mq, cmd.Mq, cmd.ShowNum, cmd.ShowNum,
		cmd.Rev, cmd.Rev, cmd.ShowTags, cmd.ShowTags,
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Identify provides the 'hg identify' command.
func (hgcl *HgClient) Identify(opts ...optionAdder) ([]byte, error) {

	// applies type defaults
	cmd := new(identifyCmd)

	// apply library defaults
	cmd.Mq = false
	cmd.Rev = ""
	cmd.ShowBookmarks = true
	cmd.ShowBranch = true
	cmd.ShowId = true
	cmd.ShowNum = true
	cmd.ShowTags = true
	cmd.Debug = false
	cmd.Profile = false
	cmd.Traceback = false

	// apply option values given by the caller
	for _, o := range opts {
		o.addOption(cmd)
	}

	hgcmd := []string{"identify"}
	if cmd.Mq {
		hgcmd = append(hgcmd, "--mq")
	}
	if cmd.Rev != "" {
		hgcmd = append(hgcmd, "-r")
		hgcmd = append(hgcmd, string(cmd.Rev))
	}
	if cmd.ShowBookmarks {
		hgcmd = append(hgcmd, "-B")
	}
	if cmd.ShowBranch {
		hgcmd = append(hgcmd, "-b")
	}
	if cmd.ShowId {
		hgcmd = append(hgcmd, "-i")
	}
	if cmd.ShowNum {
		hgcmd = append(hgcmd, "-n")
	}
	if cmd.ShowTags {
		hgcmd = append(hgcmd, "-t")
	}
	if cmd.Debug {
		hgcmd = append(hgcmd, "--debug")
	}
	if cmd.Traceback {
		hgcmd = append(hgcmd, "--traceback")
	}
	if cmd.Profile {
		hgcmd = append(hgcmd, "--profile")
	}

	data, err := hgcl.runcommand(hgcmd)
	return data, err
}
