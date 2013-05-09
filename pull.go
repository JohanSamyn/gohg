// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type pullOpts struct {
	globalOpts

	Bookmark
	Branch
	Force
	Insecure
	// Mq
	Rebase
	RemoteCmd
	Rev
	Ssh
	Update

	debugOpts
}

func (cmdOpts *pullOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewPullCmd(opts []Option, source []string) HgCmd {
	cmd, _ := NewHgCmd("pull", opts, source, new(pullOpts))
	return *cmd
}

func (hgcl *HgClient) Pull(opts []Option, source []string) ([]byte, error) {
	cmd := NewPullCmd(opts, source)
	return cmd.Exec(hgcl)
}
