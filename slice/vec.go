package slice

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
)

// Create a new list containing the arguments.
func New(item ...interface{}) immut.Seq {
	if len(item) == 0 {
		return empty{}
	}
	return slice(item)
}

// Everything below here is private

type slice []interface{}
type empty struct{}

// O(1)
func (xs slice) Len() int {
	return len(xs)
}
func (empty) Len() int { return 0 }

// O(n)
func (xs slice) Contains(x interface{}) bool {
	for _, xx := range xs {
		if xx == x {
			return true
		}
	}
	return false
}
func (empty) Contains(interface{}) bool { return false }

// O(1)
func (xs slice) Front() (interface{}, error) {
	return xs[0], nil
}
func (empty) Front() (interface{}, error) {
	return nil, fmt.Errorf("getting Front of empty seq")
}

func (xs slice) Rest() (immut.Seq, error) {
	if len(xs) == 1 {
		return empty{}, nil
	}
	return xs[1:], nil
}
func (empty) Rest() (immut.Seq, error) {
	return nil, fmt.Errorf("getting Rest of empty seq")
}

// O(1)
func (slice) IsEmpty() bool { return false }
func (empty) IsEmpty() bool { return true }

// O(n)
func (xs slice) Each(f func(interface{})) {
	for _, x := range xs {
		f(x)
	}
}
func (empty) Each(f func(interface{})) {}

// O(n)
func (xs slice) Join(sep string, buf *bytes.Buffer) {
	buf.WriteString(fmt.Sprintf("%v", xs[0]))
	for _, x := range xs[1:] {
		buf.WriteString(sep)
		buf.WriteString(fmt.Sprintf("%v", x))
	}
}
func (empty) Join(string, *bytes.Buffer) {}

func (xs slice) Reverse() immut.Seq {
	n := len(xs)
	ys := make(slice, n)
	for i := 0; i < n; i++ {
		ys[i] = xs[n-i-1]
	}
	return ys
}
func (n empty) Reverse() immut.Seq { return n }

func (xs slice) AddFront(x interface{}) immut.Seq {
	n := len(xs)
	ys := make(slice, n+1)
	ys[0] = x
	copy(ys[1:], xs)
	return ys
}
func (empty) AddFront(item interface{}) immut.Seq { return New(item) }

func (xs slice) AddBack(x interface{}) immut.Seq {
	n := len(xs)
	ys := make(slice, n+1)
	ys[n] = x
	copy(ys[:n], xs)
	return ys
}
func (n empty) AddBack(item interface{}) immut.Seq { return New(item) }

func (xs slice) AddAll(that immut.Seq) immut.Seq {
	n := len(xs)
	thatA, ok := that.(slice)
	if ok {
		m := len(thatA)
		ys := make(slice, n+m)
		copy(ys[:n], xs)
		copy(ys[n:], thatA)
		return ys
	}
	ys := make(slice, n)
	that.Each(func(x interface{}) {
		ys = append(ys, x)
	})
	return ys
}
func (n empty) AddAll(other immut.Seq) immut.Seq { return other }

func (xs slice) Forall(f func(interface{}) bool) bool {
	for _, x := range xs {
		if !f(x) {
			return false
		}
	}
	return true
}
func (empty) Forall(f func(interface{}) bool) bool { return true }

func (xs slice) Map(f func(interface{}) interface{}) immut.Seq {
	n := len(xs)
	ys := make(slice, n)
	for i := 0; i < n; i++ {
		ys[i] = f(xs[i])
	}
	return ys
}
func (n empty) Map(f func(interface{}) interface{}) immut.Seq { return n }

func (xs slice) Filter(f func(interface{}) bool) immut.Seq {
	ys := slice{}
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}
func (n empty) Filter(f func(interface{}) bool) immut.Seq { return n }

func (xs slice) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	xs.Join(",", &buf)
	buf.WriteString("]")
	return buf.String()
}
func (empty) String() string { return "[]" }

func (xs slice) Items() (ys []interface{}) {
	ys = make([]interface{}, xs.Len())
	copy(ys, xs)
	return
}
func (empty) Items() []interface{} { return []interface{}{} }

//func (xs slice) addTreeNode(x interface{}, itemS string) *tree {
//	return empty{}.addTreeNode(x, itemS)
//}
