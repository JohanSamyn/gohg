// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

// import (
// 	"fmt"
// )

type cloneOpts struct {
	Config
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
	return sprintfOpts(*cmdOpts)
}

// Clone provides the 'hg clone' command.
func (hgcl *HgClient) Clone(source string, dest string, opts ...optionAdder) error {
	cmdOpts := new(cloneOpts)
	hgcmd, err := hgcl.buildCommand("clone", cmdOpts, opts, []string{source, dest})
	if err != nil {
		return err
	}
	_, err = hgcl.runcommand(hgcmd)
	return err
}
