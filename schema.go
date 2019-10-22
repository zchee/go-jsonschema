// Copyright 2019 The go-jsonschema Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonschema

type Schema interface {
	Version() DraftVersion
}

type SchemaList []Schema
