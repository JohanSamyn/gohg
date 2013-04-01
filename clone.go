// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

import (
	"fmt"
)

type cloneOpts struct {
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

	Branch
	Insecure
	NoUpdate
	Pull
	Rev
	RemoteCmd
	Ssh
	Uncompressed
	UpdateRev

	Debug
	Profile
	Time
	Traceback
}

func (cmdOpts *cloneOpts) String() string {
	return fmt.Sprintf(
		"cloneCmd = {\n    Debug: (%T) %t\n    Profile: (%T) %t\n"+
			"   Time: (%T) %t\n    Traceback: (%T) %t\n}\n",
		cmdOpts.Debug, cmdOpts.Debug, cmdOpts.Profile, cmdOpts.Profile,
		cmdOpts.Time, cmdOpts.Time, cmdOpts.Traceback, cmdOpts.Traceback)
}

// Clone provides the 'hg clone' command.
func (hgcl *HgClient) Clone(source string, dest string, opts ...optionAdder) error {
	hgcmd, err := hgcl.buildCommand("clone", new(cloneOpts), opts, []string{source, dest})
	if err != nil {
		return err
	}
	_, err = hgcl.runcommand(hgcmd)
	return err
}
