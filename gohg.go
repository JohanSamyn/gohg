// Copyright (C) 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

// Package gohg is a Go client library for using the Mercurial dvcs
// via it's Command Server.
//
// For Mercurial see: http://mercurial/selenic.com/wiki.
//
// For the Hg Command Server see: http://mercurial.selenic.com/wiki/CommandServer.
package gohg

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

// Type hgclient will act as a (kind of) object for working with the Hg CS
// from any program using this gohg client lib.
// It will get a bunch of attributes and methods to make working with it
// as go-like as possible. It might even get a few channels for communications.
type hgclient struct {
	capabilities []string
	encoding     string
	repo         string
}

// hgMsg is what we receive from the Hg CS
type hgMsg struct {
	Ch   string
	Ln   uint
	Data string
}

// hgCmd is what we send to the Hg CS
type hgCmd struct {
	Cmd  string
	Ln   uint
	Args string
}

// The in and out pipe ends are to be considered from the point of view
// of the Hg Command Server instance.
var pout io.ReadCloser
var pin io.WriteCloser
var err error
var logfile string
var hgserver *exec.Cmd
var HgClient hgclient

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
// Arguments:
//	hgexe
//		The command to run mercurial. Optional.
//	reponame
//		The folder of the Hg repository to work on. Optional.
//		When blanc the folder where the program is run is used
//		(see function locateRepository()).
//	config
//		Configuration settings that will be added to the necessary
//		/fixed settings (see composeHgConfig() for more).
//
// Returns an error if the connection could not be established properly.
func Connect(hgexe string, reponame string, config []string) error {

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

	if hgexe == "" {
		// Let the OS determine what Mercurial to run
		// for this machine/user combination.
		hgexe = "hg"
	}

	// The Hg Command Server needs a repository.
	HgClient.repo, err = locateRepository(reponame)
	if err != nil {
		return err
	}
	if HgClient.repo == "" {
		return errors.New("could not find a Hg repository at: " + reponame)
	}

	// Maybe we can also offer the possibility of a config file?
	// f.i.: a file gohg.cfg in the same folder as the gohg.exe,
	// and a section per repo, and one "general" section.
	// Or maybe just a [gohg] section in one of the 'normal' Hg config files ?

	var hgconfig []string
	hgconfig = composeHgConfig(hgexe, HgClient.repo, config)

	hgserver = exec.Command(hgexe)
	hgserver.Args = hgconfig
	hgserver.Dir = HgClient.repo

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

	fmt.Println("Connected with Hg Command Server at: " + HgClient.repo)

	return nil

} // Connect()

// locateRepository assures we have a Mercurial repository available,
// which is required for working with the Hg CommandServer.
func locateRepository(reponame string) (string, error) {
	repo := reponame
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
// to make the connection with the Hg CS. It concerns specific Hg settings.
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
		"--cwd", repo,
		"-R", repo,
		// These arguments are fixed.
		// "--config", "ui.interactive=True",
		"--config", "ui.interactive=False",
		"--config", "extensions.color=!",
		"serve", "--cmdserver", "pipe")

	return hgconfig
} // composeHgConfig()

// readHelloMessage reads the special hello message send by the Hg CS.
//
// It has a fixed format, and contains info about the possibilities
// of the Hg CS at hand. It's also a first proof of a working connection.
func readHelloMessage() error {
	s := make([]byte, 5)
	_, err = pout.Read(s)
	if err != io.EOF && err != nil {
		return err
	}
	if len(s) == 0 {
		return errors.New("no hello message data received from Hg CommandServer")
	}
	const t1 = "hg se" // hg send: "hg serve [OPTION]"
	if string(s[0:len(t1)]) == t1 {
		return errors.New("this version of Mercurial does not support the CommandServer")
	}
	ch := string(s[0])
	if ch != "o" {
		return errors.New("received unexpected channel '" + ch +
			"' for hello message from Hg CommandServer")
	}
	var ln uint32
	ln, err = calcLengthReceivedData(s[1:5])
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	if ln <= 0 {
		return errors.New("received invalid length '" + string(ln) +
			"' for hello message from Hg CommandServer")
	}
	hello := make([]byte, ln)
	_, err = pout.Read(hello)
	if err != io.EOF && err != nil {
		return err
	}
	const t2 = "capabilities:"
	if string(hello[0:len(t2)]) != t2 {
		return errors.New("could not determine the capabilities of the Hg CommandServer")
	}
	if strings.Contains(string(hello), "runcommand") == false {
		log.Fatal("could not detect the 'runcommand' capability")
	}
	attr := strings.Split(string(hello), "\n")
	HgClient.capabilities = strings.Fields(attr[0])[1:]
	HgClient.encoding = strings.Split(attr[1], ": ")[1]
	return nil
} // readHelloMessage()

