// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

type logOpts struct {
	Config
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

	Branch
	Copies
	Date
	Follow
	Git
	Graph
	Keyword
	Limit
	// Mq
	NoMerges
	Patch
	Prune
	Removed
	Rev
	Stat
	Style
	Template
	User

	Debug
	Profile
	Time
	Traceback
}

func (cmdOpts *logOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewLogCmd(opts []Option, files []string) HgCmd {
	cmd, _ := NewHgCmd("log", opts, files, new(logOpts))
	return *cmd
}

func (hgcl *HgClient) Log(opts []Option, files []string) ([]byte, error) {
	cmd := NewLogCmd(opts, files)
	return cmd.Exec(hgcl)
}
