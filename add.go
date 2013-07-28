// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type addOpts struct {
	globalOpts

	DryRun
	Exclude
	Include
	// Mq
	SubRepos

	debugOpts
}

func (cmdOpts *addOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewAddCmd(opts []HgOption, files []string) HgCmd {
	cmd, _ := NewHgCmd("add", opts, files, new(addOpts))
	return *cmd
}

func (hgcl *HgClient) Add(opts []HgOption, files []string) ([]byte, error) {
	cmd := NewAddCmd(opts, files)
	return cmd.Exec(hgcl)
}
