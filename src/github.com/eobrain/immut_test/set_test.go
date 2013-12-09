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
	"log"
	"math/rand"
	"testing"

	"github.com/eobrain/immut"
)

var emptySet = immut.Set()
var intsSet = immut.Set(2, 3, 1)
var stringsSet = immut.Set("one", "two", "three", "four")

func ExampleSetString() {
	p(emptySet)
	p(intsSet)
	p(stringsSet)
	// Output:
	// []
	// [1,2,3]
	// [four,one,three,two]
}

func ExampleSetIsEmpty() {
	p(emptySet.IsEmpty())
	p(intsSet.IsEmpty())
	// Output:
	// true
	// false
}

func ExampleSetLength() {
	p(emptySet.Length())
	p(intsSet.Length())
	p(stringsSet.Length())
	// Output:
	// 0
	// 3
	// 4
}

func ExampleSetFirst() {

	p(stringsSet.First())
	p(intsSet.First())
	p(emptySet.First())

	// Output:
	// four <nil>
	// 1 <nil>
	// <nil> getting First of empty seq
}

func ExampleSetAddAll() {
	p(intsSet.AddAll(stringsSet))
	p(stringsSet.AddAll(intsSet))
	p(immut.Set("a", "b", "c", "e", "d", "f", "g", "h").AddAll(immut.Set("X")))
	p(immut.Set("X").AddAll(immut.Set("a", "b", "c", "d", "e", "g", "f", "h")))
	p(immut.Set("a", "b", "c", "d", "e", "f", "g", "h").AddAll(immut.Set("X", "Y")))
	p(immut.Set("X", "Y").AddAll(immut.Set("a", "b", "c", "d", "e", "f", "g", "h")))
	p(immut.Set("a", "b", "c", "d", "e", "f", "g", "h").AddAll(immut.Set("X", "Y")))
	p(immut.Set("X", "Y").AddAll(immut.Set("a", "b", "c", "d", "e", "f", "g", "h")))
	p(immut.Set("a", "b", "c", "d", "e", "f", "g", "h").AddAll(immut.Set("X", "Y", "Z")))
	p(immut.Set("X", "Y", "Z").AddAll(immut.Set("a", "b", "c", "d", "e", "f", "g", "h")))
	// Output:
	// [1,2,3,four,one,three,two]
	// [1,2,3,four,one,three,two]
	// [X,a,b,c,d,e,f,g,h]
	// [X,a,b,c,d,e,f,g,h]
	// [X,Y,a,b,c,d,e,f,g,h]
	// [X,Y,a,b,c,d,e,f,g,h]
	// [X,Y,a,b,c,d,e,f,g,h]
	// [X,Y,a,b,c,d,e,f,g,h]
	// [X,Y,Z,a,b,c,d,e,f,g,h]
	// [X,Y,Z,a,b,c,d,e,f,g,h]
}

func ExampleSetAdd() {
	p(stringsSet.Add("zero"))
	// Output:
	// [four,one,three,two,zero]

}

func ExampleSetEach() {
	intsSet.Each(func(item immut.Item) {
		i := item.(int)
		p(i * i)
	})
	// Output:
	// 1
	// 4
	// 9
}

func ExampleSetBigAllSame() {
	big := x8192(immut.Set("foo"))
	p(big.Length())
	// Output:
	// 1
}

var r = rand.New(rand.NewSource(99))

func random(n int) immut.Seq {
	if n == 0 {
		return immut.Set()
	}
	return random(n - 1).Add(r.Float64())
}

func ExampleSetBig() {

	big := random(8888)
	p(big.Length())
	// Output:
	// 8888
}

/*
func TestVeryBig(t *testing.T) {
	big := x8192(immut.Set("foo"))
	vBig := x8192(big)
	if vBig.Length() != 8192*8192 {
		t.FailNow()
	}
}
*/

func BenchmarkSetNilIsEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		emptySet.IsEmpty()
	}
}
func BenchmarkSetIntsSetIsEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intsSet.IsEmpty()
	}
}

func ExampleSetJoin() {
	p(stringsSet.Join("|"))
	p(intsSet.Join(" <--> "))
	// Output:
	// four|one|three|two
	// 1 <--> 2 <--> 3
}

func ExampleSetMap() {
	p(intsSet.Map(func(item immut.Item) immut.Item {
		i := item.(int)
		return i * i
	}))
	// Output:
	// [1,4,9]
}

func ExampleSetFilter() {
	p(intsSet.Filter(func(item immut.Item) bool {
		i := item.(int)
		return i%2 == 1
	}))
	// Output:
	// [1,3]
}

func ExampleRest() {

	p(intsSet.Rest())
	p(stringsSet.Rest())
	p(emptySet.Rest())
	// Output:
	// [2,3] <nil>
	// [one,three,two] <nil>
	// <nil> getting Rest of empty seq
}

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
}
