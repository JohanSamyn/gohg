// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type tipOpts struct {
	Template
	Debug
	Profile
	Traceback
}

func (cmd *tipOpts) String() string {
	return fmt.Sprintf(
		"tipOpts = {\n    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Tip provides the 'hg tip' command.
func (hgcl *HgClient) Tip(opts ...optionAdder) ([]byte, error) {
	hgcmd, err := hgcl.buildCommand("tip", new(tipOpts), opts, nil)
	if err != nil {
		return nil, err
	}
	data, err := hgcl.runcommand(hgcmd)
	return data, err
}
