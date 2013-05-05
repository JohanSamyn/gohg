// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

type tagsOpts struct {
	globalOpts

	All
	// Mq
	Rev

	debugOpts
}

func (cmdOpts *tagsOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewTagCmd(opts []Option, params []string) HgCmd {
	cmd, _ := NewHgCmd("tags", opts, params, new(tagsOpts))
	return *cmd
}

func (hgcl *HgClient) Tags(opts []Option, params []string) ([]byte, error) {
	cmd := NewTagCmd(opts, params)
	return cmd.Exec(hgcl)
}
