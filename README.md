# gohg - a Go client library for Mercurial

This project aims at creating a Go client library for the Mercurial dvcs,
using it's Command Server for better performance. The Command Server is
available as of Mercurial version 1.9.

It's second purpose - by no means less then the first one - is to have a real
project for working with the Go language, which really appeals to me.

So don't expect anything to be ready quickly, cause this will be a learning
process.

## Compatibility

The gohg library is developed with the Go1 releases.

## Dependencies

None so far.

## Installation

At the commandline type:
  :::bash
    go get bitbucket.org/gohg/gohg

## Example usage

  :::go
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
        var cfg []string
        hc := NewHgClient()
        if err = hc.Connect(hgexe, repo, cfg); err != nil {
            log.Fatal(err)
        }
        defer hc.Close()

        var s string
        if s, err = hc.Summary(); err != nil {
            log.Println(err)
            return
        }
        fmt.Printf("[[Summary]]:\n%s", s)
    }
