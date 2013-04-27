// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

// Package gohg is a Go client library for using the Mercurial dvcs
// via it's Command Server.
//
// For Mercurial see: http://mercurial.selenic.com.
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
	"path/filepath"
	"reflect"
	"strings"
)

// Type HgClient will act as the entrypoint through which all interaction
// with the Mercurial Command Server will take place.
type HgClient struct {
	hgServer     *exec.Cmd
	pipeIn       io.WriteCloser // the pipe that gets commands into the Hg CS
	pipeOut      io.ReadCloser  // the pipe that brings data out of the Hg CS
	hgExe        string         // which hg is used ?
	capabilities []string       // as per the hello message
	encoding     string         // as per the hello message
	repoRoot     string         // the full path to the Hg repo
	hgversion    string         // the version number only
	// config       []string
}

// NewHgClient creates a new instance of the client object for working with the
// Hg Command Server.
func NewHgClient() *HgClient {
	return new(HgClient)
}

// Connect establishes the connection with the Mercurial Command Server.
//
// Arguments:
//	hgexe
//		The command to run Mercurial. Optional.
//		The 'hg' command will be used when not provided.
//		This allows to run a specific version of Mercurial.
//	reponame
//		The folder of the Hg repository to work on. Optional.
//		When blanc the folder where the program is run is used
//		(see function locateRepository()).
//	config
//		Configuration settings that will be added to the necessary
//		fixed settings (see composeStartupConfig() for more). Optional.
//
// Returns an error if the connection could not be established properly.
func (hgcl *HgClient) Connect(hgexe string, reponame string, config []string) error {

	// for example:
	// hgcl.hgServer =
	//		exec.Command("M:/DEV/hg-stable/hg",	// the Hg command
	// 		"--cwd", "C:/DEV/go/src/golout/",	// the repo
	// 		"-R", "C:/DEV/go/src/golout/",		// the repo
	// 		"--config", "ui.interactive=False",	// mandatory settings
	// 		"--config", "extensions.color=!",	// more settings (for Windows)
	// 		"serve", "--cmdserver", "pipe")		// start the Command Server

	// Maybe accept a channel as an extra argument for sending the logging to ?
	// And if it's nil, log into a textfile in the folder of this lib.
	// Also do not override that logfile every launch.
	// Maybe even do this in the init() function ?

	if hgcl.hgServer != nil {
		return fmt.Errorf("Connect(): already running a Hg Command Server for %s", hgcl.repoRoot)
	}

	// Set some environment variables, so we can depend on a known situation.
	// HGPLAIN: Enabling this also assures Hg itself works in english,
	// so we can depend on some strings.
	os.Setenv("HGPLAIN", "True")

	// HGRCPATH: Use only the .hg/hgrc from the repo itself.
	// This one should perhaps be guarded with a passed-in option.

	// disabled (temporarily?), because otherwise a bunch of test fail
	// os.Setenv("HGRCPATH", "''")

	// This is only a 'non set' HGRCPATH, as if you do: HGRCPATH=
	// Result is as if you dd not set it.
	os.Setenv("HGRCPATH", "")

	os.Setenv("HGENCODING", "UTF-8")

	hgcl.hgExe = hgexe
	if hgcl.hgExe == "" {
		// Let the OS determine what Mercurial to run.
		hgcl.hgExe = "hg"
	}

	// The Hg Command Server needs a repository.
	var err error
	hgcl.repoRoot, err = locateRepository(reponame)
	if err != nil {
		return err
	}
	if hgcl.repoRoot == "" {
		dir := reponame
		if dir == "" {
			dir = "."
		}
		return fmt.Errorf("Connect(): could not find a Hg repository at: %s", dir)
	}

	var hgServerArgs []string
	hgServerArgs = composeStartupConfig(hgcl.hgExe, hgcl.repoRoot, config)

	hgcl.hgServer = exec.Command(hgcl.hgExe)
	hgcl.hgServer.Args = hgServerArgs
	hgcl.hgServer.Dir = hgcl.repoRoot

	hgcl.pipeOut, err = hgcl.hgServer.StdoutPipe()
	if err != nil {
		return fmt.Errorf("Connect(): could not connect StdoutPipe: %s", err)
	}
	hgcl.pipeIn, err = hgcl.hgServer.StdinPipe()
	if err != nil {
		return fmt.Errorf("Connect(): could not connect StdinPipe: %s", err)
	}

	// var p, t string
	// p = os.Getenv("path")
	// t = os.Getenv("temp")
	// fmt.Println(p, t)

	if err := hgcl.hgServer.Start(); err != nil {
		return fmt.Errorf("Connect(): could not start the Hg Command Server: %s", err)
	}

	err = hgcl.readHelloMessage()
	if err != nil {
		return err
	}

	err = hgcl.validateCapabilities()
	if err != nil {
		return err
	}

	hgcl.hgversion, err = hgcl.Version()
	if err != nil {
		return fmt.Errorf("Version(): %s", err)
	}

	return nil

} // Connect()

