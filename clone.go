// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type cloneOpts struct {
	Config
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

	Branch
	Insecure
	NoUpdate
	Pull
	Rev
	RemoteCmd
	Ssh
	Uncompressed
	UpdateRev

	Debug
	Profile
	Time
	Traceback
}

func (cmdOpts *cloneOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func (hgcl *HgClient) Clone(opts []Option, fromto []string) error {
	cmd, _ := NewHgCmd("clone", opts, fromto)
	_, err := cmd.Exec(hgcl)
	return err
}
