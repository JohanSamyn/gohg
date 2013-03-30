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
	// the global options

	Cwd            string //    --cwd
	Hidden         bool   //    --hidden
	NonInteractive bool   // -y --noninteractive
	Quiet          bool   // -q --quiet
	Repository     string // -R --repository
	Verbose        bool   // -v --verbose

	// command-specific options

	Active    bool   // -a --active
	Added     bool   // -a --added
	All       bool   //    --all
	Bookmarks bool   // -B --bookmarks
	Branch    bool   // -b --branch
	Change    bool   //    --change
	Clean     bool   // -c --clean
	Closed    bool   // -c --closed
	Config    string //    --config
	Copies    bool   // -C --copies
	Date      string // -d --date
	Deleted   bool   // -d --deleted
	DryRun    bool   // -n --dry-run (<->ShowNum)
	Exclude   string // -X --exclude
	Follow    bool   // -f --follow
	Git       bool   // -g --git
	Graph     bool   // -G --graph
	Id        bool   // -i --id
	Ignored   bool   // -i --ignored
	Include   string // -I --include
	Insecure  bool   //    --insecure
	Keyword   string // -k --keyword
	Limit     int    // -l --limit
	Modified  bool   // -m --modified
	Mq        bool   //    --mq
	NoMerges  bool   // -M --no-merges
	NoStatus  bool   // -n --no-status
	Num       bool   // -n --num (<->Dryrun)
	Patch     bool   // -p --patch
	Print0    bool   // -0 --print0
	Prune     bool   // -P --prune
	Remote    bool   //    --remote
	RemoteCmd string //    --remotecmd
	Removed   bool   // (-r) --removed
	Rev       string // -r --rev REV
	Ssh       string // -e --ssh
	Stat      bool   //    --stat
	Style     string //    --style
	SubRepos  bool   // -S --subrepos
	Tags      bool   // -t --tags (<->Topo)
	Template  string //    --template
	Topo      bool   // -t --topo (<->Tags)
	Unknown   bool   // -u --unknown
	User      string // -u --user

	// debugging and profiling options

	Debug     bool //    --debug
	Profile   bool //    --profile
	Time      bool //    --time
	Traceback bool //    --traceback
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

func (opt Added) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Added", "--added", cmdOpts, hgcmd)
}

func (opt All) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "All", "--all", cmdOpts, hgcmd)
}

func (opt Bookmarks) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Bookmarks", "--bookmarks", cmdOpts, hgcmd)
}

func (opt Branch) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Branch", "--branch", cmdOpts, hgcmd)
}

func (opt Change) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Change", "--change", cmdOpts, hgcmd)
}

func (opt Clean) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Clean", "--clean", cmdOpts, hgcmd)
}

func (opt Closed) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Closed", "--closed", cmdOpts, hgcmd)
}

func (opt Config) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Config", "--config", cmdOpts, hgcmd)
}

func (opt Copies) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Copies", "--copies", cmdOpts, hgcmd)
}

func (opt Cwd) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Cwd", "--cwd", cmdOpts, hgcmd)
}

func (opt Date) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Date", "--date", cmdOpts, hgcmd)
}

func (opt Debug) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Debug", "--debug", cmdOpts, hgcmd)
}

func (opt Deleted) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Deleted", "--deleted", cmdOpts, hgcmd)
}

func (opt DryRun) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "DryRun", "--dry-run", cmdOpts, hgcmd)
}

func (opt Follow) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Fowwol", "--follow", cmdOpts, hgcmd)
}

func (opt Exclude) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Exclude", "--exclude", cmdOpts, hgcmd)
}

func (opt Git) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Git", "--git", cmdOpts, hgcmd)
}

func (opt Graph) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Graph", "--graph", cmdOpts, hgcmd)
}

func (opt Hidden) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Hidden", "--hidden", cmdOpts, hgcmd)
}

func (opt Id) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Id", "--id", cmdOpts, hgcmd)
}

func (opt Ignored) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Ignored", "--ignored", cmdOpts, hgcmd)
}

func (opt Include) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Include", "--include", cmdOpts, hgcmd)
}

func (opt Insecure) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Insecure", "--insecure", cmdOpts, hgcmd)
}

func (opt Keyword) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Keyword", "--keyword", cmdOpts, hgcmd)
}

func (opt Limit) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addIntOpt(int(opt), "Limit", "--limit", cmdOpts, hgcmd)
}

func (opt Modified) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Modified", "--modified", cmdOpts, hgcmd)
}

func (opt Mq) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Mq", "--mq", cmdOpts, hgcmd)
}

func (opt NoMerges) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "NoMerges", "--no-merges", cmdOpts, hgcmd)
}

func (opt NonInteractive) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "NonInteractive", "--noninteractive", cmdOpts, hgcmd)
}

func (opt NoStatus) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "NoStatus", "--no-status", cmdOpts, hgcmd)
}

func (opt Num) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Num", "--num", cmdOpts, hgcmd)
}

func (opt Patch) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Patch", "--patch", cmdOpts, hgcmd)
}

func (opt Print0) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Print0", "--print0", cmdOpts, hgcmd)
}

func (opt Profile) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Profile", "--profile", cmdOpts, hgcmd)
}

func (opt Prune) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Prune", "--prune", cmdOpts, hgcmd)
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

func (opt Removed) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Removed", "--removed", cmdOpts, hgcmd)
}

func (opt Repository) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Repository", "-R", cmdOpts, hgcmd)
}

func (opt Rev) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Rev", "--rev", cmdOpts, hgcmd)
}

func (opt Ssh) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Ssh", "--ssh", cmdOpts, hgcmd)
}

func (opt Stat) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Stat", "--stat", cmdOpts, hgcmd)
}

func (opt Style) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Style", "--style", cmdOpts, hgcmd)
}

func (opt SubRepos) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "SubRepos", "--subrepos", cmdOpts, hgcmd)
}

func (opt Tags) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Tags", "--tags", cmdOpts, hgcmd)
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

func (opt Unknown) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Unknown", "--unknown", cmdOpts, hgcmd)
}

func (opt User) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "User", "--user", cmdOpts, hgcmd)
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
