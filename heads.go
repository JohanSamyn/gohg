// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

type headsOpts struct {
	Config
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

	Closed
	// Mq
	Rev
	Style
	Template
	Topo

	Debug
	Profile
	Time
	Traceback
}

func (cmdOpts *headsOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func (hgcl *HgClient) Heads(opts []Option, revs []string) ([]byte, error) {
	cmd, _ := NewHgCmd("heads", opts, revs)
	return cmd.Exec(hgcl)
}
