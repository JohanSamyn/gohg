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

// func (hgcl *HgClient) Id(source string, opts ...optionAdder) ([]byte, error) {
// 	return hgcl.Identify(source, opts...)
// }

func (hgcl *HgClient) Identify(opts []Option, source []string) ([]byte, error) {
	cmdOpts := new(identifyOpts)
	hgcmd, err := hgcl.buildCommand("identify", cmdOpts, opts, source)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("%s", cmdOpts)
	return hgcl.runcommand(hgcmd)
}
