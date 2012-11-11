# gohg - a Go client library for Mercurial

This project aims at creating a Go client library for the
[Mercurial DCVS](http://mercurial.selenic.com), using it's
[Command Server](http://mercurial.selenic.com/wiki/CommandServer)
for better performance. The Command Server is available as of Mercurial
version 1.9.

It's second purpose - by no means less then the first one - is to have a real
project for working with the Go language, which really appeals to me.

So don't expect anything to be ready quickly, as this will be a learning process.

### Compatibility

The gohg library is developed with Go1 on Windows 7.
It is also tested on Ubuntu 12.04 with Go1.
Go1 is being updated to the latest release shortly after it's appearance,
on both platforms.

### Dependencies

Only Go and it's standard library; at least for now.

### Installation

At the commandline type:

    go get bitbucket.org/gohg/gohg/gohg_lib

### Example

    :::go
    package main

    import (
        "bitbucket.org/gohg/gohg/gohg_lib"
        "fmt"
        "log"
    )

    func main() {
        var err error
        hgexe := "hg"
        repo := "/path/to/hgrepo"
        var cfg []string
        hc := gohg_lib.NewHgClient()
        if err = hc.Connect(hgexe, repo, cfg); err != nil {
            log.Fatal(err)
        }
        defer hc.Close()

        var s string
        if s, err = hc.Summary(); err != nil {
            log.Println(err)
            return
        }
        fmt.Printf("[[Summary for repo %s]]:\n%s", repo, s)
    }
