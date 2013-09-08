// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

// Package gohg is a Go client library for using the Mercurial dvcs
// via it's Command Server.
//
// For Mercurial see: http://mercurial.selenic.com.
//
// For the Hg Command Server see: http://mercurial.selenic.com/wiki/CommandServer.
package gohg

import (
	"fmt"
	"strings"
	"time"
)

// HgCmd is the type through which you can create and execute Mercurial commands.
type HgCmd struct {
	Name    string
	Options []HgOption
	Params  []string

	// I keep this field private to make prohibit tampering with the series
	// of options that is valid for a command.
	cmdOpts interface{}

	cmd []string
}

type Cset struct {
	Rev    int
	Node   string
	Tags   []string
	Branch string
	Author string
	Time   time.Time // always as UTC
	Desc   string
	Patch  []string
}

// NewHgCmd creates a new HgCmd instance for working with Mercurial commands.
func NewHgCmd(name string, opts []HgOption, params []string, cmdopts interface{}) (*HgCmd, error) {
	if name == "" {
		return nil, fmt.Errorf("give a name for the command")
	}
	hgcmd := new(HgCmd)
	hgcmd.Name = name

	if opts != nil {
		hgcmd.Options = opts
	}
	if params != nil {
		hgcmd.Params = params
	}
	if cmdopts != nil {
		hgcmd.cmdOpts = cmdopts
	}
	return hgcmd, nil
}

func (hc *HgCmd) SetOptions(opts []HgOption) {
	// Checking for double opts is a bit useless, as some options can indeed
	// be passed more than once to a Hg command.
	hc.Options = append(hc.Options, opts...)
}

func (hc *HgCmd) SetParams(params []string) {
	hc.Params = append(hc.Params, params...)
}

// Exec builds the commandline with the data in the HgCmd instance,
// and then passes the command to the Mercurial Command Server for execution,
// returning the result or an error.
func (hc *HgCmd) Exec(hgcl *HgClient) ([]byte, error) {
	if hc.Name == "" {
		return nil, fmt.Errorf("HgCmd.Exec(): no command name specified")
	}
	var err error
	hc.cmd, err = hgcl.buildCommand(hc)
	if err != nil {
		return nil, err
	}
	return hgcl.runcommand(hc.cmd)
}

// CmdLine gives you the exact commandline as it will be passed to Mercurial,
// generated with the data in the HgCmd instance. Handy for logging or showing
// in a GUI.
func (hc *HgCmd) CmdLine(hgcl *HgClient) (string, error) {
	var err error
	hc.cmd, err = hgcl.buildCommand(hc)
	if err != nil {
		return "", err
	}
	return strings.Join(hc.cmd, " "), nil
}
