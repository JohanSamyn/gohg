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
	Cwd
	Destpath
	Mq
	Debug
	Profile
	Traceback
}

func (cmd *initCmd) String() string {
	return fmt.Sprintf(
		"initCmd = {\n    filepath: (%T) %q\n    Mq: (%T) %t\n"+
			"    Cwd: (%T) %t\n"+
			"    debug: (%T) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.Destpath, cmd.Destpath, cmd.Mq, cmd.Mq, cmd.Cwd, cmd.Cwd,
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

	// apply gohg defaults
	cmd.Cwd = ""
	cmd.Destpath = "."
	cmd.Mq = false
	cmd.Debug = false
	cmd.Traceback = false
	cmd.Profile = false

	hgcmd := []string{"init"}

	var err error

	// apply option values given by the caller
	for _, o := range opts {
		err = o.addOption(cmd, &hgcmd)
		// if err == nil {
		// 	o.translateOption(&hgcmd)
		// }
	}

	var err1 error
	var fa string
	fa, err1 = filepath.Abs(string(cmd.Destpath))
	if err1 != nil {
		return fmt.Errorf("Init() -> filepath.Abs(): %s", err1)
	}
	if (cmd.Destpath == "" || cmd.Destpath == "." || fa == hgcl.RepoRoot()) && cmd.Mq == false {
		return errors.New("HgClient.Init: path for new repo must be different" +
			" from the Command Server repo path")
	}

	_, err = hgcl.runcommand(&hgcmd)
	return err
}
