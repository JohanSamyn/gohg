// Copyright (C) 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package gohg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"strings"
)

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
		case "e":
		case "d":
		case "I":
		case "L":
		default:
			log.Fatal("unexpected channel '" + ch + "' detected")
		} // switch ch
	} // for endOfRead == false
	fmt.Printf("command -> %s\nhgcmd -> %s\ndata ->\n%s\nreturncode -> %d\n",
		command, hgcmd, []byte(buf.String()), ret)
} // RunCommand()

// calcReturncode converts a 4-byte slice into a signed int
func calcReturncode(s []byte) (int32, error) {
	var ret int32
	buf := bytes.NewBuffer(s[0:4])
	err := binary.Read(buf, binary.BigEndian, &ret)
	return ret, err
}
