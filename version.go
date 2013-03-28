// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"strings"
)

// Version implements the 'hg version -q' command,
// and only returns the version number.
func (hgcl *HgClient) Version() (string, error) {
	var err error
	if hgcl.hgversion == "" {
		var data []byte
		data, err = hgcl.runcommand([]string{"version", "-q"})
		if err == nil {
			ver := strings.Split(string(data), "\n")[0]
			ver = ver[strings.LastIndex(ver, " ")+1 : len(ver)-1]
			hgcl.hgversion = ver
		}
	}
	return hgcl.hgversion, err
}
