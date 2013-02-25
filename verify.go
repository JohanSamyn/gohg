// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type verifyCmd struct {
	Mq
	hgDebugOpts
}

func (cmd *verifyCmd) String() string {
	return fmt.Sprintf(
		"verifyCmd = {\n    mq: (%T) %t\n"+
			"    debug: (%T) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Mq, cmd.Mq,
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Verify provides the 'hg verify' command.
func (hgcl *HgClient) Verify(opts ...optionAdder) ([]byte, error) {

	// applies type defaults
	cmd := new(verifyCmd)

	// apply gohg defaults
	cmd.Mq = false
	cmd.Debug = false
	cmd.Traceback = false
	cmd.Profile = false

	hgcmd := []string{"verify"}

	var err error

	// apply option values given by the caller
	for _, o := range opts {
		err = o.addOption(cmd)
		if err == nil {
			o.translateOption(&hgcmd)
		}
	}

	var data []byte
	data, err = hgcl.runcommand(&hgcmd)
	return data, err
}
