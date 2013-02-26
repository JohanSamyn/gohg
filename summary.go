// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type summaryCmd struct {
	Remote
	Mq
	Debug
	Profile
	Traceback
}

func (cmd *summaryCmd) String() string {
	return fmt.Sprintf(
		"summaryCmd = {\n    remote: (%T) %t\n    mq: (%T) %t\n"+
			"    debug: (%T) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Remote, cmd.Remote, cmd.Mq, cmd.Mq,
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Summary provides the 'hg summary' command.
func (hgcl *HgClient) Summary(opts ...optionAdder) ([]byte, error) {

	// applies type defaults
	cmd := new(summaryCmd)

	// apply gohg defaults
	cmd.Remote = false
	cmd.Mq = false
	cmd.Debug = false
	cmd.Traceback = false
	cmd.Profile = false

	hgcmd := []string{"summary"}

	var err error

	// apply option values given by the caller
	for _, o := range opts {
		err = o.addOption(cmd, &hgcmd)
		// if err == nil {
		// 	o.translateOption(&hgcmd)
		// } else {
		// 	fmt.Printf("err = ", err)
		// }
	}

	var data []byte
	data, err = hgcl.runcommand(&hgcmd)
	return data, err
}
