// Copyright (C) 2012, Johan P. P. Samyn <johan.samyn@gmail.com> All rights reserved.
// Use of this source code is governed by the Simplified BSD License
// that can be found in the LICENSE.txt file.

// Package gohg is a Go client library for using the Mercurial dvcs
// via it's Command Server.
//
// For Mercurial see: http://mercurial/selenic.com/wiki
// For the Hg Command Server see: http://mercurial.selenic.com/wiki/CommandServer
package gohg

import (
	"encoding/binary"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

var repo string
var err error
var libdir string
var logfile string

var hgserver *exec.Cmd

// The in and out pipe ends are to be considered from the point of view
// of the Hg Command Server instance. So closing pin closes the Hg CS.
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

// init takes care of some householding, namely preparing a logfile where
// all communication between this lib and the Hg CS can be logged.
func init() {
	var exedir string
	exedir = path.Dir(os.Args[0])
	exedir, err = filepath.Abs(exedir)
	if err != nil {
		log.Fatal("Could not determine path for the gohg.log logfile.")
	}
	logfile = exedir + string(os.PathSeparator) + "gohg.log"
} // init()

// Connect establishes the connection with the Mercurial CommandServer.
//
// If anything went wrong when trying to make the connection,
// an error is returned.
//
// Parameters:
//		hg			The command to run mercurial. Optional.
//		reponame	The folder of the Hg repository to work on. Optional.
//					When blanc the folder where the program is run is used
//					(see locateRepository).
//		config		Configuration settings that will be added to the necessary
//					default/fixed settings (see composeHgConfig() for more).
// Returns an error if the connection could not be established flawlessly.
func Connect(hg string, reponame string, config []string) error {

	// for example:
	// hgserver = exec.Command("M:\\DEV\\hg-stable\\hg",	// the Hg command
	// 		"-R", "C:\\DEV\\go\\src\\golout\\",				// the repo
	// 		"--config", "ui.interactive=True",				// mandatory settings
	// 		"--config", "extensions.color=!",				// more settings (for Windows)
	// 		"serve", "--cmdserver", "pipe")					// start the Command Server

	// Maybe accept a channel as an extra argument for sending the logging to ?
	// And if it's nil, log into a textfile in the folder of this lib.
	// Also do not override that logfile every launch.
	// Maybe even do this in the init() function ?

	if hgserver != nil {
		return errors.New("A Hg Command Server is already connected.")
	}

	if hg == "" {
		// Let the OS determine what Mercurial to run
		// for this machine/user combination.
		hg = "hg"
	}
	// fmt.Printf("hg: %s\n", hg)

	// The Hg Command Server needs a repository.
	repo, err = locateRepository(reponame)
	if err != nil {
		return err
	}
	if repo == "" {
		return errors.New("could not find a Hg repository at: " + reponame)
	}
	// fmt.Printf("repo: %s\n", repo)

	// Maybe we can also offer the possibility of a config file?
	// f.i.: a file gohg.cfg in the same folder as the gohg.exe,
	// and a section per repo, and one "general" section.
	// Or maybe just a [gohg] section in one of the 'normal' Hg config files ?

	var hgconfig []string
	hgconfig = composeHgConfig(hg, repo, config)
	// fmt.Printf("hgconfig: %v\n", hgconfig)

	hgserver = exec.Command(hg)
	hgserver.Args = hgconfig
	hgserver.Dir = repo
	// fmt.Printf("hgserver: %v\n", hgserver)

	pout, err = hgserver.StdoutPipe()
	if err != nil {
		return errors.New("could not connect StdoutPipe: " + err.Error())
	}
	pin, err = hgserver.StdinPipe()
	if err != nil {
		log.Fatal("could not connect StdinPipe: " + err.Error())
	}

	if err := hgserver.Start(); err != nil {
		return errors.New("could not start the Hg Command Server: " + err.Error())
	}

	err = readHelloMessage()
	if err != nil {
		return err
	}

	fmt.Println("Connected with Hg Command Server at: " + repo)

	return nil

} // Connect()

// locateRepository takes care of assuring we have a Merurial repository
// available for working on via the Hg CommandServer, as that is necessary
// to be able to use the Hg CS.
func locateRepository(reponame string) (string, error) {
	repo = reponame
	sep := string(os.PathSeparator)

	// first make a correct path from repo
	repo, err = filepath.Abs(repo)
	if err != nil {
		return "", errors.New(err.Error() +
			"\ncould not find absolute path for: " + repo)
	}
	repo = filepath.Clean(repo)

	// If we do not find a Hg repo in this dir, we search for one
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
		return "", nil
	}

	return repo, nil

} // locateRepository()

// composeHgConfig handles the different config settings that will be used
// to make the connection with the Hg CS. It concerns settings specially for Hg.
func composeHgConfig(hgcmd string, repo string, config []string) []string {
	var hgconfig []string

	// if len(config) > 0 {
	// 	var cfg string
	// 	for i := 0; i < range(config) {
	// 		cfg = cfg + "," + config[i]
	// 	}
	// 	cmd = cmd + "," + cfg
	// }

	hgconfig = append(hgconfig, hgcmd,
		"-R", repo,
		// These arguments are fixed.
		"--config", "ui.interactive=True",
		"--config", "extensions.color=!",
		"serve", "--cmdserver", "pipe")

	return hgconfig
}

// readHelloMessage reads the special hello message send by the Hg CS.
//
// It has a fixed format, and contains info about the possibilities
// of the Hg CS at hand. It's also a first proof of a working connection.
func readHelloMessage() error {
	const t1 = "capabilities:"
	const t2 = "hg serve [OPTION]"

	s := make([]byte, 1+4+1024)
	_, err = pout.Read(s)
	if err != io.EOF && err != nil {
		return err
	}
	// fmt.Printf("s: %s\n", s)
	if len(s) == 0 {
		return errors.New("no data received from Hg Command Server")
	}
	var ln uint32
	buf := bytes.NewBuffer(s[1:5])
	err = binary.Read(buf, binary.BigEndian, &ln)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	l := len(t2)
	if string(s[0:l]) == t2 {
		return errors.New("this version of Mercurial does not support the Command Server")
	}
	l = len(t1)
	if string(s[5:5+l]) != t1 {
		return errors.New("could not connect a Hg Command Server")
	}
	return nil
}

// Close ends the conection with the Mercurial CommandServer.
//
// In fact it's closing the stdin of the Hg CS that closes connection,
// as per the Hg CS documentation.
func Close() error {
	pout.Close()
	// Closing it's stdin is what really closes the Hg Command Server.
	pin.Close()
	err = hgserver.Wait()
	if err != nil {
		return err
	}
	fmt.Println("Disconnected from Hg Command Server at: " + repo)
	return nil
} // Close()
