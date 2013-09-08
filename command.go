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

// Method buildCommand builds the command string to pass to the Hg CS.
//	cmd.Name: The name of the command ("log", "status", etc.).
//	cmd.cmdOpts: A type based on struct with the valid options for the command
//				 at hand. Also contains the (gohg) default value for each option
//				 for that particular command.
//				 Used to filter only the options supported by the command.
//	cmd.Options: The options passed by the caller.
//	cmd.Params:	Any filenames, paths or other that the command needs.
//				These are to be added as the last elements of the command.
func (cmd *HgCmd) buildCommand() (hgcmd []string, err error) {
	hgcmd = []string{cmd.Name}
	for _, o := range cmd.Options {
		err = o.addOption(cmd.cmdOpts, &hgcmd)
		// Silently skip invalid options.
		// if err != nil {
		// 	fmt.Logf("err = %s", err)
		// 	// Or work out some logging system for gohg, and write the error message inthere.
		// 	log.Printf("err = %s", err)
		// 	return nil, err
		// }
	}
	for _, p := range cmd.Params {
		if p != "" {
			hgcmd = append(hgcmd, string(p))
		}
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
	hc.cmd, err = hc.buildCommand()
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
	hc.cmd, err = hc.buildCommand()
	if err != nil {
		return "", err
	}
	return strings.Join(hc.cmd, " "), nil
}
