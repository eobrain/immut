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
	"fmt"
	"github.com/eobrain/immut"
	"github.com/eobrain/immut/list"
	"testing"
)

var empty = list.New()
var ints = list.New(1, 2, 3)
var strings = list.New("one", "two", "three", "four")

func ExampleString() {
	fmt.Println(empty)
	fmt.Println(ints)
	fmt.Println(strings)
	// Output:
	// []
	// [1,2,3]
	// [one,two,three,four]
}

func BenchmarkListFront(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ints.Front()
	}
}

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
	big := x8192(list.New("foo"))
	fmt.Println(big.Len())
	// Output:
	// 8192
}

/*func TestVeryBig(t *testing.T) {
	big := x8192(list.New("foo"))
	vBig := x8192(big)
	if vBig.Len() != 8192*8192 {
		t.FailNow()
	}
}*/

func BenchmarkIntsIsEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ints.IsEmpty()
	}
}

func ExampleJoin() {
	var buf bytes.Buffer
	strings.Join("|", &buf)
	buf.WriteString("\n")
	ints.Join(" <--> ", &buf)
	fmt.Println(buf.String())
	// Output:
	// one|two|three|four
	// 1 <--> 2 <--> 3
}

func ExampleMap() {
	fmt.Println(ints.Map(func(item interface{}) interface{} {
		i := item.(int)
		return i * i
	}))
	// Output:
	// [1,4,9]
}

func ExampleFilter() {
	fmt.Println(ints.Filter(func(item interface{}) bool {
		i := item.(int)
		return i%2 == 1
	}))
	// Output:
	// [1,3]
}

// For below see http://java.ociweb.com/mark/clojure/article.html

func ExampleCount() {
	fmt.Println(list.New(19, "yellow", true).Len())
	// Output:
	// 3
}

func ExampleReverse() {
	fmt.Println(list.New(2, 4, 7).Reverse())
	// Output:
	// [7,4,2]
}

func ExampleMap2() {
	fmt.Println(list.New(2, 4, 7).Map(func(x interface{}) interface{} {
		return x.(int) + 3
	}))
	// Output:
	// [5,7,10]
}
