# Changelog

## 201y-mm-dd v0.1

* Provides the HgClient type for working with the Mercurial Command Server.
   - An instance of HgClient can connect to one Hg CS, and so work with one repo.
   - You can create as many HgClient instances as you want, to work with
     multiple repos.
* The default Hg executable is 'hg'.
* Accepts the Hg executable to use as the first argument to the Connect() method.
* The default repo to work on is the first one found in/up the current path.
* Accepts the Hg repo to work on as the second argument to the Connect() method.
* Implements commands: Init(), Summary(), Version()
* Has a testsuite that can be run by: 'cd gohg_lib; go test'
* Does not allow for interaction with Hg (yet?).
