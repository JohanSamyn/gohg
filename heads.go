// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type headsOpts struct {
	globalOpts

	Closed
	// Mq
	Rev
	Style
	Template
	Topo

	debugOpts
}

func (cmdOpts *headsOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewHeadsCmd(opts []Option, revs []string) HgCmd {
	cmd, _ := NewHgCmd("heads", opts, revs, new(headsOpts))
	return *cmd
}

func (hgcl *HgClient) Heads(opts []Option, revs []string) ([]byte, error) {
	cmd := NewHeadsCmd(opts, revs)
	return cmd.Exec(hgcl)
}
