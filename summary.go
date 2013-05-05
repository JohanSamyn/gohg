// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

// import (
// 	"fmt"
// )

type summaryOpts struct {
	globalOpts

	Remote
	// Mq

	debugOpts
}

func (cmdOpts *summaryOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewSummaryCmd(opts []Option, params []string) HgCmd {
	cmd, _ := NewHgCmd("summary", opts, params, new(summaryOpts))
	return *cmd
}

func (hgcl *HgClient) Summary(opts []Option, params []string) ([]byte, error) {

	// See commit.go for how to obtain a value from cmd.cmdOpts.
	// fmt.Printf("%s", cmdOpts)

	cmd := NewSummaryCmd(opts, params)
	return cmd.Exec(hgcl)
}
