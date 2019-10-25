// Copyright 2019 The go-jsonschema Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonschema

const (
	keySchema               = "$schema"
	keyID                   = "$id"
	keyTitle                = "title"
	keyRef                  = "$ref"
	keyComment              = "$comment"
	keyDescription          = "description"
	keyDefault              = "default"
	keyReadOnly             = "readOnly"
	keyWriteOnly            = "writeOnly"
	keyExamples             = "examples"
	keyMultipleOf           = "multipleOf"
	keyMaximum              = "maximum"
	keyExclusiveMaximum     = "exclusiveMaximum"
	keyMinimum              = "minimum"
	keyExclusiveMinimum     = "exclusiveMinimum"
	keyMaxLength            = "maxLength"
	keyMinLength            = "minLength"
	keyPattern              = "pattern"
	keyAdditionalItems      = "additionalItems"
	keyItems                = "items"
	keyMaxItems             = "maxItems"
	keyMinItems             = "minItems"
	keyUniqueItems          = "uniqueItems"
	keyContains             = "contains"
	keyMaxProperties        = "maxProperties"
	keyMinProperties        = "minProperties"
	keyRequired             = "required"
	keyAdditionalProperties = "additionalProperties"
	keyDefinitions          = "definitions"
	keyProperties           = "properties"
	keyPatternProperties    = "patternProperties"
	keyDependencies         = "dependencies"
	keyPropertyNames        = "propertyNames"
	keyConst                = "const"
	keyEnum                 = "enum"
	keyType                 = "type"
	keyFormat               = "format"
	keyContentMediaType     = "contentMediaType"
	keyContentEncoding      = "contentEncoding"
	keyIf                   = "if"
	keyThen                 = "then"
	keyElse                 = "else"
	keyAllOf                = "allOf"
	keyAnyOf                = "anyOf"
	keyOneOf                = "oneOf"
	keyNot                  = "not"
)

const (
	keyValue       = "Value"
	keyInitialized = "Initialized"
)

const (
	keySchemas     = "Schemas"
	keyHasMultiple = "HasMultiple"
	keyNames       = "Names"
)
