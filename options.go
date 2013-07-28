// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

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

	AccessLog         string // -A --accesslog FILE
	Active            bool   // -a --active
	Added             bool   // -a --added
	AddRemove         bool   // -A  --addremove
	Address           string // -a --address ADDR
	After             bool   // -A --after
	All               bool   //    --all
	AllTasks          bool   //    --all-tasks
	Amend             bool   //    --ammend
	Bookmark          string // -B --bookmark BOOKMARK [+]
	Bookmarks         bool   // -B --bookmarks
	Branch            bool   // -b --branch BRANCH [+]
	Certificate       string //    --certificate FILE
	Change            bool   // [-c] --change REV
	Changeset         bool   // -c --changeset
	Check             bool   // -c --check
	Clean             bool   // -c --clean // -C --clean (update)
	CloseBranch       bool   //    --close-branch
	Closed            bool   // -c --closed
	CmdServer         string //    --cmdserver MODE
	CompletedTasks    bool   //    --completed-tasks
	Copies            bool   // -C --copies
	Date              string // -d --date
	Daemon            bool   // -d --daemon
	DaemonPipefds     int    //    --daemon-pipefds NUM
	Deleted           bool   // -d --deleted
	DryRun            bool   // -n --dry-run
	Exclude           string // -X --exclude PATTERN [+]
	ErrorLog          string // -E --errorlog FILE
	File              bool   // -f --file
	Follow            bool   // -f --follow
	Force             bool   // -f --force
	Git               bool   // -g --git
	Graph             bool   // -G --graph
	Id                bool   // -i --id
	IgnoreAllSpace    bool   // -w --ignore-all-space
	IgnoreBlankLines  bool   // -B --ignore-blank-lines
	Ignored           bool   // -i --ignored
	IgnoreSpaceChange bool   // -b --ignore-space-change
	Include           string // -I --include PATTERN [+]
	Insecure          bool   //    --insecure
	Ipv6              bool   // -6 --ipv6
	Keyword           string // -k --keyword
	Limit             int    // -l --limit
	LineNumber        bool   // -l --line-number
	Logfile           string // -l --logfile
	Message           string // -m --message
	Modified          bool   // -m --modified
	// Mq          bool   //    --mq
	Name         string // -n --name NAME
	NewBranch    bool   //    --new-branch
	NoDates      bool   //    --nodates
	NoDecode     bool   //    --no-decode
	NoFollow     bool   //    --no-follow
	NoMerges     bool   // -M --no-merges
	NoStatus     bool   // -n --no-status
	NoUpdate     bool   // -U --noupdate
	Num          bool   // -n --num
	Number       bool   // -n --number
	Output       string // -o --output FORMAT
	Patch        bool   // -p --patch
	PidFile      string //    --pid-file FILE
	Port         int    // -p --port PORT
	Prefix       string // [-p] --prefix PREFIX
	Preview      bool   // -P --preview
	Print0       bool   // -0 --print0
	Prune        bool   // -P --prune
	Pull         bool   //    --pull
	Rebase       bool   //    --rebase
	Remote       bool   //    --remote
	RemoteCmd    string //    --remotecmd CMD
	Removed      bool   // (-r) --removed
	Rev          string // -r --rev REV [+] //    --rev REV [+]
	Reverse      bool   //    --reverse
	ShowFunction bool   // -p --show-fuction
	Similarity   int    // -s --similarity SIMILARITY
	Ssh          string // -e --ssh CMD
	Stat         bool   //    --stat
	Stdio        bool   //    --stdio
	Style        string //    --style STYLE
	SubRepos     bool   // -S --subrepos
	SwitchParent bool   //    --switch-parent
	Tags         bool   // -t --tags
	Template     string //    --template
	Templates    string // -t --templates TEMPLATE
	Text         bool   // -a --text
	Tool         string // -t --tool VALUE
	Topo         bool   // -t --topo
	Type         string // -t --type TYPE
	Uncompressed bool   //    --uncompressed
	Unified      int    // -U --unified NUM
	Unknown      bool   // -u --unknown
	Untrusted    bool   // -u --untrusted
	Update       bool   // -u --update
	UpdateRev    string // -u --updaterev
	User         string // -u --user
	WebConf      string //    --web-conf FILE

	// debugging and profiling options

	Debug     bool //    --debug
	Profile   bool //    --profile
	Time      bool //    --time
	Traceback bool //    --traceback
)

