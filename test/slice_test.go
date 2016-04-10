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
	"bytes"
	"github.com/eobrain/immut/slice"
	"testing"
)

var emptyA = slice.New()
var intsA = slice.New(1, 2, 3)
var stringsA = slice.New("one", "two", "three", "four")

func ExampleStringA() {
	p(emptyA)
	p(intsA)
	p(stringsA)
	// Output:
	// <nil>
	// [1,2,3]
	// [one,two,three,four]
}

func ExampleIsEmptyA() {
	p(emptyA.IsEmpty())
	p(intsA.IsEmpty())
	// Output:
	// true
	// false
}

func ExampleLenA() {
	p(emptyA.Len())
	p(intsA.Len())
	p(stringsA.Len())
	// Output:
	// 0
	// 3
	// 4
}

func ExampleFrontA() {

	p(stringsA.Front())
	p(intsA.Front())
	p(emptyA.Front())

	// Output:
	// one <nil>
	// 1 <nil>
	// <nil> getting Front of empty seq
}

func BenchmarkListFrontA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intsA.Front()
	}
}

func ExampleAddAllA() {
	p(intsA.AddAll(stringsA))
	// Output:
	// [1,2,3,one,two,three,four]
}

func ExampleAddA() {
	p(stringsA.AddFront("zero"))
	p(stringsA.AddBack("five"))
	// Output:
	// [zero,one,two,three,four]
	// [one,two,three,four,five]

}

func ExampleEachA() {
	intsA.Each(func(item interface{}) {
		i := item.(int)
		p(i * i)
	})
	// Output:
	// 1
	// 4
	// 9
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

func ExampleJoinA() {
	var buf bytes.Buffer
	stringsA.Join("|", &buf)
	buf.WriteString("\n")
	intsA.Join(" <--> ", &buf)
	p(buf.String())
	// Output:
	// one|two|three|four
	// 1 <--> 2 <--> 3
}

func ExampleMapA() {
	p(intsA.Map(func(item interface{}) interface{} {
		i := item.(int)
		return i * i
	}))
	// Output:
	// [1,4,9]
}

func ExampleFilterA() {
	p(intsA.Filter(func(item interface{}) bool {
		i := item.(int)
		return i%2 == 1
	}))
	// Output:
	// [1,3]
}

// For below see http://java.ociweb.com/mark/clojure/article.html

func ExampleCountA() {
	p(slice.New(19, "yellow", true).Len())
	// Output:
	// 3
}

func ExampleReverseA() {
	p(slice.New(2, 4, 7).Reverse())
	// Output:
	// [7,4,2]
}

func ExampleMap2A() {
	p(slice.New(2, 4, 7).Map(func(x interface{}) interface{} {
		return x.(int) + 3
	}))
	// Output:
	// [5,7,10]
}
