// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type headsCmd struct {
	Template
	Debug
	Profile
	Traceback
}

func (cmd *headsCmd) String() string {
	return fmt.Sprintf(
		"headsCmd = {\n    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Heads provides the 'hg heads' command.
func (hgcl *HgClient) Heads(opts ...optionAdder) ([]byte, error) {

	// applies type defaults
	cmd := new(headsCmd)

	// apply gohg defaults
	cmd.Debug = false
	cmd.Profile = false
	cmd.Traceback = false

	hgcmd := []string{"heads"}

	var err error

	// apply option values given by the caller
	for _, o := range opts {
		err = o.addOption(cmd, &hgcmd)
	}

	var data []byte
	data, err = hgcl.runcommand(&hgcmd)
	return data, err
}
