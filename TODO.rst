TODO
****

(in no particular order, and order can change anytime)

1.  DONE
    ADD TESTS !!! (good for learning to use the testing package too)

#.  DONE
    Use a dedicated repo for the tests.

#.  Add the other Hg CS channels to readToHg() and sendToHg().

#.  Make sure Init() fails gracefuly when no Hg repo avail, and it 'asks' for
    the name of a new (= unexisting) repo in it's failure message, which much
    differ from the repo the Hg CS is associated with (which would fail anyway).

#.  DONE
    Refactor Connect() and Close() into methods of HgServer.
    Or even of HgClient, and eliminate HgServer ?

#.  Catch returncode <> 0, and avoid showing the help text returned by the Hg CS
    in that case.
    Also take care of the fact that there can be an error at the Hg side, and
    that it returns a (rather large) help text instead of the expected result.

#.  Add possibility for extra configurong both using commandline arguments
    and/or a config file (where cli > configfile > defaults).
    F.i.: The default Hg command is 'hg', which could be overridden for testing
            with different Hg version.
    Could be a file gohg.cfg in the same folder as the gohg.exe, and a section
    per repo, and one "general" section.
    Or maybe just a [gohg] section in one of the 'normal' Hg config files?

#.  Add the possibility to configure (or use a command line option with) gohg so
    that it creates a dummy Hg repo first in case there is none available yet.
    Then switching repos or using a repo pool should allow to use the correct
    repo.

#.  Use the stats.py example from the python-hglib tool for testing.
    That can be a simple means of being able to verify the output
    by comparing the results of both gohg and python-hglib.

#.  I will have to make some blocking mechanism, for if Hg sends way more data
    than I can handle in one move. Some buffering is required, to capture _all_
    data from the pipe.

#.  Place all error messages together (in a separate file?) and give them a
    concise varname (perhaps 'msg0001' ?) that can be more easily used in a
    testing situation. (or: e0001 for errors, w0001 for warnings, etc.?)

#.  If the Hg CS aborts (because of an unknown command f.i.), gracefully
    re-establish the connection. Up to a max number of times, otherwise
    abort the client too.

#.  Perhaps rename gohg.go into gohglib.go, and add a new gohg.go as a command,
    so one can use the 'gohg' command from the commandline the same way one uses
    the original 'hg' one (see the chg tool Yuya developped) ?
    Used another solution: I put gohg.go into folder gohg_lib, so I can add a
    command sourcefile in folder gohg_cmd.

#.  QUESTION -
    Ask for an example of how to start the Hg CS in debug mode (logging to '-').

#.  KNOWN LIMITATION -
    Currently only UTF8 encoding is supported.

#.  KNOWN LIMITATION -
    Currently no interaction is supported (ui.interactive=False).

#.  ADVANCED -
    Maybe we can even manage multiple connections (see JavaHg) ?
    Question: can we manage more than 1 repo from the same Hg Command Server
    instance ? If not then I don't think multiple connections will work.

#.  ADVANCED -
    Add the possibility to switch to another repo then the one used to start the
    Hg CS. If possible, that is. Maybe this should be solved by adding a pool?

#.  ADVANCED -
    Is it possible to run a Hg CS as a Windows service, and connect to it ?
    See http://code.google.com/p/winsvc/ and search for it's counterpart for
    Python.
