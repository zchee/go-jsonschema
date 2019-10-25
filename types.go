// Copyright 2019 The go-jsonschema Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonschema

import (
	"encoding/json"
	"errors"
	"regexp"

	"github.com/francoispqt/gojay"
)

// Type represents a JSON Schema type keyword.
type Type int

// The list of Type keyword.
const (
	UnspecifiedType Type = iota
	ArrayType
	BooleanType
	IntegerType
	NullType
	NumberType
	ObjectType
	StringType
)

// String implements fmt.Stringer.
func (t Type) String() string {
	switch t {
	case ArrayType:
		return string(ArrayTypeName)
	case BooleanType:
		return string(BooleanTypeName)
	case IntegerType:
		return string(IntegerTypeName)
	case NullType:
		return string(NullTypeName)
	case NumberType:
		return string(NumberTypeName)
	case ObjectType:
		return string(ObjectTypeName)
	case StringType:
		return string(StringTypeName)
	default:
		return string(UnspecifiedTypeName)
	}
}

// MarshalJSON implements json.Marshaler.
func (t Type) MarshalJSON() ([]byte, error) {
	switch t {
	case
		ArrayType,
		BooleanType,
		IntegerType,
		NullType,
		NumberType,
		ObjectType,
		StringType:

		return json.Marshal(t.String())

	default:
		return nil, errors.New("unknown type")
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Type) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	x := TypeFromString(string(data))
	if x == UnspecifiedType {
		return errors.New("unspecified type")
	}
	*t = x

	return nil
}

// TypeName represents a Type keyword name.
type TypeName string

// The list of TypeName.
const (
	UnspecifiedTypeName TypeName = "<unspecified>"
	ArrayTypeName       TypeName = "array"
	BooleanTypeName     TypeName = "boolean"
	IntegerTypeName     TypeName = "integer"
	NullTypeName        TypeName = "null"
	NumberTypeName      TypeName = "number"
	ObjectTypeName      TypeName = "object"
	StringTypeName      TypeName = "string"
)

// TypeFromString returs the Type from s.
func TypeFromString(s string) Type {
	switch TypeName(s) {
	case ArrayTypeName:
		return ArrayType
	case BooleanTypeName:
		return BooleanType
	case IntegerTypeName:
		return IntegerType
	case NullTypeName:
		return NullType
	case NumberTypeName:
		return NumberType
	case ObjectTypeName:
		return ObjectType
	case StringTypeName:
		return StringType
	default:
		return UnspecifiedType
	}
}

// Types represents a list of Type.
type Types []Type

// Contains returns true if the list of types contains p.
func (ts Types) Contains(t Type) bool {
	for _, v := range ts {
		if t == v {
			return true
		}
	}

	return false
}

// Len implements sort.Interface.
func (ts Types) Len() int {
	return len(ts)
}

// Less implements sort.Interface.
func (ts Types) Less(i, j int) bool {
	return ts[i] < ts[j]
}

// Swap implements sort.Interface.
func (ts Types) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

// UnmarshalJSON implements json.Unmarshaler.
func (ts *Types) UnmarshalJSON(data []byte) error {
	if data[0] != '[' {
		var t Type
		if err := json.Unmarshal(data, &t); err != nil {
			return err
		}

		*ts = Types{t}
		return nil
	}

	var list []Type
	if err := json.Unmarshal(data, &list); err != nil {
		return err
	}
	*ts = Types(list)

	return nil
}

var (
	// compile time check whether the Types implements gojay.MarshalerJSONArray interface.
	_ gojay.MarshalerJSONArray = &Types{}
	// compile time check whether the Types implements gojay.UnmarshalerJSONArray interface.
	_ gojay.UnmarshalerJSONArray = &Types{}
	// compile time check whether the Types implements Pooler interface.
	_ Pooler = &Types{}
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (ts *Types) MarshalJSONArray(enc *gojay.Encoder) {
	for _, e := range *ts {
		enc.Int(int(e))
	}
}

// IsNil implements gojay.MarshalerJSONArray.
//
// IsNil checks if instance is nil
func (ts *Types) IsNil() bool {
	return len(*ts) == 0
}

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (ts *Types) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value Type
	v := int(value)
	if err := dec.Int(&v); err != nil {
		return err
	}
	*ts = append(*ts, Type(v))

	return nil
}

