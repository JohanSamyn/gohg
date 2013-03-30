// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type addOpts struct {
	Dryrun
	Exclude
	Include
	Mq
	Subrepos
	Debug
	Profile
	Traceback
}

func (cmd *addOpts) String() string {
	return fmt.Sprintf(
		"addOpts = {\n    mq: (%T) %t\n"+
			"    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Mq, cmd.Mq,
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Add provides the 'hg add' command.
func (hgcl *HgClient) Add(files []string, opts ...optionAdder) ([]byte, error) {
	hgcmd, err := hgcl.buildCommand("add", new(addOpts), opts, files)
	if err != nil {
		return nil, err
	}
	data, err := hgcl.runcommand(hgcmd)
	return data, err
}
