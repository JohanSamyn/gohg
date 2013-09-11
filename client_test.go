// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

import (
	"fmt"
	"log"
	"os/exec"
	"testing"
)

func TestClient_runcommand_IsConnected(t *testing.T) {
	hct := NewHgClient()
	// we deliberately do not call Connect()
	var err error
	_, err = hct.Identify(nil, nil)
	if err == nil {
		fmt.Printf("error: %s\n", err)
		log.Fatal(fmt.Errorf("%s", "Did not detect disconnect!?"))
	}
}

func TestClient_Exec(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	var expected string = "000000000000 tip\n"

	got, err := hct.ExecCmd([]string{"identify", "-v"})
	if err != nil {
		t.Error(err)
	}

	if string(got) != expected {
		t.Fatalf("Test Exec: expected:\n%s\n but got:\n%s\n", expected, got)
	}
}

// func TestConnect(*testing.T) {
// 	fmt.Printf("Hgclient.Connected = %v\n", Hgclient.Connected)
// 	var err error
// 	var repo = "."

// 	if Hgclient.Connected {
// 		Disconnect()
// 	}

// 	cfg := make([]string, 0)
// 	err = Hgclient.Connect("M:/DEV/hg-stable/hg" /*"hg"*/, repo, cfg)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func TestRun(t *testing.T) {
// 	hct := setup(t)
// 	defer teardown(t, hct)

// 	var err error
// 	// var data []byte
// 	// var ret int32
// 	// hgcmd := []string{"log", "-l", "2"}
// 	hgcmd := []string{"summary"}
// 	// data, ret, err = run(hgcmd)
// 	_, _, err = hct.run(hgcmd)
// 	if err != nil {
// 		t.Errorf("from run : %s", string(err.Error()))
// 	}
// }

// func TestDisconnect(t *testing.T) {
// 	err := Hgclient.Disconnect()
// 	if err != nil {
// 		t.Errorf("from Disconnect(): %s", string(err.Error()))
// 	}
// }

func TestShouldFailWhenNoRepoFoundAndInitrepoFalse(t *testing.T) {
	repo, err := createTempdir(t)
	if err != nil {
		t.Fatal(err)
	}
	defer destroyTempdir(repo)

	hct := NewHgClient()
	cfg := make([]string, 0)
	err = hct.Connect("hg", repo, cfg, false)
	// Maybe I should turn this into its own error type RepoNotFound,
	// so I don't have to repeat the string from the Connect() method?
	if err.Error() != "Connect(): could not find a Hg repository at: "+repo {
		t.Fatal(err)
	}
}

func TestShouldSucceedWhenRepoFoundAndInitrepoFalse(t *testing.T) {
	repo, err := createTempdir(t)
	if err != nil {
		t.Fatal(err)
	}
	defer destroyTempdir(repo)
	hct := NewHgClient()
	defer hct.Disconnect()

	hgexe := "hg"

	var cmd *exec.Cmd
	cmd = exec.Command(hgexe, "--cwd", repo, "init")
	if err = cmd.Run(); err != nil {
		t.Fatal(err)
	}

	cfg := make([]string, 0)
	err = hct.Connect(hgexe, repo, cfg, false)
	if err != nil {
		t.Fatal(err)
	}
}

func TestShouldSucceedWhenNoRepoFoundAndInitrepoTrue(t *testing.T) {
	repo, err := createTempdir(t)
	if err != nil {
		t.Fatal(err)
	}
	defer destroyTempdir(repo)
	hct := NewHgClient()
	defer hct.Disconnect()

	cfg := make([]string, 0)
	err = hct.Connect("hg", repo, cfg, true)
	if err != nil {
		t.Fatal(err)
	}
}

func TestShouldSucceedWhenRepoFoundAndInitrepoTrue(t *testing.T) {
	repo, err := createTempdir(t)
	if err != nil {
		t.Fatal(err)
	}
	defer destroyTempdir(repo)
	hct := NewHgClient()
	defer hct.Disconnect()

	hgexe := "hg"

	var cmd *exec.Cmd
	cmd = exec.Command(hgexe, "--cwd", repo, "init")
	if err = cmd.Run(); err != nil {
		t.Fatal(err)
	}

	cfg := make([]string, 0)
	err = hct.Connect(hgexe, repo, cfg, true)
	if err != nil {
		t.Fatal(err)
	}
}
