// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

type serveOpts struct {
	globalOpts

	AccessLog
	Address
	Certificate
	CmdServer
	Daemon
	DaemonPipefds
	ErrorLog
	Ipv6
	// Mq
	Name
	PidFile
	Port
	Prefix
	Stdio
	Style
	Templates
	WebConf

	debugOpts
}

func (cmdOpts *serveOpts) String() string {
	return sprintfOpts(*cmdOpts)
}

func NewServeCmd(opts []Option, params []string) HgCmd {
	// We can safely ignore any passed-in params, as 'serve' does not take any.
	cmd, _ := NewHgCmd("serve", opts, nil, new(serveOpts))
	return *cmd
}

func (hgcl *HgClient) Serve(opts []Option, params []string) ([]byte, error) {
	cmd := NewServeCmd(opts, params)
	return cmd.Exec(hgcl)
}
