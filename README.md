# gohg - a Go client library for Mercurial

This project provides a [Go](http://golang.org) client library for the
[Mercurial](http://mercurial.selenic.com) dvcs, using it's
[Command Server](http://mercurial.selenic.com/wiki/CommandServer).
The Command Server is available as of Mercurial version 1.9.

It is as much an occasion for me to learn to work with Go.

Please note that this tool is still in it's very early stages.  
If you have suggestions or requests please use the
[issue tracker](https://bitbucket.org/gohg/gohg/issues?status=new&status=open).

### Compatibility

The gohg library is developed with Go1 (go1.0.3) on Windows 7.
It is also tested on Ubuntu 12.04 with the same version of Go1.

### Dependencies

Only Go and it's standard library. Though I'm using
[gocov](https://github.com/axw/gocov) for checking test coverage.

### Installation

At the commandline type:

    go get [-u] bitbucket.org/gohg/gohg
    go install bitbucket.org/gohg/gohg
    go test -v bitbucket.org/gohg/gohg

### Example

Run this example from a folder containing a Mercurial repository.  
(You can find the source in the repo as examples/readme-test.go,
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
        if summ, err = hc.Summary(); err != nil {
            log.Println(err)
            return
        }
        fmt.Printf("[[Summary for repo %s]]:\n%s\n", hc.RepoRoot(), summ)
    }

### License

Copyright 2012, The gohg Authors. All rights reserved.  
Use of this source code is governed by a BSD style license
that can be found in the LICENSE file.
