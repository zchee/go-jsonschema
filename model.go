// Copyright 2019 The go-jsonschema Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonschema

import (
	"encoding/json"

	"github.com/francoispqt/gojay"
)

// 4.2.1. Instance Data Model
//  https://tools.ietf.org/html/draft-handrews-json-schema-01#section-4.2.1

// Array is used for ordered elements. In JSON, each element in an array may be of a different type.
type Array struct {
	Value       []interface{}
	Initialized bool
}

var (
	// compile time check whether the Array implements gojay.MarshalerJSONObject interface.
	_ gojay.MarshalerJSONObject = &Array{}
	// compile time check whether the Array implements gojay.UnmarshalerJSONObject interface.
	_ gojay.UnmarshalerJSONObject = &Array{}
	// compile time check whether the Array implements Pooler interface.
	_ Pooler = &Array{}
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (a *Array) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey("Value", (*Interfaces)(&a.Value))
	enc.BoolKey("Initialized", a.Initialized)
}

// IsNil implements gojay.MarshalerJSONObject.
//
// IsNil checks if instance is nil.
func (a *Array) IsNil() bool {
	return a == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (a *Array) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "Value":
		var ifaces = Interfaces{}
		err := dec.Array(&ifaces)
		if err == nil && len(ifaces) > 0 {
			a.Value = ifaces
		}
		return err

	case "Initialized":
		return dec.Bool(&a.Initialized)
	}

	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
//
// NKeys returns the number of keys to unmarshal.
func (*Array) NKeys() int { return 2 }

// Reset implements Pooler.
//
// Reset reset fields.
func (a *Array) Reset() {
	a.Value = nil
	a.Initialized = false
	ArrayPool.Put(a)
}

// ArrayStream represents a stream encoding and decoding to Array.
type ArrayStream chan *Array

var (
	// compile time check whether the ArrayStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*ArrayStream)(nil)
	// compile time check whether the ArrayStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*ArrayStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s ArrayStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Object(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s ArrayStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := ArrayPool.Get().(*Array)
	if err := dec.Object(o); err != nil {
		return err
	}
	s <- o

	return nil
}

// Boolean matches only two special values: true and false.
//
// Note that values that evaluate to true or false, such as 1 and 0, are not accepted by the schema.
type Boolean struct {
	Value       bool
	Default     bool
	Initialized bool
}

var (
	// compile time check whether the Boolean implements gojay.MarshalerJSONObject interface.
	_ gojay.MarshalerJSONObject = &Boolean{}
	// compile time check whether the Boolean implements gojay.UnmarshalerJSONObject interface.
	_ gojay.UnmarshalerJSONObject = &Boolean{}
	// compile time check whether the Boolean implements Pooler interface.
	_ Pooler = &Boolean{}
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (b *Boolean) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("Value", b.Value)
	enc.BoolKey("Default", b.Default)
	enc.BoolKey("Initialized", b.Initialized)
}

// IsNil implements gojay.MarshalerJSONObject.
//
// IsNil checks if instance is nil.
func (b *Boolean) IsNil() bool {
	return b == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (b *Boolean) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "Value":
		return dec.Bool(&b.Value)

	case "Default":
		return dec.Bool(&b.Default)

	case "Initialized":
		return dec.Bool(&b.Initialized)
	}

	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
//
// NKeys returns the number of keys to unmarshal.
func (*Boolean) NKeys() int { return 3 }

// Reset implements Pooler.
//
// Reset reset fields.
func (b *Boolean) Reset() {
	b.Value = false
	b.Default = false
	b.Initialized = false
	BooleanPool.Put(b)
}

// BooleanStream represents a stream encoding and decoding to Boolean.
type BooleanStream chan *Boolean

var (
	// compile time check whether the BooleanStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*BooleanStream)(nil)
	// compile time check whether the BooleanStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*BooleanStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s BooleanStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Object(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s BooleanStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := BooleanPool.Get().(*Boolean)
	if err := dec.Object(o); err != nil {
		return err
	}
	s <- o

	return nil
}

// Integer is used for integral numbers.
type Integer struct {
	Value       int64
	Initialized bool
}

