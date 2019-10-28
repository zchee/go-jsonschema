// Copyright 2019 The go-jsonschema Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package reflection wrapper of reflect package.
package reflection

import (
	"reflect"
	"unsafe"
)

const (
	KindDirectIface = 1 << 5
	KindGCProg      = 1 << 6 // Type.gc points to GC program
	KindMask        = (1 << 5) - 1
)

type StringHeader struct {
	Data unsafe.Pointer
	Len  int
}

type InterfaceHeader struct {
	Type uintptr        // 8 bytes for the pointer to the actual struct data followed by
	Word unsafe.Pointer // 8 bytes for the pointer to the type information
}

// Tflag is used by an rtype to signal what extra type information is
// available in the memory directly following the rtype value.
//
// Tflag values must be kept in sync with copies in:
//	cmd/compile/internal/gc/reflect.go
//	cmd/link/internal/ld/decodesym.go
//	runtime/type.go
type Tflag uint8

const (
	// TflagUncommon means that there is a pointer, *uncommonType,
	// just beyond the outer type structure.
	//
	// For example, if t.Kind() == Struct and t.tflag&TflagUncommon != 0,
	// then t has uncommonType data and it can be accessed as:
	//
	//	type tUncommon struct {
	//		structType
	//		u uncommonType
	//	}
	//	u := &(*tUncommon)(unsafe.Pointer(t)).u
	TflagUncommon Tflag = 1 << 0

	// TflagExtraStar means the name in the str field has an
	// extraneous '*' prefix. This is because for most types T in
	// a program, the type *T also exists and reusing the str data
	// saves binary size.
	TflagExtraStar Tflag = 1 << 1

	// TflagNamed means the type has a name.
	TflagNamed Tflag = 1 << 2

	// TflagRegularMemory means that equal and hash functions can treat
	// this type as a single region of t.size bytes.
	TflagRegularMemory Tflag = 1 << 3
)

// Method on non-interface type
type Method struct {
	Name  NameOff // name of method
	MType TypeOff // method type (without receiver)
	Ifn   TextOff // fn used in interface call (one-word receiver)
	Tfn   TextOff // fn used for normal method call
}

//go:linkname rtype reflect.rtype
type rtype unsafe.Pointer

type Rtype struct {
	size       uintptr
	Ptrdata    uintptr // number of bytes in the type that can contain pointers
	Hash       uint32  // hash of type; avoids computation in hash tables
	Tflag      Tflag   // extra type information flags
	align      uint8   // alignment of variable with this type
	fieldAlign uint8   // alignment of struct field with this type
	kind       uint8   // enumeration for C
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	Equal     func(unsafe.Pointer, unsafe.Pointer) bool
	GCdata    *byte   // garbage collection data
	Str       NameOff // string form
	PtrToThis TypeOff // type for pointer to this type, may be zero
}

// compile time check whether the RType implements reflect.Type interface.
// var _ reflect.Type = &RType{}

// UncommonType is present only for defined types or types with methods
// (if T is a defined type, the uncommonTypes for T and *T have methods).
// Using a pointer to this struct reduces the overall size required
// to describe a non-defined type with no methods.
type UncommonType struct {
	pkgPath NameOff // import path; empty for built-in types like int, string
	mcount  uint16  // number of methods
	xcount  uint16  // number of exported methods
	moff    uint32  // offset from this uncommontype to [mcount]method
	_       uint32  // unused
}

func (t *UncommonType) methods() []Method {
	if t.mcount == 0 {
		return nil
	}
	return (*[1 << 16]Method)(add(unsafe.Pointer(t), uintptr(t.moff), "t.mcount > 0"))[:t.mcount:t.mcount]
}

func (t *UncommonType) exportedMethods() []Method {
	if t.xcount == 0 {
		return nil
	}
	return (*[1 << 16]Method)(add(unsafe.Pointer(t), uintptr(t.moff), "t.xcount > 0"))[:t.xcount:t.xcount]
}

func (t *Rtype) NumMethod() int {
	if t.Kind() == reflect.Interface {
		tt := (*InterfaceType)(unsafe.Pointer(t))
		return tt.NumMethod()
	}
	return len(t.ExportedMethods())
}

