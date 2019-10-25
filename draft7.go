// Copyright 2019 The go-jsonschema Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonschema

import "github.com/francoispqt/gojay"

const (
	// Draft07SchemaURL contains the JSON Schema draft-07 URL.
	Draft7SchemaURL = "http://json-schema.org/draft-07/schema#"
)

// Draft7 represents a JSON Schema Draft 7.
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

var (
	// compile time check whether the Draft7 implements gojay.MarshalerJSONObject interface.
	_ gojay.MarshalerJSONObject = &Draft7{}
	// compile time check whether the Draft7 implements gojay.UnmarshalerJSONObject interface.
	_ gojay.UnmarshalerJSONObject = &Draft7{}
	// compile time check whether the Draft7 implements Pooler interface.
	_ Pooler = &Draft7{}
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (d *Draft7) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("$schema", &d.Schema)
	enc.ObjectKeyOmitEmpty("$id", &d.ID)
	enc.ObjectKeyOmitEmpty("title", &d.Title)
	enc.ObjectKeyOmitEmpty("$ref", &d.Ref)
	enc.ObjectKeyOmitEmpty("$comment", &d.Comment)
	enc.ObjectKeyOmitEmpty("description", &d.Description)
	enc.AddInterfaceKeyOmitEmpty("default", &d.Default)
	enc.ObjectKeyOmitEmpty("readOnly", &d.ReadOnly)
	enc.ObjectKeyOmitEmpty("writeOnly", &d.WriteOnly)
	enc.ArrayKeyOmitEmpty("examples", &d.Examples)
	enc.ObjectKeyOmitEmpty("multipleOf", &d.MultipleOf)
	enc.ObjectKeyOmitEmpty("maximum", &d.Maximum)
	enc.ObjectKeyOmitEmpty("exclusiveMaximum", &d.ExclusiveMaximum)
	enc.ObjectKeyOmitEmpty("minimum", &d.Minimum)
	enc.ObjectKeyOmitEmpty("exclusiveMinimum", &d.ExclusiveMinimum)
	enc.ObjectKeyOmitEmpty("maxLength", &d.MaxLength)
	enc.ObjectKeyOmitEmpty("minLength", &d.MinLength)
	enc.AddInterfaceKeyOmitEmpty("pattern", &d.Pattern)
	enc.ObjectKeyOmitEmpty("additionalItems", d.AdditionalItems)
	enc.ObjectKeyOmitEmpty("items", d.Items)
	enc.ObjectKeyOmitEmpty("maxItems", &d.MaxItems)
	enc.ObjectKeyOmitEmpty("minItems", &d.MinItems)
	enc.ObjectKeyOmitEmpty("uniqueItems", &d.UniqueItems)
	enc.ObjectKeyOmitEmpty("contains", d.Contains)
	enc.ObjectKeyOmitEmpty("maxProperties", &d.MaxProperties)
	enc.ObjectKeyOmitEmpty("minProperties", &d.MinProperties)
	enc.ArrayKeyOmitEmpty("required", &d.Required)
	enc.ObjectKeyOmitEmpty("additionalProperties", d.AdditionalProperties)
	enc.ObjectKeyOmitEmpty("definitions", &d.Definitions)
	enc.ObjectKeyOmitEmpty("properties", &d.Properties)
	enc.ObjectKeyOmitEmpty("patternProperties", &d.PatternProperties)
	enc.ObjectKeyOmitEmpty("dependencies", d.Dependencies)
	enc.ObjectKeyOmitEmpty("propertyNames", d.PropertyNames)
	enc.ObjectKeyOmitEmpty("const", d.Const)
	enc.ArrayKeyOmitEmpty("enum", &d.Enum)
	enc.IntKeyOmitEmpty("type", *(*int)(&d.Type))
	enc.StringKeyOmitEmpty("format", *(*string)(&d.Format))
	enc.ObjectKeyOmitEmpty("contentMediaType", &d.ContentMediaType)
	enc.ObjectKeyOmitEmpty("contentEncoding", &d.ContentEncoding)
	enc.ObjectKeyOmitEmpty("if", d.If)
	enc.ObjectKeyOmitEmpty("then", d.Then)
	enc.ObjectKeyOmitEmpty("else", d.Else)
	enc.ArrayKeyOmitEmpty("allOf", &d.AllOf)
	enc.ArrayKeyOmitEmpty("anyOf", &d.AnyOf)
	enc.ArrayKeyOmitEmpty("oneOf", &d.OneOf)
	enc.ObjectKeyOmitEmpty("not", d.Not)
}

