// Copyright 2019 The go-jsonschema Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonschema

import (
	"regexp"

	"github.com/francoispqt/gojay"
)

const (
	// Draft07SchemaURL contains the JSON Schema draft-07 URL.
	Draft7SchemaURL = "http://json-schema.org/draft-07/schema#"
)

// Schema represents a JSON Schema.
type Schema struct {
	Schema               string                     `json:"$schema"`
	ID                   string                     `json:"$id,omitempty"`
	Title                string                     `json:"title,omitempty"`
	Ref                  string                     `json:"$ref,omitempty"`
	Comment              string                     `json:"$comment,omitempty"`
	Description          string                     `json:"description,omitempty"`
	Default              interface{}                `json:"default,omitempty"`
	ReadOnly             bool                       `json:"readOnly,omitempty"`
	WriteOnly            bool                       `json:"writeOnly,omitempty"`
	Examples             SchemaList                 `json:"examples,omitempty"`
	MultipleOf           float64                    `json:"multipleOf,omitempty"` // exclusiveMinimum is 0
	Maximum              float64                    `json:"maximum,omitempty"`
	ExclusiveMaximum     bool                       `json:"exclusiveMaximum,omitempty"`
	Minimum              float64                    `json:"minimum,omitempty"`
	ExclusiveMinimum     bool                       `json:"exclusiveMinimum,omitempty"`
	MaxLength            int64                      `json:"maxLength,omitempty"` // minimum should be 0
	MinLength            int64                      `json:"minLength,omitempty"` // default should be 0
	Pattern              *regexp.Regexp             `json:"pattern,omitempty"`
	AdditionalItems      *AdditionalItems           `json:"additionalItems,omitempty"`
	Items                *Items                     `json:"items,omitempty"`
	MaxItems             int64                      `json:"maxItems,omitempty"`
	MinItems             int64                      `json:"minItems,omitempty"`
	UniqueItems          bool                       `json:"uniqueItems,omitempty"`
	Contains             *AdditionalProperties      `json:"contains,omitempty"`
	MaxProperties        int64                      `json:"maxProperties,omitempty"`
	MinProperties        int64                      `json:"minProperties,omitempty"`
	Required             StringArray                `json:"required,omitempty"`
	AdditionalProperties *AdditionalProperties      `json:"additionalProperties,omitempty"`
	Definitions          Definitions                `json:"definitions,omitempty"`
	Properties           Properties                 `json:"properties,omitempty"`
	PatternProperties    map[*regexp.Regexp]*Schema `json:"patternProperties,omitempty"`
	Dependencies         *DependencyMap             `json:"dependencies,omitempty"`
	PropertyNames        *Schema                    `json:"propertyNames,omitempty"`
	Const                *Const                     `json:"const,omitempty"`
	Enum                 Enum                       `json:"enum,omitempty"`
	Type                 Type                       `json:"type,omitempty"`
	Format               Format                     `json:"format,omitempty"`
	ContentMediaType     string                     `json:"contentMediaType,omitempty"`
	ContentEncoding      string                     `json:"contentEncoding,omitempty"`
	If                   *Schema                    `json:"if,omitempty"`
	Then                 *Schema                    `json:"then,omitempty"`
	Else                 *Schema                    `json:"else,omitempty"`
	AllOf                SchemaList                 `json:"allOf,omitempty"`
	AnyOf                SchemaList                 `json:"anyOf,omitempty"`
	OneOf                SchemaList                 `json:"oneOf,omitempty"`
	Not                  *Schema                    `json:"not,omitempty"`
}

// Version implements Schema.
func (Schema) Version() DraftVersion { return DraftVersion7 }

var (
	// compile time check whether the Draft7 implements gojay.MarshalerJSONObject interface.
	_ gojay.MarshalerJSONObject = &Schema{}
	// compile time check whether the Draft7 implements gojay.UnmarshalerJSONObject interface.
	_ gojay.UnmarshalerJSONObject = &Schema{}
	// compile time check whether the Draft7 implements Pooler interface.
	_ Pooler = &Schema{}
)

// MarshalJSON implements json.Marshaler.
func (d Schema) MarshalJSON() ([]byte, error) {
	return gojay.MarshalJSONObject(&d)
}

