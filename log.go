// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
	// "strings"
)

type logCmd struct {
	Limit
	Rev
	Debug
	Profile
	Traceback
}

func (cmd *logCmd) String() string {
	return fmt.Sprintf(
		"logCmd = {\n    limit: (%T) %t\n    rev: (%T) %t\n"+
			"    debug: (%T) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Limit, cmd.Limit, cmd.Rev, cmd.Rev,
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Log provides the 'hg log' command.
func (hgcl *HgClient) Log(opts ...optionAdder) ([]byte, error) {

	// applies type defaults
	cmd := new(logCmd)

	// apply gohg defaults
	cmd.Limit = 0
	cmd.Rev = ""
	cmd.Debug = false
	cmd.Traceback = false
	cmd.Profile = false

	hgcmd := []string{"log"}

	var data []byte
	var err error

	// apply option values given by the caller
	for _, o := range opts {
		err = o.addOption(cmd, &hgcmd)
		// if err == nil {
		// 	o.translateOption(&hgcmd)
		// 	// } else {
		// 	// Silently skip the invalid option.
		// 	// Work out some logging system for gohg,
		// 	// and write this error message inthere.
		// 	// err = fmt.Errorf("%s", strings.Replace(fmt.Sprint(err), "<cmd>", "Log", 1))
		// 	// return data, err
		// }
	}

	data, err = hgcl.runcommand(&hgcmd)
	return data, err
}
