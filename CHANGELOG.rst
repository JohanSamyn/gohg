Changes made to the gohg library
********************************

2012-??-?? v0.1
---------------

- Provides the HgClient type for working with the Mercurial Command Server.
- Default Hg executable is 'hg'.
- Accepts the Hg executable to use as the first argument to the Connect() method.
- Default repo to work on is the one found in the current path.
- Accepts the Hg repo to work on as the second argument to the Connect() method.
- Implements commands:
	Add()
	Init()
	Summary()
	Version()
- Has a testsuite that can be run by: 'cd gohg_lib; go test'
