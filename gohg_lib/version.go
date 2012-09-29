// Copyright (C) 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package gohg_lib

import (
	"errors"
	"strconv"
	"strings"
)

// Version implements the 'hg version' command.
func (hgclient) Version() (ver string, fullver string, err error) {
	var data []byte
	var ret int32
	data, ret, err = RunCommand([]string{"version"})
	if err != nil {
		return "", "", err
	}
	if ret != 0 {
		return "", "", errors.New("RunCommand(\"version\") returned: " + strconv.Itoa(int(ret)))
	}
	fullver = string(data)
	if len(fullver) > 0 {
		ver = strings.Split(fullver, "\n")[0]
		ver = ver[strings.LastIndex(ver, " ")+1 : len(ver)-1]
	}
	return ver, fullver, nil
}
