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
	"io"
)

// Create a new list containing the arguments.
func New(item ...interface{}) immut.Seq {
	if len(item) == 0 {
		return empty{}
	}
	return slice(item)
}

// Create a new slice containing n repeats of x
func Repeat(n int, x interface{}) immut.Seq {
	result := make([]interface{}, n)
	for i := 0; i < n; i++ {
		result[i] = x
	}
	return slice(result)
}

// Everything below here is private

type slice []interface{}
type empty struct{}

// O(1)
func (xs slice) Len() int {
	return len(xs)
}
func (empty) Len() int { return 0 }

// O(1)
func (xs slice) Get(i int) (interface{}, bool) {
	if i < 0 || i >= len(xs) {
		return nil, false
	}
	return xs[i], true
}
func (empty) Get(i int) (interface{}, bool) { return nil, false }

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
func (xs slice) Front() interface{} { return xs[0] }
func (empty) Front() interface{}    { panic("getting Front of empty seq") }

// O(1)
func (xs slice) Back() interface{} { return xs[len(xs)-1] }
func (empty) Back() interface{}    { panic("getting Back of empty seq") }

func (xs slice) Rest() immut.Seq {
	if len(xs) == 1 {
		return empty{}
	}
	return xs[1:]
}
func (empty) Rest() immut.Seq { panic("getting Rest of empty seq") }

// O(1)
func (slice) IsEmpty() bool { return false }
func (empty) IsEmpty() bool { return true }

// O(n)
func (xs slice) Do(f func(interface{})) {
	for _, x := range xs {
		f(x)
	}
}
func (empty) Do(f func(interface{})) {}

// O(n)
func (xs slice) DoBackwards(f func(interface{})) {
	n := len(xs)
	for i := range xs {
		f(xs[n-i-1])
	}
}
func (empty) DoBackwards(f func(interface{})) {}

// O(n)
func (xs slice) Join(sep string, out io.Writer) {
	fmt.Fprintf(out, "%v", xs[0])
	for _, x := range xs[1:] {
		fmt.Fprintf(out, "%s%v", sep, x)
	}
}
func (empty) Join(string, io.Writer) {}

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
	return slice(append([]interface{}{x}, xs...))
}
func (empty) AddFront(item interface{}) immut.Seq { return New(item) }

func (xs slice) AddBack(x interface{}) immut.Seq {
	return slice(append(xs.Items(), x))
}
func (n empty) AddBack(item interface{}) immut.Seq { return New(item) }

func (xs slice) AddAll(that immut.Seq) immut.Seq {
	return slice(append(xs.Items(), that.Items()...))
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

func (xs slice) Remove(match interface{}) immut.Seq {
	result := []interface{}{}
	for _, x := range xs {
		if x != match {
			result = append(result, x)
		}
	}
	return slice(result)
}

func (n empty) Remove(x interface{}) immut.Seq {
	return n
}
func (xs slice) Items() (ys []interface{}) {
	ys = make([]interface{}, xs.Len())
	copy(ys, xs)
	return
}
func (empty) Items() []interface{} { return []interface{}{} }

//func (xs slice) addTreeNode(x interface{}, itemS string) *tree {
//	return empty{}.addTreeNode(x, itemS)
//}
