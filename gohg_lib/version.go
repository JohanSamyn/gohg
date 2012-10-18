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
func (hgcl *HgClient) Version() (ver string, err error) {
	var data []byte
	var ret int32
	data, ret, err = hgcl.run([]string{"version", "-q"})
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
	return ver, nil
}
