// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

// import (
// 	"fmt"
// )

type summaryOpts struct {
	Config
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

	Remote
	// Mq

	Debug
	Profile
	Time
	Traceback
}

func (cmdOpts *summaryOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

// Summary provides the 'hg summary' command.
func (hgcl *HgClient) Summary(opts ...optionAdder) ([]byte, error) {
	cmdOpts := new(summaryOpts)
	hgcmd, err := hgcl.buildCommand("summary", cmdOpts, opts, nil)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("%s", cmdOpts)
	return hgcl.runcommand(hgcmd)
}
