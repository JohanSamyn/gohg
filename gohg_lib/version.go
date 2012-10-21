// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package gohg_lib

import (
	"errors"
	// "log"
	"strconv"
	"strings"
)

// Version provides the 'hg version' command.
func (hgcl *HgClient) Version(args []string) (ver string, err error) {
	var data []byte
	var ret int32
	cmd_args := []string{"version"}
	// if args != "" {
	// log.Println(len(args))
	if len(args) > 0 {
		for _, a := range args {
			if a != "" {
				cmd_args = append(cmd_args, a)
			}
		}
	}
	// log.Println(cmd_args)
	data, ret, err = hgcl.run(cmd_args)
	if err != nil {
		return "", err
	}
	if ret != 0 {
		return "", errors.New("run(\"version\") returned: " + strconv.Itoa(int(ret)))
	}
	ver = strings.Split(string(data), "\n")[0]
	if len(ver) > 0 {
		ver = strings.Split(ver, "\n")[0]
		ver = ver[strings.LastIndex(ver, " ")+1 : len(ver)-1]
	}

	// This test first disturbed the call to getHgVersion() from Connect()
	// in gohg.go, because at that moment HgClient.hgVersion is not set yet
	// (in fact that call is exactly trying to do that).
	if hgcl.hgVersion != "" && ver != hgcl.hgVersion {
		return "", errors.New("run(\"version\"): expected '" + hgcl.hgVersion + "' and got '" + ver + "'")
	}

	return ver, nil
}
