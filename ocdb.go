// Copyright 2019 The Alice-Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ocdb

//go:generate

import (
	"bytes"
	"fmt"
	"io"

	"go-hep.org/x/hep/groot/rbase"
	"go-hep.org/x/hep/groot/rbytes"
	"go-hep.org/x/hep/groot/rcont"
	"go-hep.org/x/hep/groot/root"
)

type AliCDBEntry struct {
	base  rbase.Object
	obj   root.Object     `groot:"fObject"`
	id    AliCDBId        `groot:"fId"`
	meta  *AliCDBMetaData `groot:"fMetaData"`
	owner bool            `groot:"fIsOwner"`
}

func (*AliCDBEntry) Class() string   { return "AliCDBEntry" }
func (*AliCDBEntry) RVersion() int16 { return 1 }

func (entry *AliCDBEntry) Display(w io.Writer) {
	fmt.Fprintf(w, `=== Entry ===
ID: %v
Owner: %v
`,
		entry.id, entry.owner,
	)
	if entry.meta != nil {
		fmt.Fprintf(w, "MetaData:\n")
		entry.meta.Display(w)
	}
	if entry.obj != nil {
		fmt.Fprintf(w, "Object: %T\n%v\n===\n", entry.obj, entry.obj)
	}
}

// MarshalROOT implements rbytes.Marshaler
func (o *AliCDBEntry) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(o.RVersion())

	o.base.MarshalROOT(w)
	w.WriteObjectAny(o.obj)
	o.id.MarshalROOT(w)
	w.WriteObjectAny(o.meta)
	w.WriteBool(o.owner)

	return w.SetByteCount(pos, o.Class())
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (o *AliCDBEntry) UnmarshalROOT(r *rbytes.RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	start := r.Pos()
	_, pos, bcnt := r.ReadVersion()

	if err := o.base.UnmarshalROOT(r); err != nil {
		return err
	}

	o.obj = r.ReadObjectAny()
	if err := o.id.UnmarshalROOT(r); err != nil {
		return err
	}
	o.meta = nil
	if obj := r.ReadObjectAny(); obj != nil {
		o.meta = obj.(*AliCDBMetaData)
	}
	o.owner = r.ReadBool()

	r.CheckByteCount(pos, bcnt, start, o.Class())
	return r.Err()
}

type AliCDBId struct {
	base    rbase.Object
	path    AliCDBPath     `groot:"fPath"`
	runs    AliCDBRunRange `groot:"fRunRange"`
	vers    int32          `groot:"fVersion"`
	subvers int32          `groot:"fSubVersion"`
	last    string         `groot:"fLastStorage"`
}

func (*AliCDBId) Class() string   { return "AliCDBId" }
func (*AliCDBId) RVersion() int16 { return 1 }

func (id AliCDBId) String() string {
	return fmt.Sprintf("AliCDBId{Path: %v, RunRange: %v, Version: 0x%x, SubVersion: 0x%x, Last: %q}",
		id.path, id.runs, id.vers, id.subvers, id.last,
	)
}

// MarshalROOT implements rbytes.Marshaler
func (o *AliCDBId) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(o.RVersion())

	o.base.MarshalROOT(w)
	o.path.MarshalROOT(w)
	o.runs.MarshalROOT(w)
	w.WriteI32(o.vers)
	w.WriteI32(o.subvers)
	w.WriteString(o.last)

	return w.SetByteCount(pos, o.Class())
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (o *AliCDBId) UnmarshalROOT(r *rbytes.RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	start := r.Pos()
	_, pos, bcnt := r.ReadVersion()

	if err := o.base.UnmarshalROOT(r); err != nil {
		return err
	}
	if err := o.path.UnmarshalROOT(r); err != nil {
		return err
	}
	if err := o.runs.UnmarshalROOT(r); err != nil {
		return err
	}
	o.vers = r.ReadI32()
	o.subvers = r.ReadI32()
	o.last = r.ReadString()

	r.CheckByteCount(pos, bcnt, start, o.Class())
	return r.Err()
}

type AliCDBPath struct {
	base     rbase.Object
	path     string `groot:"fPath"`
	lvl0     string `groot:"fLevel0"`
	lvl1     string `groot:"fLevel1"`
	lvl2     string `groot:"fLevel2"`
	valid    bool   `groot:"fIsValid"`
	wildcard bool   `groot:"fIsWildCard"`
}

func (p AliCDBPath) String() string {
	return fmt.Sprintf("Path{Path: %q, Level0: %q, Level1: %q, Level2: %q, Valid: %v, WildCard: %v}",
		p.path, p.lvl0, p.lvl1, p.lvl2, p.valid, p.wildcard,
	)
}

func (*AliCDBPath) Class() string   { return "AliCDBPath" }
func (*AliCDBPath) RVersion() int16 { return 1 }

