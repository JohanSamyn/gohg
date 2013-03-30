// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type headsOpts struct {
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

	Closed
	Mq
	Rev
	Style
	Template
	Topo

	Debug
	Profile
	Time
	Traceback
}

func (cmd *headsOpts) String() string {
	return fmt.Sprintf(
		"headsOpts = {\n    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Heads provides the 'hg heads' command.
func (hgcl *HgClient) Heads(revs []string, opts ...optionAdder) ([]byte, error) {
	hgcmd, err := hgcl.buildCommand("heads", new(headsOpts), opts, revs)
	if err != nil {
		return nil, err
	}
	data, err := hgcl.runcommand(hgcmd)
	return data, err
}
