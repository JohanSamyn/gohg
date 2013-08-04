# gohg - a Go client library for Mercurial

### What it is

This project provides a [Go](http://golang.org) client library for the
[Mercurial](http://mercurial.selenic.com) dvcs, using the
[Command Server](http://mercurial.selenic.com/wiki/CommandServer).

The purpose is to make working with Mercurial as transparent and Go-like as
possible. So gohg only passes your commands to Mercurial, and does not check
them for validity, as that is already done quite well by Mercurial itself.
It returns any results - and/or error messages - from Mercurial 'as is'.
Well, some results will eventually be wrapped in a more Go-like form, like
changeset info for instance.

It is as much an occasion for me to experience working with Go :) .

### Features

- Choice of what hg command/version to use (default: 'hg').
- Choice of repo to work on.
- Possibility to provide a path for creating a new repo first and then connect to it. _TODO_
- Commands implemented so far: add, addremove, annotate, archive, branches, clone, commit, diff, export, forget, heads, identify, init, log, manifest, merge, pull, push, remove, serve, showconfig, status, summary, tags, tip, update, verify, version.
- All options implemented for all implemented commands (except global --color and --print0 for status).
- Possibility to obtain the full commandstring that was passed to Mercurial.
- Possibility to pass _any_ command to Hg, allowing for use of extensions, and future commands when they are not yet implemented by gohg.
- Commands returning changeset info do that in a go-like way, using a slice of structs, where each element is a changeset. _TODO_
- Possibility to ask for 'raw' Hg output. _TODO_

### Compatibility

###### Mercurial

For Mercurial any version starting with 1.9 should be ok, cause that's the one
where the Command Server was introduced. If you send wrong options to it through
gohg, or commands not yet supported in your Hg version, you'll simply get back
an error from Hg, as gohg does not check them.
But on the other hand gohg allows issuing new commands, not yet implemented by
gohg; see the documentation.

###### Go

Currently gohg is developed with Go1.1.1. Though I started with the Go1.0
versions, I can't remember having had to change one or two minor things when
moving to Go1.1.1.

###### Platform

I'm developing and testing both on Windows 7 and Ubuntu 12.04. But I suppose
it should work on any other platform that supports Hg and Go.

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
        . "bitbucket.org/gohg/gohg"
        "fmt"
        "log"
    )

    func main() {
        var err error
        hc := NewHgClient()
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
        if l, err = hc.Log([]Option{Limit(2)}, files); err != nil {
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

Copyright 2012-2013, The gohg Authors. All rights reserved.
Use of this source code is governed by a BSD style license
that can be found in the LICENSE.md file.
