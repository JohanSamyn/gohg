gohg - a Go client library for Mercurial
****************************************

This project aims at creating a Go client library for the Mercurial dvcs,
using it's Command Server for better performance.

It's second purpose - by no means less then the first one - is to have a real
project for working with the Go language, which really appeals to me.

So don't expect anything to be ready quickly, cause this will be a learning
process.

Example usage::

    package main

    import (
        . "gohg/gohg_lib"
        "fmt"
        "log"
    )

    func main() {
        var err error
        hgexe := "hg"
        repo := ""

        hc := NewHgClient()
        var cfg []string
        if err = hc.Connect(hgexe, repo, cfg); err != nil {
            log.Fatal(err)
        }
        defer hc.Close()

        var s string
        if s, err = hc.Summary(); err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("[[Summary]]:\n%s", s)
    }
