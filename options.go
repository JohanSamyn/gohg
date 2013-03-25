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
	Active        bool   // -a --active
	All           bool   //    --all
	Closed        bool   // -c --closed
	Cwd           string //    --cwd
	Debug         bool   //    --debug
	Destpath      string // no equivalent Hg option, used by Init()
	Dryrun        bool   // -n --dry-run (<->ShowNum)
	Exclude       string // -X --exclude
	Git           bool   // -g --git
	Include       string // -I --include
	Insecure      bool   //    --insecure
	Limit         int    // -l --limit
	Mq            bool   //    --mq
	Patch         bool   // -p --patch
	Profile       bool   //    --profile
	Quiet         bool   // -q --quiet
	Remote        bool   //    --remote
	RemoteCmd     string //    --remotecmd
	Repository    string // -R --repository
	Rev           string // -r --rev REV
	ShowBookmarks bool   // -B --bookmarks
	ShowBranch    bool   // -b --branch
	ShowId        bool   // -i --id
	ShowNum       bool   // -n --num (<->Dryrun)
	ShowTags      bool   // -t --tags (<->Topo)
	Ssh           string // -e --ssh
	Style         string //    --style
	Subrepos      bool   // -S --subrepos
	Template      string //    --template
	Topo          bool   // -t --topo (<->Tags)
	Traceback     bool   //    --traceback
	Verbose       bool   // -v --verbose
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

func (o Active) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "Active", "--active", i, hgcmd)
}

func (o All) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "All", "--all", i, hgcmd)
}

func (o Closed) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "Closed", "--closed", i, hgcmd)
}

func (o Cwd) addOption(i interface{}, hgcmd *[]string) error {
	return addStringOpt(string(o), "Cwd", "--cwd", i, hgcmd)
}

func (o Debug) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "Debug", "--debug", i, hgcmd)
}

func (o Destpath) addOption(i interface{}, hgcmd *[]string) error {
	return addStringOpt(string(o), "Destpath", "", i, hgcmd)
}

func (o Dryrun) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "Dryrun", "--dry-run", i, hgcmd)
}

func (o Exclude) addOption(i interface{}, hgcmd *[]string) error {
	return addStringOpt(string(o), "Exclude", "--exclude", i, hgcmd)
}

func (o Git) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "Git", "--git", i, hgcmd)
}

func (o Include) addOption(i interface{}, hgcmd *[]string) error {
	return addStringOpt(string(o), "Include", "--include", i, hgcmd)
}

func (o Insecure) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "Insecure", "--insecure", i, hgcmd)
}

func (o Limit) addOption(i interface{}, hgcmd *[]string) error {
	return addIntOpt(int(o), "Limit", "--limit", i, hgcmd)
}

func (o Mq) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "Mq", "--mq", i, hgcmd)
}

func (o Patch) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "Patch", "--patch", i, hgcmd)
}

func (o Profile) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "Profile", "--profile", i, hgcmd)
}

func (o Quiet) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "Quite", "-q", i, hgcmd)
}

func (o Remote) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "Remote", "--remote", i, hgcmd)
}

func (o RemoteCmd) addOption(i interface{}, hgcmd *[]string) error {
	return addStringOpt(string(o), "RemoteCmd", "--remotecmd", i, hgcmd)
}

func (o Repository) addOption(i interface{}, hgcmd *[]string) error {
	return addStringOpt(string(o), "Repository", "-R", i, hgcmd)
}

func (o Rev) addOption(i interface{}, hgcmd *[]string) error {
	return addStringOpt(string(o), "Rev", "--rev", i, hgcmd)
}

func (o ShowBookmarks) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "ShowBookmarks", "--bookmarks", i, hgcmd)
}

func (o ShowBranch) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "ShowBranch", "--branch", i, hgcmd)
}

func (o ShowId) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "ShowId", "--id", i, hgcmd)
}

func (o ShowNum) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "ShowNum", "--num", i, hgcmd)
}

func (o ShowTags) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "ShowTags", "--tags", i, hgcmd)
}

func (o Ssh) addOption(i interface{}, hgcmd *[]string) error {
	return addStringOpt(string(o), "Ssh", "--ssh", i, hgcmd)
}

func (o Style) addOption(i interface{}, hgcmd *[]string) error {
	return addStringOpt(string(o), "Style", "--style", i, hgcmd)
}

func (o Subrepos) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "Subrepos", "--subrepos", i, hgcmd)
}

func (o Template) addOption(i interface{}, hgcmd *[]string) error {
	return addStringOpt(string(o), "Template", "--template", i, hgcmd)
}

func (o Topo) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "Topo", "--topo", i, hgcmd)
}

func (o Traceback) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "Traceback", "--traceback", i, hgcmd)
}

func (o Verbose) addOption(i interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(o), "Verbose", "-v", i, hgcmd)
}

func addBoolOpt(opt bool, optname string, cmd string, i interface{}, hgcmd *[]string) error {
	f := reflect.ValueOf(i).Elem().FieldByName(optname)
	if f.IsValid() && f.CanSet() {
		f.SetBool(opt)
		if opt {
			*hgcmd = append(*hgcmd, cmd)
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), optname)
	}
	return nil
}

func addIntOpt(opt int, optname string, cmd string, i interface{}, hgcmd *[]string) error {
	f := reflect.ValueOf(i).Elem().FieldByName(optname)
	if f.IsValid() && f.CanSet() {
		f.SetInt(int64(opt))
		if opt > 0 {
			*hgcmd = append(*hgcmd, cmd)
			*hgcmd = append(*hgcmd, strconv.Itoa(opt))
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), optname)
	}
	return nil
}

func addStringOpt(opt string, optname string, cmd string, i interface{}, hgcmd *[]string) error {
	f := reflect.ValueOf(i).Elem().FieldByName(optname)
	if f.IsValid() && f.CanSet() {
		f.SetString(opt)
		if cmd != "" { // think about Destpath f.i.
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
