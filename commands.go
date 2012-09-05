// Copyright (C) 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package gohg

import (
	"fmt"
)

func RunCommand() {
	err = sendToHg("runcommand", []byte("summary"))
	if err != nil {
		fmt.Println(err)
	}
	var channel string
	var alldata []byte
	var adata []byte
	var ch string
	var data []byte
	for {
		ch, data, err = readFromHg()
		if err != nil || ch == "" || ch == "r" {
			break
		}
		if ch == channel {
			var l1, l2 int
			l1 = len(alldata)
			l2 = len(data)
			adata = make([]byte, l1+l2)
			if l1 > 0 {
				copy(adata[0:l1], alldata)
			}
			copy(adata[l1:l1+l2], data)
			alldata = adata
		} else {
			alldata = data
		}
		if channel == "" {
			channel = ch
		}
	}
	fmt.Printf("channel -> %s\ndata ->\n%s\n", channel, alldata)
} // RunCommand()
