// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package serialize

import (
	"testing"

	"github.com/m3db/m3/src/x/checked"
	"github.com/m3db/m3/src/x/ident"
)

func benchmarkTags() ident.Tags {
	return ident.NewTags(
		ident.StringTag("abc", "Def"),
		ident.StringTag("ghifsdf", "andson"),
		ident.StringTag("abc", "Def"),
		ident.StringTag("ghifsdf", "andson"),
		ident.StringTag("abc", "Def"),
		ident.StringTag("ghifsdf", "andson"),
		ident.StringTag("abc", "Def"),
		ident.StringTag("ghifsdf", "andson"),
		ident.StringTag("abc", "Def"),
		ident.StringTag("ghifsdf", "andson"),
		ident.StringTag("abc", "Def"),
		ident.StringTag("ghifsdf", "andson"),
		ident.StringTag("abc", "Def"),
		ident.StringTag("ghifsdf", "andson"),
	)
}

var global checked.Bytes

func BenchmarkCustomReadWrite(b *testing.B) {
	tags := benchmarkTags()
	iter := ident.NewTagsIterator(tags)
	enc := newTagEncoder(defaultNewCheckedBytesFn, newTestEncoderOpts(), nil)
	// dec := newTagDecoder(testDecodeOpts, nil)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		// copy := iter.Duplicate()
		enc.Reset()
		enc.Encode(iter)
		data, _ := enc.Data()
		global = data
		// dec.Reset(data)
	}
}
