// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

// import (
// 	"fmt"
// )

type verifyOpts struct {
	Config
	Cwd // makes it possible to verify another repo than hgcl.Reporoot()
	Hidden
	NonInteractive
	Quiet
	Repository // makes it possible to verify another repo than hgcl.Reporoot()
	Verbose

	// Mq

	Debug
	Profile
	Time
	Traceback
}

func (cmdOpts *verifyOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

// Verify provides the 'hg verify' command.
func (hgcl *HgClient) Verify(opts ...optionAdder) ([]byte, error) {
	cmdOpts := new(verifyOpts)
	hgcmd, err := hgcl.buildCommand("verify", cmdOpts, opts, nil)
	if err != nil {
		return nil, err
	}
	return hgcl.runcommand(hgcmd)
}
