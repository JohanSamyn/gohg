// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

type tipOpts struct {
	Config
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

	Patch
	Git
	// Mq
	Style
	Template

	Debug
	Profile
	Time
	Traceback
}

func (cmdOpts *tipOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func (hgcl *HgClient) Tip(opts []Option, params []string) ([]byte, error) {
	cmd, _ := NewHgCmd("tip", opts, params)
	return cmd.Exec(hgcl)
}
