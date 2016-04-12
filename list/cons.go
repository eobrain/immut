package list

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
	return &cons{item[0], New(item[1:]...)}
}

// Create a new list containing n repeats of x
func Repeat(n int, x interface{}) (result immut.Seq) {
	result = empty{}
	for i := 0; i < n; i++ {
		result = &cons{x, result}
	}
	return result
}

// Everything below here is private

type cons struct {
	first interface{}
	rest  immut.Seq
}
type empty struct{}

func (xs *cons) check() {
	if xs == nil {
		panic("nil cons")
	}
	if xs.rest == nil {
		panic("nil rest")
	}
}

// O(n)
func (xs *cons) Len() int {
	xs.check()
	return 1 + xs.rest.Len()
}
func (empty) Len() int { return 0 }

// O(n)
func (xs *cons) Contains(x interface{}) bool {
	xs.check()
	return xs.first == x || xs.rest.Contains(x)
	//TODO make this tail recursive
}
func (empty) Contains(interface{}) bool { return false }

// O(1)
func (xs *cons) Front() (interface{}, error) {
	xs.check()
	return xs.first, nil
}
func (empty) Front() (interface{}, error) {
	return nil, fmt.Errorf("getting Front of empty seq")
}

// O(1)
func (xs *cons) Rest() (immut.Seq, error) {
	xs.check()
	return xs.rest, nil
}
func (empty) Rest() (immut.Seq, error) {
	return nil, fmt.Errorf("getting Rest of empty seq")
}

// O(1)
func (xs *cons) IsEmpty() bool {
	xs.check()
	return false
}
func (empty) IsEmpty() bool { return true }

// O(n)
func (xs *cons) Do(f func(interface{})) {
	f(xs.first)
	xs.rest.Do(f) //recursion
}
func (empty) Do(f func(interface{})) {}

// O(n)
func (xs *cons) Join(sep string, out io.Writer) {
	xs.check()
	fmt.Fprintf(out, "%v", xs.first)
	if !xs.rest.IsEmpty() {
		fmt.Fprint(out, sep)
		xs.rest.Join(sep, out)
	}
}
func (empty) Join(string, io.Writer) {}

func (xs *cons) Reverse() immut.Seq {
	xs.check()
	return xs.rest.Reverse().AddBack(xs.first)
}
func (n empty) Reverse() immut.Seq { return n }

// O(1)
func (xs *cons) AddFront(x interface{}) immut.Seq {
	xs.check()
	return &cons{x, xs}
}
func (empty) AddFront(item interface{}) immut.Seq { return New(item) }

// O(n)
func (xs *cons) AddBack(x interface{}) immut.Seq {
	xs.check()
	return &cons{xs.first, xs.rest.AddBack(x)}
}
func (n empty) AddBack(item interface{}) immut.Seq { return New(item) }

func (xs *cons) AddAll(that immut.Seq) immut.Seq {
	xs.check()
	if xs.rest.IsEmpty() {
		return &cons{xs.first, that}
	}
	return &cons{xs.first, xs.rest.AddAll(that)}
}
func (n empty) AddAll(other immut.Seq) immut.Seq { return other }

func (xs *cons) Forall(f func(interface{}) bool) bool {
	xs.check()
	return f(xs.first) && xs.rest.Forall(f)
}
func (empty) Forall(f func(interface{}) bool) bool { return true }

func (xs *cons) Map(f func(interface{}) interface{}) immut.Seq {
	xs.check()
	return &cons{f(xs.first), xs.rest.Map(f)}
}
func (n empty) Map(f func(interface{}) interface{}) immut.Seq { return n }

func (xs *cons) Filter(f func(interface{}) bool) immut.Seq {
	xs.check()
	if f(xs.first) {
		return &cons{xs.first, xs.rest.Filter(f)}
	}
	return xs.rest.Filter(f)
}
func (n empty) Filter(f func(interface{}) bool) immut.Seq { return n }

func (xs *cons) String() string {
	xs.check()
	var buf bytes.Buffer
	buf.WriteString("[")
	xs.Join(",", &buf)
	buf.WriteString("]")
	return buf.String()
}
func (empty) String() string { return "[]" }

func (xs *cons) Items() (ys []interface{}) {
	xs.check()
	ys = make([]interface{}, xs.Len())
	i := 0
	xs.Do(func(x interface{}) {
		ys[i] = x
		i++
	})
	return
}
func (empty) Items() []interface{} { return []interface{}{} }
