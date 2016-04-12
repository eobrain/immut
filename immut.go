// The immut package contains immutable structure-sharing collections
// for Go in the style of Scala or Clojure.
package immut

import (
	"bytes"
	"io"
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

// A Seq is an immutable sequence of items.
type Seq interface {

	// Len is the number of elements.
	Len() int

	// Contains is whether the item is in the Seq.
	Contains(interface{}) bool

	// Front returns the first item.
	Front() (interface{}, error)

	// Rest returns new list with all except the first item.
	Rest() (Seq, error)

	// IsEmpty is whether this is the empty seq.
	IsEmpty() bool

	// Apply the function to each item in the seq.
	Do(func(interface{}))

	// Join writes a concatenation of the string representations
	// of the items separated by sep into the Writer.
	Join(string, io.Writer)

	// AddFront returns a new seq with the item unshifted on to the beginning.
	AddFront(interface{}) Seq

	// return a new seq with the item pushed on to the end
	AddBack(interface{}) Seq

	//return a new seq that is a concatenation of this seq with the given one
	AddAll(Seq) Seq

	//return a new seq that is the reverse of this one
	Reverse() Seq

	//whether function is true for all items, or if there are no items
	Forall(func(interface{}) bool) bool

	//return a new seq where each item is the result of running
	//the function on the corresponding item of this seq
	Map(func(interface{}) interface{}) Seq

	//return a new seq with a subset of the items for which the
	//function is true
	Filter(func(interface{}) bool) Seq

	//return a newly created slice with all stored items
	Items() []interface{}
}

// Return sequence resulting from removing the item, or the sequence
// itself if item not contained in it.
func Remove(xs Seq, x interface{}) Seq {
	return xs.Filter(func(y interface{}) bool { return y != x })
}

// Return second item in sequence.
func Second(xs Seq) (interface{}, error) {
	rest, err := xs.Rest()
	if err != nil {
		return nil, err
	}
	return rest.Front()
}

// Return item number n in sequence, where immut.Nth(xs,0) is the same
// as xs.Front() and immut.Nth(xs,1) is the same as immut.Second(xs)
func Nth(xs Seq, n int) (interface{}, error) {
	if n == 0 {
		return xs.Front()
	}
	rest, err := xs.Rest()
	if err != nil {
		return nil, err
	}
	return Nth(rest, n-1)
}

// Return the last item in the sequence. O(n) or O(n^3 * log(n))
func Back(xs Seq) (interface{}, error) {
	rest, err := xs.Rest()
	if err != nil {
		return nil, err
	}
	if rest.IsEmpty() {
		return xs.Front()
	}
	return Back(rest)
}

// Return a string formed by concatenation of the string
// representations of the items separated by sep. O(n)
func Join(xs Seq, sep string) string {
	var buf bytes.Buffer
	xs.Join(sep, &buf)
	return buf.String()
}
