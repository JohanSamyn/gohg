// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"reflect"
)

// These are the options the Hg commands can take.
type (
	Debug         bool   //    --debug
	Destpath      string // no equivalent Hg option, used by Init()
	Limit         int    // -l --limit
	Mq            bool   //    --mq
	Profile       bool   //    --profile
	Quiet         bool   // -q --quiet
	Remote        bool   //    --remote
	Rev           string // -r -- rev REV
	ShowBookmarks bool   // -B --bookmarks
	ShowBranch    bool   // -b --branch
	ShowId        bool   // -i --id
	ShowNum       bool   // -n --num
	ShowTags      bool   // -t --tags
	Traceback     bool   //    --traceback
	Verbose       bool   // -v --verbose
)

type hgDebugOpts struct {
	Debug
	Profile
	Traceback
}

type optionAdder interface {
	addOption(interface{})
}

func (o Debug) addOption(i interface{}) {
	f := reflect.ValueOf(i).Elem().FieldByName("Debug")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	}
}

func (o Destpath) addOption(i interface{}) {
	f := reflect.ValueOf(i).Elem().FieldByName("Destpath")
	if f.IsValid() || f.CanSet() {
		f.SetString(string(o))
	}
}

func (o Limit) addOption(i interface{}) {
	f := reflect.ValueOf(i).Elem().FieldByName("Limit")
	if f.IsValid() || f.CanSet() {
		f.SetInt(int64(o))
	}
}

func (o Mq) addOption(i interface{}) {
	f := reflect.ValueOf(i).Elem().FieldByName("Mq")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	}
}

func (o Profile) addOption(i interface{}) {
	f := reflect.ValueOf(i).Elem().FieldByName("Profile")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	}
}

func (o Quiet) addOption(i interface{}) {
	f := reflect.ValueOf(i).Elem().FieldByName("Quiet")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	}
}

func (o Remote) addOption(i interface{}) {
	f := reflect.ValueOf(i).Elem().FieldByName("Remote")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	}
}

func (o Rev) addOption(i interface{}) {
	f := reflect.ValueOf(i).Elem().FieldByName("Rev")
	if f.IsValid() || f.CanSet() {
		f.SetString(string(o))
	}
}

func (o ShowBookmarks) addOption(i interface{}) {
	f := reflect.ValueOf(i).Elem().FieldByName("ShowBookmarks")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	}
}

func (o ShowBranch) addOption(i interface{}) {
	f := reflect.ValueOf(i).Elem().FieldByName("ShowBranch")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	}
}

func (o ShowId) addOption(i interface{}) {
	f := reflect.ValueOf(i).Elem().FieldByName("ShowId")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	}
}

func (o ShowNum) addOption(i interface{}) {
	f := reflect.ValueOf(i).Elem().FieldByName("ShowNum")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	}
}

func (o ShowTags) addOption(i interface{}) {
	f := reflect.ValueOf(i).Elem().FieldByName("ShowTags")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	}
}

func (o Traceback) addOption(i interface{}) {
	f := reflect.ValueOf(i).Elem().FieldByName("Traceback")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	}
}

func (o Verbose) addOption(i interface{}) {
	f := reflect.ValueOf(i).Elem().FieldByName("Verbose")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	}
}