// Disconnect ends the connection with the Mercurial Command Server.
//
// In fact it's closing the stdin of the Hg CS that disconnects,
// as per the Hg CS documentation.
func (hgcl *HgClient) Disconnect() error {
	if hgcl.hgServer == nil {
		log.Println("Disconnect(): no connected hgServer.")
		return nil
	}

	hgcl.pipeIn.Close()
	hgcl.pipeOut.Close()

	defer func() { hgcl.hgServer = nil }()

	err := hgcl.hgServer.Wait()
	if err != nil {
		return err
	}
	return nil
} // Disconnect()

// locateRepository assures we have a Mercurial repository available,
// which is required for working with the Hg Command Server.
func locateRepository(reponame string) (string, error) {
	repo := reponame
	if repo == "" {
		repo = "."
	}
	sep := string(os.PathSeparator)

	// first make a correct path from repo
	var err error
	repo, err = filepath.Abs(repo)
	if err != nil {
		return "", fmt.Errorf("%s\ncould not determine absolute path for: %s",
			err.Error(), repo)
	}
	repo = filepath.Clean(repo)

	// If we do not find a Hg repo in this dir, we search for one
	// up the path, in case we're deeper in it's working copy.
	for {
		_, err = os.Stat(repo + sep + ".hg")
		if err == nil && !os.IsExist(err) {
			break
		}
		var file string
		repo, file = filepath.Split(repo)
		if repo == "" || file == "" {
			break
		}
	}
	if err != nil {
		return "", err
	}

	return repo, nil

} // locateRepository()

// composeStartupConfig handles the different config settings that will be used
// to make the connection with the Hg CS. It concerns specific Hg settings.
func composeStartupConfig(hgcmd string, repo string, config []string) []string {
	var cmdline []string

	// Zoek uit hoe de inhoud van parameter config kan toegevoegd worden zonder in conflict
	// te komen met de vaste config elementen.

	// if len(config) > 0 {
	// 	var cfg string
	// 	for i := 0; i < range(config) {
	// 		cfg = cfg + "," + config[i]
	// 	}
	// 	cmd = cmd + "," + cfg
	// }

	// if len(config) > 0 {
	// 	cmdline = append(cmdline, config...)
	// }

	cmdline = append(cmdline, hgcmd,
		"--cwd", repo,
		"-R", repo,
		// These arguments are fixed.
		"--config", "ui.interactive=False",
		"--config", "extensions.color=!",
		"serve", "--cmdserver", "pipe")

	return cmdline
} // composeStartupConfig()

// readHelloMessage reads the special hello message send by the Hg CS.
//
// It has a fixed format, and contains info about the possibilities
// of the Hg CS at hand. It's also a first proof of a working connection.
func (hgcl *HgClient) readHelloMessage() error {
	var err error
	s := make([]byte, 5)
	_, err = hgcl.pipeOut.Read(s)
	if err != io.EOF && err != nil {
		return err
	}
	if len(s) == 0 {
		return errors.New("readHelloMessage(): no hello message data received from Hg Command Server")
	}
	const t1 = "hg se" // hg returned: "hg serve [OPTION]"
	if string(s[0:len(t1)]) == t1 {
		return fmt.Errorf("Fatal error: Need at least version 1.9 of Mercurial to use the Command Server.\n"+
			"Used hgexe: '%s'\n", hgcl.HgExe())
	}
	ch := string(s[0])
	if ch != "o" {
		return fmt.Errorf("readHelloMessage(): received unexpected channel '%s' for hello message from Hg Command Server",
			ch)
	}
	var ln uint32
	ln, err = calcDataLength(s[1:5])
	if err != nil {
		return fmt.Errorf("readHelloMessage(): binary.Read failed: %s", err)
	}
	if ln <= 0 {
		return fmt.Errorf("readHelloMessage(): received invalid length '%s' for hello message from Hg Command Server",
			string(ln))
	}
	hello := make([]byte, ln)
	_, err = hgcl.pipeOut.Read(hello)
	if err != io.EOF && err != nil {
		return err
	}
	const t2 = "capabilities:"
	if string(hello[0:len(t2)]) != t2 {
		return errors.New("readHelloMessage(): could not determine the capabilities of the Hg Command Server")
	}
	attr := strings.Split(string(hello), "\n")
	hgcl.capabilities = strings.Fields(attr[0])[1:]
	hgcl.encoding = strings.Split(attr[1], ": ")[1]
	return nil
} // readHelloMessage()

