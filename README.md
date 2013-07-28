# gohg - a Go client library for Mercurial

This project provides a [Go](http://golang.org) client library for the
[Mercurial](http://mercurial.selenic.com) dvcs, using it's
[Command Server](http://mercurial.selenic.com/wiki/CommandServer).
The Command Server is available as of Mercurial version 1.9.

The purpose is to make working with Mercurial as transparent and Go-like as
possible. So gohg only passes your commands to Mercurial, and does not check
them for validity, as that is done quite well by Mercurial itself.
It returns any results - and/or error messages - from Mercurial 'as is'.
Well, some results will eventually be wrapped in a more Go-like form, like
changeset info for instance.

It is as much an occasion for me to experience working with Go :) .

Please note that this tool is still in it's very early stages.
If you have suggestions or requests please use the
[issue tracker](https://bitbucket.org/gohg/gohg/issues?status=new&status=open).

### Compatibility

I started developing the gohg library with Go1 (go1.0.1) on Windows 7,
and upgraded to Go1.0.2 and Go1.0.3.
It is also tested on Ubuntu 12.04 with the same version(s) of Go.
When it was out I started using the Go1.1.1 version, on both platforms, and I
didn't have to change anything.

### Dependencies

Only Go and it's standard library. (Just using
[gocov](https://github.com/axw/gocov) for checking test coverage.)

### Installation

At the commandline type:

    go get [-u] bitbucket.org/gohg/gohg
    go test -v bitbucket.org/gohg/gohg

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
            log.Println(err)
        }
        fmt.Printf("\"summary\" for repo %s:\n%s\n", hc.RepoRoot(), summ)

        var l []byte
        files := []string{}
        if l, err = hc.Log([]Option{Limit(2)}, files); err != nil {
            fmt.Println(err)
        }
        fmt.Printf("\"log -l 2\" for repo %s:\n%s\n", hc.RepoRoot(), l)
    }

### License

Copyright 2012-2013, The gohg Authors. All rights reserved.
Use of this source code is governed by a BSD style license
that can be found in the LICENSE.md file.
