// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type addCmd struct {
	O_mq bool
	hgDebugOpts
}

func (cmd *addCmd) String() string {
	return fmt.Sprintf(
		"addCmd = {\n    mq: (%T) %t\n"+
			"    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.O_mq, cmd.O_mq,
		cmd.O_debug, cmd.O_debug, cmd.O_traceback, cmd.O_traceback, cmd.O_profile, cmd.O_profile)
}

// Add provides the 'hg add' command.
func (hgcl *HgClient) Add(opts ...optionAdder) ([]byte, error) {

	// applies type defaults
	cmd := new(addCmd)

	// apply library defaults
	cmd.O_mq = false
	cmd.O_debug = false
	cmd.O_traceback = false
	cmd.O_profile = false

	// apply option values given by the caller
	for _, o := range opts {
		o.addOption(cmd)
	}

	hgcmd := []string{"add"}
	if cmd.O_mq == true {
		hgcmd = append(hgcmd, "--mq")
	}
	if cmd.O_debug == true {
		hgcmd = append(hgcmd, "--debug")
	}
	if cmd.O_traceback == true {
		hgcmd = append(hgcmd, "--traceback")
	}
	if cmd.O_profile == true {
		hgcmd = append(hgcmd, "--profile")
	}

	data, err := command(hgcl, hgcmd)
	return data, err
}