func (t *Rtype) hasName() bool {
	return t.Tflag&TflagNamed != 0
}

func (t *Rtype) Name() string {
	if !t.hasName() {
		return ""
	}
	s := t.String()
	i := len(s) - 1
	for i >= 0 && s[i] != '.' {
		i--
	}
	return s[i+1:]
}

func (t *Rtype) ChanDir() reflect.ChanDir {
	if t.Kind() != reflect.Chan {
		panic("reflect: ChanDir of non-chan type " + t.String())
	}
	tt := (*ChanType)(unsafe.Pointer(t))
	return reflect.ChanDir(tt.Dir)
}

func (t *Rtype) PkgPath() string {
	if t.Tflag&TflagNamed == 0 {
		return ""
	}
	ut := t.Uncommon()
	if ut == nil {
		return ""
	}
	return t.NameOff(ut.pkgPath).Name()
}

func (t *Rtype) Align() int { return int(t.align) }

func (t *Rtype) FieldAlign() int { return int(t.fieldAlign) }

func (t *Rtype) Kind() reflect.Kind { return reflect.Kind(t.kind & KindMask) }

func (t *Rtype) pointers() bool { return t.Ptrdata != 0 }

func (t *Rtype) common() *rtype { return (*rtype)(unsafe.Pointer(t)) }

func (t *Rtype) ExportedMethods() []Method {
	ut := t.Uncommon()
	if ut == nil {
		return nil
	}
	return ut.exportedMethods()
}

func (t *Rtype) IsVariadic() bool {
	if t.Kind() != reflect.Func {
		panic("reflect: IsVariadic of non-func type " + t.String())
	}
	tt := (*FuncType)(unsafe.Pointer(t))
	return tt.OutCount&(1<<15) != 0
}

func (t *Rtype) Method(i int) (m reflect.Method) {
	if t.Kind() == reflect.Interface {
		tt := (*InterfaceType)(unsafe.Pointer(t))
		return tt.Method(i)
	}
	methods := t.ExportedMethods()
	if i < 0 || i >= len(methods) {
		panic("reflect: Method index out of range")
	}
	p := methods[i]
	pname := t.NameOff(p.Name)
	m.Name = pname.Name()
	fl := flag(reflect.Func)
	mtyp := t.TypeOff(p.MType)
	ft := (*FuncType)(unsafe.Pointer(mtyp))
	in := make([]reflect.Type, 0, 1+len(ft.in()))
	in = append(in, *(*reflect.Type)(unsafe.Pointer(&t)))
	for _, arg := range ft.in() {
		in = append(in, *(*reflect.Type)(unsafe.Pointer(&arg)))
	}
	out := make([]reflect.Type, 0, len(ft.out()))
	for _, ret := range ft.out() {
		out = append(out, *(*reflect.Type)(unsafe.Pointer(&ret)))
	}
	mt := reflect.FuncOf(in, out, ft.IsVariadic())
	m.Type = mt
	tfn := t.TextOff(p.Tfn)
	fn := unsafe.Pointer(&tfn)
	m.Func = *(*reflect.Value)(unsafe.Pointer(&Value{(*Rtype)(unsafe.Pointer(t)), fn, fl}))

	m.Index = i
	return m
}

func (t *Rtype) Elem() reflect.Type {
	switch t.Kind() {
	case reflect.Array:
		tt := (*ArrayType)(unsafe.Pointer(t))
		return ToType(tt.Elem)
	case reflect.Chan:
		tt := (*ChanType)(unsafe.Pointer(t))
		return ToType(tt.Elem)
	case reflect.Map:
		tt := (*MapType)(unsafe.Pointer(t))
		return ToType(tt.Elem)
	case reflect.Ptr:
		tt := (*PtrType)(unsafe.Pointer(t))
		return ToType(tt.Elem)
	case reflect.Slice:
		tt := (*SliceType)(unsafe.Pointer(t))
		return ToType(tt.Elem)
	}
	panic("reflect: Elem of invalid type " + t.String())
}

