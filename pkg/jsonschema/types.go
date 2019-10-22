// Copyright 2019 The go-jsonschema Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonschema

import (
	"regexp"
)

type StringArray []String

type Default interface{}

type Items struct {
	Schemas     SchemaList
	HasMultiple bool
}

type Pattern regexp.Regexp

type AdditionalItems struct {
	Schema
}

type AdditionalProperties struct {
	Schema
}

type Definitions map[string]Schema

type Properties map[string]Schema

type PatternProperties map[*regexp.Regexp]Schema

type DependencyMap struct {
	Names   map[string][]string
	Schemas map[string]Schema
}

type Const struct {
	Type  Type
	Value []interface{}
}

type Enum []*Const