type globalOpts struct {
	Config
	Cwd
	Hidden
	NonInteractive
	Quiet
	Repository
	Verbose
}

type debugOpts struct {
	Debug
	Profile
	Time
	Traceback
}

// Option is an interface that allows adding options to a command in a more or
// less controlled way.
type Option interface {
	addOption(interface{}, *[]string) error
}

func (opt AccessLog) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "AccessLog", "--accesslog", cmdOpts, hgcmd)
}

func (opt Address) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Address", "--address", cmdOpts, hgcmd)
}

func (opt Active) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Active", "--active", cmdOpts, hgcmd)
}

func (opt Added) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Added", "--added", cmdOpts, hgcmd)
}

func (opt AddRemove) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "AddRemove", "--addremove", cmdOpts, hgcmd)
}

func (opt After) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "After", "--after", cmdOpts, hgcmd)
}

func (opt All) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "All", "--all", cmdOpts, hgcmd)
}

func (opt AllTasks) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "AllTasks", "--all-tasks", cmdOpts, hgcmd)
}

func (opt Amend) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Amend", "--amend", cmdOpts, hgcmd)
}

func (opt Bookmark) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Bookmark", "--bookmark", cmdOpts, hgcmd)
}

func (opt Bookmarks) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Bookmarks", "--bookmarks", cmdOpts, hgcmd)
}

func (opt Branch) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Branch", "--branch", cmdOpts, hgcmd)
}

func (opt Change) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Change", "--change", cmdOpts, hgcmd)
}

func (opt Changeset) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Changeset", "--changeset", cmdOpts, hgcmd)
}

func (opt Check) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Check", "--check", cmdOpts, hgcmd)
}

func (opt Clean) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Clean", "--clean", cmdOpts, hgcmd)
}

func (opt CloseBranch) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "CloseBranch", "--close-branch", cmdOpts, hgcmd)
}

func (opt Closed) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Closed", "--closed", cmdOpts, hgcmd)
}

func (opt CmdServer) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "CmdServer", "--cmdserver", cmdOpts, hgcmd)
}

func (opt CompletedTasks) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "CompletedTasks", "--completed-tasks", cmdOpts, hgcmd)
}

func (opt Config) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Config", "--config", cmdOpts, hgcmd)
}

func (opt Copies) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Copies", "--copies", cmdOpts, hgcmd)
}

func (opt Cwd) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Cwd", "--cwd", cmdOpts, hgcmd)
}

func (opt Date) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Date", "--date", cmdOpts, hgcmd)
}

func (opt Daemon) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Daemon", "--daemon", cmdOpts, hgcmd)
}

func (opt DaemonPipefds) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addIntOption(int(opt), "DaemonPipefds", "--daemon-pipefds", cmdOpts, hgcmd)
}

func (opt Debug) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Debug", "--debug", cmdOpts, hgcmd)
}

func (opt Deleted) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Deleted", "--deleted", cmdOpts, hgcmd)
}

func (opt DryRun) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "DryRun", "--dry-run", cmdOpts, hgcmd)
}

func (opt ErrorLog) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "ErrorLog", "--errorlog", cmdOpts, hgcmd)
}

func (opt Exclude) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Exclude", "--exclude", cmdOpts, hgcmd)
}

func (opt File) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "File", "--file", cmdOpts, hgcmd)
}

func (opt Follow) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Follow", "--follow", cmdOpts, hgcmd)
}

func (opt Force) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Force", "--force", cmdOpts, hgcmd)
}

func (opt Git) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Git", "--git", cmdOpts, hgcmd)
}

func (opt Graph) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Graph", "--graph", cmdOpts, hgcmd)
}

func (opt Hidden) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Hidden", "--hidden", cmdOpts, hgcmd)
}

func (opt Id) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Id", "--id", cmdOpts, hgcmd)
}

func (opt Ignored) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Ignored", "--ignored", cmdOpts, hgcmd)
}

func (opt IgnoreAllSpace) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "IgnoreAllSpace", "--ignore-all-space", cmdOpts, hgcmd)
}

func (opt IgnoreBlankLines) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "IgnoreBlankLines", "--ignore-blank-lines", cmdOpts, hgcmd)
}

func (opt IgnoreSpaceChange) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "IgnoreSpaceChange", "--ignore-space-change", cmdOpts, hgcmd)
}

func (opt Include) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Include", "--include", cmdOpts, hgcmd)
}

