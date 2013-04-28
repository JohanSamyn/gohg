// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

type branchesOpts struct {
	Config
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

	Active
	Closed
	// Mq

	Debug
	Profile
	Time
	Traceback
}

func (cmdOpts *branchesOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func (hgcl *HgClient) Branches(opts []Option, params []string) ([]byte, error) {
	cmdOpts := new(branchesOpts)
	hgcmd, err := hgcl.buildCommand("branches", cmdOpts, opts, params)
	if err != nil {
		return nil, err
	}
	return hgcl.runcommand(hgcmd)
}
