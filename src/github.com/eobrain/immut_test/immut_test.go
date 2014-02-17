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
	"testing"

	"github.com/eobrain/immut"
)

func p(x ...interface{}) {
	fmt.Println(x...)
}

func ExampleRemove() {
	p(immut.Remove(empty, 33))
	p(immut.Remove(ints, 33))
	p(immut.Remove(ints, "foo"))
	p(immut.Remove(ints, 1))
	p(immut.Remove(ints, 2))
	p(immut.Remove(ints, 3))
	p(immut.Remove(strings, "one"))
	p(immut.Remove(strings, "two"))
	p(immut.Remove(strings, "four"))
	// Output:
	// <nil>
	// [1,2,3]
	// [1,2,3]
	// [2,3]
	// [1,3]
	// [1,2]
	// [two,three,four]
	// [one,three,four]
	// [one,two,three]
}

func str(xs immut.Item, err error) string {
	return fmt.Sprintf("%v,%v", xs, err)
}

func TestNth(t *testing.T) {
	seqs := []immut.Seq{
		empty,
		ints,
		strings,
		x8192(immut.List("foo")),
		immut.List(19, "yellow", true),
		immut.List(2, 4, 7),
		immut.List(2, 4),
		immut.List(2),
		immut.List("Moe", "Larry", "Curly", "Shemp"),
		emptySet,
		intsSet,
		stringsSet,
		immut.Set(2, 4, 3, 1),
		intsSet.AddAll(stringsSet),
		immut.Set("X", "Y", "Z").AddAll(immut.Set("a", "b", "c", "d", "e", "f", "g", "h")),
		immut.Set(1, 2),
		immut.Set(1),
	}
	for i, xs := range seqs {
		if a, b := str(immut.Nth(xs, 0)), str(xs.Front()); a != b {
			t.Errorf("%d: %s != %s", i, a, b)
		}
		if a, b := str(immut.Nth(xs, 1)), str(immut.Second(xs)); a != b {
			t.Errorf("%d: %s != %s", i, a, b)
		}
	}
}

func BenchmarkNilIsEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		empty.IsEmpty()
	}
}

func ExampleNth() {
	stooges := immut.List("Moe", "Larry", "Curly", "Shemp")
	p(stooges.Front())
	p(immut.Second(stooges))
	p(immut.Back(stooges))
	p(immut.Nth(stooges, 2))
	// Output:
	// Moe <nil>
	// Larry <nil>
	// Shemp <nil>
	// Curly <nil>
}
