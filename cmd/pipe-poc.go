// This program is for exploring how to work with pipes in Go.
// As the cmomunication with the Mercurial Command Server works via pipes,
// this is a vital part of this project.

package main

import (
	"bytes"
	"fmt"
	"encoding/binary"
	"io"
	"log"
	"os/exec"
)

func main() {

	// ========== Connecting to the Hg CS ==========

	cmd := exec.Command("M:\\DEV\\hg-stable\\hg",
		"-R", "C:\\DEV\\go\\src\\golout\\",
		"--config", "ui.interactive=True", "serve", "--cmdserver", "pipe")

	pout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("[1] ", err)
	}
	// pin, err := cmd.StdinPipe()
	// if err != nil {
	// 	log.Fatal("[2] ", err)
	// }
	if err := cmd.Start(); err != nil {
		log.Fatal("[3] ", err)
	}

	// ========== Receiving a massage drom the Hg CS ==========

	// 4101:	1 byte for the channel
	//			4 bytes for the message length
	//			up to 4096 bytes of data
	s := make([]byte, 4101)

	i, err := pout.Read(s)
	if err != io.EOF && err != nil {
		log.Fatal("[5] ", i, err)
	}

	var ui uint32
	buf := bytes.NewBuffer(s[1:5])
	err = binary.Read(buf, binary.BigEndian, &ui)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	hgm := new(hgMsg)
	hgm.Ch = string(s[0])
	hgm.Ln = uint(ui)
	hgm.Data = string(s[5 : 5+hgm.Ln])
	fmt.Printf("hgMsg={\nChannel=%s\nLength=%v\nData=%s\n}\n",
		hgm.Ch, hgm.Ln, hgm.Data)

	// ========== Sending a message to the Hg CS ==========

	msg := new(hgMsg)
	msg.Ch = "runcommand\n"
	msg.Ln = 7
	msg.Data = "summary"

	l1 := len(msg.Ch)
	// l3 := len(msg.Data)

	sm := make([]byte, len(msg.Ch)+4+len(msg.Data))
	copy(sm[0:l1-1], []byte(msg.Ch))
	// fmt.Printf("l1=%d   l3=%d   ", l1, l3)
	// fmt.Printf("len(sm)=%d   len(msg.Data)=%d\n", len(sm), len([]byte(msg.Data)))
	copy(sm[l1+4:len(sm)], []byte(msg.Data))
	fmt.Printf("sm=%s\n", sm)

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
	bs := make([]byte, 4)
	wbuf := new(bytes.Buffer)
	err = binary.Write(wbuf, binary.BigEndian, uint32(msg.Ln))
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	// How to get the content of wbuf into bs or msg.Ln ?
	// bs = wbuf
	err = binary.Read(buf, binary.BigEndian, &bs)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	// func ReadUvarint(r io.ByteReader) (uint64, error)
	var uuii uint64
	uuii, err = binary.ReadUvarint(wbuf)
	if err != nil {
		fmt.Println("binary.ReadUvarint failed:", err)
	}

	copy(sm[l1-1:l1-1+4], bs)
	fmt.Printf("bs=%v   sm=[[%s]]   uuii=%v    wbuf=%v\n", bs, sm, uuii, *wbuf)

	// i, err = pin.Write(ms)
	// if i != len(sm) {
	// 	log.Fatal("[6] ", i, err)
	// }

	// ========== Receiving the result ==========

	// ========== Closing the connection ==========

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
