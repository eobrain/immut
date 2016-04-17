package unordered

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
	"io"
)

// Note, no attempt to keep this binary tree balanced

// Create a new unordered set containing the arguments.
func New(item ...interface{}) immut.Seq {
	if len(item) == 0 {
		return empty{}
	}
	result := unordered(make(map[interface{}]bool))
	for _, x := range item {
		result[x] = true
	}
	return result
}

// A Seq implemented as a binary tree, containing at least one value
type unordered map[interface{}]bool

// An empty Seq
type empty struct{}

// Everything below here is private

// O(log n)
func (xs unordered) Len() int { return len(xs) }
func (empty) Len() int        { return 0 }

// O(n)
func (xs unordered) Get(i int) (interface{}, bool) {
	j := 0
	for x, _ := range xs {
		if j == i {
			return x, true
		}
		j++
	}
	return nil, false
}
func (empty) Get(i int) (interface{}, bool) { return nil, false }

// O(1)
func (xs unordered) Contains(x interface{}) bool { return xs[x] }
func (empty) Contains(interface{}) bool          { return false }

// O(1)
func (xs unordered) Front() interface{} {
	for x, _ := range xs {
		return x
	}
	panic("cannot happen")
}
func (empty) Front() interface{} { panic("getting Front of empty seq") }

// O(n)
func (xs unordered) Back() (x interface{}) {
	for xx, _ := range xs {
		x = xx
	}
	return
}
func (empty) Back() interface{} { panic("getting Back of empty seq") }

// O(n)
func (xs unordered) Rest() immut.Seq {
	if len(xs) == 1 {
		return empty{}
	}
	ys := make(unordered)
	first := true
	for x, _ := range xs {
		if !first {
			ys[x] = true
		}
		first = false
	}
	return ys
}
func (empty) Rest() immut.Seq {
	panic("getting Rest of empty seq")
}

// O(1)
func (unordered) IsEmpty() bool { return false }
func (empty) IsEmpty() bool     { return true }

// O(n)
func (xs unordered) Do(f func(interface{})) {
	for x, _ := range xs {
		f(x)
	}
}
func (empty) Do(f func(interface{})) {}

// O(n)
func (xs unordered) DoBackwards(f func(interface{})) {
	items := xs.Items()
	n := len(items)
	for i := range items {
		f(items[n-i-1])
	}
}
func (empty) DoBackwards(f func(interface{})) {}

// O(n)
func (xs unordered) Join(sep string, out io.Writer) {
	s := ""
	for x, _ := range xs {
		fmt.Fprintf(out, "%s%v", s, x)
		s = sep
	}
}
func (empty) Join(string, io.Writer) {}

//func (xs unordered) Join(sep string) string {
//	var buf bytes.Buffer
//	xs.join(sep, &buf)
//	return buffer.String()

/*
	//TODO: make more efficient http://stackoverflow.com/a/1766304/978525
	if xs.left.Isempty() {
		if xs.right.Isempty() {
			return xs.valueS
		}
		return xs.valueS + sep + xs.right.Join(sep)
	}
	if xs.right.Isempty() {
		return xs.left.Join(sep) + sep + xs.valueS
	}
	return xs.left.Join(sep) + sep +
		xs.valueS + sep +
		xs.right.Join(sep)
*/
//}

// Cannot reverse an unsorted set, so just return the set itself
func (xs unordered) Reverse() immut.Seq { return xs }
func (n empty) Reverse() immut.Seq      { return n }

// O(n)
func (xs unordered) AddFront(x interface{}) immut.Seq {
	ys := make(unordered)
	ys[x] = true
	for y := range xs {
		ys[y] = true
	}
	return ys
}
func (empty) AddFront(x interface{}) immut.Seq { return unordered{x: true} }

// O(n)
func (xs unordered) AddBack(x interface{}) immut.Seq {
	return xs.AddFront(x) // same
}
func (n empty) AddBack(item interface{}) immut.Seq { return New(item) }

// O(n)
func (xs unordered) AddAll(that immut.Seq) immut.Seq {
	if that.IsEmpty() {
		return xs
	}
	ys := make(unordered)
	for x := range xs {
		ys[x] = true
	}
	that.Do(func(x interface{}) {
		ys[x] = true
	})
	return ys
}
func (n empty) AddAll(other immut.Seq) immut.Seq { return other }

// O(n)
func (xs unordered) Forall(f func(interface{}) bool) bool {
	for x := range xs {
		if !f(x) {
			return false
		}
	}
	return true
}
func (empty) Forall(f func(interface{}) bool) bool { return true }

// O(n)
func (xs unordered) Map(f func(interface{}) interface{}) immut.Seq {
	ys := make(unordered, len(xs))
	for x := range xs {
		ys[f(x)] = true
	}
	return ys
}
func (n empty) Map(f func(interface{}) interface{}) immut.Seq { return n }

// O(n)
func (xs unordered) Filter(f func(interface{}) bool) immut.Seq {
	if xs.Forall(f) {
		return xs
	}
	ys := make(unordered, len(xs))
	for x := range xs {
		if f(x) {
			ys[x] = true
		}
	}
	return ys
}
func (n empty) Filter(f func(interface{}) bool) immut.Seq { return n }

func (xs unordered) String() string {
	var buf bytes.Buffer
	buf.WriteString("{")
	xs.Join(",", &buf)
	buf.WriteString("}")
	return buf.String()
}
func (empty) String() string { return "{}" }

func (xs unordered) Remove(match interface{}) immut.Seq {
	if !xs.Contains(match) {
		return xs
	}
	ys := make(unordered)
	for x := range xs {
		if x != match {
			ys[x] = true
		}
	}
	return ys
}
func (n empty) Remove(x interface{}) immut.Seq { return n }

func (xs unordered) Items() (ys []interface{}) {
	ys = make([]interface{}, xs.Len())
	i := 0
	for x := range xs {
		ys[i] = x
		i++
	}
	return
}
func (empty) Items() []interface{} { return []interface{}{} }

//func (xs unordered) String() string {
//	return fmt.Sprintf("(%v %v %v)", xs.left, xs.value, xs.right)
//}
