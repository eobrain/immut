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

func BenchmarkIntsIsEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ints.IsEmpty()
	}
}
