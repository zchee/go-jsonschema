// Copyright 2019 The go-jsonschema Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonschema

import "github.com/francoispqt/gojay"

// Schema represents a JSON Schema interface.
type Schema interface {
	Version() DraftVersion

	MarshalJSONObject(enc *gojay.Encoder)
	IsNil() bool
	UnmarshalJSONObject(dec *gojay.Decoder, k string) error
	NKeys() int
	Reset()
}

// SchemaList list of Schema.
type SchemaList []Schema
