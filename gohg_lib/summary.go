// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg_lib

import (
	"fmt"
)

// TODO	Implement the --remote flag.
//		Beware of the error when no remote repo is known.
// TODO	Put the results in an appropriate go struct for possible further treatment.

// Summary provides the 'hg summary' command.
func (hgcl *HgClient) Summary() (string, error) {
	data, hgerr, ret, err := hgcl.run([]string{"summary"})
	if err != nil {
		return "", fmt.Errorf("from hgcl.run(): %s", err)
	}
	if ret != 0 || hgerr != nil {
		return "", fmt.Errorf("Summary(): returncode=%d\nhgerr:\n%s\n", ret, string(hgerr))
	}
	return string(data), nil
}
