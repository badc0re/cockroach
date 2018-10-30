// Copyright 2018 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package tree

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

func (d *DTuple) pgwireFormat(ctx *FmtCtx) {
	// When converting a tuple to text in "postgres mode" there is
	// special behavior: values are printed in "postgres mode" then the
	// result string itself is rendered in "postgres mode".
	// Immediate NULL tuple elements are printed as the empty string.
	//
	// In this last conversion, for *tuples* the special double quote
	// and backslash characters are *doubled* (not escaped).  Other
	// special characters from C like \t \n etc are not escaped and
	// instead printed as-is. Only non-valid characters get escaped to
	// hex. So we delegate this formatting to a tuple-specific
	// string printer called pgwireFormatStringInTuple().
	ctx.WriteByte('(')
	comma := ""
	for _, v := range d.D {
		ctx.WriteString(comma)
		switch dv := v.(type) {
		case *DTuple, *DArray:
			s := AsStringWithFlags(v, ctx.flags)
			pgwireFormatStringInTuple(ctx.Buffer, s)
		case *DString:
			pgwireFormatStringInTuple(ctx.Buffer, string(*dv))
		case *DCollatedString:
			pgwireFormatStringInTuple(ctx.Buffer, dv.Contents)
		default:
			ctx.FormatNode(v)
		}
		comma = ","
	}
	ctx.WriteByte(')')
}

func pgwireFormatStringInTuple(buf *bytes.Buffer, in string) {
	quote := pgwireQuoteStringInTuple(in)
	if quote {
		buf.WriteByte('"')
	}
	// Loop through each unicode code point.
	for _, r := range in {
		if r == '"' || r == '\\' {
			// Strings in tuples double " and \.
			buf.WriteByte(byte(r))
			buf.WriteByte(byte(r))
		} else {
			buf.WriteRune(r)
		}
	}
	if quote {
		buf.WriteByte('"')
	}
}

func (d *DArray) pgwireFormat(ctx *FmtCtx) {
	// When converting an array to text in "postgres mode" there is
	// special behavior: values are printed in "postgres mode" then the
	// result string itself is rendered in "postgres mode".
	// Immediate NULL array elements are printed as "NULL".
	//
	// In this last conversion, for *arrays* the special double quote
	// and backslash characters are *escaped* (not doubled).  Other
	// special characters from C like \t \n etc are not escaped and
	// instead printed as-is. Only non-valid characters get escaped to
	// hex. So we delegate this formatting to a tuple-specific
	// string printer called pgwireFormatStringInArray().
	ctx.WriteByte('{')
	comma := ""
	for _, v := range d.Array {
		ctx.WriteString(comma)
		switch dv := v.(type) {
		case dNull:
			ctx.WriteString("NULL")
		case *DTuple, *DArray:
			s := AsStringWithFlags(v, ctx.flags)
			pgwireFormatStringInArray(ctx.Buffer, s)
		case *DString:
			pgwireFormatStringInArray(ctx.Buffer, string(*dv))
		case *DCollatedString:
			pgwireFormatStringInArray(ctx.Buffer, dv.Contents)
		case *DTimestamp, *DTimestampTZ:
			// These always contain spaces, so quote them.
			ctx.Buffer.WriteByte('"')
			ctx.FormatNode(v)
			ctx.Buffer.WriteByte('"')
		default:
			ctx.FormatNode(v)
		}
		comma = ","
	}
	ctx.WriteByte('}')
}

var tupleQuoteSet, arrayQuoteSet asciiSet

func init() {
	var ok bool
	tupleQuoteSet, ok = makeASCIISet(" \t\v\f\r\n(),\"\\")
	if !ok {
		panic("tuple asciiset")
	}
	arrayQuoteSet, ok = makeASCIISet(" \t\v\f\r\n{},\"\\")
	if !ok {
		panic("array asciiset")
	}
}

func pgwireQuoteStringInTuple(in string) bool {
	return in == "" || tupleQuoteSet.in(in)
}

func pgwireQuoteStringInArray(in string) bool {
	return in == "" || strings.ToLower(in) == "null" || arrayQuoteSet.in(in)
}

func pgwireFormatStringInArray(buf *bytes.Buffer, in string) {
	quote := pgwireQuoteStringInArray(in)
	if quote {
		buf.WriteByte('"')
	}
	// Loop through each unicode code point.
	for _, r := range in {
		if r == '"' || r == '\\' {
			// Strings in arrays escape " and \.
			buf.WriteByte('\\')
			buf.WriteByte(byte(r))
		} else {
			buf.WriteRune(r)
		}
	}
	if quote {
		buf.WriteByte('"')
	}
}

// From: https://github.com/golang/go/blob/master/src/strings/strings.go

// asciiSet is a 32-byte value, where each bit represents the presence of a
// given ASCII character in the set. The 128-bits of the lower 16 bytes,
// starting with the least-significant bit of the lowest word to the
// most-significant bit of the highest word, map to the full range of all
// 128 ASCII characters. The 128-bits of the upper 16 bytes will be zeroed,
// ensuring that any non-ASCII character will be reported as not in the set.
type asciiSet [8]uint32

// makeASCIISet creates a set of ASCII characters and reports whether all
// characters in chars are ASCII.
func makeASCIISet(chars string) (as asciiSet, ok bool) {
	for i := 0; i < len(chars); i++ {
		c := chars[i]
		if c >= utf8.RuneSelf {
			return as, false
		}
		as[c>>5] |= 1 << uint(c&31)
	}
	return as, true
}

// contains reports whether c is inside the set.
func (as *asciiSet) contains(c byte) bool {
	return (as[c>>5] & (1 << uint(c&31))) != 0
}

// in reports whether any member of the set is in s.
func (as *asciiSet) in(s string) bool {
	for i := 0; i < len(s); i++ {
		if as.contains(s[i]) {
			return true
		}
	}
	return false
}