func (opt Insecure) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Insecure", "--insecure", cmdOpts, hgcmd)
}

func (opt Ipv6) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Ipv6", "--ipv6", cmdOpts, hgcmd)
}

func (opt Keyword) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Keyword", "--keyword", cmdOpts, hgcmd)
}

func (opt Limit) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addIntOption(int(opt), "Limit", "--limit", cmdOpts, hgcmd)
}

func (opt LineNumber) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "LineNumber", "--line-number", cmdOpts, hgcmd)
}

func (opt Logfile) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Logfile", "--logfile", cmdOpts, hgcmd)
}

func (opt Message) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Message", "--message", cmdOpts, hgcmd)
}

func (opt Modified) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Modified", "--modified", cmdOpts, hgcmd)
}

// func (opt Mq) addOption(cmdOpts interface{}, hgcmd *[]string) error {
// 	return addBoolOption(bool(opt), "Mq", "--mq", cmdOpts, hgcmd)
// }

func (opt Name) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Name", "--name", cmdOpts, hgcmd)
}

func (opt NewBranch) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "NewBranch", "--new-branch", cmdOpts, hgcmd)
}

func (opt NoDates) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "NoDates", "--nodates", cmdOpts, hgcmd)
}

func (opt NoDecode) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "NoDecode", "--no-decode", cmdOpts, hgcmd)
}

func (opt NoFollow) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "NoFollow", "--no-follow", cmdOpts, hgcmd)
}

func (opt NoMerges) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "NoMerges", "--no-merges", cmdOpts, hgcmd)
}

func (opt NonInteractive) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "NonInteractive", "--noninteractive", cmdOpts, hgcmd)
}

func (opt NoStatus) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "NoStatus", "--no-status", cmdOpts, hgcmd)
}

func (opt NoUpdate) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "NoUpdate", "--noupdate", cmdOpts, hgcmd)
}

func (opt Num) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Num", "--num", cmdOpts, hgcmd)
}

func (opt Number) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Number", "--number", cmdOpts, hgcmd)
}

func (opt Output) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Output", "--output", cmdOpts, hgcmd)
}

func (opt Patch) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Patch", "--patch", cmdOpts, hgcmd)
}

func (opt PidFile) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "PidFile", "--pid-file", cmdOpts, hgcmd)
}

func (opt Port) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addIntOption(int(opt), "Port", "--port", cmdOpts, hgcmd)
}

func (opt Prefix) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Prefix", "--prefix", cmdOpts, hgcmd)
}

func (opt Preview) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Preview", "--preview", cmdOpts, hgcmd)
}

func (opt Print0) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Print0", "--print0", cmdOpts, hgcmd)
}

func (opt Profile) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Profile", "--profile", cmdOpts, hgcmd)
}

func (opt Prune) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Prune", "--prune", cmdOpts, hgcmd)
}

func (opt Pull) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Pull", "--pull", cmdOpts, hgcmd)
}

func (opt Quiet) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Quite", "-q", cmdOpts, hgcmd)
}

func (opt Rebase) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Rebase", "--rebase", cmdOpts, hgcmd)
}

func (opt Remote) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Remote", "--remote", cmdOpts, hgcmd)
}

func (opt RemoteCmd) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "RemoteCmd", "--remotecmd", cmdOpts, hgcmd)
}

func (opt Removed) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Removed", "--removed", cmdOpts, hgcmd)
}

func (opt Repository) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Repository", "-R", cmdOpts, hgcmd)
}

func (opt Rev) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Rev", "--rev", cmdOpts, hgcmd)
}

func (opt Reverse) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Reverse", "--reverse", cmdOpts, hgcmd)
}

func (opt ShowFunction) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "ShowFunction", "--show-function", cmdOpts, hgcmd)
}

func (opt Similarity) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addIntOption(int(opt), "Similarity", "--similarity", cmdOpts, hgcmd)
}

func (opt Ssh) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Ssh", "--ssh", cmdOpts, hgcmd)
}

func (opt Stat) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Stat", "--stat", cmdOpts, hgcmd)
}

func (opt Stdio) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Stdio", "--stdio", cmdOpts, hgcmd)
}

func (opt Style) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Style", "--style", cmdOpts, hgcmd)
}

func (opt SubRepos) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "SubRepos", "--subrepos", cmdOpts, hgcmd)
}

func (opt SwitchParent) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "SwitchParent", "--switch-parent", cmdOpts, hgcmd)
}

