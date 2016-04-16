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
	"github.com/eobrain/immut/set"
	"github.com/eobrain/immut/slice"
	"log"
	"testing"
)

func p(x ...interface{}) {
	fmt.Println(x...)
}

func str(xs interface{}, ok bool) string {
	return fmt.Sprintf("%v,%v", xs, ok)
}

func TestNth(t *testing.T) {
	seqs := []immut.Seq{
		// empty,
		ints,
		strings,
		x8192(list.New("foo")),
		list.New(19, "yellow", true),
		list.New(2, 4, 7),
		list.New(2, 4),
		list.New(2),
		list.New("Moe", "Larry", "Curly", "Shemp"),
		// emptyA,
		intsA,
		stringsA,
		x8192(slice.New("foo")),
		slice.New(19, "yellow", true),
		slice.New(2, 4, 7),
		slice.New(2, 4),
		slice.New(2),
		slice.New("Moe", "Larry", "Curly", "Shemp"),
		// emptySet,
		intsSet,
		stringsSet,
		set.New(2, 4, 3, 1),
		intsSet.AddAll(stringsSet),
		set.New("X", "Y", "Z").AddAll(set.New("a", "b", "c", "d", "e", "f", "g", "h")),
		set.New(1, 2),
		set.New(1),
	}
	for i, xs := range seqs {
		get0, ok := xs.Get(0)
		if !ok {
			t.Errorf("%d: unexpected false", i)
		}
		front := xs.Front()
		if get0 != front {
			t.Errorf("%d: %s != %s", i, get0, front)
		}
	}
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
		di, ok := data.Get(i)
		if !ok {
			log.Fatalf("Nth(%v,%v) -> %v, false", data, i, di)
		}
		dj, ok := data.Get(j)
		if !ok {
			panic("unexpected false")
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
	if a, b, err := findAddsTo10(list.New(2, 4, 7, 8, 10, 12)); err != nil || a != 2 || b != 8 {
		t.Errorf("got %d, %d, %q", a, b, err)
	}
}
func TestBasic2(t *testing.T) {
	if a, b, err := findAddsTo10(list.New(5, 6, 8, 22, 4)); err != nil || a != 6 || b != 4 {
		t.Errorf("got %d, %d, %q", a, b, err)
	}
}

func TestBasic3(t *testing.T) {
	if a, b, err := findAddsTo10(list.New(0, 20, 50, 100, 999, 999, 999)); err == nil {
		t.Errorf("got %d, %d, %q", a, b, err)
	}
}

func TestBasic4(t *testing.T) {
	if a, b, err := findAddsTo10(list.New(99, 99, 3, 7, 99, 99, 99)); err != nil || a != 3 || b != 7 {
		t.Errorf("got %d, %d, %q", a, b, err)
	}
}

func TestBasic5(t *testing.T) {
	if a, b, err := findAddsTo10(list.New(99, 99, 3, 99, 7, 99, 99, 99)); err != nil || a != 3 || b != 7 {
		t.Errorf("got %d, %d, %q", a, b, err)
	}
}
func TestBasic6(t *testing.T) {
	if a, b, err := findAddsTo10(list.New(3, 99, 99, 99, 99, 99, 99, 7)); err != nil || a != 3 || b != 7 {
		t.Errorf("got %d, %d, %q", a, b, err)
	}
}