// Reset implements Pooler.
//
// Reset reset fields.
func (ts *Types) Reset() {
	for i := range *ts {
		(*ts)[i] = 0
	}
	TypesPool.Put(ts)
}

// TypesStream represents a stream encoding and decoding to Types.
type TypesStream chan *Types

var (
	// compile time check whether the TypesStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*TypesStream)(nil)
	// compile time check whether the TypesStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*TypesStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s TypesStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Array(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s TypesStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := TypesPool.Get().(*Types)
	if err := dec.Array(o); err != nil {
		return err
	}
	s <- o

	return nil
}

// StringArray represents a String slice.
type StringArray []String

var (
	// compile time check whether the StringArray implements gojay.MarshalerJSONArray interface.
	_ gojay.MarshalerJSONArray = &StringArray{}
	// compile time check whether the StringArray implements gojay.UnmarshalerJSONArray interface.
	_ gojay.UnmarshalerJSONArray = &StringArray{}
	// compile time check whether the StringArray implements Pooler interface.
	_ Pooler = &StringArray{}
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (sa *StringArray) MarshalJSONArray(enc *gojay.Encoder) {
	for _, s := range *sa {
		enc.Object(&s)
	}
}

// IsNil implements gojay.MarshalerJSONArray.
//
// IsNil checks if instance is nil
func (sa *StringArray) IsNil() bool {
	return len(*sa) == 0
}

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (sa *StringArray) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value = String{}
	if err := dec.Object(&value); err != nil {
		return err
	}
	*sa = append(*sa, value)

	return nil
}

// Reset implements Pooler.
//
// Reset reset fields.
func (sa *StringArray) Reset() {
	for i := range *sa {
		(*sa)[i].Reset()
	}
	StringArrayPool.Put(sa)
}

// StringArrayStream represents a stream encoding and decoding to StringArray.
type StringArrayStream chan *StringArray

var (
	// compile time check whether the StringArrayStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*StringArrayStream)(nil)
	// compile time check whether the StringArrayStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*StringArrayStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s StringArrayStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Array(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s StringArrayStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := StringArrayPool.Get().(*StringArray)
	if err := dec.Array(o); err != nil {
		return err
	}
	s <- o

	return nil
}

// Default specifies a default value for an item.
type Default interface{}

// Schemas list of Schema.
type Schemas []*Schema

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (ss *Schemas) MarshalJSONArray(enc *gojay.Encoder) {
	for _, e := range *ss {
		enc.Object(e)
	}
}

// IsNil implements gojay.MarshalerJSONArray.
//
// IsNil checks if instance is nil
func (ss *Schemas) IsNil() bool {
	return len(*ss) == 0
}

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (ss *Schemas) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var s Schema
	if err := dec.Object(&s); err != nil {
		return err
	}
	*ss = append(*ss, &s)

	return nil
}

// Reset implements Pooler.
//
// Reset reset fields.
func (ss *Schemas) Reset() {
	for i := range *ss {
		(*ss)[i] = nil
		SchemaPool.Put((*ss)[i])
	}
}

// SchemasStream represents a stream encoding and decoding to Schemas.
type SchemasStream chan *Schemas

var (
	// compile time check whether the SchemasStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*SchemasStream)(nil)
	// compile time check whether the SchemasStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*SchemasStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s SchemasStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Array(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s SchemasStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := &Schemas{}
	if err := dec.Array(o); err != nil {
		return err
	}
	s <- o

	return nil
}

// Items the elements of the array may be anything at all.
//
// However, it’s often useful to validate the items of the array against some schema as well.
// This is done using the items, additionalItems, and contains keywords.
type Items struct {
	Schemas     SchemaList
	HasMultiple bool
}

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (i *Items) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keySchemas, &i.Schemas)
	enc.BoolKey(keyHasMultiple, i.HasMultiple)
}