// MarshalROOT implements rbytes.Marshaler
func (o *AliCDBPath) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(o.RVersion())

	o.base.MarshalROOT(w)
	w.WriteString(o.path)
	w.WriteString(o.lvl0)
	w.WriteString(o.lvl1)
	w.WriteString(o.lvl2)
	w.WriteBool(o.valid)
	w.WriteBool(o.wildcard)

	return w.SetByteCount(pos, o.Class())
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (o *AliCDBPath) UnmarshalROOT(r *rbytes.RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	start := r.Pos()
	_, pos, bcnt := r.ReadVersion()

	if err := o.base.UnmarshalROOT(r); err != nil {
		return err
	}

	o.path = r.ReadString()
	o.lvl0 = r.ReadString()
	o.lvl1 = r.ReadString()
	o.lvl2 = r.ReadString()
	o.valid = r.ReadBool()
	o.wildcard = r.ReadBool()

	r.CheckByteCount(pos, bcnt, start, o.Class())
	return r.Err()
}

type AliCDBRunRange struct {
	base  rbase.Object
	first int32 `groot:"fFirstRun"`
	last  int32 `groot:"fLastRun"`
}

func (*AliCDBRunRange) Class() string   { return "AliCDBRunRange" }
func (*AliCDBRunRange) RVersion() int16 { return 1 }

func (rr AliCDBRunRange) String() string {
	return fmt.Sprintf("RunRange{First: %d, Last: %d}", rr.first, rr.last)
}

// MarshalROOT implements rbytes.Marshaler
func (o *AliCDBRunRange) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(o.RVersion())

	o.base.MarshalROOT(w)
	w.WriteI32(o.first)
	w.WriteI32(o.last)

	return w.SetByteCount(pos, o.Class())
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (o *AliCDBRunRange) UnmarshalROOT(r *rbytes.RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	start := r.Pos()
	_, pos, bcnt := r.ReadVersion()

	if err := o.base.UnmarshalROOT(r); err != nil {
		return err
	}
	o.first = r.ReadI32()
	o.last = r.ReadI32()

	r.CheckByteCount(pos, bcnt, start, o.Class())
	return r.Err()
}

type AliCDBMetaData struct {
	base    rbase.Object
	class   string    `groot:"fObjectClassName"`
	resp    string    `groot:"fResponsible"`
	beam    uint32    `groot:"fBeamPeriod"`
	vers    string    `groot:"fAliRootVersion"`
	comment string    `groot:"fComment"`
	props   rcont.Map `groot:"fProperties"`
}

func (*AliCDBMetaData) Class() string   { return "AliCDBMetaData" }
func (*AliCDBMetaData) RVersion() int16 { return 1 }

func (meta *AliCDBMetaData) Display(w io.Writer) {
	fmt.Fprintf(w, "Class: %q\nResponsible: %q\nBeamPeriod: %d\nAliRoot Version: %q\nComment: %q\nProperties: %d\n",
		meta.class, meta.resp, meta.beam, meta.vers, meta.comment, len(meta.props.Table()),
	)
	for k, v := range meta.props.Table() {
		fmt.Fprintf(w, "  key: %v\n  val: %v\n", k, v)
	}
}

// MarshalROOT implements rbytes.Marshaler
func (o *AliCDBMetaData) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(o.RVersion())

	o.base.MarshalROOT(w)
	w.WriteString(o.class)
	w.WriteString(o.resp)
	w.WriteU32(o.beam)
	w.WriteString(o.vers)
	w.WriteString(o.comment)
	o.props.MarshalROOT(w)

	return w.SetByteCount(pos, o.Class())
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (o *AliCDBMetaData) UnmarshalROOT(r *rbytes.RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	start := r.Pos()
	_, pos, bcnt := r.ReadVersion()

	if err := o.base.UnmarshalROOT(r); err != nil {
		return err
	}

	o.class = r.ReadString()
	o.resp = r.ReadString()
	o.beam = r.ReadU32()
	o.vers = r.ReadString()
	o.comment = r.ReadString()

	if err := o.props.UnmarshalROOT(r); err != nil {
		return err
	}

	r.CheckByteCount(pos, bcnt, start, o.Class())
	return r.Err()
}

type AliMUON2DMap struct {
	base  AliMUONVStore
	exmap *AliMpExMap `groot:"fMap"`
	opt   bool        `groot:"fOptimizeForDEManu"`
}

func (*AliMUON2DMap) Class() string   { return "AliMUON2DMap" }
func (*AliMUON2DMap) RVersion() int16 { return 1 }

func (m *AliMUON2DMap) String() string {
	return fmt.Sprintf("MUON2DMap{Opt: %v, Map: %v}", m.opt, *m.exmap)
}

