// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

/*
Compatibility

The gohg client library is created with Go1 (v1.0.3). It is tested against
Mercurial 2.5.2, both on Windows 7 and Ubuntu 12.04.

Currently there is no mechanism to handle differences in possibilities between
different Mercurial versions.

Dependencies

Only Go and it's standard library. Though I'm using gocov for checking test
coverage (see https://github.com/axw/gocov).

Installation

At the commandline type:
  go get [-u] bitbucket.org/gohg/gohg
  go install bitbucket.org/gohg/gohg
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

The HgClient.Connect() method returns an error, so you can check if the
connection succeeded, and if it is safe to go on or not.

Once the work is done, you can disconnect the Hg CS. We advise to use a typical
Go idiom for this:

  err := hc.Connect("hg", "~/myrepo", nil)
  if err != nil {
      log.Fatal(err)
  }
  defer hc.Disconnect()
  // do the real work here

Commands

Now that we have a connection to a Hg CS we can do some work with the
repository. This is done with commands, implemented as methods for the HgClient
type. Each command has the same name as the corresponding Hg command, except it
starts with a capital letter.

  log, err := hc.Log(Limit(2))
  if err != nil {
      fmt.Printf(err)
  }

Commands normally return a byte slice containing the resulting data, and an
error type. But there are exceptions (see api docs hereafter). If the
returnvalue of the command indicated it did not complete successful, the
returnvalue is included in the error message. As is any error message from
Mercurial.

  log, err := hc.Log()          // log is a byte slice
  err := hc.Init("~/mynewrepo") // only returns an error
  version , err:= hc.Version()  // version is a string

Parameters and Options

In the gohg tool, parameters are used to pass-in any arguments for a command
that are not options. They are passed in first, before any options, as a string
or a string slice, depending on the command. These parameters typically contain
revisions, paths or filenames and so.

  log, err := hc.Log("myfile")
  heads, err := hc.Heads("foobranch")

Options to commands use the same name as the long form of the Mercurial option
they represent, and start with a capital letter (as do all exported symbols in
Go). An option can be of type bool, int or string. You just pass the value as
the parameter to the option. You can combine any options, just as you can on the
commandline. Options can be passed in more than once if appropriate (see the
ones marked with '[+]' in the Mercurial help).

  log, err := hc.Log(Verbose(true))
  log, err := hc.Log(Limit(2))
  log, err := hc.Log(User("John Doe"), User("me"))

In contrast with the typical commandline usage of the Hg commands, all options
have to be passed after the parameter(s) to the command, because the quantity of
options is variable. This is just a Go constraint of the implementation making
it possible to pass a variable number of options.

  log, err := hc.Log("myfile", Verbose(true), Limit(2))

The gohg tool only checks if the options the caller gives are valid for that
command. It does not check if the values are valid for the combination of that
command and that option, as that is done by Mercurial. No need to do double
work. If an option is not valid for a command, it is silently ignored, so it is
not passed to the Hg CS.

Some options are not implemented, as they seemed not relevant for use with this
tool (for instance: the global --color option, or the --print0 option for
status).

Limitations

The following config settings are fixated in the code (at least for now):
  encoding=utf-8
  ui.interactive=False
  extensions.color=!

As mentioned earlier, passing config info is not implemented yet.

Currently there is no support for any extensions to Mercurial.

Issues

If you experience any issues using the gohg tool, please register an issue using
the Bitbucket issue tracker at https://bitbucket.org/gohg/gohg/issues.

You can also register any enhancement requests or suggestions for improvement
there.

License

Copyright 2012, The gohg Authors. All rights reserved.

Use of this source code is governed by a BSD style license that can be found in
the LICENSE.md file.

*/
package gohg