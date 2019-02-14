// Copyright 2019 The Alice-Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ocdb

import (
	"reflect"

	"go-hep.org/x/hep/groot/rbytes"
	"go-hep.org/x/hep/groot/root"
	"go-hep.org/x/hep/groot/rtypes"
)

func init() {
	{
		f := func() reflect.Value {
			var o AliCDBEntry
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliCDBEntry", f)
	}
}

var (
	_ root.Object        = (*AliCDBEntry)(nil)
	_ rbytes.Marshaler   = (*AliCDBEntry)(nil)
	_ rbytes.Unmarshaler = (*AliCDBEntry)(nil)
)

func init() {
	{
		f := func() reflect.Value {
			var o AliCDBId
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliCDBId", f)
	}
}

var (
	_ root.Object        = (*AliCDBId)(nil)
	_ rbytes.Marshaler   = (*AliCDBId)(nil)
	_ rbytes.Unmarshaler = (*AliCDBId)(nil)
)

func init() {
	{
		f := func() reflect.Value {
			var o AliCDBPath
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliCDBPath", f)
	}
}

var (
	_ root.Object        = (*AliCDBPath)(nil)
	_ rbytes.Marshaler   = (*AliCDBPath)(nil)
	_ rbytes.Unmarshaler = (*AliCDBPath)(nil)
)

func init() {
	{
		f := func() reflect.Value {
			var o AliCDBRunRange
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliCDBRunRange", f)
	}
}

var (
	_ root.Object        = (*AliCDBRunRange)(nil)
	_ rbytes.Marshaler   = (*AliCDBRunRange)(nil)
	_ rbytes.Unmarshaler = (*AliCDBRunRange)(nil)
)

func init() {
	{
		f := func() reflect.Value {
			var o AliCDBMetaData
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliCDBMetaData", f)
	}
}

var (
	_ root.Object        = (*AliCDBMetaData)(nil)
	_ rbytes.Marshaler   = (*AliCDBMetaData)(nil)
	_ rbytes.Unmarshaler = (*AliCDBMetaData)(nil)
)

func init() {
	{
		f := func() reflect.Value {
			var o AliMUON2DMap
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliMUON2DMap", f)
	}
}

var (
	_ root.Object        = (*AliMUON2DMap)(nil)
	_ rbytes.Marshaler   = (*AliMUON2DMap)(nil)
	_ rbytes.Unmarshaler = (*AliMUON2DMap)(nil)
)

func init() {
	{
		f := func() reflect.Value {
			var o AliMUONVStore
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliMUONVStore", f)
	}
}

var (
	_ root.Object        = (*AliMUONVStore)(nil)
	_ rbytes.Marshaler   = (*AliMUONVStore)(nil)
	_ rbytes.Unmarshaler = (*AliMUONVStore)(nil)
)

func init() {
	{
		f := func() reflect.Value {
			var o AliMpExMap
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliMpExMap", f)
	}
}

var (
	_ root.Object        = (*AliMpExMap)(nil)
	_ rbytes.Marshaler   = (*AliMpExMap)(nil)
	_ rbytes.Unmarshaler = (*AliMpExMap)(nil)
)

func init() {
	{
		f := func() reflect.Value {
			var o AliMUONCalibParamND
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliMUONCalibParamND", f)
	}
}

var (
	_ root.Object        = (*AliMUONCalibParamND)(nil)
	_ rbytes.Marshaler   = (*AliMUONCalibParamND)(nil)
	_ rbytes.Unmarshaler = (*AliMUONCalibParamND)(nil)
)

func init() {
	{
		f := func() reflect.Value {
			var o AliMUONVCalibParam
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliMUONVCalibParam", f)
	}
}

var (
	_ root.Object        = (*AliMUONVCalibParam)(nil)
	_ rbytes.Marshaler   = (*AliMUONVCalibParam)(nil)
	_ rbytes.Unmarshaler = (*AliMUONVCalibParam)(nil)
)
