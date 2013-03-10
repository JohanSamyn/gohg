// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type statusOpts struct {
	Debug
	Profile
	Traceback
}

func (cmd *statusOpts) String() string {
	return fmt.Sprintf(
		"statusOpts = {\n    debug: (%T) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Status provides the 'hg status' command.
func (hgcl *HgClient) Status(opts ...optionAdder) ([]byte, error) {
	hgcmd, err := hgcl.buildCommand("status", new(statusOpts), opts)
	if err != nil {
		return nil, err
	}
	data, err := hgcl.runcommand(&hgcmd)
	return data, err
}
