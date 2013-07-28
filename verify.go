// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type verifyOpts struct {
	globalOpts
	// Cwd : makes it possible to verify another repo than hgcl.Reporoot()
	// Repository : makes it possible to verify another repo than hgcl.Reporoot()

	// Mq

	debugOpts
}

func (cmdOpts *verifyOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewVerifyCmd(opts []Option, params []string) HgCmd {
	cmd, _ := NewHgCmd("verify", opts, params, new(verifyOpts))
	return *cmd
}

func (hgcl *HgClient) Verify(opts []Option, params []string) ([]byte, error) {
	cmd := NewVerifyCmd(opts, params)
	return cmd.Exec(hgcl)
}
