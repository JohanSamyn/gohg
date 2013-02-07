// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type verifyCmd struct {
	Mq
	hgDebugOpts
}

func (cmd *verifyCmd) String() string {
	return fmt.Sprintf(
		"verifyCmd = {\n    mq: (%T) %t\n"+
			"    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Mq, cmd.Mq,
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Verify provides the 'hg verify' command.
func (hgcl *HgClient) Verify(opts ...optionAdder) ([]byte, error) {

	// applies type defaults
	cmd := new(verifyCmd)

	// apply library defaults
	cmd.Mq = false
	cmd.Debug = false
	cmd.Traceback = false
	cmd.Profile = false

	// apply option values given by the caller
	for _, o := range opts {
		o.addOption(cmd)
	}

	hgcmd := []string{"verify"}
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

	data, err := runcommand(hgcl, hgcmd)
	return data, err
}
