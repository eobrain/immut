package test

// Copyright 2013 Eamonn O'Brien-Strain
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"github.com/eobrain/immut/slice"
	"testing"
)

var emptyA = slice.New()
var intsA = slice.New(1, 2, 3)
var stringsA = slice.New("one", "two", "three", "four")

func BenchmarkListFrontA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intsA.Front()
	}
}

func ExampleBigA() {
	big := x8192(slice.New("foo"))
	p(big.Len())
	// Output:
	// 8192
}

/*func TestVeryBig(t *testing.T) {
	big := x8192(slice.New("foo"))
	vBig := x8192(big)
	if vBig.Len() != 8192*8192 {
		t.FailNow()
	}
}*/

func BenchmarkIntsAIsEmptyA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intsA.IsEmpty()
	}
}
