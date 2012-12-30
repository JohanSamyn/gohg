// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type verifyCmd struct {
	// O_repository string
	O_mq bool
	hgDebugOpts
}

func (cmd *verifyCmd) String() string {
	return fmt.Sprintf(
		"verifyCmd = {\n    mq: (%T) %t\n"+
			"    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.O_mq, cmd.O_mq,
		cmd.O_debug, cmd.O_debug, cmd.O_traceback, cmd.O_traceback, cmd.O_profile, cmd.O_profile)
}

// func (cmd *verifyCmd) String() string {
// 	return fmt.Sprintf(
// 		"verifyCmd = {\n    repository: (%T) %q\n    mq: (%T) %t\n"+
// 			"    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
// 		cmd.O_repository, cmd.O_repository, cmd.O_mq, cmd.O_mq,
// 		cmd.O_debug, cmd.O_debug, cmd.O_traceback, cmd.O_traceback, cmd.O_profile, cmd.O_profile)
// }

// Verify provides the 'hg verify' command.
func (hgcl *HgClient) Verify(opts ...optionAdder) ([]byte, error) {

	// applies type defaults
	cmd := new(verifyCmd)

	// apply library defaults
	// cmd.O_repository = "" // uses HgClient.RepoRoot()
	cmd.O_mq = false
	cmd.O_debug = false
	cmd.O_traceback = false
	cmd.O_profile = false

	// apply option values given by the caller
	for _, o := range opts {
		o.addOption(cmd)
	}

	hgcmd := []string{"verify"}
	// if cmd.O_repository != "" {
	// 	hgcmd = append(hgcmd, "-R")
	// 	hgcmd = append(hgcmd, cmd.O_repository)
	// }
	if cmd.O_mq == true {
		hgcmd = append(hgcmd, "--mq")
	}
	if cmd.O_debug == true {
		hgcmd = append(hgcmd, "--debug")
	}
	if cmd.O_traceback == true {
		hgcmd = append(hgcmd, "--traceback")
	}
	if cmd.O_profile == true {
		hgcmd = append(hgcmd, "--profile")
	}

	data, err := command(hgcl, hgcmd)
	return data, err
}