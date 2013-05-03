// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

type statusOpts struct {
	Config
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

	Added
	All
	Change
	Clean
	Copies
	Deleted
	Ignored
	Modified
	NoStatus
	Print0
	Removed
	SubRepos
	Unknown

	Debug
	Profile
	Time
	Traceback
}

func (cmdOpts *statusOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func (hgcl *HgClient) Status(opts []Option, files []string) ([]byte, error) {
	cmd, _ := NewHgCmd("status", opts, files)
	return cmd.Exec(hgcl)
}
