// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

type addOpts struct {
	Config
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

	DryRun
	Exclude
	Include
	// Mq
	SubRepos

	Debug
	Profile
	Time
	Traceback
}

func (cmdOpts *addOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewAddCmd(opts []Option, files []string) HgCmd {
	cmd, _ := NewHgCmd("add", opts, files, new(addOpts))
	return *cmd
}

func (hgcl *HgClient) Add(opts []Option, files []string) ([]byte, error) {
	cmd := NewAddCmd(opts, files)
	return cmd.Exec(hgcl)
}
