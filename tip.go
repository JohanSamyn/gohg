// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type tipCmd struct {
	// define all necessary options/flags

	Debug
	Profile
	Traceback
}

func (cmd *tipCmd) String() string {
	return fmt.Sprintf(
		"tipCmd = {\n    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Tip provides the 'hg tip' command.
func (hgcl *HgClient) Tip(opts ...optionAdder) ([]byte, error) {

	// applies type defaults
	cmd := new(tipCmd)

	// apply gohg defaults
	cmd.Debug = false
	cmd.Profile = false
	cmd.Traceback = false

	hgcmd := []string{"tip"}

	var err error

	// apply option values given by the caller
	for _, o := range opts {
		err = o.addOption(cmd, &hgcmd)
	}

	var data []byte
	data, err = hgcl.runcommand(&hgcmd)
	return data, err
}
