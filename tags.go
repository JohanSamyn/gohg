// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type tagsCmd struct {
	Mq
	Verbose
	Debug
	Profile
	Traceback
}

func (cmd *tagsCmd) String() string {
	return fmt.Sprintf(
		"tagsCmd = {\n    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Tags provides the 'hg tags' command.
func (hgcl *HgClient) Tags(opts ...optionAdder) ([]byte, error) {

	// applies type defaults
	cmd := new(tagsCmd)

	// apply gohg defaults
	cmd.Debug = false
	cmd.Profile = false
	cmd.Traceback = false

	hgcmd := []string{"tags"}

	var err error

	// apply option values given by the caller
	for _, o := range opts {
		err = o.addOption(cmd, &hgcmd)
	}

	var data []byte
	data, err = hgcl.runcommand(&hgcmd)
	return data, err
}
