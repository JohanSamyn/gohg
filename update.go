// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type updateOpts struct {
	globalOpts

	Check
	Clean
	Date
	// Mq
	Rev

	debugOpts
}

func (cmdOpts *updateOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewUpdateCmd(opts []Option, files []string) HgCmd {
	cmd, _ := NewHgCmd("update", opts, files, new(updateOpts))
	return *cmd
}

func (hgcl *HgClient) Update(opts []Option, params []string) ([]byte, error) {
	cmd := NewUpdateCmd(opts, params)
	return cmd.Exec(hgcl)
}