func (hgcl *HgClient) validateCapabilities() error {
	var ok bool
	for _, c := range hgcl.capabilities {
		if c == "runcommand" {
			ok = true
			break
		}
	}
	if !ok {
		return errors.New("validateCapabilities(): Fatal error: could not detect the 'runcommand' capability")
	}
	return nil
}

// receiveFromHg returns the channel and all the data read from it.
// Eventually it returns no (or empty) data but an error.
func (hgcl *HgClient) receiveFromHg() (string, []byte, error) {
	var err error
	var ch string

	// get channel and length
	data := make([]byte, 5)
	_, err = hgcl.pipeOut.Read(data)
	if err != io.EOF && err != nil {
		return "", data, err
	}
	if data == nil {
		return "", nil, errors.New("receiveFromHg(): no data read")
	}
	ch = string(data[0])
	if ch == "" {
		return "", data, errors.New("receiveFromHg(): no channel read")
	}

	// get the uint that the Hg CS sent us as the length value
	var ln uint32
	ln, err = calcDataLength(data[1:5])
	if err != nil {
		return ch, data, fmt.Errorf("receiveFromHg(): calcDataLength(): binary.Read failed: %s", err)
	}

	// now get ln bytes of data
	data = make([]byte, ln)
	_, err = hgcl.pipeOut.Read(data)
	if err != io.EOF && err != nil {
		return ch, data, err
	}

	return ch, data, nil
} // receiveFromHg()

// sendToHg writes data to the Hg CS,
// returning an error if something went wrong.
func (hgcl *HgClient) sendToHg(cmd string, args []byte) error {
	var err error

	// cmd: can only be 'runcommand' or 'getencoding' for now
	cmd = strings.TrimRight(cmd, "\n") + "\n"
	lc := len(cmd)
	la := len(args)
	l := lc
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
			return fmt.Errorf("sendToHg(): converting data length to binary failed: %s", err)
		}
		b := make([]byte, 4)
		_, err = io.ReadFull(wbuf, b)
		if err != nil {
			return fmt.Errorf("sendToHg(): writing the data length to buffer failed: %s", err)
		}
		copy(data[lc:lc+4], b)

		// send the command arguments
		copy(data[lc+4:], args)
	}

	// perform the actual send to the Hg CS
	var i int
	i, err = hgcl.pipeIn.Write(data)
	if i != len(data) {
		return fmt.Errorf("sendToHg(): writing data to Hg CS failed: %s", err)
	}

	return nil
} // sendToHg()

func sprintfOpts(opts interface{}) string {
	s := fmt.Sprintf("%T", opts) + " = {"
	t := reflect.ValueOf(opts)
	typeOfT := t.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if i > 0 {
			s = s + ", "
		}
		s = s + fmt.Sprintf("%s=%v", typeOfT.Field(i).Name, f.Interface())
	}
	s = s + "}\n"
	return s
}

// Method buildCommand builds the command string to pass to the Hg CS.
//	cmdName:	The name of the command ("log", "status", etc.).
//	cmdOpts:	A type based on struct with the valid options for the command
//				at hand. Also contains the (gohg) default value for each option
//				for that particular command.
//				Used to filter only the options supported by the command.
//	opts:		The options passed by the caller.
//	params:		Any filenames, paths or other that the command needs.
//				These are to be added as the last things of the command.
func (hgcl *HgClient) buildCommand(cmdName string, cmdOpts interface{}, opts []optionAdder, params []string) (hgcmd []string, err error) {
	hgcmd = []string{cmdName}
	for _, o := range opts {
		err = o.addOption(cmdOpts, &hgcmd)
		// Silently skip invalid options for now.
		// if err != nil {
		// 	fmt.Logf("err = %s", err)
		// 	// Or work out some logging system for gohg, and write the error message inthere.
		// 	log.Printf("err = %s", err)
		// 	return nil, err
		// }
	}
	for _, p := range params {
		if p != "" {
			hgcmd = append(hgcmd, p)
		}
	}
	return hgcmd, nil
}

