// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

// Package gohg is a Go client library for using the Mercurial dvcs
// via it's Command Server.
//
// For Mercurial see: http://mercurial/selenic.com/wiki.
//
// For the Hg Command Server see: http://mercurial.selenic.com/wiki/CommandServer.
package gohg_lib

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Type HgClient will act as the entrypoint through which all interaction
// with the Mercurial Command Server will take place.
type HgClient struct {
	hgserver *exec.Cmd
	// pin: the pipe that gets commands into the Hg CS
	pin io.WriteCloser
	// pout: the pipe that brings data out of the Hg CS
	pout         io.ReadCloser
	hgPath       string   // which hg is used ?
	capabilities []string // as per the hello message
	encoding     string   // as per the hello message
	repo         string   // the full path to the Hg repo
	hgVersion    string   // the version number only
	// config       []string
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

var err error
var logfile string

// NewHgClient creates a new instance of the client object for working with the
// Hg Command Server.
func NewHgClient() *HgClient {
	var hgclient = new(HgClient)
	return hgclient
}

// Connect establishes the connection with the Mercurial Command Server.
//
// Arguments:
//	hgexe
//		The command to run mercurial. Optional. The 'hg' command will be used
//		when not provided. This allows to run a specific version of Mercurial.
//	reponame
//		The folder of the Hg repository to work on. Optional.
//		When blanc the folder where the program is run is used
//		(see function locateRepository()).
//	config
//		Configuration settings that will be added to the necessary
//		fixed settings (see composeHgConfig() for more).
//
// Returns an error if the connection could not be established properly.
func (hgcl *HgClient) Connect(hgexe string, reponame string, config []string) error {

	// for example:
	// hgcl.hgserver =
	//		exec.Command("M:/DEV/hg-stable/hg",	// the Hg command
	// 		"-R", "C:/DEV/go/src/golout/",		// the repo
	// 		"--config", "ui.interactive=True",	// mandatory settings
	// 		"--config", "extensions.color=!",	// more settings (for Windows)
	// 		"serve", "--cmdserver", "pipe")		// start the Command Server

	// Maybe accept a channel as an extra argument for sending the logging to ?
	// And if it's nil, log into a textfile in the folder of this lib.
	// Also do not override that logfile every launch.
	// Maybe even do this in the init() function ?

	if hgcl.hgserver != nil {
		return errors.New("Connect(): already running a Hg Command Server for " + hgcl.repo)
	}

	hgcl.hgPath = hgexe
	if hgcl.hgPath == "" {
		// Let the OS determine what Mercurial to run
		// for this machine/user combination.
		hgcl.hgPath = "hg"
	}

	// The Hg Command Server needs a repository.
	hgcl.repo, err = locateRepository(reponame)
	if err != nil {
		return err
	}
	if hgcl.repo == "" {
		return errors.New("Connect(): could not find a Hg repository at: " + reponame)
	}

	var hgconfig []string
	hgconfig = composeHgConfig(hgexe, hgcl.repo, config)

	hgcl.hgserver = exec.Command(hgexe)
	hgcl.hgserver.Args = hgconfig
	hgcl.hgserver.Dir = hgcl.repo

	hgcl.pout, err = hgcl.hgserver.StdoutPipe()
	if err != nil {
		return errors.New("Connect(): could not connect StdoutPipe: " + err.Error())
	}
	hgcl.pin, err = hgcl.hgserver.StdinPipe()
	if err != nil {
		log.Fatal("Connect(): could not connect StdinPipe: " + err.Error())
	}

	if err := hgcl.hgserver.Start(); err != nil {
		return errors.New("Connect(): could not start the Hg Command Server: " + err.Error())
	}

	err = readHelloMessage(hgcl)
	if err != nil {
		return err
	}

	err = validateCapabilities(hgcl)
	if err != nil {
		return err
	}

	hgcl.hgVersion, err = hgcl.Version()
	if err != nil {
		log.Fatal("from HgVersion() : " + string(err.Error()))
	}

	return nil

} // Connect()

// Close ends the connection with the Mercurial Command Server.
//
// In fact it's closing the stdin of the Hg CS that closes the connection,
// as per the Hg CS documentation.
func (hgcl *HgClient) Close() error {
	if hgcl.hgserver == nil {
		log.Println("Close(): Trying to close a closed hgserver.")
		return nil
	}

	hgcl.pin.Close()
	hgcl.pout.Close()

	defer func() { hgcl.hgserver = nil }()

	err = hgcl.hgserver.Wait()
	if err != nil {
		return err
	}
	return nil
} // Close()

// locateRepository assures we have a Mercurial repository available,
// which is required for working with the Hg Command Server.
func locateRepository(reponame string) (string, error) {
	repo := reponame
	sep := string(os.PathSeparator)

	// first make a correct path from repo
	repo, err = filepath.Abs(repo)
	if err != nil {
		return "", errors.New(err.Error() +
			"\ncould not determine absolute path for: " + repo)
	}
	repo = filepath.Clean(repo)

	// If we do not find a Hg repo in this dir, we search for one
	// up the path, in case we're deeper in it's working copy.
	for {
		_, err = os.Stat(repo + sep + ".hg")
		if err == nil {
			break
		}
		var file string
		repo, file = filepath.Split(repo)
		if repo == "" || file == "" {
			break
		}
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
func readHelloMessage(hgcl *HgClient) error {
	s := make([]byte, 5)
	_, err = hgcl.pout.Read(s)
	if err != io.EOF && err != nil {
		return err
	}
	if len(s) == 0 {
		return errors.New("no hello message data received from Hg Command Server")
	}
	const t1 = "hg se" // hg returned: "hg serve [OPTION]"
	if string(s[0:len(t1)]) == t1 {
		log.Fatal(errors.New(
			"Need at least version 1.9 of Mercurial to use the Command Server." +
				" Used hgexe: '" + hgcl.HgPath() + "'\n"))
	}
	ch := string(s[0])
	if ch != "o" {
		return errors.New("received unexpected channel '" + ch +
			"' for hello message from Hg Command Server")
	}
	var ln uint32
	ln, err = calcDataLength(s[1:5])
	if err != nil {
		fmt.Println("readHelloMessage(): binary.Read failed:", err)
	}
	if ln <= 0 {
		return errors.New("received invalid length '" + string(ln) +
			"' for hello message from Hg Command Server")
	}
	hello := make([]byte, ln)
	_, err = hgcl.pout.Read(hello)
	if err != io.EOF && err != nil {
		return err
	}
	const t2 = "capabilities:"
	if string(hello[0:len(t2)]) != t2 {
		return errors.New("could not determine the capabilities of the Hg Command Server")
	}
	attr := strings.Split(string(hello), "\n")
	hgcl.capabilities = strings.Fields(attr[0])[1:]
	hgcl.encoding = strings.Split(attr[1], ": ")[1]
	return nil
} // readHelloMessage()

func validateCapabilities(hgcl *HgClient) error {
	var ok bool
	for _, c := range hgcl.capabilities {
		if c == "runcommand" {
			ok = true
			break
		}
	}
	if ok == false {
		log.Fatal("could not detect the 'runcommand' capability")
	}
	return nil
}

// readFromHg returns the channel and all the data read from it.
// Eventually it returns no (or empty) data but an error.
func readFromHg(hgcl *HgClient) (string, []byte, error) {
	var ch string

	// get channel and length
	data := make([]byte, 5)
	_, err = hgcl.pout.Read(data)
	if err != io.EOF && err != nil {
		return ch, data, err
	}
	if data == nil {
		return ch, data, errors.New("readFromHg(): no data read")
	}
	ch = string(data[0])
	if ch == "" {
		return ch, data, errors.New("readFromHg(): no channel read")
	}

	// get the uint that the Hg CS sent us as the length value
	var ln uint32
	ln, err = calcDataLength(data[1:5])
	if err != nil {
		return ch, data, errors.New("readFromHg(): binary.Read failed:" +
			string(err.Error()))
	}

	// now get ln bytes of data
	data = make([]byte, ln)
	_, err = hgcl.pout.Read(data)
	if err != io.EOF && err != nil {
		return ch, data, err
	}

	return ch, data, nil
} // readFromHg()

// sendToHg writes data to the Hg CS,
// returning an error if something went wrong.
func sendToHg(hgcl *HgClient, cmd string, args []byte) error {
	cmd = strings.TrimRight(cmd, "\n") + "\n"
	lc := len(cmd)
	la := len(args)
	l := lc // in case cmd == "getencoding" f.i.
	if la > 0 {
		l = l + 4 + la
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
			return errors.New("sendToHg(): binary.Write failed: " + string(err.Error()))
		}
		b := make([]byte, 4)
		_, err = io.ReadFull(wbuf, b)
		if err != nil {
			return errors.New("sendToHg(): io.ReadFull failed: " + string(err.Error()))
		}
		copy(data[lc:lc+4], b)

		// send the command arguments
		copy(data[lc+4:lc+4+la], args)
	}

	// perform the actual send to the Hg CS
	var i int
	i, err = hgcl.pin.Write(data)
	if i != len(data) {
		return errors.New("sendToHg(): writing data failed: " + string(err.Error()))
	}

	return nil
} // sendToHg()

// HgEncoding returns the servers encoding on the result channel.
// Currently only UTF8 is supported.
func (hgcl *HgClient) HgEncoding() (string, error) {
	var encoding []byte
	encoding, _, err = runInHg(hgcl, "getencoding", []string{})
	return string(encoding), err
}

// run allows to run a Mercurial command in the Hg Command Server.
// You can only run 'hg' commands that are available in this library.
func (hgcl *HgClient) run(hgcmd []string) ([]byte, int32, error) {
	var data []byte
	var ret int32
	data, ret, err = runInHg(hgcl, "runcommand", hgcmd)
	return data, ret, err
}

// runInHg sends a command to the Hg CS (using sendToHg),
// and fetches the result (using readFromHg).
func runInHg(hgcl *HgClient, command string, hgcmd []string) ([]byte, int32, error) {
	args := []byte(strings.Join(hgcmd, string(0x0)))

	err = sendToHg(hgcl, command, args)
	if err != nil {
		fmt.Println(err)
		return nil, 0, err
	}

	var data []byte
	var buf bytes.Buffer
	var ret int32

CHANNEL_LOOP:
	for true {
		var ch string
		ch, data, err = readFromHg(hgcl)
		if err != nil || ch == "" {
			log.Fatal("runInHg(): readFromHg failed: " + string(err.Error()))
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
						log.Fatal("runInHg(): binary.read failed: " + string(err.Error()))
					}
				}
				break CHANNEL_LOOP
			}
		case "I":
		case "L":
		default:
			log.Fatal("runInHg(): unexpected channel '" + ch + "' detected")
		} // switch ch
	} // for true

	return []byte(buf.String()), ret, nil

} // runInHg()

