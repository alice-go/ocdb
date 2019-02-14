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
			var o Entry
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliCDBEntry", f)
	}
}

var (
	_ root.Object        = (*Entry)(nil)
	_ rbytes.Marshaler   = (*Entry)(nil)
	_ rbytes.Unmarshaler = (*Entry)(nil)
)

func init() {
	{
		f := func() reflect.Value {
			var o ID
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliCDBId", f)
	}
}

var (
	_ root.Object        = (*ID)(nil)
	_ rbytes.Marshaler   = (*ID)(nil)
	_ rbytes.Unmarshaler = (*ID)(nil)
)

func init() {
	{
		f := func() reflect.Value {
			var o Path
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliCDBPath", f)
	}
}

var (
	_ root.Object        = (*Path)(nil)
	_ rbytes.Marshaler   = (*Path)(nil)
	_ rbytes.Unmarshaler = (*Path)(nil)
)

func init() {
	{
		f := func() reflect.Value {
			var o RunRange
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliCDBRunRange", f)
	}
}

var (
	_ root.Object        = (*RunRange)(nil)
	_ rbytes.Marshaler   = (*RunRange)(nil)
	_ rbytes.Unmarshaler = (*RunRange)(nil)
)

func init() {
	{
		f := func() reflect.Value {
			var o MetaData
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliCDBMetaData", f)
	}
}

var (
	_ root.Object        = (*MetaData)(nil)
	_ rbytes.Marshaler   = (*MetaData)(nil)
	_ rbytes.Unmarshaler = (*MetaData)(nil)
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
