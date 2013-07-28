// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type exportOpts struct {
	globalOpts

	Git
	// Mq
	NoDates
	Output
	Rev
	SwitchParent
	Text

	debugOpts
}

func (cmdOpts *exportOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewExportCmd(opts []HgOption, params []string) HgCmd {
	cmd, _ := NewHgCmd("export", opts, params, new(exportOpts))
	return *cmd
}

func (hgcl *HgClient) Export(opts []HgOption, params []string) ([]byte, error) {
	cmd := NewExportCmd(opts, params)
	return cmd.Exec(hgcl)
}
