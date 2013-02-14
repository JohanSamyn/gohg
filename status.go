// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type statusCmd struct {
	hgDebugOpts
}

func (cmd *statusCmd) String() string {
	return fmt.Sprintf(
		"statusCmd = {\n    debug: (%T) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Status provides the 'hg status' command.
func (hgcl *HgClient) Status(opts ...optionAdder) ([]byte, error) {

	// applies type defaults
	cmd := new(statusCmd)

	// apply library defaults
	cmd.Debug = false
	cmd.Traceback = false
	cmd.Profile = false

	// apply option values given by the caller
	for _, o := range opts {
		o.addOption(cmd)
	}

	hgcmd := []string{"status"}
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
