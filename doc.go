// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

/*
Compatibility

▪ Mercurial

For Mercurial any version starting from 1.9 should be ok, cause that's the one
where the Command Server was introduced. If you send wrong options to it through
gohg, or commands or options not yet supported (or obsolete) in your Hg version,
you'll simply get back an error from Hg itself, as gohg does not check them.
But on the other hand gohg allows issuing new commands, not yet implemented by
gohg; see further.

▪ Go

Currently gohg is currently developed with Go1.2.1. Though I started with the
Go1.0 versions, I can't remember having had to change one or two minor things
when moving to Go1.1.1. Updating to Go1.1.2 required no changes at all.
I had an issue though with Go1.2, on Windows only, causing some tests using
os.exec.Command to fail. I'll have to look into that further, to find out if I
should report a bug.

▪ Platform

I'm developing and testing both on Windows 7 and Ubuntu 12.04/13.04/13.10.
But I suppose it should work on any other platform that supports Hg and Go.

Dependencies

Only Go and it's standard library. And Mercurial should be installed of course.

Installation

At the commandline type:

  go get [-u] bitbucket.org/gohg/gohg
  go test [-v] bitbucket.org/gohg/gohg

to have gohg available in your GOPATH.

Import the package

Start with importing the gohg package. Examples:
  import gohg "bitbucket.org/gohg/gohg"

  or

  import hg "bitbucket.org/gohg/gohg"

Connecting the Mercurial Command Server

All interaction with the Mercurial Command Server (Hg CS from now on) happens
through the HgClient type, of which you have to create an instance:

  hgcl :=  NewHgClient()

Then you can connect the Hg CS as follows:

  err := hgcl.Connect("hg", "~/myrepo", nil, false)
   5                   1        2        3     4

1. The Hg executable:

The first parameter is the Mercurial command to use (which 'hg'). You can leave
it blanc to let the gohg tool use the default Mercurial command on the system.
Having a parameter for the Hg command allows for using a different Hg version,
for testing purposes for instance.

2. The repository path:

The second parameter is the path to the repository you want to work on. You can
leave it blanc to have gohg use the repository it can find for the current path
you are running the program in (searching upward in the folder tree eventually).

3. The config for the session:

The third parameter allows to provide extra configuration for the session.
Though this is currently not implemented yet.

4. Should gohg create a new repo before connecting?

This fourth parameter allows you to indicate that you want gohg to first create
a new Mercurial repo if it does not already exist in the path given by the
second parameter. See the documentation for more detailed info.

5. The returnvalue:

The HgClient.Connect() method eventually returns an error, so you can check if
the connection succeeded, and if it is safe to go on.

Once the work is done, you can disconnect the Hg CS using a typical Go idiom:

  err := hgcl.Connect("hg", "~/myrepo", nil)
  if err != nil {
      log.Fatal(err)
  }
  defer hgcl.Disconnect()
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
This is done with commands, and gohg offers 3 ways to use them.

1. The command methods of the HgClient type.

2. The HgCmd type.

3. The ExecCmd() method of the HgClient type.

Each of which has its own reason of existence.

Commands return a byte slice containing the resulting data, and eventually an
error. But there are a few exceptions (see api docs).

  log, err := hgcl.Log(nil, nil)       // log is a byte slice
  err := hgcl.Init(nil, "~/mynewrepo") // only returns an error eventually
  vers, err:= hgcl.Version()           // vers is a string of the form '2.4'

If a command fails, the returned error contains 5 elements: 1) the name of the
internal routine where the error was trapped, 2) the name of the HgClient
command that was run, 3) the returncode by Mercurial, 4) the full command that
was passed to the Hg CS, and 5) the eventual error message returned by Mercurial.

So the command

  idinfo, err := hgcl.Identify([]hg.HgOption{hg.Verbose(true)}, []string{"C:\\DEV\\myrepo"})

could return something like the following in the err variable when it fails:

  runcommand: Identify(): returncode=-1
  cmd: identify -v C:\DEV\myrepo
  hgerr:

The command aliases (like 'id' for 'identify') are not implemented. But there
are examples in identify.go and showconfig.go of how you can easily implement
them yourself.

Commands - HgClient command methods

This is the easiest way, a kind of convenience. And the most readable too.
A con is that as a user you cannot know the exact command that was passed to Hg,
without some extra mechanics.

Each command has the same name as the corresponding Hg command, except it starts
with a capital letter of course.

An example (also see examples/example1.go):

  log, err := hgcl.Log([]hg.HgOption{hg.Limit(2)}, []string("my-file"))
  if err != nil {
      fmt.Printf(err)
      ...
  }
  fmt.Printf("%s", log)

Note that these methods all use the HgCmd type internally. As such they are
convenience wrappers around that type. You could also consider them as a kind of
syntactic sugar. If you just want to simply issue a command, nothing more, they
are the way to go.

The only way to obtain the commandstring sent to Hg when using these command
methods, is by calling the HgClient.ShowLastCmd() method afterwards before
issuing any other commands:

  log, err := hgcl.Log([]hg.HgOption{hg.Limit(2)}, []string("my-file"))
  fmt.Printf("%s", hgcl.ShowLastCmd()) // prints: log --limit 2 my-file

Commands - the HgCmd type

Using the HgCmd type is kind of the standard way. It is a struct that you can
instantiate for any command, and for which you can set elements Name, Options
and Params (see the api docs for more details). It allows for building the
command step by step, and also to query the exact command that will be sent to
the Hg CS.

A pro of this method is that it allows you to obtain the exact command string
that will be passed to Mercurial before it is performed, by calling the
CmdLine() method of HgCmd. This could be handy for logging, or for showing
feedback to the user in a GUI program. (You could even call CmdLine() several
times, and show the building of the command step by step.)

An example (also see examples/example2.go):

  opts := make([]hg.HgOption, 2)
  var lim Limit = 2
  opts[0] = lim
  var verb Verbose = true
  opts[1] = verb
  cmd, _ := hg.NewHgCmd("log", opts, nil, new(hg.logOpts))
  cmd.SetOptions(opts)
  cmdline, _ := cmd.CmdLine(hgcl)
  fmt.Printf("%s\n", cmdline) // output: log --limit 2 -v
  cmd.Exec(hgcl)

As you can see, this way requires some more coding.

The source code will also show you that the HgCmd type is indeed used as the
underlying type for the convenience HgClient commands, in all the
New<hg-command>Cmd() constructors.

Commands - ExecCmd

The HgClient type has an extra method ExecCmd(), allowing you to pass a fully
custom built command to Hg. It accepts a string slice that is supposed to
contain all the elements of the complete command, as you would type it at the
command line.

It could be a convenient way for performing commands that are not yet
implemented in gohg, or to make use of extensions to Hg (for which gohg offers
no support (yet?)).

An example (also see examples/example3.go):

  hgcmd := []string{"log", "--limit", "2"}
  result, err := hgcl.ExecCmd(hgcmd)

Options and Parameters

Just like on the commandline, options come before parameters.

  opts := []hg.HgOption{hg.Verbose(true), hg.Limit(2)}
  params := []string{"mytool.go"}
  log, err := hgcl.Log(opts, params)

Options to commands use the same name as the long form of the Mercurial option
they represent, but start with the necessary capital letter. An options value
can be of type bool, int or string. You just pass the value as the parameter to
the option (= type conversion of the value to the option type). You can pass any
number of options, as the elements of a slice. Options can occur more than once
if appropriate (see the ones marked with '[+]' in the Mercurial help).

  log, err := hgcl.Log([]hg.HgOption{hg.Verbose(true)}, nil)                 // bool
  log, err := hgcl.Log([]hg.HgOption{hg.Limit(2)}, nil)                      // int
  log, err := hgcl.Log([]hg.HgOption{hg.User("John Doe"), hg.User("me")}, nil)  // string, repeated option

Parameters are used to provide any arguments for a command that are not options.
They are passed in as a string or a slice of strings, depending on the command.
These parameters typically contain revisions, paths or filenames and so.

  log, err := hgcl.Log(nil, []string{"myfile"})
  heads, err := hgcl.Heads(nil, []string{"foobranch"})

The gohg tool only checks if the options the caller gives are valid for that
command. It does not check if the values are valid for the combination of that
command and that option, as that is done by Mercurial. No need to implement that
again. If an option is not valid for a command, it is silently ignored, so it is
not passed to the Hg CS.

A few options are not implemented, as they seemed not relevant for use with this
tool (for instance: the global --color option, or the --print0 option for
status).

Error handling

The gohg tool only returns errors, with an as clear as possible message, and
never uses log.Fatal() nor panics, even if those may seem appropriate. It leaves
it up to the caller to do that eventually. It's not up to this library to decide
whether to do a retry or to abort the complete application.

Limitations

▪ The following config settings are fixated in the code (at least for now):
  encoding=utf-8
  ui.interactive=False
  extensions.color=!

▪ As mentioned earlier, passing config info is not implemented yet.

▪ Currently the only  support for extensions to Mercurial is through
the ExecCmd method.

▪ If multiple Hg CSs are used against the same repo, it is up to Mercurial
to handle this correctly.

▪ Mercurial is always run in english. Internationalization is not necessary
here, as the conversation with Hg is internal to the application.

Feedback

Please note that this tool is still in it's very early stages.
If you have suggestions or requests, or experience any problems, please use the
issue tracker at https://bitbucket.org/gohg/gohg/issues?status=new&status=open.
Or you could send a patch or a pull request.

License

Copyright 2012-2014, The gohg Authors. All rights reserved.

Use of this source code is governed by a BSD style license that can be found in
the LICENSE.md file.

*/
package gohg
