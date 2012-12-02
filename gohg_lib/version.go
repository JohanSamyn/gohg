// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg_lib

// import (
// 	"fmt"
// 	"strings"
// )

// // Version provides the 'hg version' command.
// func (hgcl *HgClient) Version() (string, error) {
// 	if hgcl.hgVersion == "" {
// 		data, hgerr, ret, err := hgcl.run([]string{"version", "-q"})
// 		if err != nil {
// 			return "", fmt.Errorf("from hgcl.run(): %s", err)
// 		}
// 		if ret != 0 || hgerr != nil {
// 			return "", fmt.Errorf("Version(): returncode=%d\nhgerr: %s\n", ret, string(hgerr))
// 		}
// 		ver := strings.Split(string(data), "\n")[0]
// 		ver = ver[strings.LastIndex(ver, " ")+1 : len(ver)-1]

// 		hgcl.hgVersion = ver
// 	}
// 	return hgcl.hgVersion, nil
// }
