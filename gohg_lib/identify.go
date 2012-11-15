// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg_lib

import (
	"fmt"
)

// Identify provides the 'hg identify' command.
func (hgcl *HgClient) Identify(opts []string) (string, error) {
	cmd := buildCmd([]string{"identify"}, opts)
	data, hgerr, ret, err := hgcl.run(cmd)
	if err != nil {
		return "", fmt.Errorf("from hgcl.run(): %s", err)
	}
	if ret != 0 || hgerr != nil {
		return "", fmt.Errorf("Identify(): returncode=%d\nhgerr:\n%s\n", ret, string(hgerr))
	}
	return string(data), nil
}