var (
	// compile time check whether the Integer implements gojay.MarshalerJSONObject interface.
	_ gojay.MarshalerJSONObject = &Integer{}
	// compile time check whether the Integer implements gojay.UnmarshalerJSONObject interface.
	_ gojay.UnmarshalerJSONObject = &Integer{}
	// compile time check whether the Integer implements Pooler interface.
	_ Pooler = &Integer{}
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (i *Integer) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Int64Key("Value", i.Value)
	enc.BoolKey("Initialized", i.Initialized)
}

// IsNil implements gojay.MarshalerJSONObject.
//
// IsNil checks if instance is nil.
func (i *Integer) IsNil() bool {
	return i == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (i *Integer) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "Value":
		return dec.Int64(&i.Value)

	case "Initialized":
		return dec.Bool(&i.Initialized)
	}

	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
//
// NKeys returns the number of keys to unmarshal.
func (*Integer) NKeys() int { return 2 }

// Reset implements Pooler.
//
// Reset reset fields.
func (i *Integer) Reset() {
	i.Value = 0
	i.Initialized = false
	IntegerPool.Put(i)
}

// IntegerStream represents a stream encoding and decoding to Integer.
type IntegerStream chan *Integer

var (
	// compile time check whether the IntegerStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*IntegerStream)(nil)
	// compile time check whether the IntegerStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*IntegerStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s IntegerStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Object(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s IntegerStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := IntegerPool.Get().(*Integer)
	if err := dec.Object(o); err != nil {
		return err
	}
	s <- o

	return nil
}

// Null is generally used to represent a missing value.
type Null struct {
	Value       json.RawMessage
	Initialized bool
}

