// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"reflect"
)

type (
	O_debug   bool //    --debug
	O_limit   int  // -l --limit
	O_mq      bool //    --mq
	O_profile bool //    --profile
	O_quiet   bool // -q --quiet
	O_remote  bool //    --remote

	// Maybe I should not allow the -R option,
	// so that each HgClient can only work on "it's own" repository ?
	// O_repository string

	O_traceback bool //    --traceback
	O_verbose   bool // -v --verbose
)

type hgDebugOpts struct {
	O_debug     bool
	O_traceback bool
	O_profile   bool
}

type optionAdder interface {
	addOption(interface{})
}

// addOption:
// Maybe I have to add some check using reflect.CanSet() ?
// see: http://stackoverflow.com/questions/6395076/in-golang-using-reflect-how-do-you-set-the-value-of-a-struct-field
// And eventually give an appropriate warning msg like "Command 'bla' does not accept option 'bla'.".
// But only as a warning, so still going on afterwards, just skipping the wrong option.
// So maybe this warning should be in some logfile or so.

func (o O_debug) addOption(i interface{}) {
	reflect.ValueOf(i).Elem().FieldByName("O_debug").SetBool(bool(o))
}

func (o O_limit) addOption(i interface{}) {
	reflect.ValueOf(i).Elem().FieldByName("O_limit").SetInt(int64(int(o)))
}

func (o O_mq) addOption(i interface{}) {
	reflect.ValueOf(i).Elem().FieldByName("O_mq").SetBool(bool(o))
}

func (o O_profile) addOption(i interface{}) {
	reflect.ValueOf(i).Elem().FieldByName("O_profile").SetBool(bool(o))
}

func (o O_quiet) addOption(i interface{}) {
	reflect.ValueOf(i).Elem().FieldByName("O_quiet").SetBool(bool(o))
}

func (o O_remote) addOption(i interface{}) {
	reflect.ValueOf(i).Elem().FieldByName("O_remote").SetBool(bool(o))
}

// func (o O_repository) addOption(i interface{}) {
// 	reflect.ValueOf(i).Elem().FieldByName("O_repository").SetString(string(o))
// }

func (o O_traceback) addOption(i interface{}) {
	reflect.ValueOf(i).Elem().FieldByName("O_traceback").SetBool(bool(o))
}

func (o O_verbose) addOption(i interface{}) {
	reflect.ValueOf(i).Elem().FieldByName("O_verbose").SetBool(bool(o))
}
