// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

type manifestOpts struct {
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

func (cmdOpts *manifestOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func (hgcl *HgClient) Manifest(opts []Option, params []string) ([]byte, error) {
	cmd, _ := NewHgCmd("manifest", opts, params)
	return cmd.Exec(hgcl)
}