// IsNil implements gojay.MarshalerJSONObject.
//
// IsNil checks if instance is nil.
func (d *Draft7) IsNil() bool {
	return d == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (d *Draft7) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "$schema":
		o := StringPool.Get().(*String)
		err := dec.Object(o)
		if err == nil {
			d.Schema = *o
		}
		return err

	case "$id":
		o := StringPool.Get().(*String)
		err := dec.Object(o)
		if err == nil {
			d.ID = *o
		}
		return err

	case "title":
		o := StringPool.Get().(*String)
		err := dec.Object(o)
		if err == nil {
			d.Title = *o
		}
		return err

	case "$ref":
		o := StringPool.Get().(*String)
		err := dec.Object(o)
		if err == nil {
			d.Ref = *o
		}
		return err

	case "$comment":
		o := StringPool.Get().(*String)
		err := dec.Object(o)
		if err == nil {
			d.Comment = *o
		}
		return err

	case "description":
		o := StringPool.Get().(*String)
		err := dec.Object(o)
		if err == nil {
			d.Description = *o
		}
		return err

	case "default":
		iface := d.Default.(interface{})
		return dec.Interface(&iface)

	case "readOnly":
		o := BooleanPool.Get().(*Boolean)
		err := dec.Object(o)
		if err == nil {
			d.ReadOnly = *o
		}
		return err

	case "writeOnly":
		o := BooleanPool.Get().(*Boolean)
		err := dec.Object(o)
		if err == nil {
			d.WriteOnly = *o
		}
		return err

	case "examples":
		return dec.Array(&d.Examples)

	case "multipleOf":
		o := NumberPool.Get().(*Number)
		err := dec.Object(o)
		if err == nil {
			d.MultipleOf = *o
		}
		return err

	case "maximum":
		o := NumberPool.Get().(*Number)
		err := dec.Object(o)
		if err == nil {
			d.Maximum = *o
		}
		return err

	case "exclusiveMaximum":
		o := BooleanPool.Get().(*Boolean)
		err := dec.Object(o)
		if err == nil {
			d.ExclusiveMaximum = *o
		}
		return err

	case "minimum":
		o := NumberPool.Get().(*Number)
		err := dec.Object(o)
		if err == nil {
			d.Minimum = *o
		}
		return err

	case "exclusiveMinimum":
		o := BooleanPool.Get().(*Boolean)
		err := dec.Object(o)
		if err == nil {
			d.ExclusiveMinimum = *o
		}
		return err

	case "maxLength":
		o := IntegerPool.Get().(*Integer)
		err := dec.Object(o)
		if err == nil {
			d.MaxLength = *o
		}
		return err

	case "minLength":
		o := IntegerPool.Get().(*Integer)
		err := dec.Object(o)
		if err == nil {
			d.MinLength = *o
		}
		return err

	case "pattern":
		re := d.Pattern.(interface{})
		return dec.Interface(&re)

	case "additionalItems":
		return dec.Object(d.AdditionalItems)

	case "items":
		return dec.Object(d.Items)

	case "maxItems":
		o := IntegerPool.Get().(*Integer)
		err := dec.Object(o)
		if err == nil {
			d.MaxItems = *o
		}
		return err

	case "minItems":
		o := IntegerPool.Get().(*Integer)
		err := dec.Object(o)
		if err == nil {
			d.MinItems = *o
		}
		return err

	case "uniqueItems":
		o := BooleanPool.Get().(*Boolean)
		err := dec.Object(o)
		if err == nil {
			d.UniqueItems = *o
		}
		return err

	case "contains":
		return dec.Object(d.Contains)

	case "maxProperties":
		o := IntegerPool.Get().(*Integer)
		err := dec.Object(o)
		if err == nil {
			d.MaxProperties = *o
		}
		return err

	case "minProperties":
		o := IntegerPool.Get().(*Integer)
		err := dec.Object(o)
		if err == nil {
			d.MinProperties = *o
		}
		return err

	case "required":
		o := StringArrayPool.Get().(*StringArray)
		err := dec.Array(o)
		if err == nil {
			d.Required = *o
		}
		return err

	case "additionalProperties":
		return dec.Object(d.AdditionalProperties)

	case "definitions":
		return dec.Object(&d.Definitions)

	case "properties":
		return dec.Object(&d.Properties)

	case "patternProperties":
		return dec.Object(&d.PatternProperties)

	case "dependencies":
		return dec.Object(d.Dependencies)

	case "propertyNames":
		o := Draft7Pool.Get().(*Draft7)
		err := dec.Object(o)
		if err == nil {
			d.PropertyNames = o
		}
		return err

	case "const":
		o := ConstPool.Get().(*Const)
		err := dec.Object(o)
		if err == nil {
			d.Const = o
		}
		return err

	case "enum":
		o := EnumPool.Get().(*Enum)
		err := dec.Array(o)
		if err == nil {
			d.Enum = *o
		}
		return err

	case "type":
		t := int(d.Type)
		return dec.Int(&t)

	case "format":
		f := string(d.Format)
		return dec.String(&f)

	case "contentMediaType":
		o := StringPool.Get().(*String)
		err := dec.Object(o)
		if err == nil {
			d.ContentMediaType = *o
		}
		return err

	case "contentEncoding":
		o := StringPool.Get().(*String)
		err := dec.Object(o)
		if err == nil {
			d.ContentEncoding = *o
		}
		return err

	case "if":
		o := Draft7Pool.Get().(*Draft7)
		err := dec.Object(o)
		if err == nil {
			d.If = o
		}
		return err

	case "then":
		o := Draft7Pool.Get().(*Draft7)
		err := dec.Object(o)
		if err == nil {
			d.Then = o
		}
		return err

	case "else":
		o := Draft7Pool.Get().(*Draft7)
		err := dec.Object(o)
		if err == nil {
			d.Else = o
		}
		return err

	case "allOf":
		return dec.Array(&d.AllOf)

	case "anyOf":
		return dec.Array(&d.AnyOf)

	case "oneOf":
		return dec.Array(&d.OneOf)

	case "not":
		o := Draft7Pool.Get().(*Draft7)
		err := dec.Object(o)
		if err == nil {
			d.Not = o
		}
		return err
	}

	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
