// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type cloneOpts struct {
	globalOpts

	Branch
	Insecure
	NoUpdate
	Pull
	Rev
	RemoteCmd
	Ssh
	Uncompressed
	UpdateRev

	debugOpts
}

func (cmdOpts *cloneOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewCloneCmd(opts []HgOption, fromto []string) HgCmd {
	cmd, _ := NewHgCmd("clone", opts, fromto, new(cloneOpts))
	return *cmd
}

func (hgcl *HgClient) Clone(opts []HgOption, fromto []string) error {
	cmd := NewCloneCmd(opts, fromto)
	_, err := cmd.Exec(hgcl)
	return err
}
