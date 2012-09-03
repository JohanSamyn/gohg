// Copyright (c) 2012, Johan P. P. Samyn <johan.samyn@gmail.com> All rights reserved.
// Use of this source code is governed by the Simplified BSD License
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
)

const t = "hg serve [OPTION]"

var server *exec.Cmd
var repo string

// pout and pin are to be considered from the point of view of the
// Hg Command Server instance.
var pout io.ReadCloser
var pin io.WriteCloser

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

func Connect(hg string, repo_arg string, config []string) error {

	// for example:
	// server = exec.Command("M:\\DEV\\hg-stable\\hg",	// the Hg command
	// 		"-R", "C:\\DEV\\go\\src\\golout\\",			// the repo
	// 		"--config", "ui.interactive=True",			// mandatory settings
	// 		"--config", "extensions.color=!",			// more settings (for Windows)
	// 		"serve", "--cmdserver", "pipe")				// start the Command Server

	// Maybe accept a channel as an extra argument for sending the logging to ?
	// And if it's nil, log into a textfile in the folder of this lib.
	// Also do not override that logfile every launch, but insert a timestamp
	// to mark a new run. Maybe even do this in the init() function ?

	if hg == "" {
		// Use the default Mercurial.
		hg = "hg"
	}

	var err error
	var oriRepo string
	sep := string(os.PathSeparator)
	// The Hg Command Server needs a repository.
	repo = repo_arg

	// first make a correct path from repo
	repo, err = filepath.Abs(repo)
	if err != nil {
		log.Fatal("could not find absolute path for: " + repo)
	}
	repo = filepath.Clean(repo)
	oriRepo = repo

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

	pout, err = server.StdoutPipe()
	if err != nil {
		log.Fatal("could not connect StdoutPipe: ", err)
	}
	pin, err = server.StdinPipe()
	if err != nil {
		log.Fatal("could not connect StdinPipe: ", err)
	}
	if err := server.Start(); err != nil {
		log.Fatal("could not start the Hg Command Server: ", err)
	}
	// temporarily, to avoid compilation error that pin is not used
	_, err = pin.Write(make([]byte, 0))

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
	if string(s[5:5+l]) != t {
		log.Fatal("could not connect a Hg Command Server")
	}

	fmt.Println("Connected with Hg Command Server at: " + repo)

	return nil

} // Connect()

func Close() error {
	pout.Close()
	pin.Close()
	err := server.Wait()
	if err != nil {
		return err
	}
	fmt.Println("Disconnected from Hg Command Server at: " + repo)
	return nil
} // Close()

func RunCommand() {

} // RunCommand()
