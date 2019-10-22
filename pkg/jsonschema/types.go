// Copyright 2019 The go-jsonschema Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonschema

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