// UnmarshalJSON implements json.Unmarshaler.
func (d *Schema) UnmarshalJSON(data []byte) error {
	return gojay.Unsafe.UnmarshalJSONObject(data, d)
}

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (d *Schema) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keySchema, d.Schema)
	enc.StringKeyOmitEmpty(keyID, d.ID)
	enc.StringKeyOmitEmpty(keyTitle, d.Title)
	enc.StringKeyOmitEmpty(keyRef, d.Ref)
	enc.StringKeyOmitEmpty(keyComment, d.Comment)
	enc.StringKeyOmitEmpty(keyDescription, d.Description)
	// enc.AddInterfaceKeyOmitEmpty(keyDefault, &d.Default)
	enc.BoolKeyOmitEmpty(keyReadOnly, d.ReadOnly)
	enc.BoolKeyOmitEmpty(keyWriteOnly, d.WriteOnly)
	enc.ArrayKeyOmitEmpty(keyExamples, &d.Examples)
	enc.Float64KeyOmitEmpty(keyMultipleOf, d.MultipleOf)
	enc.Float64KeyOmitEmpty(keyMaximum, d.Maximum)
	enc.BoolKeyOmitEmpty(keyExclusiveMaximum, d.ExclusiveMaximum)
	enc.Float64KeyOmitEmpty(keyMinimum, d.Minimum)
	enc.BoolKeyOmitEmpty(keyExclusiveMinimum, d.ExclusiveMinimum)
	enc.Int64KeyOmitEmpty(keyMaxLength, d.MaxLength)
	enc.Int64KeyOmitEmpty(keyMinLength, d.MinLength)
	// enc.AddInterfaceKeyOmitEmpty(keyPattern, &d.Pattern)
	enc.ObjectKeyOmitEmpty(keyAdditionalItems, d.AdditionalItems)
	enc.ObjectKeyOmitEmpty(keyItems, d.Items)
	enc.Int64KeyOmitEmpty(keyMaxItems, d.MaxItems)
	enc.Int64KeyOmitEmpty(keyMinItems, d.MinItems)
	enc.BoolKeyOmitEmpty(keyUniqueItems, d.UniqueItems)
	enc.ObjectKeyOmitEmpty(keyContains, d.Contains)
	enc.Int64KeyOmitEmpty(keyMaxProperties, d.MaxProperties)
	enc.Int64KeyOmitEmpty(keyMinProperties, d.MinProperties)
	enc.ArrayKeyOmitEmpty(keyRequired, &d.Required)
	enc.ObjectKeyOmitEmpty(keyAdditionalProperties, d.AdditionalProperties)
	enc.ObjectKeyOmitEmpty(keyDefinitions, d.Definitions)
	enc.ObjectKeyOmitEmpty(keyProperties, d.Properties)
	// enc.ObjectKeyOmitEmpty(keyPatternProperties, d.PatternProperties)
	enc.ObjectKeyOmitEmpty(keyDependencies, d.Dependencies)
	enc.ObjectKeyOmitEmpty(keyPropertyNames, d.PropertyNames)
	enc.ObjectKeyOmitEmpty(keyConst, d.Const)
	enc.ArrayKeyOmitEmpty(keyEnum, &d.Enum)
	enc.IntKeyOmitEmpty(keyType, *(*int)(&d.Type))
	enc.StringKeyOmitEmpty(keyFormat, *(*string)(&d.Format))
	enc.StringKeyOmitEmpty(keyContentMediaType, d.ContentMediaType)
	enc.StringKeyOmitEmpty(keyContentEncoding, d.ContentEncoding)
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
func (d *Schema) IsNil() bool {
	return d == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (d *Schema) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keySchema:
		// o := StringPool.Get().(*String)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.Schema = *o
		// }
		// return err
		return dec.String(&d.Schema)

	case keyID:
		// o := StringPool.Get().(*String)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.ID = *o
		// }
		// return err
		return dec.String(&d.ID)

	case keyTitle:
		// o := StringPool.Get().(*String)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.Title = *o
		// }
		// return err
		return dec.String(&d.Title)

	case keyRef:
		// o := StringPool.Get().(*String)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.Ref = *o
		// }
		// return err
		return dec.String(&d.Ref)

	case keyComment:
		// o := StringPool.Get().(*String)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.Comment = *o
		// }
		// return err
		return dec.String(&d.Comment)

	case keyDescription:
		// o := StringPool.Get().(*String)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.Description = *o
		// }
		// return err
		return dec.String(&d.Description)

	case keyDefault:
		iface := d.Default.(interface{})
		return dec.Interface(&iface)

	case keyReadOnly:
		// o := BooleanPool.Get().(*Boolean)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.ReadOnly = *o
		// }
		// return err
		return dec.Bool(&d.ReadOnly)

	case keyWriteOnly:
		// o := BooleanPool.Get().(*Boolean)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.WriteOnly = *o
		// }
		// return err
		return dec.Bool(&d.WriteOnly)

	case keyExamples:
		return dec.Array(&d.Examples)

	case keyMultipleOf:
		// o := NumberPool.Get().(*Number)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.MultipleOf = *o
		// }
		// return err
		dec.Float64(&d.MultipleOf)

	case keyMaximum:
		// o := NumberPool.Get().(*Number)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.Maximum = *o
		// }
		// return err
		dec.Float64(&d.Maximum)

	case keyExclusiveMaximum:
		// o := BooleanPool.Get().(*Boolean)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.ExclusiveMaximum = *o
		// }
		// return err
		return dec.Bool(&d.ExclusiveMaximum)

	case keyMinimum:
		// o := NumberPool.Get().(*Number)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.Minimum = *o
		// }
		// return err
		dec.Float64(&d.Minimum)

	case keyExclusiveMinimum:
		// o := BooleanPool.Get().(*Boolean)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.ExclusiveMinimum = *o
		// }
		// return err
		return dec.Bool(&d.ExclusiveMinimum)

	case keyMaxLength:
		// o := IntegerPool.Get().(*Integer)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.MaxLength = *o
		// }
		// return err
		dec.Int64(&d.MaxLength)

	case keyMinLength:
		// o := IntegerPool.Get().(*Integer)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.MinLength = *o
		// }
		// return err
		dec.Int64(&d.MinLength)

	case keyPattern:
		// re := d.Pattern.(regexpinterface.Regexp)
		// return dec.Interface(&re)

	case keyAdditionalItems:
		if d.AdditionalItems == nil {
			d.AdditionalItems = &AdditionalItems{Schema: new(Schema)}
		}
		return dec.Object(d.AdditionalItems)

	case keyItems:
		if d.Items == nil {
			d.Items = &Items{}
		}
		return dec.Object(d.Items)

	case keyMaxItems:
		// o := IntegerPool.Get().(*Integer)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.MaxItems = *o
		// }
		// return err
		dec.Int64(&d.MaxItems)

	case keyMinItems:
		// o := IntegerPool.Get().(*Integer)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.MinItems = *o
		// }
		// return err
		dec.Int64(&d.MinItems)

	case keyUniqueItems:
		// o := BooleanPool.Get().(*Boolean)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.UniqueItems = *o
		// }
		// return err
		return dec.Bool(&d.UniqueItems)

	case keyContains:
		if d.Contains == nil {
			d.Contains = &AdditionalProperties{Schema: new(Schema)}
		}
		return dec.Object(d.Contains)

	case keyMaxProperties:
		// o := IntegerPool.Get().(*Integer)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.MaxProperties = *o
		// }
		// return err
		dec.Int64(&d.MaxProperties)

	case keyMinProperties:
		// o := IntegerPool.Get().(*Integer)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.MinProperties = *o
		// }
		// return err
		dec.Int64(&d.MinProperties)

	case keyRequired:
		// o := StringArrayPool.Get().(*StringArray)
		// err := dec.Array(o)
		// if err == nil {
		// 	d.Required = *o
		// }
		// return err
		dec.Array(&d.Required)

	case keyAdditionalProperties:
		if d.AdditionalProperties == nil {
			d.AdditionalProperties = &AdditionalProperties{Schema: new(Schema)}
		}
		return dec.Object(d.AdditionalProperties)

	case keyDefinitions:
		// return dec.Object(&d.Definitions)

	case keyProperties:
		// if d.Properties == nil {
		// 	d.Properties = make(Properties)
		// }
		// return dec.Object(d.Properties)

	case keyPatternProperties:
		// if d.PatternProperties == nil {
		// 	d.PatternProperties = PatternProperties{
		// 		&lazyregexp.Regexp{}: &Draft7{},
		// 	}
		// }
		// return dec.Object(d.PatternProperties)

	case keyDependencies:
		if d.Dependencies == nil {
			d.Dependencies = &DependencyMap{}
		}
		return dec.Object(d.Dependencies)

	case keyPropertyNames:
		if d.PropertyNames == nil {
			d.PropertyNames = &Schema{}
		}
		o := SchemaPool.Get().(*Schema)
		err := dec.Object(o)
		if err == nil {
			d.PropertyNames = o
		}
		return err

	case keyConst:
		if d.Const == nil {
			d.Const = &Const{}
		}
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
		// o := StringPool.Get().(*String)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.ContentMediaType = *o
		// }
		// return err
		dec.String(&d.ContentMediaType)

	case keyContentEncoding:
		// o := StringPool.Get().(*String)
		// err := dec.Object(o)
		// if err == nil {
		// 	d.ContentEncoding = *o
		// }
		// return err
		dec.String(&d.ContentEncoding)

	case keyIf:
		if d.If == nil {
			d.If = &Schema{}
		}
		o := SchemaPool.Get().(*Schema)
		err := dec.Object(o)
		if err == nil {
			d.If = o
		}
		return err

	case keyThen:
		if d.Then == nil {
			d.Then = &Schema{}
		}
		o := SchemaPool.Get().(*Schema)
		err := dec.Object(o)
		if err == nil {
			d.Then = o
		}
		return err

	case keyElse:
		if d.Else == nil {
			d.Else = &Schema{}
		}
		o := SchemaPool.Get().(*Schema)
		err := dec.Object(o)
		if err == nil {
			d.Else = o
		}
		return err

	case keyAllOf:
		// return dec.Array(d.AllOf)

	case keyAnyOf:
		// return dec.Array(d.AnyOf)

	case keyOneOf:
		// return dec.Array(d.OneOf)

	case keyNot:
		if d.Not == nil {
			d.Not = &Schema{}
		}
		o := SchemaPool.Get().(*Schema)
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
func (*Schema) NKeys() int { return 46 }

// Reset implements Pooler.
//
// Reset reset fields.
func (d *Schema) Reset() {
	// d.Schema.Reset()
	// StringPool.Put(&d.Schema)

	// d.ID.Reset()
	// StringPool.Put(&d.ID)

	// d.Title.Reset()
	// StringPool.Put(&d.Title)

	// d.Ref.Reset()
	// StringPool.Put(&d.Ref)

	// d.Comment.Reset()
	// StringPool.Put(&d.ID)

	// d.Description.Reset()
	// StringPool.Put(&d.Description)

	// d.Default.Reset()

	// d.ReadOnly.Reset()
	// BooleanPool.Put(&d.ReadOnly)

	// d.WriteOnly.Reset()
	// BooleanPool.Put(&d.WriteOnly)

	// d.Examples.Reset()

	// d.MultipleOf.Reset()
	// NumberPool.Put(&d.MultipleOf)

	// d.Maximum.Reset()
	// NumberPool.Put(&d.Maximum)

	// d.ExclusiveMaximum.Reset()
	// BooleanPool.Put(&d.ExclusiveMaximum)

	// d.Minimum.Reset()
	// NumberPool.Put(&d.Minimum)

	// d.ExclusiveMinimum.Reset()
	// BooleanPool.Put(&d.ExclusiveMinimum)

	// d.MaxLength.Reset()
	// IntegerPool.Put(&d.MaxLength)

	// d.MinLength.Reset()
	// IntegerPool.Put(&d.MinLength)

	// d.Pattern = nil

	// d.AdditionalItems.Reset()
	// SchemaPool.Put(&d.AdditionalItems.Schema)

	// d.Items.Reset()
	// ItemsPool.Put(&d.Items)

	// d.MaxItems.Reset()
	// IntegerPool.Put(&d.MaxItems)

	// d.MinItems.Reset()
	// IntegerPool.Put(&d.MinItems)

	// d.UniqueItems.Reset()
	// BooleanPool.Put(&d.UniqueItems)

	// d.Contains.Reset()
	// ConstPool.Put(&d.Contains)

	// d.MaxProperties.Reset()
	// IntegerPool.Put(&d.MaxProperties)

	// d.MinProperties.Reset()
	// IntegerPool.Put(&d.MinProperties)

	// d.Required.Reset()
	// StringArrayPool.Put(&d.Required)

	// d.AdditionalProperties.Reset()
	// SchemaPool.Put(&d.AdditionalProperties.Schema)

	// for i := range d.Definitions {
	// 	d.Definitions[i].Reset()
	// 	SchemaPool.Put(d.Definitions[i])
	// }

	// for i := range d.Properties {
	// 	d.Properties[i].Reset()
	// 	SchemaPool.Put(d.Properties[i])
	// }

	// for i := range d.Properties {
	// 	d.Properties[i].Reset()
	// 	SchemaPool.Put(d.Properties[i])
	// }

	// for i := range d.PatternProperties {
	// 	d.PatternProperties[i].Reset()
	// 	SchemaPool.Put(d.PatternProperties[i])
	// }

	// d.Dependencies.Reset()

	// vpropertyName := d.PropertyNames
	// if vpropertyName != nil {
	// 	vpropertyName.Reset()
	// 	SchemaPool.Put(vpropertyName)
	// }

	// d.Const.Reset()
	// ConstPool.Put(&d.Const)

	// d.Enum.Reset()
	// EnumPool.Put(&d.Enum)

	// d.Type = 0

	// d.Format = ""

	// d.ContentMediaType.Reset()
	// StringPool.Put(&d.ContentMediaType)

	// d.ContentEncoding.Reset()
	// StringPool.Put(&d.ContentEncoding)

	// vif := d.If
	// if vif != nil {
	// 	vif.Reset()
	// 	SchemaPool.Put(vif)
	// }

	// vthen := d.Then
	// if vthen != nil {
	// 	vthen.Reset()
	// 	SchemaPool.Put(vthen)
	// }

	// velse := d.Else
	// if velse != nil {
	// 	velse.Reset()
	// 	SchemaPool.Put(velse)
	// }

	// for i := range d.AllOf {
	// 	d.AllOf[i].Reset()
	// 	SchemaPool.Put(d.AllOf[i])
	// }

	// for i := range d.AnyOf {
	// 	d.AnyOf[i].Reset()
	// 	SchemaPool.Put(d.AnyOf[i])
	// }

	// for i := range d.OneOf {
	// 	d.AllOf[i].Reset()
	// 	SchemaPool.Put(d.AllOf[i])
	// }

	// not := d.Not
	// if not != nil {
	// 	not.Reset()
	// 	SchemaPool.Put(not)
	// }
}

// SchemaStream represents a stream encoding and decoding to Draft7.
type SchemaStream chan *Schema

var (
	// compile time check whether the Draft7Stream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*SchemaStream)(nil)
	// compile time check whether the Draft7Stream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*SchemaStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s SchemaStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Object(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s SchemaStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := SchemaPool.Get().(*Schema)
	if err := dec.Object(o); err != nil {
		return err
	}
	s <- o

	return nil
}

// SchemaList list of Schema.
type SchemaList []*Schema

var (
	// compile time check whether the Draft7List implements gojay.MarshalerJSONArray interface.
	_ gojay.MarshalerJSONArray = &SchemaList{}
	// compile time check whether the Draft7List implements gojay.UnmarshalerJSONArray interface.
	_ gojay.UnmarshalerJSONArray = &SchemaList{}
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (s *SchemaList) MarshalJSONArray(enc *gojay.Encoder) {
	for _, e := range *s {
		enc.Object(e)
	}
}

// IsNil implements gojay.MarshalerJSONArray.
//
// IsNil checks if instance is nil.
func (s *SchemaList) IsNil() bool {
	return len(*s) == 0
}

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (s *SchemaList) UnmarshalJSONArray(dec *gojay.Decoder) error {
	o := SchemaPool.Get().(*Schema)
	if err := dec.Object(o); err != nil {
		return err
	}
	*s = append(*s, o)

	return nil
}

// Draft7ListStream represents a stream encoding and decoding to Draft7List.
type SchemaListStream chan *SchemaList

var (
	// compile time check whether the Draft7Stream implements gojay.MarshalerStream interface.
	_ gojay.MarshalerStream = (*SchemaListStream)(nil)
	// compile time check whether the Draft7Stream implements gojay.UnmarshalerStream interface.
	_ gojay.UnmarshalerStream = (*SchemaListStream)(nil)
)

// MarshalStream implements gojay.MarshalerStream.
func (s SchemaListStream) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case o := <-s:
		enc.Array(o)

	case <-enc.Done():
		return
	}
}

// UnmarshalStream implements gojay.UnmarshalerStream.
func (s SchemaListStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	o := &SchemaList{}
	if err := dec.Array(o); err != nil {
		return err
	}
	s <- o

	return nil
}
