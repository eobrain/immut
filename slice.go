package immut

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
)

// Create a new list containing the arguments.
func Slice(item ...interface{}) Seq {
	if len(item) == 0 {
		return null{}
	}
	return slice(item)
}

// Everything below here is private

type slice []interface{}

func (xs slice) Len() int {
	return len(xs)
}

func (xs slice) Contains(x interface{}) bool {
	for _, xx := range xs {
		if xx == x {
			return true
		}
	}
	return false
}

func (xs slice) Front() (interface{}, error) {
	return xs[0], nil
}

func (xs slice) Rest() (Seq, error) {
	if len(xs) == 1 {
		return null{}, nil
	}
	return xs[1:], nil
}

func (slice) IsEmpty() bool {
	return false
}

func (xs slice) Each(f func(interface{})) {
	for _, x := range xs {
		f(x)
	}
}

// O(n)
func (xs slice) Join(sep string, buf *bytes.Buffer) {
	buf.WriteString(fmt.Sprintf("%v", xs[0]))
	for _, x := range xs[1:] {
		buf.WriteString(sep)
		buf.WriteString(fmt.Sprintf("%v", x))
	}
}

func (xs slice) Reverse() Seq {
	n := len(xs)
	ys := make(slice, n)
	for i := 0; i < n; i++ {
		ys[i] = xs[n-i-1]
	}
	return ys
}

// Add to beginning
func (xs slice) AddFront(x interface{}) Seq {
	n := len(xs)
	ys := make(slice, n+1)
	ys[0] = x
	copy(ys[1:], xs)
	return ys
}

func (xs slice) AddBack(x interface{}) Seq {
	n := len(xs)
	ys := make(slice, n+1)
	ys[n] = x
	copy(ys[:n], xs)
	return ys
}

func (xs slice) AddAll(that Seq) Seq {
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

func (xs slice) Forall(f func(interface{}) bool) bool {
	for _, x := range xs {
		if !f(x) {
			return false
		}
	}
	return true
}

func (xs slice) Map(f func(interface{}) interface{}) Seq {
	n := len(xs)
	ys := make(slice, n)
	for i := 0; i < n; i++ {
		ys[i] = f(xs[i])
	}
	return ys
}

func (xs slice) Filter(f func(interface{}) bool) Seq {
	ys := slice{}
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}

func (xs slice) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	xs.Join(",", &buf)
	buf.WriteString("]")
	return buf.String()
}

func (xs slice) addTreeNode(x interface{}, itemS string) *tree {
	return null{}.addTreeNode(x, itemS)
}
