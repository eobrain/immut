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
	"github.com/eobrain/immut/set"
	"log"
	"math/rand"
	"testing"
)

var emptySet = set.New()
var intsSet = set.New(2, 3, 1)
var stringsSet = set.New("one", "two", "three", "four")

func ExampleSetString() {
	fmt.Println(emptySet)
	fmt.Println(intsSet)
	fmt.Println(stringsSet)
	fmt.Println(set.New(2, 4, 3, 1))
	// Output:
	// {}
	// {1,2,3}
	// {four,one,three,two}
	// {1,2,3,4}
}

func BenchmarkSetFront(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intsSet.Front()
	}
}

func ExampleSetInitAnyOrder() {
	fmt.Println(set.New(1, 2))
	fmt.Println(set.New(2, 1))
	// Output:
	// {1,2}
	// {1,2}
}

func ExampleSetBigAllSame() {
	big := x8192(set.New("foo"))
	fmt.Println(big.Len())
	// Output:
	// 1
}

var r = rand.New(rand.NewSource(99))

func random(n int) immut.Seq {
	if n == 0 {
		return set.New()
	}
	return random(n - 1).AddFront(r.Float64())
}

func ExampleSetBig() {

	big := random(8888)
	fmt.Println(big.Len())
	// Output:
	// 8888
}

/*
func TestVeryBig(t *testing.T) {
	big := x8192(set.New("foo"))
	vBig := x8192(big)
	if vBig.Len() != 8192*8192 {
		t.FailNow()
	}
}
*/

func BenchmarkSetNilIsEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		emptySet.IsEmpty()
	}
}
func BenchmarkSetIsEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intsSet.IsEmpty()
	}
}

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
}