// Close ends the connection with the Mercurial CommandServer.
//
// In fact it's closing the stdin of the Hg CS that closes the connection,
// as per the Hg CS documentation.
func Close() error {
	pout.Close()
	pin.Close()
	err = hgserver.Wait()
	if err != nil {
		return err
	}
	fmt.Println("Disconnected from Hg Command Server at: " + HgClient.repo)
	return nil
} // Close()

// readFromHg returns the channel and all the data read from it.
// Eventually it returns no (or empty) data but an error.
func readFromHg() (string, []byte, error) {
	var ch string

	// get channel and length
	data := make([]byte, 5)
	_, err = pout.Read(data)
	if err != io.EOF && err != nil {
		return ch, data, err
	}
	if data == nil {
		return ch, data, errors.New("no data read")
	}
	ch = string(data[0])
	if ch == "" {
		return ch, data, errors.New("no channel read")
	}

	// get the uint that the Hg CS sent us as the length value
	var ln uint32
	ln, err = calcLengthReceivedData(data[1:5])
	if err != nil {
		return ch, data, errors.New("binary.Read failed:" + string(err.Error()))
	}

	// now get ln bytes of data
	data = make([]byte, ln)
	_, err = pout.Read(data)
	if err != io.EOF && err != nil {
		return ch, data, err
	}

	return ch, data, nil
} // readFromHg()

// sendToHg writes data to the Hg CS,
// returning an error if something went wrong.
func sendToHg(cmd string, args []byte) error {
	cmd = strings.TrimRight(cmd, "\n") + "\n"
	lc := len(cmd)
	la := len(args)
	var l int
	if la > 0 {
		l = lc + 4 + la
	} else {
		// in case cmd == "getencoding" f.i.
		l = lc
	}
	data := make([]byte, l)

	// send the command
	copy(data[0:lc], cmd)

	if la > 0 {
		// send the length of the command arguments
		ln := uint32(len(args))
		wbuf := new(bytes.Buffer)
		err = binary.Write(wbuf, binary.BigEndian, ln)
		if err != nil {
			return errors.New("binary.Write failed: " + string(err.Error()))
		}
		b := make([]byte, 4)
		_, err = io.ReadFull(wbuf, b)
		if err != nil {
			return errors.New("io.ReadFull failed: " + string(err.Error()))
		}
		copy(data[lc:lc+4], b)

		// send the command arguments
		copy(data[lc+4:lc+4+la], args)
	}

	// perform the actual send to the Hg CS
	var i int
	i, err = pin.Write(data)
	if i != len(data) {
		return errors.New("writing length of data failed: " +
			string(err.Error()))
	}

	return nil
} // sendToHg()

func RunCommand(hgcmd []string) {
	command := "runcommand"
	// command := "getencoding"
	args := []byte(strings.Join(hgcmd, string(0x0)))

	err = sendToHg(command, args)
	if err != nil {
		fmt.Println(err)
		return
	}

	var data []byte
	var buf bytes.Buffer
	var ret int32
	endOfRead := false
	for endOfRead == false {
		var ch string
		ch, data, err = readFromHg()
		if err != nil || ch == "" {
			log.Fatal("readFromHg failed: " + string(err.Error()))
		}
		switch ch {
		case "d":
		case "e":
		case "o":
			buf.WriteString(string(data))
		case "r":
			{
				if command == "getencoding" {
					buf.WriteString(string(data))
				} else {
					ret, err = calcReturncode(data[0:4])
					if err != nil {
						log.Fatal("binary.read failed: " + string(err.Error()))
					}
				}
				endOfRead = true
			}
		case "I":
		case "L":
		default:
			log.Fatal("unexpected channel '" + ch + "' detected")
		} // switch ch
	} // for endOfRead == false
	fmt.Printf("command -> %s\nhgcmd -> %s\ndata ->\n%s\nreturncode -> %d\n",
		command, hgcmd, []byte(buf.String()), ret)
} // RunCommand()

// calcLengthDataReceived converts a 4-byte slice into an unsigned int
func calcLengthReceivedData(s []byte) (uint32, error) {
	var ln int32
	ln, err = calcIntFromBytes(s)
	return uint32(ln), err
}

// calcReturncode converts a 4-byte slice into a signed int
func calcReturncode(s []byte) (int32, error) {
	var ret int32
	ret, err = calcIntFromBytes(s)
	return ret, err
}

// calcIntFromBytes performs the real conversion
func calcIntFromBytes(s []byte) (int32, error) {
	var i int32
	buf := bytes.NewBuffer(s[0:4])
	err := binary.Read(buf, binary.BigEndian, &i)
	return i, err
}
