// this file was generated by gomacro command: import "image/jpeg"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "image/jpeg"
	. "reflect"
)

func init() {
	Binds["image/jpeg"] = map[string]Value{
		"Decode":	ValueOf(pkg.Decode),
		"DecodeConfig":	ValueOf(pkg.DecodeConfig),
		"DefaultQuality":	ValueOf(pkg.DefaultQuality),
		"Encode":	ValueOf(pkg.Encode),
	}
	Types["image/jpeg"] = map[string]Type{
		"FormatError":	TypeOf((*pkg.FormatError)(nil)).Elem(),
		"Options":	TypeOf((*pkg.Options)(nil)).Elem(),
		"Reader":	TypeOf((*pkg.Reader)(nil)).Elem(),
		"UnsupportedError":	TypeOf((*pkg.UnsupportedError)(nil)).Elem(),
	}
}