func (opt Tags) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Tags", "--tags", cmdOpts, hgcmd)
}

func (opt Template) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Template", "--template", cmdOpts, hgcmd)
}

func (opt Templates) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Templates", "--templates", cmdOpts, hgcmd)
}

func (opt Text) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Text", "--text", cmdOpts, hgcmd)
}

func (opt Time) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Time", "--time", cmdOpts, hgcmd)
}

func (opt Tool) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Tool", "--tool", cmdOpts, hgcmd)
}

func (opt Topo) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Topo", "--topo", cmdOpts, hgcmd)
}

func (opt Traceback) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Traceback", "--traceback", cmdOpts, hgcmd)
}

func (opt Type) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "Type", "--type", cmdOpts, hgcmd)
}

func (opt Uncompressed) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Uncompressed", "--uncompressed", cmdOpts, hgcmd)
}

func (opt Unknown) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Unknown", "--unknown", cmdOpts, hgcmd)
}

func (opt Update) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Update", "--update", cmdOpts, hgcmd)
}

func (opt UpdateRev) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "UpdateRev", "--updaterev", cmdOpts, hgcmd)
}

func (opt Unified) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addIntOption(int(opt), "Unified", "--unified", cmdOpts, hgcmd)
}

func (opt User) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "User", "--user", cmdOpts, hgcmd)
}

func (opt Verbose) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addBoolOption(bool(opt), "Verbose", "-v", cmdOpts, hgcmd)
}

func (opt WebConf) addOption(cmdOpts interface{}, hgcmd *[]string) error {
	return addStringOption(string(opt), "WebConf", "--web-conf", cmdOpts, hgcmd)
}

// An add<Type>Option method does 2 things:
// 1. It checks if the option is valid for a given command,
//    by checking its presence in that commands option struct (logOpts, etc.).
// 2. If the option is valid, it then adds the option and any values for it
//    to the command's commandline (a []string that is joined to a \0-separated
//    string later in method runInHg()).
// It also returns an error stating the option is invalid for the command,
// but these error messages are silently ignored for the moment.

func addBoolOption(optval bool, optname string, optstr string, cmdOpts interface{}, hgcmd *[]string) error {
	fld := reflect.ValueOf(cmdOpts).Elem().FieldByName(optname)
	if fld.IsValid() && fld.CanSet() {
		fld.SetBool(optval) // add the value to the <cmd>Opts struct
		*hgcmd = append(*hgcmd, optstr)
	} else {
		return fmt.Errorf("adBoolOption(): command %s has no option %s", strings.Title((*hgcmd)[0]), optname)
	}
	return nil
}

func addIntOption(optval int, optname string, optstr string, cmdOpts interface{}, hgcmd *[]string) error {
	fld := reflect.ValueOf(cmdOpts).Elem().FieldByName(optname)
	if fld.IsValid() && fld.CanSet() {
		fld.SetInt(int64(optval)) // add the value to the <cmd>Opts struct
		*hgcmd = append(*hgcmd, optstr, strconv.Itoa(optval))
	} else {
		return fmt.Errorf("addIntOption(): command %s has no option %s", strings.Title((*hgcmd)[0]), optname)
	}
	return nil
}

func addStringOption(optval string, optname string, optstr string, cmdOpts interface{}, hgcmd *[]string) error {
	fld := reflect.ValueOf(cmdOpts).Elem().FieldByName(optname)
	if fld.IsValid() && fld.CanSet() {
		fld.SetString(optval) // add the value to the <cmd>Opts struct
		*hgcmd = append(*hgcmd, optstr, optval)
	} else {
		return fmt.Errorf("addStringOption(): command %s has no option %s", strings.Title((*hgcmd)[0]), optname)
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
		typeOfF := f.Type()
		typeOfFS := fmt.Sprintf("%s", f.Type())
		if strings.Contains(typeOfFS, "gohg.globalOpts") || strings.Contains(typeOfFS, "gohg.debugOpts") {
			for j := 0; j < f.NumField(); j++ {
				g := f.Field(j)
				if j == 0 {
					// s = s + typeOfFS + " = {"
				} else {
					s = s + ", "
				}
				s = s + fmt.Sprintf("%s=%v", typeOfF.Field(j).Name, g.Interface())
			}
			// s = s + "}"
		} else {
			s = s + fmt.Sprintf("%s=%v", typeOfT.Field(i).Name, f.Interface())
		}
	}
	s = s + "}"
	return s
}
