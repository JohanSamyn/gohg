// Copyright 2012, The gohg Authors. All rights reserved.
// Use of this source code is governed by a BSD style license
// that can be found in the LICENSE file.

package gohg

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var errstr = "command %s has no option %s"

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
	Template      string //    --template
	Traceback     bool   //    --traceback
	Verbose       bool   // -v --verbose
)

type optionAdder interface {
	addOption(interface{}, *[]string) error
}

func (o Debug) addOption(i interface{}, hgcmd *[]string) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Debug")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
		if bool(o) {
			*hgcmd = append(*hgcmd, "--debug")
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), "Debug")
	}
	return nil
}

func (o Destpath) addOption(i interface{}, hgcmd *[]string) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Destpath")
	if f.IsValid() || f.CanSet() {
		f.SetString(string(o))
		if string(o) != "" {
			*hgcmd = append(*hgcmd, string(o))
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), "Destpath")
	}
	return nil
}

func (o Limit) addOption(i interface{}, hgcmd *[]string) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Limit")
	if f.IsValid() || f.CanSet() {
		f.SetInt(int64(o))
		if int64(o) > 0 {
			*hgcmd = append(*hgcmd, "-l")
			*hgcmd = append(*hgcmd, strconv.Itoa(int(int64(o))))
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), "Limit")
	}
	return nil
}

func (o Mq) addOption(i interface{}, hgcmd *[]string) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Mq")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
		if bool(o) {
			*hgcmd = append(*hgcmd, "--mq")
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), "Mq")
	}
	return nil
}

func (o Profile) addOption(i interface{}, hgcmd *[]string) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Profile")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
		if bool(o) {
			*hgcmd = append(*hgcmd, "--profile")
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), "Profile")
	}
	return nil
}

func (o Quiet) addOption(i interface{}, hgcmd *[]string) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Quiet")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
		if bool(o) {
			*hgcmd = append(*hgcmd, "-q")
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), "Quiet")
	}
	return nil
}

func (o Remote) addOption(i interface{}, hgcmd *[]string) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Remote")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
		if bool(o) {
			*hgcmd = append(*hgcmd, "--remote")
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), "Remote")
	}
	return nil
}

func (o Rev) addOption(i interface{}, hgcmd *[]string) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Rev")
	if f.IsValid() || f.CanSet() {
		f.SetString(string(o))
		if string(o) != "" {
			*hgcmd = append(*hgcmd, "-r")
			*hgcmd = append(*hgcmd, string(o))
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), "Rev")
	}
	return nil
}

func (o ShowBookmarks) addOption(i interface{}, hgcmd *[]string) error {
	f := reflect.ValueOf(i).Elem().FieldByName("ShowBookmarks")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
		if bool(o) {
			*hgcmd = append(*hgcmd, "--bookmarks")
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), "ShowBookmarks")
	}
	return nil
}

func (o ShowBranch) addOption(i interface{}, hgcmd *[]string) error {
	f := reflect.ValueOf(i).Elem().FieldByName("ShowBranch")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
		if bool(o) {
			*hgcmd = append(*hgcmd, "--branch")
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), "ShowBranch")
	}
	return nil
}

func (o ShowId) addOption(i interface{}, hgcmd *[]string) error {
	f := reflect.ValueOf(i).Elem().FieldByName("ShowId")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
		// fmt.Printf("ShowId.transl %v\n", bool(o))
		if bool(o) {
			*hgcmd = append(*hgcmd, "--id")
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), "ShowId")
	}
	return nil
}

func (o ShowNum) addOption(i interface{}, hgcmd *[]string) error {
	f := reflect.ValueOf(i).Elem().FieldByName("ShowNum")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
		// fmt.Printf("ShowNum.transl %v\n", bool(o))
		if bool(o) {
			*hgcmd = append(*hgcmd, "--num")
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), "ShowNum")
	}
	return nil
}

func (o ShowTags) addOption(i interface{}, hgcmd *[]string) error {
	f := reflect.ValueOf(i).Elem().FieldByName("ShowTags")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
		if bool(o) {
			*hgcmd = append(*hgcmd, "--tags")
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), "ShowTags")
	}
	return nil
}

func (o Template) addOption(i interface{}, hgcmd *[]string) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Template")
	if f.IsValid() || f.CanSet() {
		f.SetString(string(o))
		if string(o) != "" {
			*hgcmd = append(*hgcmd, "--template")
			*hgcmd = append(*hgcmd, string(o))
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), "Traceback")
	}
	return nil
}

func (o Traceback) addOption(i interface{}, hgcmd *[]string) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Traceback")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
		if bool(o) {
			*hgcmd = append(*hgcmd, "--traceback")
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), "Traceback")
	}
	return nil
}

func (o Verbose) addOption(i interface{}, hgcmd *[]string) error {
	f := reflect.ValueOf(i).Elem().FieldByName("Verbose")
	if f.IsValid() || f.CanSet() {
		f.SetBool(bool(o))
		if bool(o) {
			*hgcmd = append(*hgcmd, "-v")
		}
	} else {
		return fmt.Errorf(errstr, strings.Title((*hgcmd)[0]), "Verbose")
	}
	return nil
}
