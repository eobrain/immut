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
	"log"
	"testing"
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

func ExampleRemoveA() {
	p(immut.Remove(emptyA, 33))
	p(immut.Remove(intsA, 33))
	p(immut.Remove(intsA, "foo"))
	p(immut.Remove(intsA, 1))
	p(immut.Remove(intsA, 2))
	p(immut.Remove(intsA, 3))
	p(immut.Remove(stringsA, "one"))
	p(immut.Remove(stringsA, "two"))
	p(immut.Remove(stringsA, "four"))
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

func BenchmarkListRemove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		immut.Remove(ints, 2)
	}
}

func BenchmarkListRemoveA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		immut.Remove(intsA, 2)
	}
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
		emptyA,
		intsA,
		stringsA,
		x8192(immut.Slice("foo")),
		immut.Slice(19, "yellow", true),
		immut.Slice(2, 4, 7),
		immut.Slice(2, 4),
		immut.Slice(2),
		immut.Slice("Moe", "Larry", "Curly", "Shemp"),
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

func BenchmarkListNth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		immut.Nth(ints, 2)
	}
}

func BenchmarkListNthA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		immut.Nth(intsA, 2)
	}
}

func BenchmarkInstsSetNth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		immut.Nth(intsSet, 2)
	}
}

func BenchmarkMakeBigList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x8192(ints)
	}
}

func BenchmarkMakeBigListA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x8192(intsA)
	}
}

func BenchmarkBigListNth(b *testing.B) {
	b.StopTimer()
	big := x8192(immut.List(1))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		immut.Nth(big, 8000)
	}
}

func BenchmarkBigListNthA(b *testing.B) {
	b.StopTimer()
	big := x8192(immut.Slice(1))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		immut.Nth(big, 8000)
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

func ExampleNthA() {
	stooges := immut.Slice("Moe", "Larry", "Curly", "Shemp")
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

// Returns the frst two elements of the integer data vector that sum
// to 10. I believe this is asymptotically O(n) assuming that
// distribution of integer values stays statistically the same as the
// data array grows in size.
func findAddsTo10(data immut.Seq) (int, int, error) {
	n := data.Len()

	var loop func(int, int) (int, int, error)
	loop = func(indexSum, i int) (int, int, error) {
		j := indexSum - i
		di, err := immut.Nth(data, i)
		if err != nil {
			log.Fatalf("Nth(%v,%v) -> %v, %v", data, i, di, err)
		}
		dj, err := immut.Nth(data, j)
		if err != nil {
			panic(err)
		}

		if di.(int)+dj.(int) == 10 {
			// found result
			return di.(int), dj.(int), nil
		}

		if (i + 1) < indexSum {
			// increment inner loop
			return loop(indexSum, i+1)
		}

		if indexSum < n {
			// increment outer loop
			ii := (indexSum + 1) - n + 1
			if ii < 0 {
				ii = 0
			}
			return loop(indexSum+1, ii)
		}

		return -1, -1, fmt.Errorf("No elements add to 10")
	}

	return loop(1, 0)
}

func TestBasic1(t *testing.T) {
	if a, b, err := findAddsTo10(immut.List(2, 4, 7, 8, 10, 12)); err != nil || a != 2 || b != 8 {
		t.Errorf("got %d, %d, %q", a, b, err)
	}
}
func TestBasic2(t *testing.T) {
	if a, b, err := findAddsTo10(immut.List(5, 6, 8, 22, 4)); err != nil || a != 6 || b != 4 {
		t.Errorf("got %d, %d, %q", a, b, err)
	}
}

func TestBasic3(t *testing.T) {
	if a, b, err := findAddsTo10(immut.List(0, 20, 50, 100, 999, 999, 999)); err == nil {
		t.Errorf("got %d, %d, %q", a, b, err)
	}
}

func TestBasic4(t *testing.T) {
	if a, b, err := findAddsTo10(immut.List(99, 99, 3, 7, 99, 99, 99)); err != nil || a != 3 || b != 7 {
		t.Errorf("got %d, %d, %q", a, b, err)
	}
}

func TestBasic5(t *testing.T) {
	if a, b, err := findAddsTo10(immut.List(99, 99, 3, 99, 7, 99, 99, 99)); err != nil || a != 3 || b != 7 {
		t.Errorf("got %d, %d, %q", a, b, err)
	}
}
func TestBasic6(t *testing.T) {
	if a, b, err := findAddsTo10(immut.List(3, 99, 99, 99, 99, 99, 99, 7)); err != nil || a != 3 || b != 7 {
		t.Errorf("got %d, %d, %q", a, b, err)
	}
}
