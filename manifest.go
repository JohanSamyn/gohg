// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type manifestCmd struct {
	Debug
	Profile
	Traceback
}

func (cmd *manifestCmd) String() string {
	return fmt.Sprintf(
		"manifestCmd = {\n    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Manifest provides the 'hg manifest' command.
func (hgcl *HgClient) Manifest(opts ...optionAdder) ([]byte, error) {

	// applies type defaults
	cmd := new(manifestCmd)

	// apply gohg defaults
	cmd.Debug = false
	cmd.Profile = false
	cmd.Traceback = false

	hgcmd := []string{"manifest"}

	var err error

	// apply option values given by the caller
	for _, o := range opts {
		err = o.addOption(cmd, &hgcmd)
	}

	var data []byte
	data, err = hgcl.runcommand(&hgcmd)
	return data, err
}
