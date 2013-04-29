// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

/*
Compatibility

The gohg client library is created with Go1 (v1.0.3). It is tested against
Mercurial 2.5.2, both on Windows 7 and Ubuntu 12.04.

Currently there is no mechanism to handle differences in possibilities between
different Mercurial versions. The errors returned by Mercurial are your only
help here.

Dependencies

Only Go and it's standard library. Though I'm using gocov for checking test
coverage (see https://github.com/axw/gocov).

Installation

At the commandline type:
  go get [-u] bitbucket.org/gohg/gohg
  go test -v bitbucket.org/gohg/gohg

Import the package

Start with importing the gohg package:
  import . "bitbucket.org/gohg/gohg"

Connecting the Mercurial Command Server

All interaction with the Hg CS happens through the HgClient type, of which you
have to create an instance:

  hc :=  NewHgClient()

Then you can connect the Hg CS:

  err := hc.Connect("hg", "~/myrepo", nil)
   4                 1        2        3

1. The Hg executable:

The first parameter is the Mercurial command to use (which 'hg'). You can leave
it blanc to let the gohg tool use the default Mercurial command on the system.
Having a parameter for the Hg command allows for using a different Hg version,
for testing purposes for instance.

2. The repository path:

The second parameter is the path to the repository you want to work on. You can
leave it blanc to have gohg use the repository it can find for the current path
(searching upward in the folder tree eventually).

3. The config for the session:

The third parameter allows to provide extra configuration for the session.
Though this is currently not implemented yet.

4. The returnvalue:

The HgClient.Connect() method eventually returns an error, so you can check if
the connection succeeded, and if it is safe to go on or not.

Once the work is done, you can disconnect the Hg CS. We advise to use a typical
Go idiom for this:

  err := hc.Connect("hg", "~/myrepo", nil)
  if err != nil {
      log.Fatal(err)
  }
  defer hc.Disconnect()
  // do the real work here

Config

The gohg tool sets some environment variables for the Hg CS session, to ensure
it's good working:
  // ensure Hg works in english
  HGPLAIN=True
  // Use only the .hg/hgrc from the repo itself.
  HGRCPATH=''
  HGENCODING=UTF-8

Commands

Once we have a connection to a Hg CS we can do some work with the repository.
This is done with commands, implemented as methods for the HgClient type. Each
command has the same name as the corresponding Hg command, except it starts with
a capital letter of course.

  log, err := hc.Log([]Option{Limit(2)}, nil)
  if err != nil {
      fmt.Printf(err)
      ...
  }
  fmt.Printf("%s", log)

Commands return a byte slice containing the resulting data, and eventually an
error. But there are a few exceptions (see api docs).

  log, err := hc.Log(nil, nil)       // log is a byte slice
  err := hc.Init(nil, "~/mynewrepo") // only returns an error eventually
  vers, err:= hc.Version()           // vers is a string of the form '2.4'

If a command fails, the returned error contains 3 elements: 1) the returncode
by Mercurial, 2) the full command that was passed to the Hg CS, and 3) the
eventual error message returned by Mercurial.

So the command

  idinfo, err := hct.Identify([]Option{Verbose(true)}, []string{"C:\\DEV\\myrepo"})

could return something like the following in the err variable when it fails:

  runcommand: Identify(): returncode=-1
  cmd: identify -v C:\DEV\myrepo
  hgerr:

// TODO: make up my mind about the notes below

Note: I could have implemented the command aliases too, but that would cost you
an extra function call (to go from Ci to Commit f.i.), so I did not do it. And
having to use the original commands makes your code clearly readable too. (But
there are examples of how to do it in identify.go and showconfig.go.)

Note: All aliases for commands that are mentioned in the Mercurial help will
work too. So you can call either Commit() or Ci() for example. (But it will cost
you an extra function call, to call Commit from Ci.)

Options and Parameters

As on the commandline, options come before parameters.

  opts := []Option{Verbose(true), Limit(2)}
  params := []string{"mytool.go"}
  log, err := hc.Log(opts, params)

Options to commands use the same name as the long form of the Mercurial option
they represent, but start with a capital letter (as do all exported symbols in
Go). An options value can be of type bool, int or string. You just pass the
value as the parameter to the option (= type conversion of the value to the
option type). You can pass any number of options, as the elements of a slice.
Options can occur more than once if appropriate (see the ones marked with '[+]'
in the Mercurial help).

  log, err := hc.Log([]Option{Verbose(true)}, nil)
  log, err := hc.Log([]Option{Limit(2)}, nil)
  log, err := hc.Log([]Option{User("John Doe"), User("me")}, nil)

Parameters are used to provide any arguments for a command that are not options.
They are passed in as a string or a string slice, depending on the command.
These parameters typically contain revisions, paths or filenames and so.

  log, err := hc.Log(nil, []string{"myfile"})
  heads, err := hc.Heads(nil, []string{"foobranch"})

The gohg tool only checks if the options the caller gives are valid for that
command. It does not check if the values are valid for the combination of that
command and that option, as that is done by Mercurial. No need to implement that
again. If an option is not valid for a command, it is silently ignored, so it is
not passed to the Hg CS.

Some options are not implemented, as they seemed not relevant for use with this
tool (for instance: the global --color option, or the --print0 option for
status).

Error handling

The gohg tool only returns errors, with an as clear as possible message, and
never uses log.Fatal() nor panics, even if those may seem appropriate. It leaves
it up to the caller to do that eventually. It's not up to this library to decide
whether to do a retry or to abort the complete application.

Limitations

* The following config settings are fixated in the code (at least for now):
  encoding=utf-8
  ui.interactive=False
  extensions.color=!

* As mentioned earlier, passing config info is not implemented yet.

* Currently there is no support for any extensions to Mercurial.

* If multiple Hg CSers are used against the same repo, it is up to Mercurial
to handle this correctly.

* Mercurial is always run in english. No internationalization yet.

Issues

If you experience any problems using the gohg tool, please register an issue
using the Bitbucket issue tracker at https://bitbucket.org/gohg/gohg/issues.

You can also register any enhancement requests or suggestions for improvement
there.

License

Copyright 2012, The gohg Authors. All rights reserved.

Use of this source code is governed by a BSD style license that can be found in
the LICENSE.md file.

*/
package gohg
