// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE.md file.

package gohg

import (
	"os"
	"strings"
	"testing"
)

func TestHgClient_ShowConfig(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	// We have to disconnect and later reconnect again, because inbetween we
	// have to write the config file for the repo. As there is no config file
	// after hg init yet, the connection otherwise thinks there is no content
	// in the repo config file, and the test fails.
	he := hct.HgExe()
	rr := hct.RepoRoot()
	hct.Disconnect()

	f, err := os.Create(hct.RepoRoot() + "/.hg/hgrc")
	_, _ = f.WriteString("[paths]\ndefault=the-default-path\n")
	f.Sync()
	f.Close()

	// Now we reconnect again, now that there is a config file for the testrepo.
	hct.Connect(he, rr, nil, false)

	var expected string = hct.RepoRoot() + string(os.PathSeparator) + "the-default-path\n"

	got, err := hct.ShowConfig(nil, []string{"paths.default"})
	if err != nil {
		t.Error(err)
	}

	if string(got) != expected {
		t.Fatalf("Test ShowConfig: expected:\n%s\n but got:\n%s\n", expected, got)
	}
}

func TestHgClient_ShowConfig_Debug(t *testing.T) {
	hct := setup(t)
	defer teardown(t, hct)

	// We have to disconnect and later reconnect again, because inbetween we
	// have to write the config file for the repo. As there is no config file
	// after hg init yet, the connection otherwise thinks there is no content
	// in the repo config file, and the test fails.
	he := hct.HgExe()
	rr := hct.RepoRoot()
	hct.Disconnect()

	f, err := os.Create(hct.RepoRoot() + "/.hg/hgrc")
	_, _ = f.WriteString("[paths]\ndefault=the-default-path\n")
	f.Sync()
	f.Close()

	// Now we reconnect again, now that there is a config file for the testrepo.
	hct.Connect(he, rr, nil, false)

	var expected string = "none: " + hct.RepoRoot() + string(os.PathSeparator) + "the-default-path\n"

	gotraw, err := hct.ShowConfig([]HgOption{Debug(true)}, []string{"paths.default"})
	if err != nil {
		t.Error(err)
	}

	var got string
	g := strings.Split(string(gotraw), "\n")
	for _, e := range g {
		if strings.HasPrefix(e, "read config from: ") {
			continue
		}
		if len(got) > 0 {
			got += "\n"
		}
		got += e
	}

	if got != expected {
		t.Fatalf("Test ShowConfig: expected:\n%s\n but got:\n%s\n", expected, got)
	}
}
