// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type summaryCmd struct {
	O_remote bool
	O_mq     bool
	hgDebugOpts
}

func (cmd *summaryCmd) String() string {
	return fmt.Sprintf(
		"summaryCmd = {\n    remote: (%T) %t\n    mq: (%T) %t\n"+
			"    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.O_remote, cmd.O_remote, cmd.O_mq, cmd.O_mq,
		cmd.O_debug, cmd.O_debug, cmd.O_traceback, cmd.O_traceback, cmd.O_profile, cmd.O_profile)
}

// Summary provides the 'hg summary' command.
func (hgcl *HgClient) Summary(opts ...optionAdder) ([]byte, error) {

	// applies type defaults
	cmd := new(summaryCmd)

	// apply library defaults
	cmd.O_remote = false
	cmd.O_mq = false
	cmd.O_debug = false
	cmd.O_traceback = false
	cmd.O_profile = false

	// apply option values given by the caller
	for _, o := range opts {
		o.addOption(cmd)
	}

	hgcmd := []string{"summary"}
	if cmd.O_remote == true {
		hgcmd = append(hgcmd, "--remote")
	}
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
