// Copyright 2019 The go-jsonschema Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonschema

import (
	"encoding/json"
	"errors"
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

// Len returns the length of the list of types.
func (ts Types) Len() int {
	return len(ts)
}

// Less returns true if the i-th element in the list is listed before the j-th element.
func (ts Types) Less(i, j int) bool {
	return ts[i] < ts[j]
}

// Swap swaps the elements in positions i and j.
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

type Regexp interface {
	FindSubmatch(s []byte) [][]byte
	FindStringSubmatch(s string) []string
	FindStringSubmatchIndex(s string) []int
	ReplaceAllString(src, repl string) string
	FindString(s string) string
	FindAllString(s string, n int) []string
	MatchString(s string) bool
	SubexpNames() []string
}

type StringArray []String

type Default interface{}

type Items struct {
	Schemas     SchemaList
	HasMultiple bool
}

type Pattern Regexp

type AdditionalItems struct {
	Schema
}

type AdditionalProperties struct {
	Schema
}

type Definitions map[string]Schema

type Properties map[string]Schema

type PatternProperties map[Regexp]Schema

type DependencyMap struct {
	Names   map[string][]string
	Schemas map[string]Schema
}

type Const struct {
	Type  Type
	Value []interface{}
}

type Enum []*Const
