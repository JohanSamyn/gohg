// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

// import (
// 	"fmt"
// )

type manifestOpts struct {
	Config
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

	All
	// Mq
	Rev

	Debug
	Profile
	Time
	Traceback
}

func (cmdOpts *manifestOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

// Manifest provides the 'hg manifest' command.
func (hgcl *HgClient) Manifest(opts ...optionAdder) ([]byte, error) {
	cmdOpts := new(manifestOpts)
	hgcmd, err := hgcl.buildCommand("manifest", cmdOpts, opts, nil)
	if err != nil {
		return nil, err
	}
	return hgcl.runcommand(hgcmd)
}
