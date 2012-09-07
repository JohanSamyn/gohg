TODO
****

(in no particular order, and order can change anytime)

1.  ADD TESTS !!!

#.  Use the stats.py example from the python-hglib tool for testing.
    That can be a simple means of being able to verify the output
    by comparing the results of both gohg and python-hglib.

#.  Take care of the fact that there can be an error at the Hg side, and that it
    returns a (rather large) help text instead of the expected result.

#.  I will have to make some blocking mechanism, for if Hg sends way more data than
    I can handle in one move. Some buffering is required, to capture _all_ data
    from the pipe.

#.  Maybe I should rename into gohglib ? Mmm, or maybe not.

#.  Place all error messages together (in a separate file?) and give them a
    concise varname that can be more easily used in a testing situation.

#.  If the Hg CS aborts (because of an unknown command f.i.), gracefully
    re-establish the connection. Up to a max number of times, otherwise
    abort the client too.

#.  KNOWN LIMITATION
    Currently only UTF8 encoding is supported.

#.  KNOWN LIMITATION
    Currently no interaction is supported (ui.interactive=False).

#.  ADVANCED
    Maybe we can even manage multiple connections ?
    Question: can we manage more than 1 repo from the same Hg Command Server
    instance ? If not then I don't think multiple connections will work.

#.  ADVANCED
    Is it possible to run a Hg CS as a Windows service, and connect to it ?
    See http://code.google.com/p/winsvc/ and search for it's counterpart for Python.
