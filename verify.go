// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type verifyOpts struct {
	Mq
	Debug
	Profile
	Traceback
}

func (cmd *verifyOpts) String() string {
	return fmt.Sprintf(
		"verifyOpts = {\n    mq: (%T) %t\n"+
			"    debug: (%T) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Mq, cmd.Mq,
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Verify provides the 'hg verify' command.
func (hgcl *HgClient) Verify(opts ...optionAdder) ([]byte, error) {
	hgcmd, err := hgcl.buildCommand("verify", new(verifyOpts), opts)
	if err != nil {
		return nil, err
	}
	data, err := hgcl.runcommand(hgcmd)
	return data, err
}
