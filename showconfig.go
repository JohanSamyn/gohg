// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type showconfigOpts struct {
	Config
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose

	// Mq
	Untrusted

	Debug
	Profile
	Time
	Traceback
}

func (cmdOpts *showconfigOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func (hgcl *HgClient) DebugConfig(sections []string, opts ...optionAdder) ([]byte, error) {
	return hgcl.ShowConfig(sections, opts...)
}

// ShowConfig provides the 'hg showconfig' command.
func (hgcl *HgClient) ShowConfig(sections []string, opts ...optionAdder) ([]byte, error) {
	cmdOpts := new(showconfigOpts)

	hgcmd, err := hgcl.buildCommand("showconfig", cmdOpts, opts, sections)
	if err != nil {
		return nil, err
	}
	return hgcl.runcommand(hgcmd)
}
