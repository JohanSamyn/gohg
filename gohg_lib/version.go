// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg_lib

import (
	"errors"
	"strconv"
	"strings"
)

// Version provides the 'hg version' command.
func (hgcl *HgClient) Version() (string, error) {
	if hgcl.hgVersion == "" {
		var data []byte
		var ret int32
		data, ret, err = hgcl.run([]string{"version", "-q"})
		if err != nil {
			return "", err
		}
		if ret != 0 {
			return "", errors.New("Mercurial returned: " + strconv.Itoa(int(ret)))
		}
		ver := strings.Split(string(data), "\n")[0]
		ver = ver[strings.LastIndex(ver, " ")+1 : len(ver)-1]

		hgcl.hgVersion = ver
	}

	return hgcl.hgVersion, nil
}
