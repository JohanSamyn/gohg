// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg_lib

func buildCmd(cmd []string, opts []string) []string {
	if len(opts) == 0 {
		return cmd
	}
	fullcmd := make([]string, len(cmd)+len(opts))
	copy(fullcmd, cmd)
	copy(fullcmd[len(cmd):], opts)
	return fullcmd
}
