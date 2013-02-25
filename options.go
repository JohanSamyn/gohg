// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
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
	addOption(interface{}) error
	translateOption(*[]string)
}

func (o Debug) addOption(i interface{}) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Debug")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	} else {
		return errors.New("command <cmd> has no option Debug")
	}
	return nil
}

func (o Debug) translateOption(hgcmd []string) {
	if bool(o) {
		hgcmd = append(hgcmd, "--debug")
	}
}

func (o Destpath) addOption(i interface{}) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Destpath")
	if f.IsValid() || f.CanSet() {
		f.SetString(string(o))
	} else {
		return errors.New("command <cmd> has no option Destpath")
	}
	return nil
}

func (o Destpath) translateOption(hgcmd *[]string) {
	if string(o) != "" {
		*hgcmd = append(*hgcmd, string(o))
	}
}

func (o Limit) addOption(i interface{}) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Limit")
	if f.IsValid() || f.CanSet() {
		f.SetInt(int64(o))
	} else {
		return errors.New("command <cmd> has no option Limit")
	}
	return nil
}

func (o Limit) translateOption(hgcmd *[]string) {
	if int64(o) > 0 {
		*hgcmd = append(*hgcmd, "-l")
		*hgcmd = append(*hgcmd, strconv.Itoa(int(int64(o))))
	}
}

func (o Mq) addOption(i interface{}) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Mq")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	} else {
		return errors.New("command <cmd> has no option Mq")
	}
	return nil
}

func (o Mq) translateOption(hgcmd *[]string) {
	if bool(o) {
		*hgcmd = append(*hgcmd, "--mq")
	}
}

func (o Profile) addOption(i interface{}) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Profile")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	} else {
		return errors.New("command <cmd> has no option Profile")
	}
	return nil
}

func (o Profile) translateOption(hgcmd *[]string) {
	if bool(o) {
		*hgcmd = append(*hgcmd, "--profile")
	}
}

func (o Quiet) addOption(i interface{}) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Quiet")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	} else {
		return errors.New("command <cmd> has no option Quiet")
	}
	return nil
}

func (o Quiet) translateOption(hgcmd *[]string) {
	if bool(o) {
		*hgcmd = append(*hgcmd, "-q")
	}
}

func (o Remote) addOption(i interface{}) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Remote")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	} else {
		return errors.New("command <cmd> has no option Remote")
	}
	return nil
}

func (o Remote) translateOption(hgcmd *[]string) {
	if bool(o) {
		*hgcmd = append(*hgcmd, "--remote")
	}
}

func (o Rev) addOption(i interface{}) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Rev")
	if f.IsValid() || f.CanSet() {
		f.SetString(string(o))
	} else {
		return errors.New("command <cmd> has no option Rev")
	}
	return nil
}

func (o Rev) translateOption(hgcmd *[]string) {
	if string(o) != "" {
		*hgcmd = append(*hgcmd, "-r")
		*hgcmd = append(*hgcmd, string(o))
	}
}

func (o ShowBookmarks) addOption(i interface{}) error {
	f := reflect.ValueOf(i).Elem().FieldByName("ShowBookmarks")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	} else {
		return errors.New("command <cmd> has no option ShowBookmarks")
	}
	return nil
}

func (o ShowBookmarks) translateOption(hgcmd *[]string) {
	if bool(o) {
		*hgcmd = append(*hgcmd, "--bookmarks")
	}
}

func (o ShowBranch) addOption(i interface{}) error {
	f := reflect.ValueOf(i).Elem().FieldByName("ShowBranch")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	} else {
		return errors.New("command <cmd> has no option ShowBranch")
	}
	return nil
}

func (o ShowBranch) translateOption(hgcmd *[]string) {
	if bool(o) {
		*hgcmd = append(*hgcmd, "--branch")
	}
}

func (o ShowId) addOption(i interface{}) error {
	f := reflect.ValueOf(i).Elem().FieldByName("ShowId")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	} else {
		return errors.New("command <cmd> has no option ShowId")
	}
	return nil
}

func (o ShowId) translateOption(hgcmd *[]string) {
	fmt.Printf("ShowId.transl %v\n", bool(o))
	if bool(o) {
		*hgcmd = append(*hgcmd, "--id")
	}
}

func (o ShowNum) addOption(i interface{}) error {
	f := reflect.ValueOf(i).Elem().FieldByName("ShowNum")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	} else {
		return errors.New("command <cmd> has no option ShowNum")
	}
	return nil
}

func (o ShowNum) translateOption(hgcmd *[]string) {
	fmt.Printf("ShowNum.transl %v\n", bool(o))
	if bool(o) {
		*hgcmd = append(*hgcmd, "--num")
	}
}

func (o ShowTags) addOption(i interface{}) error {
	f := reflect.ValueOf(i).Elem().FieldByName("ShowTags")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	} else {
		return errors.New("command <cmd> has no option ShowTags")
	}
	return nil
}

func (o ShowTags) translateOption(hgcmd *[]string) {
	if bool(o) {
		*hgcmd = append(*hgcmd, "--tags")
	}
}

func (o Traceback) addOption(i interface{}) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Traceback")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	} else {
		return errors.New("command <cmd> has no option Traceback")
	}
	return nil
}

func (o Traceback) translateOption(hgcmd *[]string) {
	if bool(o) {
		*hgcmd = append(*hgcmd, "--traceback")
	}
}

func (o Verbose) addOption(i interface{}) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Verbose")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
	} else {
		return errors.New("command <cmd> has no option Traceback")
	}
	return nil
}

func (o Verbose) translateOption(hgcmd *[]string) {
	if bool(o) {
		*hgcmd = append(*hgcmd, "-v")
	}
}
