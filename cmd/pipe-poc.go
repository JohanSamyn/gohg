// This program is for exploring how to work with pipes in Go.
// As the communication with the Mercurial Command Server works via pipes,
// this is a vital part of this project.

package main

import (
	"bytes"
	"fmt"
	"encoding/binary"
	"io"
	"log"
	"os"
	"os/exec"
	// "strconv"
	"strings"
)

func main() {

	// ========== Connecting to the Hg CS ==========
	cmd := "M:\\DEV\\hg-stable\\hg" // hg
	// cmd := "hg"
	server := exec.Command(cmd)
	server.Args = append(server.Args, "-R", "C:\\DEV\\go\\src\\golout\\")
	server.Args = append(server.Args,
		"--config", "ui.interactive=True", "--config",
		"extensions.color=False", "serve", "--cmdserver", "pipe")
	fmt.Printf("server=[[%v]]\n", server)

	// cmd := exec.Command("M:\\DEV\\hg-stable\\hg",
	// 	"-R", "C:\\DEV\\go\\src\\golout\\",
	// 	"--config", "ui.interactive=True",
	// 	"--config", "extensions.color=False",
	// 	"serve", "--cmdserver", "pipe")

	pout, err := server.StdoutPipe()
	if err != nil {
		log.Fatal("[1] ", err)
	}
	pin, err := server.StdinPipe()
	if err != nil {
		log.Fatal("[2] ", err)
	}
	if err := server.Start(); err != nil {
		log.Fatal("[3] ", err)
	}

	// ========== Receiving a message from the Hg CS ==========

	// 1029:	1 byte for the channel
	//			4 bytes for the message length
	//			up to 1024 bytes of data
	s := make([]byte, 1+4+1024)

	// change this, so we can read only 5 bytes first, and so determine the
	// channel and the length of the data. If channel = I or L, Hg is asking
	// for input, not sending us data. Otherwise we can go on reading 'length'
	// bytes. See the simple example on the Hg CS page.

	i, err := pout.Read(s)
	if err != io.EOF && err != nil {
		log.Fatal("[5] ", i, err)
	}
	// Als s begint met "hg server [OPTIONS]" dan is er iets mis.
	// Hoogstwaarschijnlijk heeft de Hg versie geen Command Server capability.
	var t string
	t = "hg serve [OPTION]"
	if string(s[0:len(t)]) == t {
		log.Fatal("This version of Mercurial does not have the Command Server capability.")
	}
	fmt.Printf("len(s)=%d    s=%s\n", len(s), s)
	if len(s) == 0 {
		log.Fatal("no data received")
	}
	// fmt.Printf("s=%v\n", s[0:100])
	// fmt.Printf("s=[[%s]]\n", s[0:100])

	var l uint32
	// var l64 uint64
	// l64, err = strconv.ParseUint(string(s[1:5]), 16, 32)
	buf := bytes.NewBuffer(s[1:5])
	err = binary.Read(buf, binary.BigEndian, &l)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	// l = uint32(l64)
	// fmt.Printf("buf=%v\n", *buf)

	hgm0 := new(hgMsg)
	hgc0 := new(hgCmd)

	hgm := new(hgMsg)
	hgm.Ch = string(s[0])
	hgm.Ln = uint(l)
	fmt.Printf("hgm.Ln=%d\n", hgm.Ln)
	hgm.Data = string(s[5 : 5+hgm.Ln])

	hgm0.Data = strings.Replace(hgm.Data, "\n", "\\n", -1)
	fmt.Printf("hgMsg={Channel: \"%s\"    Length: %v\n       Data: \"%s\"}\n",
		hgm.Ch, hgm.Ln, hgm0.Data)

	// ========== Sending a message to the Hg CS ==========

	msg := new(hgCmd)
	msg.Cmd = "runcommand\n"
	msg.Args = "summary"
	msg.Ln = uint(len(msg.Args))
	hgc0.Cmd = strings.Replace(msg.Cmd, "\n", "\\n", -1)
	hgc0.Args = strings.Replace(msg.Args, "\n", "\\n", -1)
	fmt.Printf("hgCmd={Command: \"%s\"    Length: %v\n       Args: \"%s\"}\n",
		hgc0.Cmd, msg.Ln, hgc0.Args)

	l1 := len(msg.Cmd)
	l3 := len(msg.Args)

	sm := make([]byte, l1+4+l3)
	copy(sm[0:l1], msg.Cmd)
	// fmt.Printf("l1=%d   l3=%d   ", l1, l3)
	// fmt.Printf("len(sm)=%d   len(msg.Data)=%d\n", len(sm), len([]byte(msg.Data)))
	copy(sm[l1+4:len(sm)], msg.Args)
	// fmt.Printf("sm=%v    len(sm)=%d\n", sm, len(sm))
	// fmt.Printf("sm=%s    len(sm)=%d\n", sm, len(sm))

	wbuf := new(bytes.Buffer)
	err = binary.Write(wbuf, binary.BigEndian, uint32(msg.Ln))
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	bf := make([]byte, 4)
	// fmt.Printf("len(bf)=%d\n", len(bf))
	// var n int
	_, err = io.ReadFull(wbuf, bf)
	if err != nil {
		fmt.Println("io.ReadFull failed:", err)
	}
	// fmt.Printf("n=%d    bf=%v\n", n, bf)

	copy(sm[l1:l1+4], bf)
	// fmt.Printf("sm=%v\n", sm)
	// fmt.Printf("bf=%v   sm=[[%s]]    len(sm)=%d\nwbuf=%v\n", bf, sm, len(sm), *wbuf)
	// copy(sm[l1-1:l1-1+4], bs)
	// fmt.Printf("bs=%v   sm=[[%s]]   uuii=%v\nwbuf=%v\n", bs, sm, uuii, *wbuf)

	i, err = pin.Write(sm)
	if i != len(sm) {
		log.Fatal("[6] ", i, err)
	}

	// ========== Receiving the result ==========

	/*
		while True:
		  28     channel, val = readchannel(server)
		  29     if channel == 'o':
		  30         print "output:", repr(val)
		  31     elif channel == 'e':
		  32         print "error:", repr(val)
		  33     elif channel == 'r':
		  34         print "exit code:", struct.unpack(">l", val)[0]
		  35         break
		  36     elif channel == 'L':
		  37         print "(line read request)"
		  38         writeblock(sys.stdin.readline(val))
		  39     elif channel == 'I':
		  40         print "(block read request)"
		  41         writeblock(sys.stdin.read(val))
		  42     else:
		  43         print "unexpected channel:", channel, val
		  44         if channel.isupper(): # required?
		  45             break
	*/

	var hgm2 hgMsg
	var show bool
	for {
		hgm2 = receiveFromHg(pout)
		show = true
		switch {
		case hgm2.Ch == "o":
			{
			}
		case hgm2.Ch == "e":
			{
			}
		case hgm2.Ch == "r":
			{
			}
		case hgm2.Ch == "I":
			{
			}
		case hgm2.Ch == "L":
			{
			}
		default:
			{
				show = false
				// fmt.Printf("unexpected channel: %s %d\n", hgm2.Ch, hgm2.Ln)
				if hgm2.Ch == strings.ToUpper(hgm2.Ch) {
					os.Exit(1)
				}
			}
		} // switch
		if show == true {
			// hgm0.Data = strings.Replace(hgm2.Data, "\n", "\\n", -1)
			hgm0.Data = hgm2.Data
			fmt.Printf("hgMsg={Channel: \"%s\"    Length: %v\n       Data: \"%s\"}\n",
				hgm2.Ch, hgm2.Ln, hgm0.Data)
		}
	} // for

	// ========== Closing the connection ==========

	pout.Close()
	pin.Close()
	if err := server.Wait(); err != nil {
		log.Fatal("[4] ", err)
	}

} // main()

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

