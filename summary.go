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

func (hgcl *HgClient) Summary(opts []Option, params []string) ([]byte, error) {

	// See commit.go for how to obtain a value from cmd.cmdOpts.
	// fmt.Printf("%s", cmdOpts)

	cmd, _ := NewHgCmd("summary", opts, params)
	return cmd.Exec(hgcl)
}
