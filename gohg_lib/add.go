// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg_lib

// import (
// 	"fmt"
// )

// // Add provides the 'hg add' command.
// func (hgcl *HgClient) Add(opts []string) ([]byte, error) {
// 	cmd := buildCmd("add", opts)
// 	res, hgerr, ret, err := hgcl.run(cmd)
// 	if err != nil {
// 		return nil, fmt.Errorf("from hgcl.run(): %s", err)
// 	}
// 	if ret != 0 || hgerr != nil {
// 		return nil, fmt.Errorf("Add(): returncode=%d\nhgerr:\n%s\n", ret, string(hgerr))
// 	}
// 	return res, nil
// }
