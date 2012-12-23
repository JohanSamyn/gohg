// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type SummaryCmd struct {
	// define all necessary options/flags
	repository string
	cwd        string
	remote     bool
	mq         bool
}

func NewSummary() *SummaryCmd {
	// apply type defaults
	cmd := new(SummaryCmd)

	// apply application defaults
	cmd.repository = ""
	cmd.remote = false
	cmd.mq = false

	return cmd
}

// returning the struct from all methods allows working with a fluent interface

func (cmd *SummaryCmd) SetRepo(repo string) *SummaryCmd {
	cmd.repository = repo
	return cmd
}

func (cmd *SummaryCmd) SetRemote(r bool) *SummaryCmd {
	cmd.remote = r
	return cmd
}

func (cmd *SummaryCmd) SetMq(b bool) *SummaryCmd {
	cmd.mq = b
	return cmd
}

func (cmd *SummaryCmd) String() string {
	// print it out nicely
	return fmt.Sprintf(
		"SummaryCmd = {\n    repository: (%T) %q\n    remote: (%T) %t\n    mq: (%T) %t\n}\n",
		cmd.repository, cmd.repository, cmd.remote, cmd.remote, cmd.mq, cmd.mq)
}

// Summary provides the 'hg summary' command.
func (hgcl *HgClient) Summary(scmd *SummaryCmd) ([]byte, error) {
	opts := []string{"summary"}
	if scmd.repository != "" {
		opts = append(opts, "-R")
		opts = append(opts, scmd.repository)
	}
	if scmd.remote == true {
		opts = append(opts, "--remote")
	}
	if scmd.mq == true {
		opts = append(opts, "--mq")
	}
	// fmt.Printf("SummaryOpts %v\n", opts)

	// data, err := command(hgcl, "summary", opts)
	data, err := command(hgcl, opts)
	return data, err
}
