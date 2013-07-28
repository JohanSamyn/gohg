// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type addremoveOpts struct {
	globalOpts

	DryRun
	Exclude
	Include
	// Mq
	Similarity

	debugOpts
}

func (cmdOpts *addremoveOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewAddRemoveCmd(opts []HgOption, files []string) HgCmd {
	cmd, _ := NewHgCmd("addremove", opts, files, new(addremoveOpts))
	return *cmd
}

func (hgcl *HgClient) AddRemove(opts []HgOption, files []string) ([]byte, error) {
	cmd := NewAddRemoveCmd(opts, files)
	return cmd.Exec(hgcl)
}
