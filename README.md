# gohg - a Go client library for Mercurial

This project provides a Go client library for the
[Mercurial DCVS](http://mercurial.selenic.com), using it's
[Command Server](http://mercurial.selenic.com/wiki/CommandServer)
for better performance. The Command Server is available as of Mercurial
version 1.9.

### Compatibility

The gohg library is developed with Go1 on Windows 7.
It is also tested on Ubuntu 12.04 with Go1.
The used Go1 is being updated to the latest release shortly after it's
appearance, on both platforms.

### Dependencies

Only Go and it's standard library.

### Installation

At the commandline type:

    go get bitbucket.org/gohg/gohg

### Example

    :::go
    package main

    import (
        "bitbucket.org/gohg/gohg"
        "fmt"
        "log"
    )

    func main() {
        var err error
        hgexe := "hg"
        repo := "/path/to/hgrepo"
        var cfg []string
        hc := gohg.NewHgClient()
        if err = hc.Connect(hgexe, repo, cfg); err != nil {
            log.Fatal(err)
        }
        defer hc.Close()

        var summ []byte
        scmd := gohg.NewSummary()
        if summ, err = hc.Summary(scmd); err != nil {
            log.Println(err)
            return
        }
        fmt.Printf("[[Summary for repo %s]]:\n%s\n", repo, summ)
    }
