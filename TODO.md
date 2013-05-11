# TODO

(in no particular order, and order can change anytime)

* Deal with the output of the Update() command (see python-hglib).
This could be useful info. In fact, I should deal with all results returned by
any Hg command, beyond the error returns.

* Deal with the possible prompting Hg does for Update() (see test-update.py in
python-hglib).

* Return the results of commands like Log() in a more user friendly way (a slice
or map of csets, ...). Maybe even in a special format via an extra option
Json[Out](true) or Xml[Out](true)?

* Use the short options instead of the long ones. Could become of importance
when commands get rather long, and they start to hit the boundaries of the OS.
PROBLEM: this is not always possible, cause certain options are used for
different commands, and do not always have the same short option. So that would
require us to still use the long option for those. It would have to be commented
in the code too.

* Panic when we cannot find a Mercurial version >= 1.9.
Or maybe make those situation a dedicated ConnectionError, making it clear to
the caller _he_ should panic.

* If an option is silently dropped, because it is not valid for the command,
mention that in the logfile.

* Is the forced HGPLAIN=True setting really necessary?

* Should I always add a '--' as the option closer before passing arguments?
How does TortoiseHg do this? And how do the other client libs do this ?

* Tests should make use of the returncode of commands too. And of the fact that
function runcommand() returns nil as data when an error was detected.

* The program examples/example1.go is now practicaly the same as the one in the
README. Change it, so that is is worth keeping, or get rid of it.

* Make the return of the commands in gohg a slice-of-strings (separated by the
linefeed in the Hg output). That way you already avoid having trouble with
multi-byte chars (runes) when parsing. It will also make tests easier (slice
only comparable to nil). After all, that's what the results are: lines of text,
as they are also printed from the commandline.

* Find out how to handle functionality that belongs to extensions.
Think of the -mq option for instance. So we'll have to find a way to detect what
extensions are active, and conditionaly allow their options.

* Encoding (in the hellomessage) should be UTF-8.