// IsNil implements gojay.MarshalerJSONObject.
//
// IsNil checks if instance is nil.
func (i *Items) IsNil() bool {
	return i == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (i *Items) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keySchemas:
		var ss SchemaList
		err := dec.Array(&ss)
		if err == nil && len(ss) > 0 {
			i.Schemas = ss
		}
		return err

	case keyHasMultiple:
		return dec.Bool(&i.HasMultiple)

	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
//
// NKeys returns the number of keys to unmarshal.
func (i *Items) NKeys() int { return 2 }

// Reset implements Pooler.
//
// Reset reset fields.
func (i *Items) Reset() {
	for j := range i.Schemas {
		SchemaPool.Put(&i.Schemas[j])
	}
	i.Schemas = nil
	i.HasMultiple = false
	ItemsPool.Put(i)
}

// ItemsStream represents a stream encoding and decoding to Items.
type ItemsStream chan *Items

var (
	// compile time check whether the ItemsStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*ItemsStream)(nil)
	// compile time check whether the ItemsStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*ItemsStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s ItemsStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Object(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s ItemsStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := ItemsPool.Get().(*Items)
	if err := dec.Object(o); err != nil {
		return err
	}
	s <- o

	return nil
}

// Pattern represents a use regular expressions to express constraints.
type Pattern *regexp.Regexp

// AdditionalItems represents a JSON Schema AdditionalItems type.
//
// The elements of the array may be anything at all.
type AdditionalItems struct {
	*Schema
}

var (
	// compile time check whether the AdditionalItems implements gojay.MarshalerJSONObject interface.
	_ gojay.MarshalerJSONObject = &AdditionalItems{}
	// compile time check whether the AdditionalItems implements gojay.UnmarshalerJSONObject interface.
	_ gojay.UnmarshalerJSONObject = &AdditionalItems{}
	// compile time check whether the AdditionalItems implements Pooler interface.
	_ Pooler = &AdditionalItems{}
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (ai *AdditionalItems) MarshalJSONObject(enc *gojay.Encoder) {
	switch ai.Schema.Version() {
	case DraftVersion7:
		enc.Object(ai)
	}
}

// IsNil implements gojay.MarshalerJSONObject.
//
// IsNil checks if instance is nil.
func (ai *AdditionalItems) IsNil() bool {
	return ai == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (ai *AdditionalItems) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch ai.Schema.Version() {
	case DraftVersion7:
		return dec.Object(ai)
	}

	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
//
// NKeys returns the number of keys to unmarshal.
func (*AdditionalItems) NKeys() int { return 0 }

// Reset implements Pooler.
//
// Reset reset fields.
func (ai *AdditionalItems) Reset() {
	switch ai.Schema.Version() {
	case DraftVersion7:
		ai.Reset()
		SchemaPool.Put(ai)
	}
}

// AdditionalItemsStream represents a stream encoding and decoding to AdditionalItems.
type AdditionalItemsStream chan *AdditionalItems

var (
	// compile time check whether the AdditionalItemsStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*AdditionalItemsStream)(nil)
	// compile time check whether the AdditionalItemsStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*AdditionalItemsStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s AdditionalItemsStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Object(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s AdditionalItemsStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := SchemaPool.Get().(*Schema)
	if err := dec.Object(o); err != nil {
		return err
	}
	s <- &AdditionalItems{Schema: o}

	return nil
}

// AdditionalPropertieies is used to control the handling of extra stuff, that is, properties whose names are not listed in the properties keyword.
// By default any additional properties are allowed.
//
// The additionalProperties keyword may be either a boolean or an object. If additionalProperties is a boolean and set to false, no additional properties will be allowed.
type AdditionalProperties struct {
	*Schema
}

var (
	// compile time check whether the AdditionalProperties implements gojay.MarshalerJSONObject interface.
	_ gojay.MarshalerJSONObject = &AdditionalProperties{}
	// compile time check whether the AdditionalProperties implements gojay.UnmarshalerJSONObject interface.
	_ gojay.UnmarshalerJSONObject = &AdditionalProperties{}
	// compile time check whether the AdditionalProperties implements Pooler interface.
	_ Pooler = &AdditionalProperties{}
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (ap *AdditionalProperties) MarshalJSONObject(enc *gojay.Encoder) {
	switch ap.Version() {
	case DraftVersion7:
		enc.Object(ap)
	}
}

// IsNil implements gojay.MarshalerJSONObject.
//
// IsNil checks if instance is nil.
func (ap *AdditionalProperties) IsNil() bool {
	return ap == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (ap *AdditionalProperties) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch ap.Schema.Version() {
	case DraftVersion7:
		return dec.Object(ap)
	}

	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
//
// NKeys returns the number of keys to unmarshal.
func (*AdditionalProperties) NKeys() int { return 0 }

// Reset implements Pooler.
//
// Reset reset fields.
func (ap *AdditionalProperties) Reset() {
	switch ap.Schema.Version() {
	case DraftVersion7:
		ap.Reset()
		SchemaPool.Put(ap)
	}
}

// AdditionalPropertiesStream represents a stream encoding and decoding to AdditionalProperties.
type AdditionalPropertiesStream chan *AdditionalProperties

var (
	// compile time check whether the AdditionalPropertiesStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*AdditionalPropertiesStream)(nil)
	// compile time check whether the AdditionalPropertiesStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*AdditionalPropertiesStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s AdditionalPropertiesStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Object(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s AdditionalPropertiesStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := SchemaPool.Get().(*Schema)
	if err := dec.Object(o); err != nil {
		return err
	}
	s <- &AdditionalProperties{Schema: o}

	return nil
}

// Definitions provides a standardized location for schema authors
// to inline re-usable JSON Schemas into a more general schema. The
// keyword does not directly affect the validation result.
type Definitions map[string]Schema

var (
	// compile time check whether the Definitions implements gojay.MarshalerJSONObject interface.
	_ gojay.MarshalerJSONObject = &Definitions{}
	// compile time check whether the Definitions implements gojay.UnmarshalerJSONObject interface.
	_ gojay.UnmarshalerJSONObject = &Definitions{}
	// compile time check whether the Definitions implements Pooler interface.
	_ Pooler = &Definitions{}
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (d Definitions) MarshalJSONObject(enc *gojay.Encoder) {
	for k, v := range d {
		enc.ObjectKey(k, &v)
	}
}

// IsNil implements gojay.MarshalerJSONObject.
//
// IsNil checks if instance is nil.
func (d Definitions) IsNil() bool {
	return d == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (d Definitions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var s Schema
	err := dec.Object(&s)
	if err != nil {
		return err
	}
	d[k] = s
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
//
// NKeys returns the number of keys to unmarshal.
func (Definitions) NKeys() int { return 0 }

// Reset implements Pooler.
//
// Reset reset fields.
func (d Definitions) Reset() {
	// for i := range d {
	// 	d[i].Reset()
	// 	SchemaPool.Put(d[i])
	// }
}

// DefinitionsStream represents a stream encoding and decoding to Definitions.
type DefinitionsStream chan *Definitions

var (
	// compile time check whether the DefinitionsStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*DefinitionsStream)(nil)
	// compile time check whether the DefinitionsStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*DefinitionsStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s DefinitionsStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Object(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s DefinitionsStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	var o Definitions
	if err := dec.Object(&o); err != nil {
		return err
	}
	s <- &o

	return nil
}

// Properties (key-value pairs) on an object are defined using the properties keyword.
//
// The value of properties is an object, where each key is the name of a property and each value is a JSON schema used to validate that property.
type Properties map[string]Schema

var (
	// compile time check whether the Properties implements gojay.MarshalerJSONObject interface.
	_ gojay.MarshalerJSONObject = &Properties{}
	// compile time check whether the Properties implements gojay.UnmarshalerJSONObject interface.
	_ gojay.UnmarshalerJSONObject = &Properties{}
	// compile time check whether the Properties implements Pooler interface.
	_ Pooler = &Properties{}
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (p Properties) MarshalJSONObject(enc *gojay.Encoder) {
	for k, v := range p {
		enc.ObjectKey(k, &v)
	}
}

// IsNil implements gojay.MarshalerJSONObject.
//
// IsNil checks if instance is nil.
func (p Properties) IsNil() bool {
	return p == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (p Properties) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var s Schema
	err := dec.Object(&s)
	if err != nil {
		return err
	}
	p[k] = s
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
//
// NKeys returns the number of keys to unmarshal.
func (Properties) NKeys() int { return 0 }

// Reset implements Pooler.
//
// Reset reset fields.
func (p Properties) Reset() {
	// for i := range p {
	// 	switch p[i].Version() {
	// 	case DraftVersion7:
	// 		p[i].Reset()
	// 		SchemaPool.Put(p[i])
	// 	}
	// }
}

// PropertiesStream represents a stream encoding and decoding to Properties.
type PropertiesStream chan *Properties

var (
	// compile time check whether the PropertiesStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*PropertiesStream)(nil)
	// compile time check whether the PropertiesStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*PropertiesStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s PropertiesStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Object(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s PropertiesStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	var o Properties
	if err := dec.Object(&o); err != nil {
		return err
	}
	s <- &o

	return nil
}

// PatternProperties is the each property name of this object SHOULD be a valid regular expression, according to the ECMA 262 regular expression dialect.
// Each property value of this object MUST be a valid JSON Schema.
type PatternProperties map[*regexp.Regexp]*Schema

var (
	// compile time check whether the PatternProperties implements gojay.MarshalerJSONObject interface.
	_ gojay.MarshalerJSONObject = &PatternProperties{}
	// compile time check whether the PatternProperties implements gojay.UnmarshalerJSONObject interface.
	_ gojay.UnmarshalerJSONObject = &PatternProperties{}
	// compile time check whether the PatternProperties implements Pooler interface.
	_ Pooler = &PatternProperties{}
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (pp PatternProperties) MarshalJSONObject(enc *gojay.Encoder) {
	for k, v := range pp {
		enc.ObjectKey(k.String(), v)
	}
}

// IsNil implements gojay.MarshalerJSONObject.
//
// IsNil checks if instance is nil.
func (pp PatternProperties) IsNil() bool {
	return pp == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (pp PatternProperties) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var s Schema
	err := dec.Object(&s)
	if err != nil {
		return err
	}
	re := regexp.MustCompile(k)
	if re != nil {
		pp[re] = &s
	}

	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
//
// NKeys returns the number of keys to unmarshal.
func (PatternProperties) NKeys() int { return 0 }

// Reset implements Pooler.
//
// Reset reset fields.
func (pp PatternProperties) Reset() {
	for i := range pp {
		switch pp[i].Version() {
		case DraftVersion7:
			pp[i].Reset()
			SchemaPool.Put(pp[i])
		}
	}
}

// PatternPropertiesStream represents a stream encoding and decoding to PatternProperties.
type PatternPropertiesStream chan *PatternProperties

var (
	// compile time check whether the PatternPropertiesStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*PatternPropertiesStream)(nil)
	// compile time check whether the PatternPropertiesStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*PatternPropertiesStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s PatternPropertiesStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Object(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s PatternPropertiesStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	var o PatternProperties
	if err := dec.Object(&o); err != nil {
		return err
	}
	s <- &o

	return nil
}

// NameMap represents a sttring key map of []string.
type NameMap map[string][]string

var (
	// compile time check whether the NameMap implements gojay.MarshalerJSONObject interface.
	_ gojay.MarshalerJSONObject = &NameMap{}
	// compile time check whether the NameMap implements gojay.UnmarshalerJSONObject interface.
	_ gojay.UnmarshalerJSONObject = &NameMap{}
	// compile time check whether the NameMap implements Pooler interface.
	_ Pooler = &NameMap{}
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (nm NameMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, v := range nm {
		enc.SliceStringKey(k, v)
	}
}

// IsNil implements gojay.MarshalerJSONObject.
//
// IsNil checks if instance is nil.
func (nm NameMap) IsNil() bool {
	return nm == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (nm NameMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var ss []string
	err := dec.SliceString(&ss)
	if err != nil {
		return err
	}
	nm[k] = ss
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
//
// NKeys returns the number of keys to unmarshal.
func (NameMap) NKeys() int { return 0 }

// Reset implements Pooler.
//
// Reset reset fields.
func (nm NameMap) Reset() {
	for i := range nm {
		nm[i] = nil
	}
}

// NameMapStream represents a stream encoding and decoding to NameMap.
type NameMapStream chan *NameMap

var (
	// compile time check whether the NameMapStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*NameMapStream)(nil)
	// compile time check whether the NameMapStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*NameMapStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s NameMapStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Object(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s NameMapStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := &NameMap{}
	if err := dec.Object(o); err != nil {
		return err
	}
	s <- o

	return nil
}

// SchemaMap represents a sttring key map of Schema.
type SchemaMap map[string]*Schema

var (
	// compile time check whether the Schemas implements gojay.MarshalerJSONObject interface.
	_ gojay.MarshalerJSONObject = &SchemaMap{}
	// compile time check whether the Schemas implements gojay.UnmarshalerJSONObject interface.
	_ gojay.UnmarshalerJSONObject = &SchemaMap{}
	// compile time check whether the SchemaMap implements Pooler interface.
	_ Pooler = &SchemaMap{}
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (sm SchemaMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, v := range sm {
		enc.ObjectKey(k, v)
	}
}

// IsNil implements gojay.MarshalerJSONObject.
//
// IsNil checks if instance is nil.
func (sm SchemaMap) IsNil() bool {
	return sm == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (sm SchemaMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var s Schema
	err := dec.Object(&s)
	if err != nil {
		return err
	}
	sm[k] = &s

	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
//
// NKeys returns the number of keys to unmarshal.
func (SchemaMap) NKeys() int { return 0 }

// Reset implements Pooler.
//
// Reset reset fields.
func (sm SchemaMap) Reset() {
	for i := range sm {
		sm[i] = nil
	}
}

// SchemaMapStream represents a stream encoding and decoding to SchemaMap.
type SchemaMapStream chan *SchemaMap

var (
	// compile time check whether the SchemaMapStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*SchemaMapStream)(nil)
	// compile time check whether the SchemaMapStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*SchemaMapStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s SchemaMapStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Object(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s SchemaMapStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := &SchemaMap{}
	if err := dec.Object(o); err != nil {
		return err
	}
	s <- o

	return nil
}

// DependencyMap represents a Dependencies map.
type DependencyMap struct {
	Names   map[string][]string
	Schemas map[string]*Schema
}

var (
	// compile time check whether the DependencyMap implements gojay.MarshalerJSONObject interface.
	_ gojay.MarshalerJSONObject = &DependencyMap{}
	// compile time check whether the DependencyMap implements gojay.UnmarshalerJSONObject interface.
	_ gojay.UnmarshalerJSONObject = &DependencyMap{}
	// compile time check whether the DependencyMap implements Pooler interface.
	_ Pooler = &DependencyMap{}
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (dm *DependencyMap) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyNames, NameMap(dm.Names))
	enc.ObjectKey(keySchemas, SchemaMap(dm.Schemas))
}

// IsNil implements gojay.MarshalerJSONObject.
//
// IsNil checks if instance is nil.
func (dm *DependencyMap) IsNil() bool {
	return dm == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (dm *DependencyMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyNames:
		return dec.Object(NameMap(dm.Names))

	case keySchemas:
		return dec.Object(SchemaMap(dm.Schemas))
	}

	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
//
// NKeys returns the number of keys to unmarshal.
func (DependencyMap) NKeys() int { return 0 }

// Reset implements Pooler.
//
// Reset reset fields.
func (dm *DependencyMap) Reset() {
	for i := range dm.Names {
		dm.Names[i] = nil
	}
	for i := range dm.Schemas {
		dm.Schemas[i] = nil
	}
}

// DependencyMapStream represents a stream encoding and decoding to DependencyMap.
type DependencyMapStream chan *DependencyMap

var (
	// compile time check whether the DependencyMapStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*DependencyMapStream)(nil)
	// compile time check whether the DependencyMapStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*DependencyMapStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s DependencyMapStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Object(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s DependencyMapStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := &DependencyMap{}
	if err := dec.Object(o); err != nil {
		return err
	}
	s <- o

	return nil
}

// Const is used to restrict a value to a single value.
type Const struct {
	Type  Type
	Value []interface{}
}

var (
	// compile time check whether the Const implements gojay.MarshalerJSONObject interface.
	_ gojay.MarshalerJSONObject = &Const{}
	// compile time check whether the Const implements gojay.UnmarshalerJSONObject interface.
	_ gojay.UnmarshalerJSONObject = &Const{}
	// compile time check whether the Const implements Pooler interface.
	_ Pooler = &Const{}
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (c *Const) MarshalJSONObject(enc *gojay.Encoder) {
	enc.IntKey(keyType, int(c.Type))
	enc.AddInterfaceKey(keyValue, Interfaces(c.Value))
}

// IsNil implements gojay.MarshalerJSONObject.
//
// IsNil checks if instance is nil.
func (c *Const) IsNil() bool {
	return c == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (c *Const) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyType:
		itype := int(c.Type)
		err := dec.Int(&itype)
		return err

	case keyValue:
		ifaces := Interfaces(c.Value)
		return dec.Array(&ifaces)
	}

	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
//
// NKeys returns the number of keys to unmarshal
func (*Const) NKeys() int { return 2 }

// Reset implements Pooler.
//
// Reset reset fields.
func (c *Const) Reset() {
	c.Type = 0
	c.Value = nil
	ConstPool.Put(c)
}

// ConstStream represents a stream encoding and decoding to Const.
type ConstStream chan *Const

var (
	// compile time check whether the ConstStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*ConstStream)(nil)
	// compile time check whether the ConstStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*ConstStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s ConstStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Object(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s ConstStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := ConstPool.Get().(*Const)
	if err := dec.Object(o); err != nil {
		return err
	}
	s <- o

	return nil
}

// Enum is used to restrict a value to a fixed set of values. It must be an array with at least one element, where each element is unique.
type Enum []*Const

var (
	// compile time check whether the Enum implements gojay.MarshalerJSONArray interface.
	_ gojay.MarshalerJSONArray = &Enum{}
	// compile time check whether the Enum implements gojay.UnmarshalerJSONArray interface.
	_ gojay.UnmarshalerJSONArray = &Enum{}
	// compile time check whether the Enum implements Pooler interface.
	_ Pooler = &Enum{}
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (e *Enum) MarshalJSONArray(enc *gojay.Encoder) {
	for _, t := range *e {
		enc.Object(t)
	}
}

// IsNil implements gojay.MarshalerJSONArray.
//
// IsNil checks if instance is nil.
func (e *Enum) IsNil() bool {
	return e == nil
}

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (e *Enum) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var t Const
	if err := dec.Object(&t); err != nil {
		return err
	}
	*e = append(*e, &t)

	return nil
}

// Reset implements Pooler.
//
// Reset reset fields.
func (e *Enum) Reset() {
	for i := range *e {
		(*e)[i].Reset()
		ConstPool.Put((*e)[i])
	}
}

// EnumStream represents a stream encoding and decoding to Enum.
type EnumStream chan *Enum

var (
	// compile time check whether the EnumStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*EnumStream)(nil)
	// compile time check whether the EnumStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*EnumStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s EnumStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Array(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s EnumStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := EnumPool.Get().(*Enum)
	if err := dec.Array(o); err != nil {
		return err
	}
	s <- o

	return nil
}

// Interfaces represents a slice of interface.
type Interfaces []interface{}

var (
	// compile time check whether the Enum implements gojay.MarshalerJSONArray interface.
	_ gojay.MarshalerJSONArray = &Interfaces{}
	// compile time check whether the Enum implements gojay.UnmarshalerJSONArray interface.
	_ gojay.UnmarshalerJSONArray = &Interfaces{}
	// compile time check whether the Enum implements Pooler interface.
	_ Pooler = &Interfaces{}
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (v *Interfaces) MarshalJSONArray(enc *gojay.Encoder) {
	for _, t := range *v {
		enc.AddInterface(t)
	}
}

// IsNil implements gojay.MarshalerJSONArray.
func (v *Interfaces) IsNil() bool { return len(*v) == 0 }

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (v *Interfaces) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var t interface{}
	if err := dec.Interface(&t); err != nil {
		return err
	}
	*v = append(*v, t)

	return nil
}

// Reset implements Pooler.
//
// Reset reset fields.
func (v *Interfaces) Reset() {
	for i := range *v {
		(*v)[i] = nil
	}
}

// InterfacesStream represents a stream encoding and decoding to Interfaces.
type InterfacesStream chan *Interfaces

var (
	// compile time check whether the InterfacesStream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*InterfacesStream)(nil)
	// compile time check whether the InterfacesStream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*InterfacesStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s InterfacesStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Array(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s InterfacesStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := &Interfaces{}
	if err := dec.Array(o); err != nil {
		return err
	}
	s <- o

	return nil
}
