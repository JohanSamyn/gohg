// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type diffOpts struct {
	globalOpts

	Change
	Exclude
	Git
	IgnoreAllSpace
	IgnoreBlankLines
	IgnoreSpaceChange
	Include
	// Mq
	NoDates
	Rev
	Reverse
	ShowFunction
	Stat
	SubRepos
	Text
	Unified

	debugOpts
}

func (cmdOpts *diffOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewDiffCmd(opts []HgOption, files []string) HgCmd {
	cmd, _ := NewHgCmd("diff", opts, files, new(diffOpts))
	return *cmd
}

func (hgcl *HgClient) Diff(opts []HgOption, files []string) ([]byte, error) {
	cmd := NewDiffCmd(opts, files)
	return cmd.Exec(hgcl)
}
