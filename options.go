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

// These are the options the Hg commands can take.

type (
	// the global options

	Config         string //    --config
	Cwd            string //    --cwd
	Hidden         bool   //    --hidden
	NonInteractive bool   // -y --noninteractive
	Quiet          bool   // -q --quiet
	Repository     string // -R --repository
	Verbose        bool   // -v --verbose

	// command-specific options

	Active      bool   // -a --active
	Added       bool   // -a --added
	AddRemove   bool   // -A  --addremove
	All         bool   //    --all
	Amend       bool   //    --ammend
	Bookmarks   bool   // -B --bookmarks
	Branch      bool   // -b --branch
	Change      bool   //    --change
	Clean       bool   // -c --clean
	CloseBranch bool   //    --close-branch
	Closed      bool   // -c --closed
	Copies      bool   // -C --copies
	Date        string // -d --date
	Deleted     bool   // -d --deleted
	DryRun      bool   // -n --dry-run
	Exclude     string // -X --exclude
	Follow      bool   // -f --follow
	Git         bool   // -g --git
	Graph       bool   // -G --graph
	Id          bool   // -i --id
	Ignored     bool   // -i --ignored
	Include     string // -I --include
	Insecure    bool   //    --insecure
	Keyword     string // -k --keyword
	Limit       int    // -l --limit
	Logfile     string // -l --logfile
	Message     string // -m --message
	Modified    bool   // -m --modified
	// Mq          bool   //    --mq
	NoMerges     bool   // -M --no-merges
	NoStatus     bool   // -n --no-status
	NoUpdate     bool   // -U --noupdate
	Num          bool   // -n --num
	Patch        bool   // -p --patch
	Print0       bool   // -0 --print0
	Prune        bool   // -P --prune
	Pull         bool   //    --pull
	Remote       bool   //    --remote
	RemoteCmd    string //    --remotecmd
	Removed      bool   // (-r) --removed
	Rev          string // -r --rev REV
	Ssh          string // -e --ssh
	Stat         bool   //    --stat
	Style        string //    --style
	SubRepos     bool   // -S --subrepos
	Tags         bool   // -t --tags
	Template     string //    --template
	Topo         bool   // -t --topo
	Uncompressed bool   //    --uncompressed
	Unknown      bool   // -u --unknown
	Untrusted    bool   // -u --untrusted
	UpdateRev    string // -u --updaterev
	User         string // -u --user

	// debugging and profiling options

	Debug     bool //    --debug
	Profile   bool //    --profile
	Time      bool //    --time
	Traceback bool //    --traceback
)

// type optionAdder interface {
// 	addOption(interface{}, *[]string) error
// }

type Option interface {
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

func (opt AddRemove) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "AddRemove", "--addremove", cmdOpts, hgcmd)
}

func (opt All) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "All", "--all", cmdOpts, hgcmd)
}

func (opt Amend) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Amend", "--amend", cmdOpts, hgcmd)
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

func (opt CloseBranch) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "CloseBranch", "--close-branch", cmdOpts, hgcmd)
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

func (opt Logfile) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Logfile", "--logfile", cmdOpts, hgcmd)
}

func (opt Message) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "Message", "--message", cmdOpts, hgcmd)
}

func (opt Modified) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Modified", "--modified", cmdOpts, hgcmd)
}

// func (opt Mq) addOption(cmdOpts interface{}, hgcmd *[]string) error {
// 	return addBoolOpt(bool(opt), "Mq", "--mq", cmdOpts, hgcmd)
// }

func (opt NoMerges) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "NoMerges", "--no-merges", cmdOpts, hgcmd)
}

func (opt NonInteractive) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "NonInteractive", "--noninteractive", cmdOpts, hgcmd)
}

func (opt NoStatus) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "NoStatus", "--no-status", cmdOpts, hgcmd)
}

func (opt NoUpdate) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "NoUpdate", "--noupdate", cmdOpts, hgcmd)
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

func (opt Pull) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Pull", "--pull", cmdOpts, hgcmd)
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

func (opt Uncompressed) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Uncompressed", "--uncompressed", cmdOpts, hgcmd)
}

func (opt Unknown) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Unknown", "--unknown", cmdOpts, hgcmd)
}

func (opt UpdateRev) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "UpdateRev", "--updaterev", cmdOpts, hgcmd)
}

func (opt User) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOpt(string(opt), "User", "--user", cmdOpts, hgcmd)
}

func (opt Verbose) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOpt(bool(opt), "Verbose", "-v", cmdOpts, hgcmd)
}

func addBoolOpt(optval bool, optname string, optstr string, cmdOpts interface{}, hgcmd *[]string) error {
	fld := reflect.ValueOf(cmdOpts).Elem().FieldByName(optname)
	if fld.IsValid() && fld.CanSet() {
		fld.SetBool(optval) // add the value to the <cmd>Opts struct
		*hgcmd = append(*hgcmd, optstr)
	} else {
		return fmt.Errorf("adBoolOpt(): command %s has no option %s", strings.Title((*hgcmd)[0]), optname)
	}
	return nil
}

func addIntOpt(optval int, optname string, optstr string, cmdOpts interface{}, hgcmd *[]string) error {
	fld := reflect.ValueOf(cmdOpts).Elem().FieldByName(optname)
	if fld.IsValid() && fld.CanSet() {
		fld.SetInt(int64(optval)) // add the value to the <cmd>Opts struct
		*hgcmd = append(*hgcmd, optstr, strconv.Itoa(optval))
	} else {
		return fmt.Errorf("addIntOpt(): command %s has no option %s", strings.Title((*hgcmd)[0]), optname)
	}
	return nil
}

func addStringOpt(optval string, optname string, optstr string, cmdOpts interface{}, hgcmd *[]string) error {
	fld := reflect.ValueOf(cmdOpts).Elem().FieldByName(optname)
	if fld.IsValid() && fld.CanSet() {
		fld.SetString(optval) // add the value to the <cmd>Opts struct
		*hgcmd = append(*hgcmd, optstr, optval)
	} else {
		return fmt.Errorf("addStringOpt(): command %s has no option %s", strings.Title((*hgcmd)[0]), optname)
	}
	return nil
}

func sprintfOpts(opts interface{}) string {
	s := fmt.Sprintf("%T", opts) + " = {"
	t := reflect.ValueOf(opts)
	typeOfT := t.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if i > 0 {
			s = s + ", "
		}
		s = s + fmt.Sprintf("%s=%v", typeOfT.Field(i).Name, f.Interface())
	}
	s = s + "}\n"
	return s
}
