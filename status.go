// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

type statusOpts struct {
	Config
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

	Added
	All
	Change
	Clean
	Copies
	Deleted
	Ignored
	Modified
	NoStatus
	Print0
	Removed
	SubRepos
	Unknown

	Debug
	Profile
	Time
	Traceback
}

func (cmdOpts *statusOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func (hgcl *HgClient) Status(files []string, opts ...optionAdder) ([]byte, error) {
	cmdOpts := new(statusOpts)
	hgcmd, err := hgcl.buildCommand("status", cmdOpts, opts, files)
	if err != nil {
		return nil, err
	}
	return hgcl.runcommand(hgcmd)
}
