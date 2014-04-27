# gohg - a Go client library for Mercurial

[![Build Status](https://drone.io/bitbucket.org/gohg/gohg/status.png)](https://drone.io/bitbucket.org/gohg/gohg/latest)

### What it is

This project provides a [Go](http://golang.org) client library for the
[Mercurial](http://mercurial.selenic.com) dvcs, using its
[Command Server](http://mercurial.selenic.com/wiki/CommandServer).

The purpose is to make working with Mercurial as transparent and Go-like as
possible. So gohg only passes your commands to Mercurial, and does not check
them for validity, as that is already done quite well by Mercurial itself.
It returns any results - and/or error messages - from Mercurial 'as is'.
Well, some results will eventually be wrapped in a more Go-like form, like
changeset info for instance.

It is as much an occasion for me to experience working with Go :) .


### Features

- Choice of what hg command/version to use (default = 'hg').
- Choice of repo to work on (default = '.').
- First create a new repo on the fly before connecting to it (see the 4th param to Connect()).
- Commands implemented so far: add, addremove, annotate, archive, branches, clone, commit, diff, export, forget, heads, identify, init, log, manifest, merge, pull, push, remove, serve, showconfig, status, summary, tags, tip, update, verify, version.
- Options implemented so far: all options for all implemented commands (except global --color and --print0 for status).
- Obtain the full commandstring that was passed to Mercurial (for showing in the GUI f.i.).
- Pass _any_ command to Hg, allowing for use of extensions, and future commands when they are not yet implemented by gohg.
- Commands returning changeset info do that in a Go-like way, using a slice of structs, where each element is a changeset. _TODO_
- Ask for 'raw' Hg output (as shown on stdout when issuing a hg command from a terminal). _TODO_

### Compatibility

###### Mercurial

For Mercurial any version starting from 1.9 should be ok, cause that's the one
where the Command Server was introduced. If you send wrong options to it through
gohg, or commands or options not yet supported (or obsolete) in your Hg version,
you'll simply get back an error from Hg itself, as gohg does not check them.
But on the other hand gohg allows issuing new commands, not yet implemented
by gohg; see the documentation.

###### Go

Currently gohg is developed with Go1.2.1. Though I started with the
Go1.0 versions, I can't remember having had to change more than one or two minor
things when moving to Go1.1.1. Updating to Go1.1.2 required no changes at all.
I had an issue though with Go1.2, on Windows only, causing some tests using
os.exec.Command to fail. I'll have to look into that further, to find out if I
should report a bug.

###### Platform

I'm developing and testing both on Windows 7 and Ubuntu 12.04/13.04/13.10. But I suppose
it should work on any other platform that supports Mercurial and Go.

### Dependencies

Only Go and it's standard library. And Mercurial should be installed of course.

### Installation

At the commandline type:

    go get [-u] bitbucket.org/gohg/gohg
    go test [-v] bitbucket.org/gohg/gohg

to have gohg available in your GOPATH.

### Example

Run this example program in a terminal from a folder containing a Mercurial
repository. Or pass the repo of your choice as the second parameter for
Connect(). (You can find the source in the \examples folder as readme-test.go,
along with a few others.)

    :::go
    package main

    import (
        hg "bitbucket.org/gohg/gohg"
        "fmt"
        "log"
    )

    func main() {
        var err error
        hc := hg.NewHgClient()
        if err = hc.Connect("", "", nil); err != nil {
            log.Fatal(err)
        }
        defer hc.Disconnect()

        var summ []byte
        if summ, err = hc.Summary(nil, nil); err != nil {
            fmt.Println(err)
        }
        fmt.Printf("\"summary\" for repo %s:\n%s\n", hc.RepoRoot(), summ)

        var l []byte
        files := []string{}
        if l, err = hc.Log([]hg.Option{hg.Limit(2)}, files); err != nil {
            fmt.Println(err)
        }
        fmt.Printf("\"log -l 2\" for repo %s:\n%s\n", hc.RepoRoot(), l)
    }

### Documentation

For more details on how to use gohg have a look at
[the documentation](http://godoc.org/bitbucket.org/gohg/gohg).

### Feedback

Please note that this tool is still in it's very early stages.
If you have suggestions or requests, or experience any problems, please use the
[issue tracker](https://bitbucket.org/gohg/gohg/issues?status=new&status=open).
Or you could send a patch or a pull request.

### License

Copyright 2012-2014, The gohg Authors. All rights reserved.
Use of this source code is governed by a BSD style license
that can be found in the LICENSE.md file.