// runcommand allows to run a Mercurial command in the Hg Command Server.
// You can only run 'hg' commands that are available in this library.
func (hgcl *HgClient) runcommand(cmd []string) (data []byte, err error) {

	if !hgcl.IsConnected() {
		return nil, fmt.Errorf("%s", "runcommand: Cannot execute "+
			strings.Title(cmd[0])+": no Hg CS connected!")
	}

	// boilerplate code for all commands

	// fmt.Printf("cmd = %s\nopts = %v\n", (*cmd)[0], (*cmd)[1:])

	data, hgerr, ret, err := hgcl.runInHg("runcommand", cmd)
	if err != nil {
		return nil, fmt.Errorf("runcommand: %s", err)
	}
	// Maybe make this 2 checks, to differentiate between ret and hgerr?
	if ret != 0 || hgerr != nil {
		return nil, fmt.Errorf("runcommand: %s(): returncode=%d\nhgerr:\n%s\n",
			strings.Title(cmd[0]), ret, string(hgerr))
	}
	return data, nil
} // runcommand()

// runInHg sends a command to the Hg CS (using sendToHg),
// and fetches the result (using receiveFromHg).
func (hgcl *HgClient) runInHg(command string, hgcmd []string) ([]byte, []byte, int32, error) {

	if command == "" || hgcmd == nil {
		return nil, nil, 0, fmt.Errorf("runInHg(): Received invalid empty or blank params.")
	}

	args := []byte(strings.Join(hgcmd, string(0x0)))
	// fmt.Printf("args: %s\n", strings.Replace(string(args), string(0x0), " ", -1))

	var err error
	err = hgcl.sendToHg(command, args)
	if err != nil {
		fmt.Println(err)
		return nil, nil, 0, err
	}

	var buf bytes.Buffer
	var errbuf bytes.Buffer
	var ret int32

CHANNEL_LOOP:
	for true {
		var ch string
		var data []byte
		ch, data, err = hgcl.receiveFromHg()
		if err != nil || ch == "" {
			return nil, nil, 0, fmt.Errorf("runInHg(): receiveFromHg() failed: %s", err)
		}
		switch ch {
		case "d":
		case "e":
			errbuf.Write(data)
		case "o":
			buf.Write(data)
		case "r":
			{
				if command == "getencoding" {
					buf.Write(data)
				} else {
					ret, err = calcReturncode(data[0:4])
					if err != nil {
						return nil, nil, 0, fmt.Errorf("runInHg(): calcReturncode() failed: %s", err)
					}
				}
				break CHANNEL_LOOP
			}
		case "I":
		case "L":
		default:
			return nil, nil, 0, fmt.Errorf("runInHg(): unexpected channel '%s' detected", ch)
		} // switch ch
	} // for true

	return buf.Bytes(), errbuf.Bytes(), ret, nil

} // runInHg()

// calcDataLength converts a 4-byte slice into an unsigned int
func calcDataLength(s []byte) (uint32, error) {
	i, err := calcIntFromBytes(s)
	if err != nil {
		return 0, err
	}
	return uint32(i), nil
}

// calcReturncode converts a 4-byte slice into a signed int
func calcReturncode(s []byte) (int32, error) {
	return calcIntFromBytes(s)
}

// calcIntFromBytes performs the real conversion of a 4-byte-slice into an int
func calcIntFromBytes(s []byte) (i int32, err error) {
	err = binary.Read(bytes.NewBuffer(s[0:4]), binary.BigEndian, &i)
	return
}

// HgExe returns the path of the Mercurial executable used in the Hg CS.
func (hgcl *HgClient) HgExe() string {
	return hgcl.hgExe
}

// hgVersion returns the Mercurial version of the connected Hg CS.
func (hgcl *HgClient) hgVersion() string {
	return hgcl.hgversion
}

// RepoRoot returns the root of the repository the connected Hg CS is working on.
func (hgcl *HgClient) RepoRoot() string {
	return hgcl.repoRoot
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
	return hgcl.hgServer != nil
}
