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
	p(emptySet)
	p(intsSet)
	p(stringsSet)
	p(set.New(2, 4, 3, 1))
	// Output:
	// <nil>
	// {1,2,3}
	// {four,one,three,two}
	// {1,2,3,4}
}

func ExampleSetRemove() {
	p(immut.Remove(emptySet, 33))
	p(immut.Remove(intsSet, 33))
	p(immut.Remove(intsSet, "foo"))
	p(immut.Remove(intsSet, 1))
	p(immut.Remove(intsSet, 2))
	p(immut.Remove(intsSet, 3))
	p(stringsSet, "- one  =", immut.Remove(stringsSet, "one"))
	p(stringsSet, "- two  =", immut.Remove(stringsSet, "two"))
	p(stringsSet, "- four =", immut.Remove(stringsSet, "four"))
	p(immut.Remove(set.New(2, 4, 3, 1), 2))
	p(immut.Remove(set.New(2, 4, 3, 1), 4))
	p(immut.Remove(set.New(2, 4, 3, 1), 1))
	// Output:
	// <nil>
	// {1,2,3}
	// {1,2,3}
	// {2,3}
	// {1,3}
	// {1,2}
	// {four,one,three,two} - one  = {four,three,two}
	// {four,one,three,two} - two  = {four,one,three}
	// {four,one,three,two} - four = {one,three,two}
	// {1,3,4}
	// {1,2,3}
	// {2,3,4}
}

func ExampleSetIsEmpty() {
	p(emptySet.IsEmpty())
	p(intsSet.IsEmpty())
	// Output:
	// true
	// false
}

func ExampleSetLen() {
	p(emptySet.Len())
	p(intsSet.Len())
	p(stringsSet.Len())
	// Output:
	// 0
	// 3
	// 4
}

func ExampleSetFirst() {

	p(stringsSet.Front())
	p(intsSet.Front())
	p(emptySet.Front())

	// Output:
	// four <nil>
	// 1 <nil>
	// <nil> getting Front of empty seq
}

func BenchmarkSetFront(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intsSet.Front()
	}
}

func ExampleSetAddAll() {
	p(intsSet.AddAll(stringsSet))
	p(stringsSet.AddAll(intsSet))
	p(set.New("a", "b", "c", "e", "d", "f", "g", "h").AddAll(set.New("X")))
	p(set.New("X").AddAll(set.New("a", "b", "c", "d", "e", "g", "f", "h")))
	p(set.New("a", "b", "c", "d", "e", "f", "g", "h").AddAll(set.New("X", "Y")))
	p(set.New("X", "Y").AddAll(set.New("a", "b", "c", "d", "e", "f", "g", "h")))
	p(set.New("a", "b", "c", "d", "e", "f", "g", "h").AddAll(set.New("X", "Y")))
	p(set.New("X", "Y").AddAll(set.New("a", "b", "c", "d", "e", "f", "g", "h")))
	p(set.New("a", "b", "c", "d", "e", "f", "g", "h").AddAll(set.New("X", "Y", "Z")))
	p(set.New("X", "Y", "Z").AddAll(set.New("a", "b", "c", "d", "e", "f", "g", "h")))
	// Output:
	// {1,2,3,four,one,three,two}
	// {1,2,3,four,one,three,two}
	// {X,a,b,c,d,e,f,g,h}
	// {X,a,b,c,d,e,f,g,h}
	// {X,Y,a,b,c,d,e,f,g,h}
	// {X,Y,a,b,c,d,e,f,g,h}
	// {X,Y,a,b,c,d,e,f,g,h}
	// {X,Y,a,b,c,d,e,f,g,h}
	// {X,Y,Z,a,b,c,d,e,f,g,h}
	// {X,Y,Z,a,b,c,d,e,f,g,h}
}

func ExampleSetAdd() {
	p(stringsSet.AddFront("zero"))
	p(stringsSet.AddBack("zero"))
	// Output:
	// {four,one,three,two,zero}
	// {four,one,three,two,zero}

}

func ExampleSetAddAnyOrder() {
	p(set.New(1).AddFront(2))
	p(set.New(2).AddFront(1))
	p(set.New("aaa").AddFront("bbb"))
	p(set.New("bbb").AddFront("aaa"))
	// Output:
	// {1,2}
	// {1,2}
	// {aaa,bbb}
	// {aaa,bbb}
}

func ExampleSetInitAnyOrder() {
	p(set.New(1, 2))
	p(set.New(2, 1))
	// Output:
	// {1,2}
	// {1,2}
}

func ExampleSetisSet() {
	p(intsSet.AddFront(1))
	p(intsSet.AddFront(2))
	p(intsSet.AddFront(3))
	p(intsSet.AddFront(0))
	p(intsSet.AddFront(4))
	// Output:
	// {1,2,3}
	// {1,2,3}
	// {1,2,3}
	// {0,1,2,3}
	// {1,2,3,4}
}

func ExampleSetEach() {
	intsSet.Each(func(item interface{}) {
		i := item.(int)
		p(i * i)
	})
	// Output:
	// 1
	// 4
	// 9
}

func ExampleSetBigAllSame() {
	big := x8192(set.New("foo"))
	p(big.Len())
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
	p(big.Len())
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

func ExampleSetJoin() {
	var buf bytes.Buffer
	stringsSet.Join("|", &buf)
	buf.WriteString("\n")
	intsSet.Join(" <--> ", &buf)
	p(buf.String())
	// Output:
	// four|one|three|two
	// 1 <--> 2 <--> 3
}

func ExampleSetMap() {
	p(intsSet.Map(func(item interface{}) interface{} {
		i := item.(int)
		return i * i
	}))
	// Output:
	// {1,4,9}
}

func ExampleSetFilter() {
	p(intsSet.Filter(func(item interface{}) bool {
		i := item.(int)
		return i%2 == 1
	}))
	// Output:
	// {1,3}
}

func ExampleRest() {
	p(intsSet.Rest())
	p(stringsSet.Rest())
	p(emptySet.Rest())
	// Output:
	// {2,3} <nil>
	// {one,three,two} <nil>
	// <nil> getting Rest of empty seq
}

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
}
