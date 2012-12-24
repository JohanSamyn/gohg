// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
)

type SummaryOpt interface {
	AddSummaryOpt(*SummaryCmd)
}

func (o O_repository) AddSummaryOpt(s *SummaryCmd) {
	s.O_repository = string(o)
}

func (o O_remote) AddSummaryOpt(s *SummaryCmd) {
	s.O_remote = bool(o)
}

func (o O_mq) AddSummaryOpt(s *SummaryCmd) {
	s.O_mq = bool(o)
}

func (o O_debug) AddSummaryOpt(s *SummaryCmd) {
	s.O_debug = bool(o)
}

func (o O_traceback) AddSummaryOpt(s *SummaryCmd) {
	s.O_traceback = bool(o)
}

func (o O_profile) AddSummaryOpt(s *SummaryCmd) {
	s.O_profile = bool(o)
}

type SummaryCmd struct {
	// define all necessary options/flags
	O_repository string
	O_remote     bool
	O_mq         bool
	O_debug      bool
	O_traceback  bool
	O_profile    bool
}

func NewSummary(opts ...SummaryOpt) *SummaryCmd {
	// applies type defaults
	cmd := new(SummaryCmd)

	// apply application defaults
	cmd.O_repository = ""
	cmd.O_remote = false
	cmd.O_mq = false
	cmd.O_debug = false
	cmd.O_traceback = false
	cmd.O_profile = false

	// apply option values given by the user
	for _, o := range opts {
		o.AddSummaryOpt(cmd)
	}

	return cmd
}

// returning the struct from all methods allows working with a fluent interface

func (cmd *SummaryCmd) SetRepo(repo string) *SummaryCmd {
	cmd.O_repository = repo
	return cmd
}

func (cmd *SummaryCmd) SetRemote(r bool) *SummaryCmd {
	cmd.O_remote = r
	return cmd
}

func (cmd *SummaryCmd) SetMq(b bool) *SummaryCmd {
	cmd.O_mq = b
	return cmd
}

func (cmd *SummaryCmd) SetDebug(b bool) *SummaryCmd {
	cmd.O_debug = b
	return cmd
}

func (cmd *SummaryCmd) SetTraceback(b bool) *SummaryCmd {
	cmd.O_traceback = b
	return cmd
}

func (cmd *SummaryCmd) SetProfile(b bool) *SummaryCmd {
	cmd.O_profile = b
	return cmd
}

func (cmd *SummaryCmd) String() string {
	// print it out nicely
	return fmt.Sprintf(
		"SummaryCmd = {\n    repository: (%T) %q\n    remote: (%T) %t\n    mq: (%T) %t\n"+
			"    debug: (%t) %t\n    traceback: (%T) %t\n    profile: (%T) %t\n}\n",
		cmd.O_repository, cmd.O_repository, cmd.O_remote, cmd.O_remote, cmd.O_mq, cmd.O_mq,
		cmd.O_debug, cmd.O_debug, cmd.O_traceback, cmd.O_traceback, cmd.O_profile, cmd.O_profile)
}

// Summary provides the 'hg summary' command.
func (hgcl *HgClient) Summary(scmd *SummaryCmd) ([]byte, error) {
	hgcmd := []string{"summary"}
	if scmd.O_repository != "" {
		hgcmd = append(hgcmd, "-R")
		hgcmd = append(hgcmd, scmd.O_repository)
	}
	if scmd.O_remote == true {
		hgcmd = append(hgcmd, "--remote")
	}
	if scmd.O_mq == true {
		hgcmd = append(hgcmd, "--mq")
	}
	if scmd.O_debug == true {
		hgcmd = append(hgcmd, "--debug")
	}
	if scmd.O_traceback == true {
		hgcmd = append(hgcmd, "--traceback")
	}
	if scmd.O_profile == true {
		hgcmd = append(hgcmd, "--profile")
	}

	data, err := command(hgcl, hgcmd)
	return data, err
}
