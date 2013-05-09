// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type removeOpts struct {
	globalOpts

	After
	Exclude
	Force
	Include
	// Mq

	debugOpts
}

func (cmdOpts *removeOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewRemoveCmd(opts []Option, files []string) HgCmd {
	cmd, _ := NewHgCmd("remove", opts, files, new(removeOpts))
	return *cmd
}

func (hgcl *HgClient) Remove(opts []Option, files []string) ([]byte, error) {
	cmd := NewRemoveCmd(opts, files)
	return cmd.Exec(hgcl)
}
