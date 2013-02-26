// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type addCmd struct {
	Mq
	Debug
	Profile
	Traceback
}

func (cmd *addCmd) String() string {
	return fmt.Sprintf(
		"addCmd = {\n    mq: (%T) %t\n"+
			"    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Mq, cmd.Mq,
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
}

// Add provides the 'hg add' command.
func (hgcl *HgClient) Add(opts ...optionAdder) ([]byte, error) {

	// applies type defaults
	cmd := new(addCmd)

	// apply gohg defaults
	cmd.Mq = false
	cmd.Debug = false
	cmd.Traceback = false
	cmd.Profile = false

	hgcmd := []string{"add"}

	var err error

	// // apply option values given by the caller
	// if len(opts) > 0 {
	// 	for _, o := range opts {
	// 		err = o.addOption(cmd, &hgcmd)
	// 		if err == nil {
	// 			o.translateOption(&hgcmd)
	// 		}
	// 	}
	// }

	var data []byte
	data, err = hgcl.runcommand(&hgcmd)
	return data, err
}
