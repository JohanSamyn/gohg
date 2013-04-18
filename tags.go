// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

// import (
// 	"fmt"
// )

type tagsOpts struct {
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

func (cmdOpts *tagsOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

// Tags provides the 'hg tags' command.
func (hgcl *HgClient) Tags(opts ...optionAdder) ([]byte, error) {
	cmdOpts := new(tagsOpts)
	hgcmd, err := hgcl.buildCommand("tags", cmdOpts, opts, nil)
	if err != nil {
		return nil, err
	}
	return hgcl.runcommand(hgcmd)
}
