# TODO

(in no particular order, and order can change anytime)

* Verify() should return the data _and_ the error info returned by the Hg CS,
as this is important info for the caller. So Command() should return both the
data and the error info returned by the Hg CS.
Maybe we should allow all commands to return the hgerr info?

* Encoding (in the hellomessage) should be UTF-8.

* Make the return of the commands in gohg a slice-of-strings (separated by the
linefeed in the Hg output). That way you already avoid having trouble with
multi-byte chars (runes) when parsing. It will also make tests easier (slice
only comparable to nil).

* Tests should make use of the returncode of commands too. And of the fact that
function command() returns nil as data when an error was detected.

* log.Fatal should only be used at the topmost level (and probably even never
in a lib !?)

* Add methods to add options and flags to commands, so no syntax errors can be
made. See [JavaHg](https://bitbucket.org/aragost/javahg) for an example.
Maybe add a struct containing all possible options and flags as booleans, and
let the caller activate them, and pass-in data for them when appropriate.

* Set HGRCPATH (explicitely) to "" and HGPLAIN to "True".
The first assures that only the hgrc file from the repo itself is used.
The second assures all Mercurial output is with default values, including
internationalization. Use HGPLAINEXCEPT="i18n" to keep internationalization,
what can be useful for error messages f.i.. (see CommandServerFactory.scala in
[Meutrino](http://code.google.com/p/meutrino))
DO NOT use HGPLAINEXCEPT for keeping the language, as gohg should shield the
caller from those mesages, and return "it's own" error messages. Or not ?

* Assure that Disconnect() cannot terminate the connection when a command is
still running, unless a kind of --force option is passed-in.

* The Verify() command should simply whether the result channel returns 0 (=ok)
or 1 (=error). In case of 0 it simply returns something like "Repo %s is healthy".
In the case of 1 it should pass thru the errors produced by Hg.

* hct.Version() : test and compare with a "commandline call to the same hg"
and capturing cmd.Output

* Package gohg should never panic itself, nor use *Fatal*. That is to be left
to the using software. It should simply always return errors.

* Is it really necessary to have a seperate method for each command ?
As they are all very alike, maybe I can come up with some generic command
method, and lead all commands to it ?
Having a seperate method for each one does offer the possibility to add some
more things, like converting the result of certain commands into a more Go-like
format. F.i. returning log info as a map of changesets (key = hash), so each
element of a cset is easily addressable without the client having to parse the
gohg result.

* Add the other Hg CS channels to readFromHg() and sendToHg().

* Allow to start the Hg CS in debug mode (logging to '-').

* Find out if it is possible to interrupt the Hg CS when it is producing a
lengthy result ? F.i. if someone commanded "hg log", with no limitation, for a
rather big repo, one should be able to interrupt it. Like you can type ctrl-c
on the commandline. Such output should be accepted and passed thru in a buffered
way too I guess.

* Make sure init() fails gracefuly when no Hg repo avail, and it 'asks' for
the name of a new (= unexisting) repo in it's failure message, which much
differ from the repo the Hg CS is associated with (which would fail anyway).

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

* Use the stats.py example from the python-hglib tool for testing. That can be
a simple means of being able to verify the output by comparing the results of
both gohg and python-hglib.

* I will have to make some blocking mechanism, for if Hg sends way more data
than I can handle in one move. Some buffering is required, to capture _all_ data
from the pipe.

* Place all error messages together (in a separate file?) and give them a
concise varname (perhaps 'msg0001' ?) that can be more easily used in a
testing situation. (or: e0001 for errors, w0001 for warnings, etc.?)
This seems not to be the standard Go practice though. Only when some error
condition occurs in different places it seems they are turned into some global
constant form and moved into a common place.

* KNOWN LIMITATION - Currently only UTF8 encoding is supported.

* KNOWN LIMITATION - Currently no interaction is supported (ui.interactive=False).


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


* DONE - Remove one of Encoding() and HgEncoding() in client.go.

* DONE - ADD TESTS !!! (good for learning to use the testing package too)

* DONE - Use a dedicated repo for the tests.

* DONE - Refactor Connect() and Disconnect() into methods of HgServer.
Or even of HgClient, and eliminate HgServer ?

* DONE - Add the possibility to use options with commands.


* WONTFIX (at least not in the near future) -
Perhaps rename gohg.go into gohglib.go, and add a new gohg.go as a command,
so one can use the 'gohg' command from the commandline the same way one uses
the original 'hg' one (see the chg tool Yuya developped) ? Used another
solution: I put gohg.go into folder gohg, so I can add a command sourcefile
in folder gohg_cmd.

* WONTFIX - ADVANCED -
Is it possible to run a Hg CS as a Windows service, and connect to it ?
See http://code.google.com/p/winsvc/ and search for it's counterpart for Python.

* WONTFIX (gohg takes care of only using the known commands into the Hg CS) -
If the Hg CS aborts (because of an unknown command f.i.), gracefully
re-establish the connection. Up to a max number of times, otherwise abort the
client too.

* WONTFIX Make sure nothing in gohg depends on translatable pieces in Hg.
This is not necessary, as we assure HGPLAIN=True, so all that Hg returns is
in english.
