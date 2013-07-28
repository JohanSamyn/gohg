// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type tipOpts struct {
	globalOpts

	Patch
	Git
	// Mq
	Style
	Template

	debugOpts
}

func (cmdOpts *tipOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewTipCmd(opts []HgOption, params []string) HgCmd {
	cmd, _ := NewHgCmd("tip", opts, params, new(tipOpts))
	return *cmd
}

func (hgcl *HgClient) Tip(opts []HgOption, params []string) ([]byte, error) {
	cmd := NewTipCmd(opts, params)
	return cmd.Exec(hgcl)
}
