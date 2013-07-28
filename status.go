// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type statusOpts struct {
	globalOpts

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

	debugOpts
}

func (cmdOpts *statusOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewStatusCmd(opts []HgOption, files []string) HgCmd {
	cmd, _ := NewHgCmd("status", opts, files, new(statusOpts))
	return *cmd
}

func (hgcl *HgClient) Status(opts []HgOption, files []string) ([]byte, error) {
	cmd := NewStatusCmd(opts, files)
	return cmd.Exec(hgcl)
}
