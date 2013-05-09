// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type forgetOpts struct {
	globalOpts

	Exclude
	Include
	// Mq

	debugOpts
}

func (cmdOpts *forgetOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewForgetCmd(opts []Option, files []string) HgCmd {
	cmd, _ := NewHgCmd("forget", opts, files, new(forgetOpts))
	return *cmd
}

func (hgcl *HgClient) Forget(opts []Option, files []string) ([]byte, error) {
	cmd := NewForgetCmd(opts, files)
	return cmd.Exec(hgcl)
}
