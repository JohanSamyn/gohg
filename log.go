// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
	"strconv"
)

type logCmd struct {
	Limit
	Rev
	hgDebugOpts
}

func (cmd *logCmd) String() string {
	return fmt.Sprintf(
		"logCmd = {\n    limit: (%T) %t\n    rev: (%T) %t\n"+
			"    debug: (%T) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Limit, cmd.Limit, cmd.Rev, cmd.Rev,
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Log provides the 'hg log' command.
func (hgcl *HgClient) Log(opts ...optionAdder) ([]byte, error) {

	// applies type defaults
	cmd := new(logCmd)

	// apply library defaults
	cmd.Limit = 0
	cmd.Rev = ""
	cmd.Debug = false
	cmd.Traceback = false
	cmd.Profile = false

	// apply option values given by the caller
	for _, o := range opts {
		o.addOption(cmd)
	}

	hgcmd := []string{"log"}
	if cmd.Limit > 0 {
		hgcmd = append(hgcmd, "-l")
		hgcmd = append(hgcmd, strconv.Itoa(int(cmd.Limit)))
	}
	if cmd.Rev != "" {
		hgcmd = append(hgcmd, "-r")
		hgcmd = append(hgcmd, string(cmd.Rev))
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
