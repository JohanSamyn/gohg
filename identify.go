// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

// import (
// 	"fmt"
// )

type identifyOpts struct {
	Config
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

	Insecure
	// Mq
	RemoteCmd
	Rev
	Bookmarks
	Branch
	Id
	Num
	Tags
	Ssh

	Debug
	Profile
	Time
	Traceback
}

func (cmdOpts *identifyOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewIdentifyCmd(opts []Option, source []string) HgCmd {
	cmd, _ := NewHgCmd("identify", opts, source, new(identifyOpts))
	return *cmd
}

// func (hgcl *HgClient) Id(source string, opts ...optionAdder) ([]byte, error) {
// 	return hgcl.Identify(source, opts...)
// }

func (hgcl *HgClient) Identify(opts []Option, source []string) ([]byte, error) {
	cmd := NewIdentifyCmd(opts, source)
	// cmd.cmd, _ = hgcl.buildCommand(cmd)
	// fmt.Printf("opts -> %s\ncmd.Options -> %s\ncmd.cmdOpts -> %s\n", opts, cmd.Options, cmd.cmdOpts)
	// fmt.Printf("cmd.cmd = %s\n", cmd.cmd)
	return cmd.Exec(hgcl)
}
