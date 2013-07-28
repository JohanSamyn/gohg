// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type mergeOpts struct {
	globalOpts

	Force
	// Mq
	Preview
	Rev
	Tool

	debugOpts
}

func (cmdOpts *mergeOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewMergeCmd(opts []HgOption, rev []string) HgCmd {
	cmd, _ := NewHgCmd("merge", opts, rev, new(mergeOpts))
	return *cmd
}

// Maybe for this one we could add some extra checking:
// - there can only be one rev in the slice
// - if rev is not nil then there should not be a Rev option

func (hgcl *HgClient) Merge(opts []HgOption, rev []string) ([]byte, error) {
	cmd := NewMergeCmd(opts, rev)
	return cmd.Exec(hgcl)
}
