// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"reflect"
)

// These are the options the Hg commands can take.
type (
	Debug     bool   //    --debug
	Destpath  string // no equivalent Hg option, used by Init()
	Limit     int    // -l --limit
	Mq        bool   //    --mq
	Profile   bool   //    --profile
	Quiet     bool   // -q --quiet
	Remote    bool   //    --remote
	Traceback bool   //    --traceback
	Verbose   bool   // -v --verbose
)

type hgDebugOpts struct {
	Debug     bool
	Profile   bool
	Traceback bool
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
		f.SetInt(int64(int(o)))
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
