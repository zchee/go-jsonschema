// Copyright 2019 The go-jsonschema Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonschema

// DraftVersion represents a Draft version.
type DraftVersion string

// The list of Draft version.
const (
	DraftVersion4      DraftVersion = "draft04"
	DraftVersion6      DraftVersion = "draft06"
	DraftVersion7      DraftVersion = "draft07"
	DraftVersion201909 DraftVersion = "2019-09"
)

// MediaType represents a media type used for a JSON Schema.
//
// 4.3. JSON Schema Documents
//  https://tools.ietf.org/html/draft-handrews-json-schema-01#section-4.3
const MediaType = "application/schema+json"

// Type represents a JSON Schema primitive type.
type Type int

// The list of simple types.
const (
	UnspecifiedType Type = iota
	ArrayType
	BooleanType
	IntegerType
	NullType
	NumberType
	ObjectType
	StringType
)

// Types represents a list of Type.
type Types []Type

// Format allows for basic semantic validation on certain kinds of string values that are commonly used.
// This allows values to be constrained beyond what the other tools in JSON Schema, including Regular Expressions can do.
type Format string

// The following is the list of built-in formats specified in the JSON Schema specification.
const (
	// FormatDateTime is the Date and time together. for example, 2018-11-13T20:20:39+00:00.
	FormatDateTime Format = "date-time"

	// FormatTime is the Time. for example, 20:20:39+00:00.
	//
	// New in draft 7.
	FormatTime Format = "time"

	// FormatDate is the Date. for example, 2018-11-13.
	//
	// New in draft 7.
	FormatDate Format = "date"

	// FormatEmail Internet email address, see RFC 5322, section 3.4.1.
	//
	// RFC 5322, section 3.4.1:
	//   http://tools.ietf.org/html/rfc5322#section-3.4.1
	FormatEmail Format = "email"

	// FormatIDNEmail the internationalized form of an Internet email address, see RFC 6531.
	//
	// RFC 6531:
	//   https://tools.ietf.org/html/rfc6531
	//
	// New in draft 7.
	FormatIDNEmail Format = "idn-email"

	// FormatHostname Internet host name, see RFC 1034, section 3.1.
	//
	// RFC 1034, section 3.1:
	//   http://tools.ietf.org/html/rfc1034#section-3.1
	FormatHostname Format = "hostname"

	// FormatIDNHostname an internationalized Internet host name, see RFC5890, section 2.3.2.3.
	//
	// RFC5890, section 2.3.2.3:
	//   https://tools.ietf.org/html/rfc5890#section-2.3.2.3
	//
	// New in draft 7.
	FormatIDNHostname Format = "idn-hostname"

	// FormatIPv4 IPv4 address, according to dotted-quad ABNF syntax as defined in RFC 2673, section 3.2.
	//
	// RFC 2673, section 3.2:
	//   http://tools.ietf.org/html/rfc2673#section-3.2
	FormatIPv4 Format = "ipv4"

	// FormatIPv6 IPv6 address, as defined in RFC 2373, section 2.2.
	//
	// RFC 2373, section 2.2:
	//   http://tools.ietf.org/html/rfc2373#section-2.2
	FormatIPv6 Format = "ipv6"

	// FormatURI a universal resource identifier (URI), according to RFC3986.
	//
	// Draft 4 only includes "uri", not "uri-reference". Therefore, there is some ambiguity around whether "uri" should accept relative paths.
	//
	// RFC3986:
	//   http://tools.ietf.org/html/rfc3986
	FormatURI Format = "uri"

	// FormatURIReference a URI Reference (either a URI or a relative-reference), according to RFC3986, section 4.1.
	//
	// RFC3986, section 4.1:
	//   http://tools.ietf.org/html/rfc3986#section-4.1
	//
	// New in draft 6.
	FormatURIReference Format = "uri-reference"

	// FormatIRI the internationalized equivalent of a “uri”, according to RFC3987.
	//
	// RFC3987:
	//   https://tools.ietf.org/html/rfc3987
	//
	// New in draft 7.
	FormatIRI Format = "iri"

	// FormatIRIReference the internationalized equivalent of a “uri-reference”, according to RFC3987.
	//
	// RFC3987:
	//   https://tools.ietf.org/html/rfc3987
	//
	// New in draft 7.
	FormatIRIReference Format = "iri-reference"

	// FormatURITemplate a URI Template (of any level) according to RFC6570. If you don’t already know what a URI Template is, you probably don’t need this value.
	//
	// RFC6570:
	//   https://tools.ietf.org/html/rfc6570
	//
	// New in draft 6.
	FormatURITemplate Format = "uri-template"

	// FormatJSONPointer a JSON Pointer, according to RFC6901.
	// There is more discussion on the use of JSON Pointer within JSON Schema in Structuring a complex schema.
	//
	// Note that this should be used only when the entire string contains only JSON Pointer content, e.g. /foo/bar.
	// JSON Pointer URI fragments, e.g. #/foo/bar/ should use "uri-reference".
	//
	// RFC6901:
	//   https://tools.ietf.org/html/rfc6901
	//
	// New in draft 6.
	FormatJSONPointer Format = "json-pointer"

	// FormatRelativeJSONPointer a relative JSON pointer.
	//
	// New in draft 7.
	FormatRelativeJSONPointer Format = "relative-json-pointer"

	// FormatRegex a regular expression, which should be valid according to the ECMA 262 dialect.
	//
	// ECMA 262:
	//   http://www.ecma-international.org/publications/files/ECMA-ST/Ecma-262.pdf
	//
	// New in draft 7.
	FormatRegex Format = "regex"
)
