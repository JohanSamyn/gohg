// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type summaryCmd struct {
	Remote
	Mq
	hgDebugOpts
}

func (cmd *summaryCmd) String() string {
	return fmt.Sprintf(
		"summaryCmd = {\n    remote: (%T) %t\n    mq: (%T) %t\n"+
			"    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Remote, cmd.Remote, cmd.Mq, cmd.Mq,
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Summary provides the 'hg summary' command.
func (hgcl *HgClient) Summary(opts ...optionAdder) ([]byte, error) {

	// applies type defaults
	cmd := new(summaryCmd)

	// apply library defaults
	cmd.Remote = false
	cmd.Mq = false
	cmd.Debug = false
	cmd.Traceback = false
	cmd.Profile = false

	// apply option values given by the caller
	for _, o := range opts {
		o.addOption(cmd)
	}

	hgcmd := []string{"summary"}
	if cmd.Remote == true {
		hgcmd = append(hgcmd, "--remote")
	}
	if cmd.Mq == true {
		hgcmd = append(hgcmd, "--mq")
	}
	if cmd.Debug == true {
		hgcmd = append(hgcmd, "--debug")
	}
	if cmd.Traceback == true {
		hgcmd = append(hgcmd, "--traceback")
	}
	if cmd.Profile == true {
		hgcmd = append(hgcmd, "--profile")
	}

	data, err := command(hgcl, hgcmd)
	return data, err
}