* Set HGRCPATH (explicitely) to "" and HGPLAIN to "True".
The first assures that only the hgrc file from the repo itself is used.
The second assures all Mercurial output is with default values, excluding
internationalization. Use HGPLAINEXCEPT="i18n" to keep internationalization,
what can be useful for error messages f.i.. (see CommandServerFactory.scala in
[Meutrino](http://code.google.com/p/meutrino))
DO NOT use HGPLAINEXCEPT for keeping the language, as gohg should shield the
caller from those messages, and return "it's own" error messages. Or not ?

* Assure that Disconnect() cannot terminate the connection when a command is
still running, unless a kind of --force option is passed-in.

* The Verify() command should simply whether the result channel returns 0 (=ok)
or 1 (=error). In case of 0 it simply returns something like "Repo %s is healthy".
In the case of 1 it should pass thru the errors produced by Hg.
Why not let _always_ it just return what Hg returns?

* hct.Version() : test and compare with a "commandline call to the same hg"
and capturing cmd.Output

* Is it really necessary to have a seperate method for each command ?
As they are all very alike, maybe I can come up with some generic command
method, and lead all commands to it ?
Having a seperate method for each one does offer the possibility to add some
more things, like converting the result of certain commands into a more Go-like
format. F.i. returning log info as a map of changesets (key = hash), so each
element of a cset is easily addressable without the client having to parse the
gohg result.

* Add the other Hg CS channels to receiveFromHg() and sendToHg().

* Allow to start the Hg CS in debug mode (logging to '-').

* Find out if it is possible to interrupt the Hg CS when it is producing a
lengthy result ? F.i. if someone commanded "hg log", with no limitation, for a
rather big repo, one should be able to interrupt it. Like you can type ctrl-c
on the commandline. Such output should be accepted and passed thru in a buffered
way too I guess.

* Catch returncode <> 0, and avoid showing the help text returned by the Hg CS
in that case. Also take care of the fact that there can be an error at the Hg
side, and that it returns a (rather large) help text instead of the expected
result.

* Add possibility for extra configuring both using commandline arguments and/or
a config file (where cli > configfile > defaults).
F.i.: The default Hg command is 'hg', which could be overridden for testing
with different Hg version.
Could be a file gohg.cfg in the same folder as the gohg.exe, and a section per
repo, and one "general" section. Or maybe just a [gohg] section in one of the
'normal' Hg config files?

* Add the possibility to configure (or use a command line option with) gohg so
that it creates a dummy Hg repo first in case there is none available yet. Then
switching repos or using a repo pool should allow to use the correct repo.

* I will have to make some blocking mechanism, for if Hg sends way more data
than I can handle in one move. Some buffering is required, to capture _all_ data
from the pipe.


* KNOWN LIMITATION - Currently only UTF8 encoding is supported.

* KNOWN LIMITATION - Currently no interaction is supported (ui.interactive=False).

* KNOWN LIMITATION - Currently no config settings are accepted.


* ADVANCED - Maybe we can even manage multiple connections (see JavaHg) ?
Question: can we manage more than 1 repo from the same Hg Command Server
instance ? If not then I don't think multiple connections will work.
But maybe we can manage more than one Hg CS instance, each to a different repo ?
Further thoughts on this:
Think about offering one entrypoint into gohg, always passing in the reporoot
and the command to run. Then gohg checks if there is already a task for that
reporoot or not. If not, your command is run, otherwise it has to wait. So we
can queue the commands for one reporoot (perhaps in a buffered channel). Maybe
that will require the caller (client) also to wait on a (unbuffered) channel for
the result ? This would be a kind of hosting version of gohg.

* ADVANCED - Add the possibility to switch to another repo then the one used to
start the Hg CS. If possible, that is. Maybe this should be solved by adding a pool?


* DONE - Turn the commands into a 'command' type ? So we can f.i. add a method to print
the resulting command string. Maybe I'll have to turn the 'hgcmd []string' into
a type for that to be possible, so I can give it a
'func (*hgcmd []string) String() string' method.

* DONE - Find a way to add a (global ?) option so that we can log/consult the built
command. This should be done right after the call to buildCommand(). (see also
function sprintfOpts() in client.go)

* DONE - If a boolean option is passed-in (explicitely), also explicitely set the
corresponding option to false if that is what is passed in, because that can be
a way to override a setting from the repo hgrc for instance. So the way
addBoolOption now handles the passed-in values is not really correct.

* DONE - Make an Exec() method for HgClient, to make it possible for the user to pass
in a full commandline. This could be convenient for if the gohg tool is not
updated yet, and one wants to use a new Hg command.

* DONE - Put the command string (= the result from buildCommand()) into the error
messages, as follows: 'cmd: log -limit 2'. Prepend this before the 'err:' and
'hgerr:' lines in the error message.

* DONE - Add a check in the right place (runCommand() ?) so that if there could not be
made a connection, but the caller ignored the error returned by Connect(), no
other command will proceed, but will also return an error indicating there is no
connection.

* DONE - log.Fatal should only be used at the topmost level (and probably even almost
never in a lib !?)

* DONE - Package gohg should never panic itself, nor use *Fatal*. That is to be left
to the using software. It should simply always return errors.

* DONE - MUST the first argument of runcommand be an explicit pointer to a []byte ?
Isn't a []byte a reference type, meaning it is an address already (as its
content is mutable) ?

* DONE - Use the stats.py example from the python-hglib tool for testing. That can be
a simple means of being able to verify the output by comparing the results of
both gohg and python-hglib.

* DONE - Add methods to add options and flags to commands, so no syntax errors can be
made. See [JavaHg](https://bitbucket.org/aragost/javahg) for an example.
Maybe add a struct containing all possible options and flags as booleans, and
let the caller activate them, and pass-in data for them when appropriate.

* DONE - Verify() should return the data _and_ the error info returned by the Hg CS,
as this is important info for the caller. So Command() should return both the
data and the error info returned by the Hg CS.
Maybe we should allow all commands to return the hgerr info? <- Done too.

* DONE - Remove one of Encoding() and HgEncoding() in client.go.

* DONE - ADD TESTS !!! (good for learning to use the testing package too)

* DONE - Use a dedicated repo for the tests.

* DONE - Refactor Connect() and Disconnect() into methods of HgServer.
Or even of HgClient, and eliminate HgServer ?

* DONE - Add the possibility to use options with commands.


* WONTFIX - Maybe I'll reorganize commands into less files, combining them according to
type: query/info (log, branches, heads, grep, ...), updates (init, add, commit,
...), config (showconfig, ...), exchange (push, pull, import, archive, bundle,
...), etc.
Wontfix: because it comes much more difficult if not impossible to test a single
command with the gt.bat script, as I experienced.

* DONE/WONTFIX - Maybe I should make it possible for callers to set options directly in the
<cmd>Opts struct ? That would require those structs to be made public, and they
would need to be coupled to their command too, which seems less obvious at the
moment. Maybe a method injectOpts(cmdOpts interface{}) to receiver 'func Log()'
f.i. could be a solution ?
Example:
  hc.Log.logOpts.Limit = 2

* WONTFIX - Maybe we should make a special option "Params" for parameters? Then the order
(params first, options last) wouldn't matter anymore. (But buildCommand() would
have to take care of their order when composing the Hg command.)

* WONTFIX (apparently this is not necessary) - Shouldn't I put (double) quotes
around certain string params and arguments ?

* WONTFIX - Make sure init() fails gracefuly when no Hg repo avail, and it 'asks' for
the name of a new (= unexisting) repo in it's failure message, which much
differ from the repo the Hg CS is associated with (which would fail anyway).

* WONTFIX - (at least not in the near future) -
Perhaps rename gohg.go into gohglib.go, and add a new gohg.go as a command,
so one can use the 'gohg' command from the commandline the same way one uses
the original 'hg' one (see the chg tool Yuya developped) ? Used another
solution: I put gohg.go into folder gohg, so I can add a command sourcefile
in folder gohg_cmd.

* WONTFIX - ADVANCED -
Is it possible to run a Hg CS as a Windows service, and connect to it ?
See http://code.google.com/p/winsvc/ and search for it's counterpart for Python.

* WONTFIX - (gohg takes care of only using the known commands into the Hg CS) -
If the Hg CS aborts (because of an unknown command f.i.), gracefully
re-establish the connection. Up to a max number of times, otherwise abort the
client too.

* WONTFIX - Make sure nothing in gohg depends on translatable pieces in Hg.
This is not necessary, as we assure HGPLAIN=True, so all that Hg returns is
in english.

* WONTFIX - Simplify the handling of options?
Maybe we need a type that allows to pass an option (like '--branch' or '-R')
and an optional value for it. Then we have enough with 3 types: one for bool,
int and string. Then we could still use something like the respective <cmd>Opts
structs (or something else) to verify what options are valid for what commands.
Example:
  data, err := hc.Log(HgOpt("--verbose", true), hgOpt("-R", "M:/DEV/hg-stable"), hgOpt("--limit", "2"))
  data, err := hcLog([]string{"--verbose", true, "-R", "M:/DEV/hg-stable"}, "--limit", "2")
Hmmm, looks a bit more verbose in writing, no ? Naa, I don't think this is a
good idea. The burden should be for the implementor, not for the user.

* WONTFIX - Place all error messages together (in a separate file?) and give them a
concise varname (perhaps 'msg0001' ?) that can be more easily used in a
testing situation. (or: e0001 for errors, w0001 for warnings, etc.?)
This seems not to be the standard Go practice though. Only when some error
condition occurs in different places it seems they are turned into some global
constant form and moved into a common place.

* WONTFIX - Why not also always return the Hg returnvalue to the command, instead of only
checking it in runCommand ? Certain callers might be insterested in that value.
Wontfix: because if there is a non-zero returnvalue it is reported back already.