//
// NKeys returns the number of keys to unmarshal.
func (*Draft7) NKeys() int { return 46 }

// Reset implements Pooler.
//
// Reset reset fields.
func (d *Draft7) Reset() {
	d.Schema.Reset()
	StringPool.Put(&d.Schema)

	d.ID.Reset()
	StringPool.Put(&d.ID)

	d.Title.Reset()
	StringPool.Put(&d.Title)

	d.Ref.Reset()
	StringPool.Put(&d.Ref)

	d.Comment.Reset()
	StringPool.Put(&d.ID)

	d.Description.Reset()
	StringPool.Put(&d.Description)

	// d.Default.Reset()

	d.ReadOnly.Reset()
	BooleanPool.Put(&d.ReadOnly)

	d.WriteOnly.Reset()
	BooleanPool.Put(&d.WriteOnly)

	// d.Examples.Reset()

	d.MultipleOf.Reset()
	NumberPool.Put(&d.MultipleOf)

	d.Maximum.Reset()
	NumberPool.Put(&d.Maximum)

	d.ExclusiveMaximum.Reset()
	BooleanPool.Put(&d.ExclusiveMaximum)

	d.Minimum.Reset()
	NumberPool.Put(&d.Minimum)

	d.ExclusiveMinimum.Reset()
	BooleanPool.Put(&d.ExclusiveMinimum)

	d.MaxLength.Reset()
	IntegerPool.Put(&d.MaxLength)

	d.MinLength.Reset()
	IntegerPool.Put(&d.MinLength)

	d.Pattern = nil

	d.AdditionalItems.Reset()
	Draft7Pool.Put(&d.AdditionalItems.Schema)

	d.Items.Reset()
	ItemsPool.Put(&d.Items)

	d.MaxItems.Reset()
	IntegerPool.Put(&d.MaxItems)

	d.MinItems.Reset()
	IntegerPool.Put(&d.MinItems)

	d.UniqueItems.Reset()
	BooleanPool.Put(&d.UniqueItems)

	d.Contains.Reset()
	ConstPool.Put(&d.Contains)

	d.MaxProperties.Reset()
	IntegerPool.Put(&d.MaxProperties)

	d.MinProperties.Reset()
	IntegerPool.Put(&d.MinProperties)

	d.Required.Reset()
	StringArrayPool.Put(&d.Required)

	d.AdditionalProperties.Reset()
	Draft7Pool.Put(&d.AdditionalProperties.Schema)

	for i := range d.Definitions {
		d.Definitions[i].Reset()
		Draft7Pool.Put(d.Definitions[i])
	}

	for i := range d.Properties {
		d.Properties[i].Reset()
		Draft7Pool.Put(d.Properties[i])
	}

	for i := range d.Properties {
		d.Properties[i].Reset()
		Draft7Pool.Put(d.Properties[i])
	}

	for i := range d.PatternProperties {
		d.PatternProperties[i].Reset()
		Draft7Pool.Put(d.PatternProperties[i])
	}

	d.Dependencies.Reset()

	vpropertyName := d.PropertyNames
	if vpropertyName != nil {
		vpropertyName.Reset()
		Draft7Pool.Put(vpropertyName)
	}

	d.Const.Reset()
	ConstPool.Put(&d.Const)

	d.Enum.Reset()
	EnumPool.Put(&d.Enum)

	d.Type = 0

	d.Format = ""

	d.ContentMediaType.Reset()
	StringPool.Put(&d.ContentMediaType)

	d.ContentEncoding.Reset()
	StringPool.Put(&d.ContentEncoding)

	vif := d.If
	if vif != nil {
		vif.Reset()
		Draft7Pool.Put(vif)
	}

	vthen := d.Then
	if vthen != nil {
		vthen.Reset()
		Draft7Pool.Put(vthen)
	}

	velse := d.Else
	if velse != nil {
		velse.Reset()
		Draft7Pool.Put(velse)
	}

	for i := range d.AllOf {
		d.AllOf[i].Reset()
		Draft7Pool.Put(d.AllOf[i])
	}

	for i := range d.AnyOf {
		d.AnyOf[i].Reset()
		Draft7Pool.Put(d.AnyOf[i])
	}

	for i := range d.OneOf {
		d.AllOf[i].Reset()
		Draft7Pool.Put(d.OneOf[i])
	}

	not := d.Not
	if not != nil {
		not.Reset()
		Draft7Pool.Put(not)
	}
}

