// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var errstr = "command %s has no option %s"

// These are the options the Hg commands can take.

type (
	Active         bool   // -a --active
	All            bool   //    --all
	Closed         bool   // -c --closed
	Config         string //    --config
	Cwd            string //    --cwd
	Debug          bool   //    --debug
	Destpath       string // no equivalent Hg option, used by Init()
	Dryrun         bool   // -n --dry-run (<->ShowNum)
	Exclude        string // -X --exclude
	Git            bool   // -g --git
	Hidden         bool   //    --hidden
	Include        string // -I --include
	Insecure       bool   //    --insecure
	Limit          int    // -l --limit
	Mq             bool   //    --mq
	NonInteractive bool   // -y --noninteractive
	Patch          bool   // -p --patch
	Profile        bool   //    --profile
	Quiet          bool   // -q --quiet
	Remote         bool   //    --remote
	RemoteCmd      string //    --remotecmd
	Repository     string // -R --repository
	Rev            string // -r --rev REV
	ShowBookmarks  bool   // -B --bookmarks
	ShowBranch     bool   // -b --branch
	ShowId         bool   // -i --id
	ShowNum        bool   // -n --num (<->Dryrun)
	ShowTags       bool   // -t --tags (<->Topo)
	Ssh            string // -e --ssh
	Style          string //    --style
	Subrepos       bool   // -S --subrepos
	Template       string //    --template
	Time           bool   //    --time
	Topo           bool   // -t --topo (<->Tags)
	Traceback      bool   //    --traceback
	Verbose        bool   // -v --verbose
)

type optionAdder interface {
	addOption(interface{}, *[]string) error
}

// An addOption method does 2 things:
// 1. It checks if the option is valid for a given command,
//    by checking its presence in that commands option struct (logOpts, etc.).
// 2. If the option is valid, it then adds the option and any values for it
//    to the command's commandline (a []string that is joined to a \0-separated
//    string later in method runInHg()).
// It also returns an error stating the option is invalid for the command,
// but these error messages are silently ignored for the moment.

func (opt Active) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Active", "--active", cmdOpts, hgcmd)
}

func (opt All) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "All", "--all", cmdOpts, hgcmd)
}

func (opt Closed) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Closed", "--closed", cmdOpts, hgcmd)
}

func (opt Config) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Config", "--config", cmdOpts, hgcmd)
}

func (opt Cwd) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Cwd", "--cwd", cmdOpts, hgcmd)
}

func (opt Debug) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Debug", "--debug", cmdOpts, hgcmd)
}

func (opt Destpath) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Destpath", "", cmdOpts, hgcmd)
}

func (opt Dryrun) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Dryrun", "--dry-run", cmdOpts, hgcmd)
}

func (opt Exclude) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Exclude", "--exclude", cmdOpts, hgcmd)
}

func (opt Git) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Git", "--git", cmdOpts, hgcmd)
}

func (opt Hidden) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Hidden", "--hidden", cmdOpts, hgcmd)
}

func (opt Include) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Include", "--include", cmdOpts, hgcmd)
}

func (opt Insecure) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Insecure", "--insecure", cmdOpts, hgcmd)
}

func (opt Limit) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addIntOpt(int(opt), "Limit", "--limit", cmdOpts, hgcmd)
}

func (opt Mq) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Mq", "--mq", cmdOpts, hgcmd)
}

func (opt NonInteractive) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "NonInteractive", "--noninteractive", cmdOpts, hgcmd)
}

func (opt Patch) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Patch", "--patch", cmdOpts, hgcmd)
}

func (opt Profile) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Profile", "--profile", cmdOpts, hgcmd)
}

func (opt Quiet) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Quite", "-q", cmdOpts, hgcmd)
}

func (opt Remote) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Remote", "--remote", cmdOpts, hgcmd)
}

func (opt RemoteCmd) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "RemoteCmd", "--remotecmd", cmdOpts, hgcmd)
}

func (opt Repository) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Repository", "-R", cmdOpts, hgcmd)
}

func (opt Rev) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Rev", "--rev", cmdOpts, hgcmd)
}

func (opt ShowBookmarks) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "ShowBookmarks", "--bookmarks", cmdOpts, hgcmd)
}

func (opt ShowBranch) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "ShowBranch", "--branch", cmdOpts, hgcmd)
}

func (opt ShowId) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "ShowId", "--id", cmdOpts, hgcmd)
}

func (opt ShowNum) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "ShowNum", "--num", cmdOpts, hgcmd)
}

func (opt ShowTags) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "ShowTags", "--tags", cmdOpts, hgcmd)
}

func (opt Ssh) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Ssh", "--ssh", cmdOpts, hgcmd)
}

func (opt Style) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Style", "--style", cmdOpts, hgcmd)
}

func (opt Subrepos) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Subrepos", "--subrepos", cmdOpts, hgcmd)
}

func (opt Template) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Template", "--template", cmdOpts, hgcmd)
}

func (opt Time) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Time", "--time", cmdOpts, hgcmd)
}

func (opt Topo) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Topo", "--topo", cmdOpts, hgcmd)
}

func (opt Traceback) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Traceback", "--traceback", cmdOpts, hgcmd)
}

func (opt Verbose) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Verbose", "-v", cmdOpts, hgcmd)
}

func addBoolOpt(opt bool, optname string, cmd string, cmdOpts interface{}, hgcmd *[]string) error {
	fld := reflect.ValueOf(cmdOpts).Elem().FieldByName(optname)
	if fld.IsValid() && fld.CanSet() {
		fld.SetBool(opt) // add the value to the <cmd>Opts struct
		if opt {
			*hgcmd = append(*hgcmd, cmd)
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), optname)
	}
	return nil
}

func addIntOpt(opt int, optname string, cmd string, cmdOpts interface{}, hgcmd *[]string) error {
	fld := reflect.ValueOf(cmdOpts).Elem().FieldByName(optname)
	if fld.IsValid() && fld.CanSet() {
		fld.SetInt(int64(opt)) // add the value to the <cmd>Opts struct
		if opt > 0 {
			*hgcmd = append(*hgcmd, cmd)
			*hgcmd = append(*hgcmd, strconv.Itoa(opt))
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), optname)
	}
	return nil
}

func addStringOpt(opt string, optname string, cmd string, cmdOpts interface{}, hgcmd *[]string) error {
	fld := reflect.ValueOf(cmdOpts).Elem().FieldByName(optname)
	if fld.IsValid() && fld.CanSet() {
		fld.SetString(opt) // add the value to the <cmd>Opts struct
		if cmd != "" {     // think of Destpath f.i.
			*hgcmd = append(*hgcmd, cmd)
		}
		if opt != "" {
			*hgcmd = append(*hgcmd, opt)
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), optname)
	}
	return nil
}
