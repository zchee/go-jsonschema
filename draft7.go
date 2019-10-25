// Copyright 2019 The go-jsonschema Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonschema

import (
	"github.com/francoispqt/gojay"
)

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

// MarshalJSON implements json.Marshaler.
func (d Draft7) MarshalJSON() ([]byte, error) {
	return gojay.MarshalJSONObject(&d)
}

// UnmarshalJSON implements json.Unmarshaler.
func (d *Draft7) UnmarshalJSON(data []byte) error {
	return gojay.Unsafe.UnmarshalJSONObject(data, d)
}

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (d *Draft7) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keySchema, &d.Schema)
	enc.ObjectKeyOmitEmpty(keyID, &d.ID)
	enc.ObjectKeyOmitEmpty(keyTitle, &d.Title)
	enc.ObjectKeyOmitEmpty(keyRef, &d.Ref)
	enc.ObjectKeyOmitEmpty(keyComment, &d.Comment)
	enc.ObjectKeyOmitEmpty(keyDescription, &d.Description)
	enc.AddInterfaceKeyOmitEmpty(keyDefault, &d.Default)
	enc.ObjectKeyOmitEmpty(keyReadOnly, &d.ReadOnly)
	enc.ObjectKeyOmitEmpty(keyWriteOnly, &d.WriteOnly)
	enc.ArrayKeyOmitEmpty(keyExamples, &d.Examples)
	enc.ObjectKeyOmitEmpty(keyMultipleOf, &d.MultipleOf)
	enc.ObjectKeyOmitEmpty(keyMaximum, &d.Maximum)
	enc.ObjectKeyOmitEmpty(keyExclusiveMaximum, &d.ExclusiveMaximum)
	enc.ObjectKeyOmitEmpty(keyMinimum, &d.Minimum)
	enc.ObjectKeyOmitEmpty(keyExclusiveMinimum, &d.ExclusiveMinimum)
	enc.ObjectKeyOmitEmpty(keyMaxLength, &d.MaxLength)
	enc.ObjectKeyOmitEmpty(keyMinLength, &d.MinLength)
	enc.AddInterfaceKeyOmitEmpty(keyPattern, &d.Pattern)
	enc.ObjectKeyOmitEmpty(keyAdditionalItems, d.AdditionalItems)
	enc.ObjectKeyOmitEmpty(keyItems, d.Items)
	enc.ObjectKeyOmitEmpty(keyMaxItems, &d.MaxItems)
	enc.ObjectKeyOmitEmpty(keyMinItems, &d.MinItems)
	enc.ObjectKeyOmitEmpty(keyUniqueItems, &d.UniqueItems)
	enc.ObjectKeyOmitEmpty(keyContains, d.Contains)
	enc.ObjectKeyOmitEmpty(keyMaxProperties, &d.MaxProperties)
	enc.ObjectKeyOmitEmpty(keyMinProperties, &d.MinProperties)
	enc.ArrayKeyOmitEmpty(keyRequired, &d.Required)
	enc.ObjectKeyOmitEmpty(keyAdditionalProperties, d.AdditionalProperties)
	enc.ObjectKeyOmitEmpty(keyDefinitions, &d.Definitions)
	enc.ObjectKeyOmitEmpty(keyProperties, &d.Properties)
	enc.ObjectKeyOmitEmpty(keyPatternProperties, &d.PatternProperties)
	enc.ObjectKeyOmitEmpty(keyDependencies, d.Dependencies)
	enc.ObjectKeyOmitEmpty(keyPropertyNames, d.PropertyNames)
	enc.ObjectKeyOmitEmpty(keyConst, d.Const)
	enc.ArrayKeyOmitEmpty(keyEnum, &d.Enum)
	enc.IntKeyOmitEmpty(keyType, *(*int)(&d.Type))
	enc.StringKeyOmitEmpty(keyFormat, *(*string)(&d.Format))
	enc.ObjectKeyOmitEmpty(keyContentMediaType, &d.ContentMediaType)
	enc.ObjectKeyOmitEmpty(keyContentEncoding, &d.ContentEncoding)
	enc.ObjectKeyOmitEmpty(keyIf, d.If)
	enc.ObjectKeyOmitEmpty(keyThen, d.Then)
	enc.ObjectKeyOmitEmpty(keyElse, d.Else)
	enc.ArrayKeyOmitEmpty(keyAllOf, &d.AllOf)
	enc.ArrayKeyOmitEmpty(keyAnyOf, &d.AnyOf)
	enc.ArrayKeyOmitEmpty(keyOneOf, &d.OneOf)
	enc.ObjectKeyOmitEmpty(keyNot, d.Not)
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
	case keySchema:
		o := StringPool.Get().(*String)
		err := dec.Object(o)
		if err == nil {
			d.Schema = *o
		}
		return err

	case keyID:
		o := StringPool.Get().(*String)
		err := dec.Object(o)
		if err == nil {
			d.ID = *o
		}
		return err

	case keyTitle:
		o := StringPool.Get().(*String)
		err := dec.Object(o)
		if err == nil {
			d.Title = *o
		}
		return err

	case keyRef:
		o := StringPool.Get().(*String)
		err := dec.Object(o)
		if err == nil {
			d.Ref = *o
		}
		return err

	case keyComment:
		o := StringPool.Get().(*String)
		err := dec.Object(o)
		if err == nil {
			d.Comment = *o
		}
		return err

	case keyDescription:
		o := StringPool.Get().(*String)
		err := dec.Object(o)
		if err == nil {
			d.Description = *o
		}
		return err

	case keyDefault:
		iface := d.Default.(interface{})
		return dec.Interface(&iface)

	case keyReadOnly:
		o := BooleanPool.Get().(*Boolean)
		err := dec.Object(o)
		if err == nil {
			d.ReadOnly = *o
		}
		return err

	case keyWriteOnly:
		o := BooleanPool.Get().(*Boolean)
		err := dec.Object(o)
		if err == nil {
			d.WriteOnly = *o
		}
		return err

	case keyExamples:
		return dec.Array(&d.Examples)

	case keyMultipleOf:
		o := NumberPool.Get().(*Number)
		err := dec.Object(o)
		if err == nil {
			d.MultipleOf = *o
		}
		return err

	case keyMaximum:
		o := NumberPool.Get().(*Number)
		err := dec.Object(o)
		if err == nil {
			d.Maximum = *o
		}
		return err

	case keyExclusiveMaximum:
		o := BooleanPool.Get().(*Boolean)
		err := dec.Object(o)
		if err == nil {
			d.ExclusiveMaximum = *o
		}
		return err

	case keyMinimum:
		o := NumberPool.Get().(*Number)
		err := dec.Object(o)
		if err == nil {
			d.Minimum = *o
		}
		return err

	case keyExclusiveMinimum:
		o := BooleanPool.Get().(*Boolean)
		err := dec.Object(o)
		if err == nil {
			d.ExclusiveMinimum = *o
		}
		return err

	case keyMaxLength:
		o := IntegerPool.Get().(*Integer)
		err := dec.Object(o)
		if err == nil {
			d.MaxLength = *o
		}
		return err

	case keyMinLength:
		o := IntegerPool.Get().(*Integer)
		err := dec.Object(o)
		if err == nil {
			d.MinLength = *o
		}
		return err

	case keyPattern:
		re := d.Pattern.(interface{})
		return dec.Interface(&re)

	case keyAdditionalItems:
		return dec.Object(d.AdditionalItems)

	case keyItems:
		return dec.Object(d.Items)

	case keyMaxItems:
		o := IntegerPool.Get().(*Integer)
		err := dec.Object(o)
		if err == nil {
			d.MaxItems = *o
		}
		return err

	case keyMinItems:
		o := IntegerPool.Get().(*Integer)
		err := dec.Object(o)
		if err == nil {
			d.MinItems = *o
		}
		return err

	case keyUniqueItems:
		o := BooleanPool.Get().(*Boolean)
		err := dec.Object(o)
		if err == nil {
			d.UniqueItems = *o
		}
		return err

	case keyContains:
		return dec.Object(d.Contains)

	case keyMaxProperties:
		o := IntegerPool.Get().(*Integer)
		err := dec.Object(o)
		if err == nil {
			d.MaxProperties = *o
		}
		return err

	case keyMinProperties:
		o := IntegerPool.Get().(*Integer)
		err := dec.Object(o)
		if err == nil {
			d.MinProperties = *o
		}
		return err

	case keyRequired:
		o := StringArrayPool.Get().(*StringArray)
		err := dec.Array(o)
		if err == nil {
			d.Required = *o
		}
		return err

	case keyAdditionalProperties:
		return dec.Object(d.AdditionalProperties)

	case keyDefinitions:
		return dec.Object(&d.Definitions)

	case keyProperties:
		return dec.Object(&d.Properties)

	case keyPatternProperties:
		return dec.Object(&d.PatternProperties)

	case keyDependencies:
		return dec.Object(d.Dependencies)

	case keyPropertyNames:
		o := Draft7Pool.Get().(*Draft7)
		err := dec.Object(o)
		if err == nil {
			d.PropertyNames = o
		}
		return err

	case keyConst:
		o := ConstPool.Get().(*Const)
		err := dec.Object(o)
		if err == nil {
			d.Const = o
		}
		return err

	case keyEnum:
		o := EnumPool.Get().(*Enum)
		err := dec.Array(o)
		if err == nil {
			d.Enum = *o
		}
		return err

	case keyType:
		t := int(d.Type)
		return dec.Int(&t)

	case keyFormat:
		f := string(d.Format)
		return dec.String(&f)

	case keyContentMediaType:
		o := StringPool.Get().(*String)
		err := dec.Object(o)
		if err == nil {
			d.ContentMediaType = *o
		}
		return err

	case keyContentEncoding:
		o := StringPool.Get().(*String)
		err := dec.Object(o)
		if err == nil {
			d.ContentEncoding = *o
		}
		return err

	case keyIf:
		o := Draft7Pool.Get().(*Draft7)
		err := dec.Object(o)
		if err == nil {
			d.If = o
		}
		return err

	case keyThen:
		o := Draft7Pool.Get().(*Draft7)
		err := dec.Object(o)
		if err == nil {
			d.Then = o
		}
		return err

	case keyElse:
		o := Draft7Pool.Get().(*Draft7)
		err := dec.Object(o)
		if err == nil {
			d.Else = o
		}
		return err

	case keyAllOf:
		return dec.Array(&d.AllOf)

	case keyAnyOf:
		return dec.Array(&d.AnyOf)

	case keyOneOf:
		return dec.Array(&d.OneOf)

	case keyNot:
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
