// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type showconfigOpts struct {
	globalOpts

	// Mq
	Untrusted

	debugOpts
}

func (cmdOpts *showconfigOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewShowConfigCmd(opts []Option, configitems []string) HgCmd {
	cmd, _ := NewHgCmd("showconfig", opts, configitems, new(showconfigOpts))
	return *cmd
}

// func (hgcl *HgClient) DebugConfig(configitems []string, opts ...optionAdder) ([]byte, error) {
// 	return hgcl.ShowConfig(configitems, opts...)
// }

func (hgcl *HgClient) ShowConfig(opts []Option, configitems []string) ([]byte, error) {
	cmd := NewShowConfigCmd(opts, configitems)
	return cmd.Exec(hgcl)
}
