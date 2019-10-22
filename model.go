// Copyright 2019 The go-jsonschema Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonschema

import "encoding/json"

// 4.2.1. Instance Data Model
//  https://tools.ietf.org/html/draft-handrews-json-schema-01#section-4.2.1

// Array is used for ordered elements. In JSON, each element in an array may be of a different type.
type Array struct {
	Value []interface{}
}

// Boolean matches only two special values: true and false. Note that values that evaluate to true or false, such as 1 and 0, are not accepted by the schema.
type Boolean struct {
	Value bool
}

// Integer is used for integral numbers.
type Integer struct {
	Value int64
}

// Null is generally used to represent a missing value.
type Null struct {
	Value json.RawMessage
}

// Number used for any numeric type, either integers or floating point numbers.
type Number struct {
	Value float64
}

// Object is the mapping type in JSON. They map 'keys' to 'values'.
//
// In JSON, the 'keys' must always be strings. Each of these pairs is conventionally referred to as a 'property'.
type Object struct {
	Value interface{}
}

// String is used for strings of text. It may contain Unicode characters.
type String struct {
	Value string
}
