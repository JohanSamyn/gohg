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
	"encoding/binary"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const t = "hg serve [OPTION]"

var server *exec.Cmd
var pout io.ReadCloser
var pin io.ReadCloser

type hgMsg struct {
	Ch   string
	Ln   uint
	Data string
}

type hgCmd struct {
	Cmd  string
	Ln   uint
	Args string
}

func init() {
	// fmt.Println("Hello from gohg!")
} // init()

func Connect(hg string, repo string, config []string) error {

	// for example:
	// server = exec.Command("M:\\DEV\\hg-stable\\hg",	// the Hg command
	// 		"-R", "C:\\DEV\\go\\src\\golout\\",			// the repo
	// 		"--config", "ui.interactive=True",			// mandatory settings
	// 		"--config", "extensions.color=!",			// more settings (for Windows)
	// 		"serve", "--cmdserver", "pipe")				// start the Command Server

	if hg == "" {
		// Use the default Mercurial.
		hg = "hg"
	}

	var err error
	var oriRepo string
	sep := string(os.PathSeparator)
	// The Hg Command Server needs a repository.
	if repo == "" {
		if repo == "" {
			repo, err = os.Getwd()
			oriRepo = repo
		} else {
			repo = strings.TrimRight(repo, sep)
		}
	}
	// If we do not find a Hg repo in the current dir, we search for one
	// up the path, in case we're deeper in it's working copy.
	for {
		_, err = os.Stat(repo + sep + ".hg")
		if err == nil {
			break
		}
		var dir, file string
		dir, file = filepath.Split(repo)
		if dir == "" || file == "" {
			repo = ""
			break
		}
		repo = dir
	}
	if err != nil || repo == "" {
		log.Fatal("could not find a Hg repository at: " + oriRepo)
	}

	// if len(config) > 0 {
	// 	var cfg string
	// 	for i := 0; i < range(config) {
	// 		cfg = cfg + "," + config[i]
	// 	}
	// 	cmd = cmd + "," + cfg
	// }

	server = exec.Command(hg)
	server.Args = append(server.Args, "-R", repo)
	server.Args = append(server.Args,
		// These arguments are fixed.
		"--config", "ui.interactive=True",
		"--config", "extensions.color=!",
		"serve", "--cmdserver", "pipe")

	var pout io.ReadCloser
	pout, err = server.StdoutPipe()
	if err != nil {
		log.Fatal("could not connect StdoutPipe: ", err)
	}
	var pin io.WriteCloser
	pin, err = server.StdinPipe()
	if err != nil {
		log.Fatal("could not connect StdinPipe: ", err)
	}
	if err := server.Start(); err != nil {
		log.Fatal("could not start the Hg Command Server: ", err)
	}
	// temporarily, fo avoid compilation error that pin is not used
	_, err = pin.Write(make([]byte, 0))
	// fmt.Printf("pout=%v,    pin=%v\n", pout, pin)

	s := make([]byte, 1+4+1024)
	_, err = pout.Read(s)
	if err != io.EOF && err != nil {
		log.Fatal(err)
	}
	if len(s) == 0 {
		log.Fatal("no data received from Hg Command Server")
	}
	var ln uint32
	buf := bytes.NewBuffer(s[1:5])
	err = binary.Read(buf, binary.BigEndian, &ln)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	t := "capabilities:"
	l := len(t)
	// fmt.Println("[[" + string(s[5:5+l]) + "]]")
	if string(s[5:5+l]) != t {
		log.Fatal("could not connect a Hg Command Server")
	}

	fmt.Println("Connection established with Hg Command Server at: " + repo)

	return nil

} // Connect()

func Close() error {
	fmt.Println("start of Close()")
	// pout.Close()
	// pin.Close()
	err := server.Wait()
	if err != nil {
		return err
	}
	fmt.Println("before normal return of Close()")
	return nil
} // Close()

func RunCommand() {

} // RunCommand()
