TODO
****

#   Use the stats.py example from the python-hglib tool for testing.
    That can be a simple means of being able to verify the output
    by comparing the results of both gohg and python-hglib.

#   Maybe we can even manage multiple connections ?
    Question: can we manage more than 1 repo from the same Hg Command Server
    instance ? If not then I don't think multiple connections will work.

#   Take care of the fact that there can be an error at the Hg side, and that it
    returns a (rather large) help text instead of the expected result.

#   Is it possible to run a Hg CS as a Windows service, and connect to it ?
    See http://code.google.com/p/winsvc/ and search for it's counterpart for Python.

#   I will have to make some blocking mechanism, for if Hg sends way more data than
    I can handle in one move. Some buffering is required, to capture _all_ data
    from the pipe.