// Draft7Stream represents a stream encoding and decoding to Draft7.
type Draft7Stream chan *Draft7

var (
	// compile time check whether the Draft7Stream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*Draft7Stream)(nil)
	// compile time check whether the Draft7Stream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*Draft7Stream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s Draft7Stream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Object(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s Draft7Stream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := Draft7Pool.Get().(*Draft7)
	if err := dec.Object(o); err != nil {
		return err
	}
	s <- o

	return nil
}

// Draft7List list of Draft7.
type Draft7List []*Draft7

var (
	// compile time check whether the Draft7List implements gojay.MarshalerJSONArray interface.
	_ gojay.MarshalerJSONArray = &Draft7List{}
	// compile time check whether the Draft7List implements gojay.UnmarshalerJSONArray interface.
	_ gojay.UnmarshalerJSONArray = &Draft7List{}
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (s *Draft7List) MarshalJSONArray(enc *gojay.Encoder) {
	for _, e := range *s {
		enc.Object(e)
	}
}

// IsNil implements gojay.MarshalerJSONArray.
//
// IsNil checks if instance is nil.
func (s *Draft7List) IsNil() bool {
	return len(*s) == 0
}

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (s *Draft7List) UnmarshalJSONArray(dec *gojay.Decoder) error {
	o := Draft7Pool.Get().(*Draft7)
	if err := dec.Object(o); err != nil {
		return err
	}
	*s = append(*s, o)

	return nil
}

// Draft7ListStream represents a stream encoding and decoding to Draft7List.
type Draft7ListStream chan *Draft7List

var (
	// compile time check whether the Draft7Stream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*Draft7Stream)(nil)
	// compile time check whether the Draft7Stream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*Draft7Stream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s Draft7ListStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Array(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s Draft7ListStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := &Draft7List{}
	if err := dec.Array(o); err != nil {
		return err
	}
	s <- o

	return nil
}
