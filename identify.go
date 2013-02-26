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
	Debug
	Profile
	Traceback
}

func (cmd *identifyCmd) String() string {
	return fmt.Sprintf(
		"identifyCmd = {\n    bookmarks: (%T) %t\n    branch: (%T) %t\n    id: (%T) %t\n"+
			"    mq: (%T) %t\n    num: (%T) %t\n    rev: (%T) %t\n    tags: (%T) %t\n"+
			"    debug: (%T) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.ShowBookmarks, cmd.ShowBookmarks, cmd.ShowBranch, cmd.ShowBranch,
		cmd.ShowId, cmd.ShowId, cmd.Mq, cmd.Mq, cmd.ShowNum, cmd.ShowNum,
		cmd.Rev, cmd.Rev, cmd.ShowTags, cmd.ShowTags,
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Identify provides the 'hg identify' command.
func (hgcl *HgClient) Identify(opts ...optionAdder) ([]byte, error) {

	// applies type defaults
	cmd := new(identifyCmd)

	// apply gohg defaults
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

	hgcmd := []string{"identify"}

	var err error

	// apply option values given by the caller
	for _, o := range opts {
		err = o.addOption(cmd)
		if err == nil {
			o.translateOption(&hgcmd)
		}
	}

	var data []byte
	data, err = hgcl.runcommand(&hgcmd)
	return data, err
}