// calcDataLength converts a 4-byte slice into an unsigned int
func calcDataLength(s []byte) (uint32, error) {
	var ln int32
	ln, err = calcIntFromBytes(s)
	return uint32(ln), err
}

// calcReturncode converts a 4-byte slice into a signed int
func calcReturncode(s []byte) (int32, error) {
	var rc int32
	rc, err = calcIntFromBytes(s)
	return rc, err
}

// calcIntFromBytes performs the real conversion
func calcIntFromBytes(s []byte) (int32, error) {
	var i int32
	buf := bytes.NewBuffer(s[0:4])
	err := binary.Read(buf, binary.BigEndian, &i)
	return i, err
}

// HgPath returns the path of the Mercurial executable used in the Hg CS.
func (hgcl *HgClient) HgPath() string {
	return hgcl.hgPath
}

// HgVersion returns the Mercurial version of the connected Hg CS.
func (hgcl *HgClient) HgVersion() string {
	return hgcl.hgVersion
}

// Repo returns the repo root of the connected Hg CS
func (hgcl *HgClient) Repo() string {
	return hgcl.repo
}

// Capabilities returns the capabilities of the connected Hg CS.
func (hgcl *HgClient) Capabilities() []string {
	return hgcl.capabilities
}

// Encoding returns the encoding for the connected Hg CS.
func (hgcl *HgClient) Encoding() string {
	return hgcl.encoding
}

// IsConnected tells if there is a connection to a Hg CS.
func (hgcl *HgClient) IsConnected() bool {
	return hgcl.hgserver != nil
}
