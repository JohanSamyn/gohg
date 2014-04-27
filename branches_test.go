// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

import (
	// "fmt"
	"os"
	"os/exec"
	// "path/filepath"
	"strings"
	"testing"
)

func TestHgClient_Branches(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	// dropped the revision after the colon for more independent testing
	var expected string = "newbranch                      1:\n" +
		"default                        0:\n"

	f, err := os.Create(hct.RepoRoot() + "/a")
	_, _ = f.Write([]byte{'a', 'a', 'a'})
	f.Sync()
	f.Close()

	cmd := exec.Command(hct.HgExe(), "-R", hct.RepoRoot(), "ci", "-Am\"test\"")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}
	cmd = exec.Command(hct.HgExe(), "-R", hct.RepoRoot(), "branch", "newbranch")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}
	f, err = os.Create(hct.RepoRoot() + "/b")
	_, _ = f.Write([]byte{'b', 'b', 'b'})
	f.Sync()
	f.Close()
	cmd = exec.Command(hct.HgExe(), "-R", hct.RepoRoot(), "ci", "-Am\"test2\"")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	got1, err := hct.Branches(nil, nil)
	if err != nil {
		t.Error(err)
	}
	got := extractBranchInfo(got1)
	if string(got) != expected {
		t.Fatalf("Test Branches: expected:\n%s\n but got:\n%s\n", expected, got)
	}

	// test Active option

	expected = "newbranch                      1:\n"
	got1, err = hct.Branches([]HgOption{Active(true)}, nil)
	if err != nil {
		t.Error(err)
	}
	got = extractBranchInfo(got1)
	if string(got) != expected {
		t.Fatalf("Test Branches Active: expected:\n%s\n but got:\n%s\n", expected, got)
	}

	// This test was disabled because of a problem on drone.io.
	// dron.io uses Mercurial v2.0.2.
	// To be investigated.
	// test Closed option

	cmd = exec.Command(hct.HgExe(), "-R", hct.RepoRoot(), "update", "default")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}
	cmd = exec.Command(hct.HgExe(), "-R", hct.RepoRoot(), "ci", "--close-branch",
		"-m\"closed branch newbranch\"")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}
	expected = "newbranch                      1:\n" +
		"default                        2:\n"
	got1, err = hct.Branches([]HgOption{Closed(true)}, nil)
	if err != nil {
		t.Error(err)
	}
	got = extractBranchInfo(got1)
	if string(got) != expected {
		t.Fatalf("Test Branches Closed: expected:\n%s\n but got:\n%s\n", expected, got)
	}

	// // test Mq option

	// // for some reason this method produces returnvalue 255, at least in Linux
	// // fmt.Printf("reporoot: %s\n", hct.RepoRoot())
	// // // cmd = exec.Command(hct.HgExe(), "init --cwd "+hct.RepoRoot()+" --mq")
	// // cmd = exec.Command(hct.HgExe(), "init", "--mq")
	// // if err := cmd.Run(); err != nil {
	// // 	t.Fatal(err)
	// // }

	// // this method does not create files .hgignore and series however,
	// // at least on Linux
	// // // err = hct.Init(Mq(true), Cwd(hct.RepoRoot()))
	// // path, err := filepath.Abs(hct.RepoRoot() + "/.hg/patches")
	// // if err != nil {
	// // 	t.Error(err)
	// // }
	// // err = hct.Init(Destpath(path))
	// err = hct.Init(Mq(true))
	// if err != nil {
	// 	t.Error(err)
	// }
	// // return

	// // and this one then fails on Win 7
	// // cmd = exec.Command(hct.HgExe(), "-R", hct.RepoRoot(), "--mq", "branch", "newmqbranch")
	// cmd = exec.Command(hct.HgExe(), "branch", "newmqbranch", "--mq")
	// if err := cmd.Run(); err != nil {
	// 	t.Fatal(err)
	// }
	// // commit files .hgignore and series
	// cmd = exec.Command(hct.HgExe(), "-R", hct.RepoRoot(), "ci", "--mq", "-Am\"testmq\"")
	// if err := cmd.Run(); err != nil {
	// 	t.Fatal(err)
	// }
	// expected = "newmqbranch                    1:\n" +
	// 	"default                        0:\n"
	// got1, err = hct.Branches(Mq(true))
	// if err != nil {
	// 	t.Error(err)
	// }
	// got = extractBranchInfo(got1)
	// if string(got) != expected {
	// 	t.Fatalf("Test Branches Mq: expected:\n%s\n but got:\n%s\n", expected, got)
	// }
}

func extractBranchInfo(branches []byte) string {
	got := ""
	got2 := strings.Split(string(branches), "\n")
	for _, b := range got2 {
		if len(b) == 0 {
			continue
		}
		got = got + strings.SplitN(string(b), ":", 2)[0] + ":\n"
	}
	return got
}