func receiveFromHg(pout io.ReadCloser) hgMsg {

	// change this, so we can read only 5 bytes first, and so determine the
	// channel and the length of the data. If channel = I or L, Hg is asking
	// for input, not sending us data. Otherwise we can go on reading 'length'
	// bytes. See the simple example on the Hg CS page.

	s1 := make([]byte, 1+4)
	i, err := pout.Read(s1)
	if err != io.EOF && err != nil {
		log.Fatal("[5] ", i, err)
	}
	// fmt.Printf("s=%v\n", s[0:100])
	// fmt.Printf("s=[[%s]]\n", s[0:100])

	hgm := new(hgMsg)

	if s1[0] == byte("r"[0]) {
		return *hgm
	}
	var l uint32
	buf := bytes.NewBuffer(s1[1:5])
	err = binary.Read(buf, binary.BigEndian, &l)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	// fmt.Printf("buf=%v\n", *buf)

	hgm.Ch = string(s1[0])
	hgm.Ln = uint(l)

	if hgm.Ln == 0 {
		return *hgm
	}
	s2 := make([]byte, hgm.Ln)
	i, err = pout.Read(s2)
	if err != io.EOF && err != nil {
		log.Fatal("[5] ", i, err)
	}

	// hgm.Data = string(s2[5 : 5+hgm.Ln])
	hgm.Data = string(s2[:])
	// fmt.Printf("hgMsg={\nChannel=%s\nLength=%v\nData=%s\n}\n",
	// 	hgm.Ch, hgm.Ln, hgm.Data)

	return *hgm
}
