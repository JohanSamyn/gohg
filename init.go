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
	Destpath
	hgDebugOpts
}

func (cmd *initCmd) String() string {
	return fmt.Sprintf(
		"initCmd = {\n    filepath: (%T) %q\n"+
			"    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Destpath, cmd.Destpath,
		cmd.Debug, cmd.Debug, cmd.Traceback, cmd.Traceback, cmd.Profile, cmd.Profile)
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
	cmd.Destpath = "."
	cmd.Debug = false
	cmd.Traceback = false
	cmd.Profile = false

	// apply option values given by the caller
	for _, o := range opts {
		o.addOption(cmd)
	}

	hgcmd := []string{"init"}
	if cmd.Destpath != "" {
		hgcmd = append(hgcmd, string(cmd.Destpath))
	}
	if cmd.Debug == true {
		hgcmd = append(hgcmd, "--debug")
	}
	if cmd.Traceback == true {
		hgcmd = append(hgcmd, "--traceback")
	}
	if cmd.Profile == true {
		hgcmd = append(hgcmd, "--profile")
	}

	// // Is not shown with 'gt init' ??
	// fmt.Printf("%v\n", cmd)

	var err1 error
	var fa string
	fa, err1 = filepath.Abs(string(cmd.Destpath))
	if err1 != nil {
		return fmt.Errorf("Init() -> filepath.Abs(): %s", err1)
	}
	if cmd.Destpath == "" || cmd.Destpath == "." || fa == hgcl.RepoRoot() {
		return errors.New("HgClient.Init: path for new repo must be different" +
			" from the Command Server repo path")
	}

	_, err := runcommand(hgcl, hgcmd)
	return err
}
