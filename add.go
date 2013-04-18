// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

type addOpts struct {
	Config
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

	DryRun
	Exclude
	Include
	// Mq
	SubRepos

	Debug
	Profile
	Time
	Traceback
}

func (cmdOpts *addOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func (hgcl *HgClient) Add(files []string, opts ...optionAdder) ([]byte, error) {
	cmdOpts := new(addOpts)
	hgcmd, err := hgcl.buildCommand("add", cmdOpts, opts, files)
	if err != nil {
		return nil, err
	}
	return hgcl.runcommand(hgcmd)
}
