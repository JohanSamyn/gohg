// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type summaryOpts struct {
	Remote
	Mq
	Debug
	Profile
	Traceback
}

func (cmd *summaryOpts) String() string {
	return fmt.Sprintf(
		"summaryOpts = {\n    remote: (%T) %t\n    mq: (%T) %t\n"+
			"    debug: (%T) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Remote, cmd.Remote, cmd.Mq, cmd.Mq,
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Summary provides the 'hg summary' command.
func (hgcl *HgClient) Summary(opts ...optionAdder) ([]byte, error) {
	hgcmd, err := hgcl.buildCommand("summary", new(summaryOpts), opts, nil)
	if err != nil {
		return nil, err
	}
	data, err := hgcl.runcommand(hgcmd)
	return data, err
}
