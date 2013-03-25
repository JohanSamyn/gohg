// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type headsOpts struct {
	Closed
	Mq
	Rev
	Style
	Topo
	Template
	Debug
	Profile
	Traceback
}

func (cmd *headsOpts) String() string {
	return fmt.Sprintf(
		"headsOpts = {\n    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Heads provides the 'hg heads' command.
func (hgcl *HgClient) Heads(opts ...optionAdder) ([]byte, error) {
	hgcmd, err := hgcl.buildCommand("heads", new(headsOpts), opts)
	if err != nil {
		return nil, err
	}
	data, err := hgcl.runcommand(&hgcmd)
	return data, err
}
