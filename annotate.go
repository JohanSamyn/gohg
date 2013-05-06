// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type annotateOpts struct {
	globalOpts

	Changeset
	Date
	Exclude
	File
	IgnoreAllSpace
	IgnoreBlankLines
	IgnoreSpaceChange
	Include
	LineNumber
	NoFollow
	Number
	// Mq
	Rev
	Text
	User

	debugOpts
}

func (cmdOpts *annotateOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewAnnotateCmd(opts []Option, files []string) HgCmd {
	cmd, _ := NewHgCmd("annotate", opts, files, new(annotateOpts))
	return *cmd
}

func (hgcl *HgClient) Annotate(opts []Option, files []string) ([]byte, error) {
	cmd := NewAnnotateCmd(opts, files)
	return cmd.Exec(hgcl)
}
