// Copyright 2019 The go-jsonschema Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonschema

const (
	// Draft07SchemaURL contains the JSON Schema draft-07 URL.
	Draft7SchemaURL = "http://json-schema.org/draft-07/schema#"
)

type Draft7 struct {
	Schema               String                `json:"$schema"`
	ID                   String                `json:"$id,omitempty"`
	Title                String                `json:"title,omitempty"`
	Ref                  String                `json:"$ref,omitempty"`
	Comment              String                `json:"$comment,omitempty"`
	Description          String                `json:"description,omitempty"`
	Default              Default               `json:"default,omitempty"`
	ReadOnly             Boolean               `json:"readOnly,omitempty"`
	WriteOnly            Boolean               `json:"writeOnly,omitempty"`
	Examples             Draft7List            `json:"examples,omitempty"`
	MultipleOf           Number                `json:"multipleOf,omitempty"` // exclusiveMinimum is 0
	Maximum              Number                `json:"maximum,omitempty"`
	ExclusiveMaximum     Boolean               `json:"exclusiveMaximum,omitempty"`
	Minimum              Number                `json:"minimum,omitempty"`
	ExclusiveMinimum     Boolean               `json:"exclusiveMinimum,omitempty"`
	MaxLength            Integer               `json:"maxLength,omitempty"` // minimum should be 0
	MinLength            Integer               `json:"minLength,omitempty"` // default should be 0
	Pattern              Pattern               `json:"pattern,omitempty"`
	AdditionalItems      *AdditionalItems      `json:"additionalItems,omitempty"`
	Items                *Items                `json:"items,omitempty"`
	MaxItems             Integer               `json:"maxItems,omitempty"`
	MinItems             Integer               `json:"minItems,omitempty"`
	UniqueItems          Boolean               `json:"uniqueItems,omitempty"`
	Contains             *AdditionalProperties `json:"contains,omitempty"`
	MaxProperties        Integer               `json:"maxProperties,omitempty"`
	MinProperties        Integer               `json:"minProperties,omitempty"`
	Required             StringArray           `json:"required,omitempty"`
	AdditionalProperties *AdditionalProperties `json:"additionalProperties,omitempty"`
	Definitions          Definitions           `json:"definitions,omitempty"`
	Properties           Properties            `json:"properties,omitempty"`
	PatternProperties    PatternProperties     `json:"patternProperties,omitempty"`
	Dependencies         *DependencyMap        `json:"dependencies,omitempty"`
	PropertyNames        *Draft7               `json:"propertyNames,omitempty"`
	Const                *Const                `json:"const,omitempty"`
	Enum                 Enum                  `json:"enum,omitempty"`
	Type                 Type                  `json:"type,omitempty"`
	Format               Format                `json:"format,omitempty"`
	ContentMediaType     String                `json:"contentMediaType,omitempty"`
	ContentEncoding      String                `json:"contentEncoding,omitempty"`
	If                   *Draft7               `json:"if,omitempty"`
	Then                 *Draft7               `json:"then,omitempty"`
	Else                 *Draft7               `json:"else,omitempty"`
	AllOf                Draft7List            `json:"allOf,omitempty"`
	AnyOf                Draft7List            `json:"anyOf,omitempty"`
	OneOf                Draft7List            `json:"oneOf,omitempty"`
	Not                  *Draft7               `json:"not,omitempty"`
}

// compile time check whether the SchemaDraft7 implements Schema interface.
var _ Schema = &Draft7{}

// Version implements Schema.
func (Draft7) Version() DraftVersion { return DraftVersion7 }

type Draft7List []*Draft7
