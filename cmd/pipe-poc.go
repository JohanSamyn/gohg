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
	"os/exec"
	"strings"
)

func main() {

	// ========== Connecting to the Hg CS ==========

	cmd := exec.Command("M:\\DEV\\hg-stable\\hg",
		"-R", "C:\\DEV\\go\\src\\golout\\",
		"--config", "ui.interactive=True",
		"--config", "color=False",
		"serve", "--cmdserver", "pipe")

	pout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("[1] ", err)
	}
	pin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal("[2] ", err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal("[3] ", err)
	}

	// ========== Receiving a massage drom the Hg CS ==========

	// 4101:	1 byte for the channel
	//			4 bytes for the message length
	//			up to 4096 bytes of data
	s := make([]byte, 4101)

	// change this, so we can read only 5 bytes first, and so determine the
	// channel and the length of the data. If channel = I or L, Hg is asking
	// for input, not sending us data. Otherwise we can go on reading 'length'
	// bytes. See the simple example on the Hg CS page.

	i, err := pout.Read(s)
	if err != io.EOF && err != nil {
		log.Fatal("[5] ", i, err)
	}
	fmt.Printf("s=%v\n", s[0:100])
	fmt.Printf("s=[[%s]]\n", s[0:100])

	var ui uint32
	buf := bytes.NewBuffer(s[1:5])
	err = binary.Read(buf, binary.BigEndian, &ui)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Printf("buf=%v\n", *buf)

	hgm := new(hgMsg)
	hgm.Ch = string(s[0])
	hgm.Ln = uint(ui)
	hgm.Data = string(s[5 : 5+hgm.Ln])
	fmt.Printf("hgMsg={\nChannel=%s\nLength=%v\nData=%s\n}\n",
		hgm.Ch, hgm.Ln, hgm.Data)

	// ========== Sending a message to the Hg CS ==========

	msg := new(hgMsg)
	msg.Ch = "runcommand\n"
	msg.Data = "summary"
	msg.Ln = uint(len(msg.Data))

	l1 := len(msg.Ch)
	l3 := len(msg.Data)

	sm := make([]byte, l1+4+l3)
	copy(sm[0:l1], msg.Ch)
	// fmt.Printf("l1=%d   l3=%d   ", l1, l3)
	// fmt.Printf("len(sm)=%d   len(msg.Data)=%d\n", len(sm), len([]byte(msg.Data)))
	copy(sm[l1+4:len(sm)], msg.Data)
	fmt.Printf("sm=%v    len(sm)=%d\n", sm, len(sm))
	fmt.Printf("sm=%s    len(sm)=%d\n", sm, len(sm))

	/*
		func Write(w io.Writer, order ByteOrder, data interface{}) error

		Write writes the binary representation of data into w.
		Data must be a fixed-size value or a slice of fixed-size values,
		or a pointer to such data. Bytes written to w are encoded using
		the specified byte order and read from successive fields of the data.

		buf := new(bytes.Buffer)
		var pi float64 = math.Pi
		err := binary.Write(buf, binary.LittleEndian, pi)
		if err != nil {
		    fmt.Println("binary.Write failed:", err)
		}
		fmt.Printf("% x", buf.Bytes())
	*/
	// bs := make([]byte, 4)
	wbuf := new(bytes.Buffer)
	err = binary.Write(wbuf, binary.BigEndian, uint32(msg.Ln))
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	// How to get the content of wbuf into bs or msg.Ln ?
	// bs = wbuf
	// err = binary.Read(buf, binary.BigEndian, &bs)
	// if err != nil {
	// 	fmt.Println("binary.Read failed:", err)
	// }

	bf := make([]byte, 4)
	// fmt.Printf("len(bf)=%d\n", len(bf))
	var n int
	n, err = io.ReadFull(wbuf, bf)
	if err != nil {
		fmt.Println("io.ReadFull failed:", err)
	}
	fmt.Printf("n=%d    bf=%v\n", n, bf)

	// // func ReadUvarint(r io.ByteReader) (uint64, error)
	// var uuii uint64
	// uuii, err = binary.ReadUvarint(wbuf)
	// if err != nil {
	// 	fmt.Println("binary.ReadUvarint failed:", err)
	// }

	copy(sm[l1:l1+4], bf)
	fmt.Printf("sm=%v\n", sm)
	fmt.Printf("bf=%v   sm=[[%s]]    len(sm)=%d\nwbuf=%v\n", bf, sm, len(sm), *wbuf)
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

	// var br []byte
	var hgm2 hgMsg
	for {
		hgm2 = readFromHg(pout)
		hgm2.Data = strings.Replace(hgm2.Data, "\n", "\\n", -1)
		fmt.Printf("hgm2=%s\n", hgm2)
		// fmt.Printf("hgm2=%v\nhgm2=%s", hgm2, hgm2)
		// switch {
		// 	case
		// }
		if hgm2.Ch == "r" {
			break
		}
	}

	// ========== Closing the connection ==========

	pout.Close()
	pin.Close()
	if err := cmd.Wait(); err != nil {
		log.Fatal("[4] ", err)
	}

} // main()

type hgMsg struct {
	Ch   string
	Ln   uint
	Data string
}

// func decodeHgMsg(s string) (hgMsg, error) {
// 	return "", nil
// }

func readFromHg(pout io.ReadCloser) hgMsg {
	s := make([]byte, 4101)

	// change this, so we can read only 5 bytes first, and so determine the
	// channel and the length of the data. If channel = I or L, Hg is asking
	// for input, not sending us data. Otherwise we can go on reading 'length'
	// bytes. See the simple example on the Hg CS page.

	i, err := pout.Read(s)
	if err != io.EOF && err != nil {
		log.Fatal("[5] ", i, err)
	}
	// fmt.Printf("s=%v\n", s[0:100])
	// fmt.Printf("s=[[%s]]\n", s[0:100])

	var ui uint32
	buf := bytes.NewBuffer(s[1:5])
	err = binary.Read(buf, binary.BigEndian, &ui)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	// fmt.Printf("buf=%v\n", *buf)

	hgm := new(hgMsg)
	hgm.Ch = string(s[0])
	hgm.Ln = uint(ui)
	hgm.Data = string(s[5 : 5+hgm.Ln])
	// fmt.Printf("hgMsg={\nChannel=%s\nLength=%v\nData=%s\n}\n",
	// 	hgm.Ch, hgm.Ln, hgm.Data)

	return *hgm
}
