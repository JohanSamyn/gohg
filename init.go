// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"errors"
	"fmt"
	"path/filepath"
)

type initCmd struct {
	O_filepath string
	hgDebugOpts
}

func (cmd *initCmd) String() string {
	return fmt.Sprintf(
		"initCmd = {\n    filepath: (%T) %q\n"+
			"    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.O_filepath, cmd.O_filepath,
		cmd.O_debug, cmd.O_debug, cmd.O_traceback, cmd.O_traceback, cmd.O_profile, cmd.O_profile)
}

// TODO	Implement the flags for hg init.

// Init provides the 'hg init' command.
//
// Be aware of the fact that it cannot be used to initialize the repo you want
// the (current) Hg CS to work on, as the Hg CS requires an existing repo.
// But Init() can be used to create any new repo outside the one the Hg CS is
// running for.
func (hgcl *HgClient) Init(opts ...optionAdder) error {

	// applies type defaults
	cmd := new(initCmd)

	// apply library defaults
	cmd.O_filepath = "."
	cmd.O_debug = false
	cmd.O_traceback = false
	cmd.O_profile = false

	// apply option values given by the caller
	for _, o := range opts {
		o.addOption(cmd)
	}

	hgcmd := []string{"init"}
	if cmd.O_filepath != "" {
		hgcmd = append(hgcmd, cmd.O_filepath)
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

	// // Is not shown with 'gt init' ??
	// fmt.Printf("%v\n", cmd)

	var err1 error
	var fa string
	fa, err1 = filepath.Abs(cmd.O_filepath)
	if err1 != nil {
		return fmt.Errorf("Init() -> filepath.Abs(): %s", err1)
	}
	if cmd.O_filepath == "" || cmd.O_filepath == "." || fa == hgcl.RepoRoot() {
		return errors.New("HgClient.Init: path for new repo must be different" +
			" from the Command Server repo path")
	}

	_, err := command(hgcl, hgcmd)
	return err
}
