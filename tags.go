// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type tagsOpts struct {
	Git
	Mq
	Patch
	Style
	Template
	Verbose
	Debug
	Profile
	Traceback
}

func (cmd *tagsOpts) String() string {
	return fmt.Sprintf(
		"tagsOpts = {\n    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Tags provides the 'hg tags' command.
func (hgcl *HgClient) Tags(opts ...optionAdder) ([]byte, error) {
	hgcmd, err := hgcl.buildCommand("tags", new(tagsOpts), opts, nil)
	if err != nil {
		return nil, err
	}
	data, err := hgcl.runcommand(hgcmd)
	return data, err
}
