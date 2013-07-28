// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type archiveOpts struct {
	globalOpts

	Exclude
	Include
	// Mq
	NoDecode
	Prefix
	Rev
	SubRepos
	Type

	debugOpts
}

func (cmdOpts *archiveOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewArchiveCmd(opts []HgOption, dest []string) HgCmd {
	cmd, _ := NewHgCmd("archive", opts, dest, new(archiveOpts))
	return *cmd
}

func (hgcl *HgClient) Archive(opts []HgOption, dest []string) ([]byte, error) {
	cmd := NewArchiveCmd(opts, dest)
	return cmd.Exec(hgcl)
}
