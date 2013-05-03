// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

type verifyOpts struct {
	Config
	Cwd // makes it possible to verify another repo than hgcl.Reporoot()
	Hidden
	NonInteractive
	Quiet
	Repository // makes it possible to verify another repo than hgcl.Reporoot()
	Verbose

	// Mq

	Debug
	Profile
	Time
	Traceback
}

func (cmdOpts *verifyOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func (hgcl *HgClient) Verify(opts []Option, params []string) ([]byte, error) {
	cmd, _ := NewHgCmd("verify", opts, params)
	return cmd.Exec(hgcl)
}
