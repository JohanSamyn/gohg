// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

type manifestOpts struct {
	globalOpts

	All
	// Mq
	Rev

	debugOpts
}

func (cmdOpts *manifestOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewManifestCmd(opts []Option, files []string) HgCmd {
	cmd, _ := NewHgCmd("manifest", opts, files, new(manifestOpts))
	return *cmd
}

func (hgcl *HgClient) Manifest(opts []Option, params []string) ([]byte, error) {
	cmd := NewManifestCmd(opts, params)
	return cmd.Exec(hgcl)
}
