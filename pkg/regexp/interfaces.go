// Copyright 2019 The go-jsonschema Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package regexp

// Regexp represents a regexp interface.
type Regexp interface {
	String() string
	Copy() Regexp
	FindSubmatch(s []byte) [][]byte
	FindStringSubmatch(s string) []string
	FindStringSubmatchIndex(s string) []int
	ReplaceAllString(src, repl string) string
	FindString(s string) string
	FindAllString(s string, n int) []string
	MatchString(s string) bool
	SubexpNames() []string
}
