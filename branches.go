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

func NewBranchesCmd(opts []Option, files []string) HgCmd {
	cmd, _ := NewHgCmd("branches", opts, files, new(branchesOpts))
	return *cmd
}

func (hgcl *HgClient) Branches(opts []Option, params []string) ([]byte, error) {
	cmd := NewBranchesCmd(opts, params)
	return cmd.Exec(hgcl)
}
