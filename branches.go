// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type branchesCmd struct {
	Active
	Closed
	Mq
	Debug
	Profile
	Traceback
}

func (cmd *branchesCmd) String() string {
	return fmt.Sprintf(
		"branchesCmd = {\n    "+
			"Active: (%T) %t\n    Closed: (%T) %t\n    Mq: (%T) %t\n"+
			"debug: (%T) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Active, cmd.Active, cmd.Closed, cmd.Closed, cmd.Mq, cmd.Mq,
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Branches provides the 'hg branches' command.
func (hgcl *HgClient) Branches(opts ...optionAdder) ([]byte, error) {

	// applies type defaults
	cmd := new(branchesCmd)

	// apply gohg defaults
	cmd.Active = false
	cmd.Closed = false
	cmd.Mq = false
	cmd.Debug = false
	cmd.Profile = false
	cmd.Traceback = false

	hgcmd := []string{"branches"}

	var err error

	// apply option values given by the caller
	for _, o := range opts {
		err = o.addOption(cmd, &hgcmd)
	}

	var data []byte
	data, err = hgcl.runcommand(&hgcmd)
	return data, err
}