func (t *Rtype) MethodByName(name string) (m reflect.Method, ok bool) {
	if t.Kind() == reflect.Interface {
		tt := (*InterfaceType)(unsafe.Pointer(t))
		return tt.MethodByName(name)
	}
	ut := t.Uncommon()
	if ut == nil {
		return reflect.Method{}, false
	}
	// TODO(mdempsky): Binary search.
	for i, p := range ut.exportedMethods() {
		if t.NameOff(p.Name).Name() == name {
			return t.Method(i), true
		}
	}
	return reflect.Method{}, false
}

func (t *Rtype) Uncommon() *UncommonType {
	if t.Tflag&TflagUncommon == 0 {
		return nil
	}
	switch t.Kind() {
	case reflect.Struct:
		return &(*StructTypeUncommon)(unsafe.Pointer(t)).U
	case reflect.Ptr:
		type u struct {
			PtrType
			u UncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	case reflect.Func:
		type u struct {
			FuncType
			u UncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	case reflect.Slice:
		type u struct {
			SliceType
			u UncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	case reflect.Array:
		type u struct {
			ArrayType
			u UncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	case reflect.Chan:
		type u struct {
			ChanType
			u UncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	case reflect.Map:
		type u struct {
			MapType
			u UncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	case reflect.Interface:
		type u struct {
			InterfaceType
			u UncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	default:
		type u struct {
			Rtype
			u UncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	}
}

func (t *Rtype) String() string {
	s := t.NameOff(t.Str).Name()
	if t.Tflag&TflagExtraStar != 0 {
		return s[1:]
	}
	return s
}

func (t *Rtype) Size() uintptr { return t.size }

func (t *Rtype) Bits() int {
	if t == nil {
		panic("reflect: Bits of nil Type")
	}
	k := t.Kind()
	if k < reflect.Int || k > reflect.Complex128 {
		panic("reflect: Bits of non-arithmetic Type " + t.String())
	}
	return int(t.size) * 8
}

func (t *Rtype) Implements(u reflect.Type) bool {
	if u == nil {
		panic("reflect: nil type passed to Type.Implements")
	}
	if u.Kind() != reflect.Interface {
		panic("reflect: non-interface type passed to Type.Implements")
	}
	uu := (*rtype)(unsafe.Pointer(&u))
	tt := (*rtype)(unsafe.Pointer(&t))
	return implements(uu, tt)
}

func (t *Rtype) AssignableTo(u reflect.Type) bool {
	if u == nil {
		panic("reflect: nil type passed to Type.AssignableTo")
	}
	uu := (*rtype)(unsafe.Pointer(&u))
	tt := (*rtype)(unsafe.Pointer(&t))
	return directlyAssignable(uu, tt) || implements(uu, tt)
}

func (t *Rtype) ConvertibleTo(u reflect.Type) bool {
	if u == nil {
		panic("reflect: nil type passed to Type.ConvertibleTo")
	}
	return convertOp(unsafe.Pointer(&u), unsafe.Pointer(&t)) != nil
}

func (t *Rtype) Comparable() bool {
	return t.Equal != nil
}

func (t *Rtype) Field(i int) reflect.StructField {
	if t.Kind() != reflect.Struct {
		panic("reflect: Field of non-struct type " + t.String())
	}
	tt := (*StructType)(unsafe.Pointer(t))
	return tt.Field(i)
}

func (t *Rtype) FieldByIndex(index []int) reflect.StructField {
	if t.Kind() != reflect.Struct {
		panic("reflect: FieldByIndex of non-struct type " + t.String())
	}
	tt := (*StructType)(unsafe.Pointer(t))
	return tt.FieldByIndex(index)
}

func (t *Rtype) FieldByName(name string) (reflect.StructField, bool) {
	if t.Kind() != reflect.Struct {
		panic("reflect: FieldByName of non-struct type " + t.String())
	}
	tt := (*StructType)(unsafe.Pointer(t))
	return tt.FieldByName(name)
}

func (t *Rtype) FieldByNameFunc(match func(string) bool) (reflect.StructField, bool) {
	if t.Kind() != reflect.Struct {
		panic("reflect: FieldByNameFunc of non-struct type " + t.String())
	}
	tt := (*StructType)(unsafe.Pointer(t))
	return tt.FieldByNameFunc(match)
}

func (t *Rtype) In(i int) reflect.Type {
	if t.Kind() != reflect.Func {
		panic("reflect: In of non-func type " + t.String())
	}
	tt := (*FuncType)(unsafe.Pointer(t))
	return ToType(tt.in()[i])
}

func (t *Rtype) Key() reflect.Type {
	if t.Kind() != reflect.Map {
		panic("reflect: Key of non-map type " + t.String())
	}
	tt := (*MapType)(unsafe.Pointer(t))
	return ToType(tt.Key)
}

func (t *Rtype) Len() int {
	if t.Kind() != reflect.Array {
		panic("reflect: Len of non-array type " + t.String())
	}
	tt := (*ArrayType)(unsafe.Pointer(t))
	return int(tt.Len)
}

func (t *Rtype) NumField() int {
	if t.Kind() != reflect.Struct {
		panic("reflect: NumField of non-struct type " + t.String())
	}
	tt := (*StructType)(unsafe.Pointer(t))
	return len(tt.Fields)
}

func (t *Rtype) NumIn() int {
	if t.Kind() != reflect.Func {
		panic("reflect: NumIn of non-func type " + t.String())
	}
	tt := (*FuncType)(unsafe.Pointer(t))
	return int(tt.InCount)
}

func (t *Rtype) NumOut() int {
	if t.Kind() != reflect.Func {
		panic("reflect: NumOut of non-func type " + t.String())
	}
	tt := (*FuncType)(unsafe.Pointer(t))
	return len(tt.out())
}

func (t *Rtype) Out(i int) reflect.Type {
	if t.Kind() != reflect.Func {
		panic("reflect: Out of non-func type " + t.String())
	}
	tt := (*FuncType)(unsafe.Pointer(t))
	return ToType(tt.out()[i])
}

// implements reports whether the type V implements the interface type T.
func implements(T, V *rtype) bool {
	if (*Rtype)(unsafe.Pointer(T)).Kind() != reflect.Interface {
		return false
	}
	t := (*InterfaceType)(unsafe.Pointer(T))
	if len(t.Methods) == 0 {
		return true
	}

	// The same algorithm applies in both cases, but the
	// method tables for an interface type and a concrete type
	// are different, so the code is duplicated.
	// In both cases the algorithm is a linear scan over the two
	// lists - T's methods and V's methods - simultaneously.
	// Since method tables are stored in a unique sorted order
	// (alphabetical, with no duplicate method names), the scan
	// through V's methods must hit a match for each of T's
	// methods along the way, or else V does not implement T.
	// This lets us run the scan in overall linear time instead of
	// the quadratic time  a naive search would require.
	// See also ../runtime/iface.go.
	if (*Rtype)(unsafe.Pointer(V)).Kind() == reflect.Interface {
		v := (*InterfaceType)(unsafe.Pointer(V))
		i := 0
		for j := 0; j < len(v.Methods); j++ {
			tm := &t.Methods[i]
			tmName := t.NameOff(tm.Name)
			vm := &v.Methods[j]
			vmName := (*Rtype)(unsafe.Pointer(V)).NameOff(vm.Name)
			if vmName.Name() == tmName.Name() && (*Rtype)(unsafe.Pointer(V)).TypeOff(vm.Type) == t.TypeOff(tm.Type) {
				if !tmName.IsExported() {
					tmPkgPath := tmName.PkgPath()
					if tmPkgPath == "" {
						tmPkgPath = t.PkgPath.Name()
					}
					vmPkgPath := vmName.PkgPath()
					if vmPkgPath == "" {
						vmPkgPath = v.PkgPath.Name()
					}
					if tmPkgPath != vmPkgPath {
						continue
					}
				}
				if i++; i >= len(t.Methods) {
					return true
				}
			}
		}
		return false
	}

	v := (*Rtype)(unsafe.Pointer(V)).Uncommon()
	if v == nil {
		return false
	}
	i := 0
	vmethods := v.methods()
	for j := 0; j < int(v.mcount); j++ {
		tm := &t.Methods[i]
		tmName := t.NameOff(tm.Name)
		vm := vmethods[j]
		vmName := (*Rtype)(unsafe.Pointer(V)).NameOff(vm.Name)
		if vmName.Name() == tmName.Name() && (*Rtype)(unsafe.Pointer(V)).TypeOff(vm.MType) == t.TypeOff(tm.Type) {
			if !tmName.IsExported() {
				tmPkgPath := tmName.PkgPath()
				if tmPkgPath == "" {
					tmPkgPath = t.PkgPath.Name()
				}
				vmPkgPath := vmName.PkgPath()
				if vmPkgPath == "" {
					vmPkgPath = (*Rtype)(unsafe.Pointer(V)).NameOff(v.pkgPath).Name()
				}
				if tmPkgPath != vmPkgPath {
					continue
				}
			}
			if i++; i >= len(t.Methods) {
				return true
			}
		}
	}
	return false
}

// directlyAssignable reports whether a value x of type V can be directly
// assigned (using memmove) to a value of type T.
// https://golang.org/doc/go_spec.html#Assignability
// Ignoring the interface rules (implemented elsewhere)
// and the ideal constant rules (no ideal constants at run time).
func directlyAssignable(T, V *rtype) bool {
	// x's type V is identical to T?
	if T == V {
		return true
	}

	// Otherwise at least one of T and V must not be defined
	// and they must have the same kind.
	if (*Rtype)(unsafe.Pointer(T)).hasName() && (*Rtype)(unsafe.Pointer(V)).hasName() || (*Rtype)(unsafe.Pointer(T)).Kind() != (*Rtype)(unsafe.Pointer(V)).Kind() {
		return false
	}

	// x's type T and V must  have identical underlying types.
	return haveIdenticalUnderlyingType((*rtype)(unsafe.Pointer(T)), (*rtype)(unsafe.Pointer(V)), true)
}

func haveIdenticalType(T, V *rtype, cmpTags bool) bool {
	if cmpTags {
		return T == V
	}

	if (*Rtype)(unsafe.Pointer(T)).Name() != (*Rtype)(unsafe.Pointer(V)).Name() || (*Rtype)(unsafe.Pointer(T)).Kind() != (*Rtype)(unsafe.Pointer(V)).Kind() {
		return false
	}

	return haveIdenticalUnderlyingType((*Rtype)(unsafe.Pointer(T)).common(), (*Rtype)(unsafe.Pointer(V)).common(), false)
}

func haveIdenticalUnderlyingType(T, V *rtype, cmpTags bool) bool {
	if T == V {
		return true
	}

	kind := (*Rtype)(unsafe.Pointer(T)).Kind()
	if kind != (*Rtype)(unsafe.Pointer(V)).Kind() {
		return false
	}

	// Non-composite types of equal kind have same underlying type
	// (the predefined instance of the type).
	if reflect.Bool <= kind && kind <= reflect.Complex128 || kind == reflect.String || kind == reflect.UnsafePointer {
		return true
	}

	// Composite types.
	switch kind {
	case reflect.Array:
		telem := (*Rtype)(unsafe.Pointer(T)).Elem()
		velem := (*Rtype)(unsafe.Pointer(V)).Elem()
		return (*Rtype)(unsafe.Pointer(T)).Len() == (*Rtype)(unsafe.Pointer(V)).Len() && haveIdenticalType((*rtype)(unsafe.Pointer(&telem)), (*rtype)(unsafe.Pointer(&velem)), cmpTags)

	case reflect.Chan:
		// Special case:
		// x is a bidirectional channel value, T is a channel type,
		// and x's type V and T have identical element types.
		telem := (*Rtype)(unsafe.Pointer(T)).Elem()
		velem := (*Rtype)(unsafe.Pointer(V)).Elem()
		if (*Rtype)(unsafe.Pointer(V)).ChanDir() == reflect.BothDir && haveIdenticalType((*rtype)(unsafe.Pointer(&telem)), (*rtype)(unsafe.Pointer(&velem)), cmpTags) {
			return true
		}

		// Otherwise continue test for identical underlying type.
		return (*Rtype)(unsafe.Pointer(V)).ChanDir() == (*Rtype)(unsafe.Pointer(T)).ChanDir() && haveIdenticalType((*rtype)(unsafe.Pointer(&telem)), (*rtype)(unsafe.Pointer(&velem)), cmpTags)

	case reflect.Func:
		t := (*FuncType)(unsafe.Pointer(T))
		v := (*FuncType)(unsafe.Pointer(V))
		if t.OutCount != v.OutCount || t.InCount != v.InCount {
			return false
		}
		for i := 0; i < t.NumIn(); i++ {
			tin := t.In(i)
			vin := v.In(i)
			if !haveIdenticalType((*rtype)(unsafe.Pointer(&tin)), (*rtype)(unsafe.Pointer(&vin)), cmpTags) {
				return false
			}
		}
		for i := 0; i < t.NumOut(); i++ {
			tout := t.Out(i)
			vout := v.Out(i)
			if !haveIdenticalType((*rtype)(unsafe.Pointer(&tout)), (*rtype)(unsafe.Pointer(&vout)), cmpTags) {
				return false
			}
		}
		return true

	case reflect.Interface:
		t := (*InterfaceType)(unsafe.Pointer(T))
		v := (*InterfaceType)(unsafe.Pointer(V))
		if len(t.Methods) == 0 && len(v.Methods) == 0 {
			return true
		}
		// Might have the same methods but still
		// need a run time conversion.
		return false

	case reflect.Map:
		tkey := (*Rtype)(unsafe.Pointer(T)).Key()
		vkey := (*Rtype)(unsafe.Pointer(V)).Key()
		telem := (*Rtype)(unsafe.Pointer(T)).Elem()
		velem := (*Rtype)(unsafe.Pointer(V)).Elem()
		return haveIdenticalType((*rtype)(unsafe.Pointer(&tkey)), (*rtype)(unsafe.Pointer(&vkey)), cmpTags) && haveIdenticalType((*rtype)(unsafe.Pointer(&telem)), (*rtype)(unsafe.Pointer(&velem)), cmpTags)

	case reflect.Ptr, reflect.Slice:
		telem := (*Rtype)(unsafe.Pointer(T)).Elem()
		velem := (*Rtype)(unsafe.Pointer(V)).Elem()
		return haveIdenticalType((*rtype)(unsafe.Pointer(&telem)), (*rtype)(unsafe.Pointer(&velem)), cmpTags)

	case reflect.Struct:
		t := (*StructType)(unsafe.Pointer(T))
		v := (*StructType)(unsafe.Pointer(V))
		if len(t.Fields) != len(v.Fields) {
			return false
		}
		if t.PkgPath.Name() != v.PkgPath.Name() {
			return false
		}
		for i := range t.Fields {
			tf := &t.Fields[i]
			vf := &v.Fields[i]
			if tf.Name.Name() != vf.Name.Name() {
				return false
			}
			if !haveIdenticalType((*rtype)(unsafe.Pointer(&tf.Type)), (*rtype)(unsafe.Pointer(&vf.Type)), cmpTags) {
				return false
			}
			if cmpTags && tf.Name.Tag() != vf.Name.Tag() {
				return false
			}
			if tf.OffsetEmbed != vf.OffsetEmbed {
				return false
			}
		}
		return true
	}

	return false
}

// StructField
type StructField struct {
	Name        Name    // name is always non-empty
	Type        *Rtype  // type of field
	OffsetEmbed uintptr // byte offset of field<<1 | isEmbedded
}

// StructType represents a struct type.
type StructType struct {
	Rtype
	PkgPath Name
	Fields  []StructField // sorted by offset
}

// add returns p+x.
//
// The whySafe string is ignored, so that the function still inlines
// as efficiently as p+x, but all call sites should use the string to
// record why the addition is safe, which is to say why the addition
// does not cause x to advance to the very end of p's allocation
// and therefore point incorrectly at the next block in memory.
func add(p unsafe.Pointer, x uintptr, whySafe string) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}

// Name is an encoded type Name with optional extra data.
//
// The first byte is a bit field containing:
//
//	1<<0 the Name is exported
//	1<<1 tag data follows the Name
//	1<<2 pkgPath nameOff follows the Name and tag
//
// The next two bytes are the data length:
//
//	 l := uint16(data[1])<<8 | uint16(data[2])
//
// Bytes [3:3+l] are the string data.
//
// If tag data follows then bytes 3+l and 3+l+1 are the tag length,
// with the data following.
//
// If the import path follows, then 4 bytes at the end of
// the data form a nameOff. The import path is only set for concrete
// methods that are defined in a different package than their type.
//
// If a Name starts with "*", then the exported bit represents
// whether the pointed to type is exported.
type Name struct {
	bytes *byte
}

func (n Name) Data(off int, whySafe string) *byte {
	return (*byte)(add(unsafe.Pointer(n.bytes), uintptr(off), whySafe))
}

func (n Name) IsExported() bool {
	return (*n.bytes)&(1<<0) != 0
}

func (n Name) NameLen() int {
	return int(uint16(*n.Data(1, "name len field"))<<8 | uint16(*n.Data(2, "name len field")))
}

func (n Name) TagLen() int {
	if *n.Data(0, "name flag field")&(1<<1) == 0 {
		return 0
	}
	off := 3 + n.NameLen()
	return int(uint16(*n.Data(off, "name taglen field"))<<8 | uint16(*n.Data(off+1, "name taglen field")))
}

func (n Name) Name() (s string) {
	if n.bytes == nil {
		return
	}
	b := (*[4]byte)(unsafe.Pointer(n.bytes))

	hdr := (*StringHeader)(unsafe.Pointer(&s))
	hdr.Data = unsafe.Pointer(&b[3])
	hdr.Len = int(b[1])<<8 | int(b[2])
	return s
}

func (n Name) Tag() (s string) {
	tl := n.TagLen()
	if tl == 0 {
		return ""
	}
	nl := n.NameLen()
	hdr := (*StringHeader)(unsafe.Pointer(&s))
	hdr.Data = unsafe.Pointer(n.Data(3+nl+2, "non-empty string"))
	hdr.Len = tl
	return s
}

func (n Name) PkgPath() string {
	if n.bytes == nil || *n.Data(0, "name flag field")&(1<<2) == 0 {
		return ""
	}
	off := 3 + n.NameLen()
	if tl := n.TagLen(); tl > 0 {
		off += 2 + tl
	}
	var nameOff int32
	// Note that this field may not be aligned in memory,
	// so we cannot use a direct int32 assignment here.
	copy((*[4]byte)(unsafe.Pointer(&nameOff))[:], (*[4]byte)(unsafe.Pointer(n.Data(off, "name offset field")))[:])
	pkgPathName := Name{(*byte)(ResolveTypeOff(unsafe.Pointer(n.bytes), nameOff))}
	return pkgPathName.Name()
}

func NewName(n, tag string, exported bool) Name {
	if len(n) > 1<<16-1 {
		panic("reflect.nameFrom: name too long: " + n)
	}
	if len(tag) > 1<<16-1 {
		panic("reflect.nameFrom: tag too long: " + tag)
	}

	var bits byte
	l := 1 + 2 + len(n)
	if exported {
		bits |= 1 << 0
	}
	if len(tag) > 0 {
		l += 2 + len(tag)
		bits |= 1 << 1
	}

	b := make([]byte, l)
	b[0] = bits
	b[1] = uint8(len(n) >> 8)
	b[2] = uint8(len(n))
	copy(b[3:], n)
	if len(tag) > 0 {
		tb := b[3+len(n):]
		tb[0] = uint8(len(tag) >> 8)
		tb[1] = uint8(len(tag))
		copy(tb[2:], tag)
	}

	return Name{bytes: &b[0]}
}

// resolveNameOff resolves a name offset from a base pointer.
// The (*rtype).nameOff method is a convenience wrapper for this function.
// Implemented in the runtime package.
func ResolveNameOff(ptrInModule unsafe.Pointer, off int32) unsafe.Pointer

// ResolveTypeOff resolves an *rtype offset from a base type.
// The (*rtype).typeOff method is a convenience wrapper for this function.
// Implemented in the runtime package.
func ResolveTypeOff(rtype unsafe.Pointer, off int32) unsafe.Pointer

// resolveTextOff resolves an function pointer offset from a base type.
// The (*rtype).textOff method is a convenience wrapper for this function.
// Implemented in the runtime package.
func ResolveTextOff(rtype unsafe.Pointer, off int32) unsafe.Pointer

type NameOff int32 // offset to a name
type TypeOff int32 // offset to an *rtype
type TextOff int32 // offset from top of text section

func (t *Rtype) NameOff(off NameOff) Name {
	return Name{(*byte)(ResolveNameOff(unsafe.Pointer(t), int32(off)))}
}

func (t *Rtype) TypeOff(off TypeOff) *Rtype {
	return (*Rtype)(ResolveTypeOff(unsafe.Pointer(t), int32(off)))
}

func (t *Rtype) TextOff(off TextOff) unsafe.Pointer {
	return ResolveTextOff(unsafe.Pointer(t), int32(off))
}

// ArrayType represents a fixed array type.
type ArrayType struct {
	Rtype
	Elem  *Rtype // array element type
	Slice *Rtype // slice type
	Len   uintptr
}

// ChanType represents a channel type.
type ChanType struct {
	Rtype
	Elem *Rtype  // channel element type
	Dir  uintptr // channel direction (ChanDir)
}

// FuncType represents a function type.
//
// A *rtype for each in and out parameter is stored in an array that
// directly follows the FuncType (and possibly its uncommonType). So
// a function type with one method, one input, and one output is:
//
//	struct {
//		FuncType
//		uncommonType
//		[2]*rtype    // [0] is in, [1] is out
//	}
type FuncType struct {
	Rtype
	InCount  uint16
	OutCount uint16 // top bit is set if last input parameter is ...
}

func (t *FuncType) in() []*Rtype {
	uadd := unsafe.Sizeof(*t)
	if t.Tflag&TflagUncommon != 0 {
		uadd += unsafe.Sizeof(UncommonType{})
	}
	if t.InCount == 0 {
		return nil
	}
	return (*[1 << 20]*Rtype)(add(unsafe.Pointer(t), uadd, "t.inCount > 0"))[:t.InCount:t.InCount]
}

func (t *FuncType) out() []*Rtype {
	uadd := unsafe.Sizeof(*t)
	if t.Tflag&TflagUncommon != 0 {
		uadd += unsafe.Sizeof(UncommonType{})
	}
	outCount := t.OutCount & (1<<15 - 1)
	if outCount == 0 {
		return nil
	}
	return (*[1 << 20]*Rtype)(add(unsafe.Pointer(t), uadd, "outCount > 0"))[t.InCount : t.InCount+outCount : t.InCount+outCount]
}

// IMethod represents a method on an interface type
type IMethod struct {
	Name NameOff // name of method
	Type TypeOff // .(*FuncType) underneath
}

// InterfaceType represents an interface type.
type InterfaceType struct {
	Rtype
	PkgPath Name      // import path
	Methods []IMethod // sorted by hash
}

// NumMethod returns the number of interface methods in the type's method set.
func (t *InterfaceType) NumMethod() int { return len(t.Methods) }

// MapType represents a map type.
type MapType struct {
	Rtype
	Key    *Rtype // map key type
	Elem   *Rtype // map element (value) type
	Bucket *Rtype // internal bucket structure
	// function for hashing keys (ptr to key, seed) -> hash
	Hasher     func(unsafe.Pointer, uintptr) uintptr
	Keysize    uint8  // size of key slot
	Valuesize  uint8  // size of value slot
	Bucketsize uint16 // size of bucket
	Flags      uint32
}

// PtrType represents a pointer type.
type PtrType struct {
	Rtype
	Elem *Rtype // pointer element (pointed at) type
}

// SliceType represents a slice type.
type SliceType struct {
	Rtype
	Elem *Rtype // slice element type
}

type StructTypeUncommon struct {
	StructType
	U UncommonType
}

// ToType converts from a *rtype to a Type that can be returned
// to the client of package reflect. In gc, the only concern is that
// a nil *rtype must be replaced by a nil Type, but in gccgo this
// function takes care of ensuring that multiple *rtype for the same
// type are coalesced into a single Type.
func ToType(t *Rtype) reflect.Type {
	if t == nil {
		return nil
	}
	return *(*reflect.Type)(unsafe.Pointer(t))
}
