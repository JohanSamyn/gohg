// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.txt file.

package gohg_lib

import (
	"errors"
	"strconv"
)

// TODO	Implement the --remote flag.
//		Beware of the error when no remote repo is known.

// Summary provides the 'hg summary' command.
func (hgcl *HgClient) Summary() (string, error) {
	var data []byte
	var ret int32
	data, ret, err = hgcl.run([]string{"summary"})
	if err != nil {
		return "", err
	}
	if ret != 0 {
		return "", errors.New("run(\"summary\") returned: " + strconv.Itoa(int(ret)))
	}
	return string(data), nil
}