// MarshalROOT implements rbytes.Marshaler
func (o *AliMUON2DMap) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(o.RVersion())

	o.base.MarshalROOT(w)
	w.WriteObjectAny(o.exmap)
	w.WriteBool(o.opt)

	return w.SetByteCount(pos, o.Class())
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (o *AliMUON2DMap) UnmarshalROOT(r *rbytes.RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	start := r.Pos()
	_, pos, bcnt := r.ReadVersion()

	if err := o.base.UnmarshalROOT(r); err != nil {
		return err
	}

	o.exmap = nil
	if obj := r.ReadObjectAny(); obj != nil {
		o.exmap = obj.(*AliMpExMap)
	}
	o.opt = r.ReadBool()

	r.CheckByteCount(pos, bcnt, start, o.Class())
	return r.Err()
}

type AliMUONVStore struct {
	base rbase.Object
}

func (*AliMUONVStore) Class() string   { return "AliMUONVStore" }
func (*AliMUONVStore) RVersion() int16 { return 1 }

// MarshalROOT implements rbytes.Marshaler
func (o *AliMUONVStore) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(o.RVersion())

	o.base.MarshalROOT(w)

	return w.SetByteCount(pos, o.Class())
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (o *AliMUONVStore) UnmarshalROOT(r *rbytes.RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	start := r.Pos()
	_, pos, bcnt := r.ReadVersion()

	if err := o.base.UnmarshalROOT(r); err != nil {
		return err
	}

	r.CheckByteCount(pos, bcnt, start, o.Class())
	return r.Err()
}

type AliMpExMap struct {
	base rbase.Object
	objs rcont.ObjArray `groot:"fObjects"`
	keys rcont.ArrayL64 `groot:"fKeys"`
}

func (exmap AliMpExMap) String() string {
	o := new(bytes.Buffer)
	fmt.Fprintf(o, "ExMap{Objs: [")
	for i := 0; i < exmap.objs.Len(); i++ {
		if i > 0 {
			fmt.Fprintf(o, ", ")
		}
		fmt.Fprintf(o, "%v", exmap.objs.At(i))
	}
	fmt.Fprintf(o, "], Keys: %v}", exmap.keys.Data)
	return o.String()
}

func (*AliMpExMap) Class() string   { return "AliMpExMap" }
func (*AliMpExMap) RVersion() int16 { return 1 }

// MarshalROOT implements rbytes.Marshaler
func (o *AliMpExMap) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(o.RVersion())

	o.base.MarshalROOT(w)
	o.objs.MarshalROOT(w)
	o.keys.MarshalROOT(w)

	return w.SetByteCount(pos, o.Class())
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (o *AliMpExMap) UnmarshalROOT(r *rbytes.RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	start := r.Pos()
	_, pos, bcnt := r.ReadVersion()

	if err := o.base.UnmarshalROOT(r); err != nil {
		return err
	}

	if err := o.objs.UnmarshalROOT(r); err != nil {
		return err
	}

	if err := o.keys.UnmarshalROOT(r); err != nil {
		return err
	}

	r.CheckByteCount(pos, bcnt, start, o.Class())
	return r.Err()
}

type AliMUONCalibParamND struct {
	base AliMUONVCalibParam
	dim  int32     `groot:"fDimension"`
	size int32     `groot:"fSize"`
	n    int32     `groot:"fN"`
	vs   []float64 `groot:"fValues"`
}

func (*AliMUONCalibParamND) Class() string   { return "AliMUONCalibParamND" }
func (*AliMUONCalibParamND) RVersion() int16 { return 1 }

// MarshalROOT implements rbytes.Marshaler
func (o *AliMUONCalibParamND) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(o.RVersion())

	o.base.MarshalROOT(w)
	w.WriteI32(o.dim)
	w.WriteI32(o.size)
	w.WriteI32(o.n)
	w.WriteI8(1) // FIXME(sbinet)
	w.WriteFastArrayF64(o.vs)

	return w.SetByteCount(pos, o.Class())
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (o *AliMUONCalibParamND) UnmarshalROOT(r *rbytes.RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	start := r.Pos()
	_, pos, bcnt := r.ReadVersion()

	if err := o.base.UnmarshalROOT(r); err != nil {
		return err
	}

	o.dim = r.ReadI32()
	o.size = r.ReadI32()
	o.n = r.ReadI32()
	_ = r.ReadI8() // FIXME(sbinet)
	o.vs = r.ReadFastArrayF64(int(o.n))

	r.CheckByteCount(pos, bcnt, start, o.Class())
	return r.Err()
}

type AliMUONVCalibParam struct {
	base rbase.Object
}

func (*AliMUONVCalibParam) Class() string   { return "AliMUONVCalibParam" }
func (*AliMUONVCalibParam) RVersion() int16 { return 1 }

// MarshalROOT implements rbytes.Marshaler
func (o *AliMUONVCalibParam) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(o.RVersion())

	o.base.MarshalROOT(w)

	return w.SetByteCount(pos, o.Class())
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (o *AliMUONVCalibParam) UnmarshalROOT(r *rbytes.RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	start := r.Pos()
	_, pos, bcnt := r.ReadVersion()

	if err := o.base.UnmarshalROOT(r); err != nil {
		return err
	}

	r.CheckByteCount(pos, bcnt, start, o.Class())
	return r.Err()
}
