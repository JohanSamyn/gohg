// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type pushOpts struct {
	globalOpts

	AllTasks
	CompletedTasks
	Bookmark
	Branch
	Force
	Insecure
	// Mq
	NewBranch
	RemoteCmd
	Rev
	Ssh

	debugOpts
}

func (cmdOpts *pushOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewPushCmd(opts []Option, dest []string) HgCmd {
	cmd, _ := NewHgCmd("push", opts, dest, new(pushOpts))
	return *cmd
}

func (hgcl *HgClient) Push(opts []Option, dest []string) ([]byte, error) {
	cmd := NewPushCmd(opts, dest)
	return cmd.Exec(hgcl)
}
