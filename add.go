// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type addCmd struct {
	Mq
	hgDebugOpts
}

func (cmd *addCmd) String() string {
	return fmt.Sprintf(
		"addCmd = {\n    mq: (%T) %t\n"+
			"    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Mq, cmd.Mq,
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Add provides the 'hg add' command.
func (hgcl *HgClient) Add(opts ...optionAdder) ([]byte, error) {

	// applies type defaults
	cmd := new(addCmd)

	// apply library defaults
	cmd.Mq = false
	cmd.Debug = false
	cmd.Traceback = false
	cmd.Profile = false

	// apply option values given by the caller
	if len(opts) > 0 {
		for _, o := range opts {
			o.addOption(cmd)
		}
	}

	hgcmd := []string{"add"}
	if cmd.Mq {
		hgcmd = append(hgcmd, "--mq")
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

	data, err := runcommand(hgcl, hgcmd)
	return data, err
}
