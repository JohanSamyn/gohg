// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type verifyOpts struct {
	Cwd // makes it possible to verify another repo than hgcl.Reporoot()
	Hidden
	NonInteractive
	Quiet
	Repository // makes it possible to verify another repo than hgcl.Reporoot()
	Verbose

	// Mq

	Debug
	Profile
	Time
	Traceback
}

func (cmd *verifyOpts) String() string {
	return fmt.Sprintf(
		// "verifyOpts = {\n    mq: (%T) %t\n"+
		"verifyOpts = {\n"+
			"    debug: (%T) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		// cmd.Mq, cmd.Mq,
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Verify provides the 'hg verify' command.
func (hgcl *HgClient) Verify(opts ...optionAdder) ([]byte, error) {
	hgcmd, err := hgcl.buildCommand("verify", new(verifyOpts), opts, nil)
	if err != nil {
		return nil, err
	}
	data, err := hgcl.runcommand(hgcmd)
	return data, err
}
