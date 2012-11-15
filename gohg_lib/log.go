// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg_lib

import (
	"fmt"
)

// Log provides the 'hg log' command.
func (hgcl *HgClient) Log(opts []string) ([]byte, error) {
	cmd := buildCmd([]string{"log"}, opts)
	data, hgerr, ret, err := hgcl.run(cmd)
	if err != nil {
		return nil, fmt.Errorf("from hgcl.run(): %s", err)
	}
	if ret != 0 || hgerr != nil {
		return nil, fmt.Errorf("Log(): returncode=%d\nhgerr:\n%s\n", ret, string(hgerr))
	}
	return data, nil
}