var (
	// compile time check whether the Null implements gojay.MarshalerJSONObject interface.
	_ gojay.MarshalerJSONObject = &Null{}
	// compile time check whether the Null implements gojay.UnmarshalerJSONObject interface.
	_ gojay.UnmarshalerJSONObject = &Null{}
	// compile time check whether the Null implements Pooler interface.
	_ Pooler = &Null{}
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (n *Null) MarshalJSONObject(enc *gojay.Encoder) {
	raw := gojay.EmbeddedJSON(n.Value)
	enc.AddEmbeddedJSONKey("Value", &raw)
	enc.BoolKey("Initialized", n.Initialized)
}

// IsNil implements gojay.MarshalerJSONObject.
//
// IsNil checks if instance is nil.
func (n *Null) IsNil() bool {
	return n == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (n *Null) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "Value":
		raw := gojay.EmbeddedJSON{}
		err := dec.AddEmbeddedJSON(&raw)
		if err == nil && len(raw) > 0 {
			n.Value = []byte(raw)
		}
		return err

	case "Initialized":
		return dec.Bool(&n.Initialized)
	}

	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
//
// NKeys returns the number of keys to unmarshal.
func (*Null) NKeys() int { return 2 }

// Reset implements Pooler.
//
// Reset reset fields.
func (n *Null) Reset() {
	n.Value = nil
	n.Initialized = false
	NullPool.Put(n)
}

// NullStream represents a stream encoding and decoding to Null.
type NullStream chan *Null

var (
	// compile time check whether the NullStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*NullStream)(nil)
	// compile time check whether the NullStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*NullStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s NullStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Object(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s NullStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := NullPool.Get().(*Null)
	if err := dec.Object(o); err != nil {
		return err
	}
	s <- o

	return nil
}

// Number used for any numeric type, either integers or floating point numbers.
type Number struct {
	Value       float64
	Initialized bool
}

var (
	// compile time check whether the Number implements gojay.MarshalerJSONObject interface.
	_ gojay.MarshalerJSONObject = &Number{}
	// compile time check whether the Number implements gojay.UnmarshalerJSONObject interface.
	_ gojay.UnmarshalerJSONObject = &Number{}
	// compile time check whether the Number implements Pooler interface.
	_ Pooler = &Number{}
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (n *Number) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Float64Key("Value", n.Value)
	enc.BoolKey("Initialized", n.Initialized)
}

// IsNil implements gojay.MarshalerJSONObject.
//
// IsNil checks if instance is nil.
func (n *Number) IsNil() bool {
	return n == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (n *Number) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "Value":
		return dec.Float64(&n.Value)

	case "Initialized":
		return dec.Bool(&n.Initialized)
	}

	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
//
// NKeys returns the number of keys to unmarshal.
func (*Number) NKeys() int { return 2 }

// Reset implements Pooler.
//
// Reset reset fields.
func (n *Number) Reset() {
	n.Value = 0.0
	n.Initialized = false
	NumberPool.Put(n)
}

// NumberStream represents a stream encoding and decoding to Number.
type NumberStream chan *Number

var (
	// compile time check whether the NumberStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*NumberStream)(nil)
	// compile time check whether the NumberStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*NumberStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s NumberStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Object(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s NumberStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := NumberPool.Get().(*Number)
	if err := dec.Object(o); err != nil {
		return err
	}
	s <- o

	return nil
}

// Object is the mapping type in JSON. They map 'keys' to 'values'.
//
// In JSON, the 'keys' must always be strings. Each of these pairs is conventionally referred to as a 'property'.
type Object struct {
	Value       interface{}
	Initialized bool
}

var (
	// compile time check whether the Object implements gojay.MarshalerJSONObject interface.
	_ gojay.MarshalerJSONObject = &Object{}
	// compile time check whether the Object implements gojay.UnmarshalerJSONObject interface.
	_ gojay.UnmarshalerJSONObject = &Object{}
	// compile time check whether the Object implements Pooler interface.
	_ Pooler = &Object{}
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (o *Object) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddInterfaceKey("Value", o.Value)
	enc.BoolKey("Initialized", o.Initialized)
}

// IsNil implements gojay.MarshalerJSONObject.
//
// IsNil checks if instance is nil.
func (o *Object) IsNil() bool {
	return o == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (o *Object) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "Value":
		return dec.Interface(&o.Value)

	case "Initialized":
		return dec.Bool(&o.Initialized)
	}

	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
//
// NKeys returns the number of keys to unmarshal.
func (*Object) NKeys() int { return 2 }

// Reset implements Pooler.
//
// Reset reset fields.
func (o *Object) Reset() {
	o.Value = nil
	o.Initialized = false
	ObjectPool.Put(o)
}

// ObjectStream represents a stream encoding and decoding to Object.
type ObjectStream chan *Object

var (
	// compile time check whether the ObjectStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*ObjectStream)(nil)
	// compile time check whether the ObjectStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*ObjectStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s ObjectStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Object(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s ObjectStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := ObjectPool.Get().(*Object)
	if err := dec.Object(o); err != nil {
		return err
	}
	s <- o

	return nil
}

// String is used for strings of text. It may contain Unicode characters.
type String struct {
	Value       string
	Initialized bool
}

var (
	// compile time check whether the String implements gojay.MarshalerJSONObject interface.
	_ gojay.MarshalerJSONObject = &String{}
	// compile time check whether the String implements gojay.UnmarshalerJSONObject interface.
	_ gojay.UnmarshalerJSONObject = &String{}
	// compile time check whether the String implements Pooler interface.
	_ Pooler = &String{}
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (s *String) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("Value", s.Value)
	enc.BoolKey("Initialized", s.Initialized)
}

// IsNil implements gojay.MarshalerJSONObject.
//
// IsNil checks if instance is nil.
func (s *String) IsNil() bool {
	return s == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (s *String) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "Value":
		return dec.String(&s.Value)

	case "Initialized":
		return dec.Bool(&s.Initialized)
	}

	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
//
// NKeys returns the number of keys to unmarshal.
func (*String) NKeys() int { return 2 }

// Reset implements Pooler.
//
// Reset reset fields.
func (s *String) Reset() {
	s.Value = ""
	s.Initialized = false
	StringPool.Put(s)
}

// StringStream represents a stream encoding and decoding to String.
type StringStream chan *String

var (
	// compile time check whether the StringStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*StringStream)(nil)
	// compile time check whether the StringStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*StringStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s StringStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Object(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s StringStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := StringPool.Get().(*String)
	if err := dec.Object(o); err != nil {
		return err
	}
	s <- o

	return nil
}
