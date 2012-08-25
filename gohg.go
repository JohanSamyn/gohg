// Copyright (c) 2012, Johan P. P. Samyn <johan.samyn@gmail.com> All rights reserved.
// Use of this source code is governed by the BSD-2-clause license
// that can be found in the LICENSE.txt file.

// Package gohg is a Go client library for using the Mercurial dvcs
// using it's Command Server for better performance.
//
// For Mercurial see: http://mercurial/selenic.com/wiki.
// For the Hg Command Server see: http://mercurial.selenic.com/wiki/CommandServer.
package gohg

import (
	"fmt"
)

func init() {
	fmt.Println("Hello from gohg!")
}

func Connect() error {
	// Maybe we can even manage multiple connections ?
	// Question: can we manage more than 1 repo from the same Hg Command Server
	// instance ? If not then I don't think multiple connections will work.
}

func Close() {

}

func RunCommand() {

}
