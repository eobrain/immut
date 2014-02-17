// The immut package contains immutable structure-sharing collections
// in the style of Scala or Clojure.
package immut

import (
	"bytes"
)

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

// An Item is an element in a Seq.
type Item interface{}

// A Seq is an immutable sequence of Items.  Where multiple O(...)
// given, first is for list, second is for set (average case, assuming
// it is a balanced tree),
type Seq interface {

	// Len is the number of elements. O(n) or O(log n)
	Len() int

	// Contains is whether the Item is in the Seq. O(n) or O(log n)
	Contains(Item) bool

	// Front returns the first item. O(1) or O(log n)
	Front() (Item, error)

	// Rest returns new list with all except the first item. O(1) or  O(n^2 * log(n))
	Rest() (Seq, error)

	// IsEmpty is whether this is the empty seq. O(1)
	IsEmpty() bool

	// Each Apply the function to each item in the seq. O(n)
	Each(func(Item))

	// Join writes a concatentaion of the string representations
	// of the items separated by sep into the Writer. O(n)
	Join(string, *bytes.Buffer)

	// AddFront returns a new seq with the item added on to the beginning. O(1) or O(log n)
	AddFront(Item) Seq

	//O(n) or O(1) return a new seq with the item added on to the end
	AddBack(Item) Seq

	//return a new seq that is a concatenation of this seq with the given one
	AddAll(Seq) Seq

	//return a new seq that is the reverse of this one
	Reverse() Seq

	//whether function is true for all items, or if there are no items
	Forall(func(Item) bool) bool

	//return a new seq where each item is the result of running
	//the function on the corresponding item of this seq
	Map(func(Item) Item) Seq

	//return a new seq with a subset of the items for which the
	//function is true
	Filter(func(Item) bool) Seq

	addTreeNode(Item, string) *tree
}

// Return sequence resulting from removing the item, or the sequence
// itself if item not contained in it.
func Remove(xs Seq, x Item) Seq {
	return xs.Filter(func(y Item) bool { return y != x })
}

// Return second item in sequence.
func Second(xs Seq) (Item, error) {
	rest, err := xs.Rest()
	if err != nil {
		return nil, err
	}
	return rest.Front()
}

// Return item number n in sequence, where immut.Nth(xs,0) is the same
// as xs.Front() and immut.Nth(xs,1) is the same as immut.Second(xs)
func Nth(xs Seq, n uint) (Item, error) {
	if n == 0 {
		return xs.Front()
	}
	rest, err := xs.Rest()
	if err != nil {
		return nil, err
	}
	return Nth(rest, n-1)
}

func Back(xs Seq) (Item, error) {
	rest, err := xs.Rest()
	if err != nil {
		return nil, err
	}
	if rest.IsEmpty() {
		return xs.Front()
	}
	return Back(rest)
}

func Join(xs Seq, sep string) string {
	var buf bytes.Buffer
	xs.Join(sep, &buf)
	return buf.String()
}
