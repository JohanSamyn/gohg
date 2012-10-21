// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package gohg_lib

import (
	"errors"
	"strconv"
	"strings"
)

// Version provides the 'hg version' command.
func (hgcl *HgClient) Version() (string, error) {
	var data []byte
	var ret int32
	data, ret, err = hgcl.run([]string{"-q"})
	if err != nil {
		return "", err
	}
	if ret != 0 {
		return "", errors.New("run(\"version\") returned: " + strconv.Itoa(int(ret)))
	}
	ver := strings.Split(string(data), "\n")[0]
	ver = ver[strings.LastIndex(ver, " ")+1 : len(ver)-1]

	// This test first disturbed the call to getHgVersion() from Connect()
	// in gohg.go, because at that moment HgClient.hgVersion is not set yet
	// (in fact that call is exactly trying to do that).
	if hgcl.hgVersion != "" && ver != hgcl.hgVersion {
		return "", errors.New("run(\"version\"): expected '" + hgcl.hgVersion + "' and got '" + ver + "'")
	}

	return ver, nil
}
