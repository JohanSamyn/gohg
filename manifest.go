// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type manifestOpts struct {
	All
	Mq
	rev
	Debug
	Profile
	Traceback
}

func (cmd *manifestOpts) String() string {
	return fmt.Sprintf(
		"manifestOpts = {\n    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Manifest provides the 'hg manifest' command.
func (hgcl *HgClient) Manifest(opts ...optionAdder) ([]byte, error) {
	hgcmd, err := hgcl.buildCommand("manifest", new(manifestOpts), opts)
	if err != nil {
		return nil, err
	}
	data, err := hgcl.runcommand(&hgcmd)
	return data, err
}
