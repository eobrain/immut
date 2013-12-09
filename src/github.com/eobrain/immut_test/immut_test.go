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

import "github.com/eobrain/immut"
import "testing"
import "fmt"

var empty = immut.List()
var ints = immut.List(1, 2, 3)
var strings = immut.List("one", "two", "three", "four")

func p(x ...interface{}) {
	fmt.Println(x...)
}

func ExampleString() {
	p(empty)
	p(ints)
	p(strings)
	// Output:
	// []
	// [1,2,3]
	// [one,two,three,four]
}

func ExampleIsEmpty() {
	p(empty.IsEmpty())
	p(ints.IsEmpty())
	// Output:
	// true
	// false
}

func ExampleLength() {
	p(empty.Length())
	p(ints.Length())
	p(strings.Length())
	// Output:
	// 0
	// 3
	// 4
}

func ExampleFirst() {

	p(strings.First())
	p(ints.First())
	p(empty.First())

	// Output:
	// one <nil>
	// 1 <nil>
	// <nil> getting First of empty seq
}

func ExampleAddAll() {
	p(ints.AddAll(strings))
	// Output:
	// [1,2,3,one,two,three,four]
}

func ExampleAdd() {
	p(strings.Add("zero"))
	// Output:
	// [zero,one,two,three,four]

}

func ExampleEach() {
	ints.Each(func(item immut.Item) {
		i := item.(int)
		p(i * i)
	})
	// Output:
	// 1
	// 4
	// 9
}

/*func ExampleReverse() {
	p(strings.Reverse())
	// Output:
	// [four,three,two,one]
}*/

func x8192(x immut.Seq) (result immut.Seq) {
	x2 := x.AddAll(x)
	x4 := x2.AddAll(x2)
	x8 := x4.AddAll(x4)
	x16 := x8.AddAll(x8)
	x32 := x16.AddAll(x16)
	x64 := x32.AddAll(x32)
	x128 := x64.AddAll(x64)
	x256 := x128.AddAll(x128)
	x512 := x256.AddAll(x256)
	x1024 := x512.AddAll(x512)
	x2048 := x1024.AddAll(x1024)
	x4096 := x2048.AddAll(x2048)
	result = x4096.AddAll(x4096)
	return
}

func ExampleBig() {
	big := x8192(immut.List("foo"))
	p(big.Length())
	// Output:
	// 8192
}

/*
func TestVeryBig(t *testing.T) {
	big := x8192(immut.List("foo"))
	vBig := x8192(big)
	if vBig.Length() != 8192*8192 {
		t.FailNow()
	}
}
*/

func BenchmarkNilIsEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		empty.IsEmpty()
	}
}
func BenchmarkIntsIsEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ints.IsEmpty()
	}
}

func ExampleJoin() {
	p(strings.Join("|"))
	p(ints.Join(" <--> "))
	// Output:
	// one|two|three|four
	// 1 <--> 2 <--> 3
}

func ExampleMap() {
	p(ints.Map(func(item immut.Item) immut.Item {
		i := item.(int)
		return i * i
	}))
	// Output:
	// [1,4,9]
}

func ExampleFilter() {
	p(ints.Filter(func(item immut.Item) bool {
		i := item.(int)
		return i%2 == 1
	}))
	// Output:
	// [1,3]
}
