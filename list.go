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
func List(item ...interface{}) Seq {
	if len(item) == 0 {
		return null{}
	}
	return &cons{item[0], List(item[1:]...)}
}

// Everything below here is private

type cons struct {
	first interface{}
	rest  Seq
}

func (xs *cons) Len() int {
	return 1 + xs.rest.Len()
}

func (xs *cons) Contains(x interface{}) bool {
	return xs.first == x || xs.rest.Contains(x)
	//TODO make this tail recursive
}

func (xs *cons) Front() (interface{}, error) {
	return xs.first, nil
}

func (xs *cons) Rest() (Seq, error) {
	return xs.rest, nil
}

func (cons) IsEmpty() bool {
	return false
}

func (xs *cons) Each(f func(interface{})) {
	f(xs.first)
	xs.rest.Each(f) //recursion
}

// O(n)
func (xs *cons) Join(sep string, buf *bytes.Buffer) {
	buf.WriteString(fmt.Sprintf("%v", xs.first))
	if !xs.rest.IsEmpty() {
		buf.WriteString(sep)
		xs.rest.Join(sep, buf)
	}
}

func (xs *cons) Reverse() Seq {
	return xs.rest.Reverse().AddBack(xs.first)
}

// Add to beginning
func (xs *cons) AddFront(x interface{}) Seq {
	return &cons{x, xs}
}

func (xs *cons) AddBack(x interface{}) Seq {
	return &cons{xs.first, xs.rest.AddBack(x)}
}

func (xs *cons) AddAll(that Seq) Seq {
	//fmt.Printf("[%d].AddAll([%d])\n", xs.Len(), that.Len())
	return &cons{xs.first, xs.rest.AddAll(that)}
}

func (xs *cons) Forall(f func(interface{}) bool) bool {
	return f(xs.first) && xs.rest.Forall(f)
}

func (xs *cons) Map(f func(interface{}) interface{}) Seq {
	return &cons{f(xs.first), xs.rest.Map(f)}
}

func (xs *cons) Filter(f func(interface{}) bool) Seq {
	if f(xs.first) {
		return &cons{xs.first, xs.rest.Filter(f)}
	}
	return xs.rest.Filter(f)
}

func (xs *cons) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	xs.Join(",", &buf)
	buf.WriteString("]")
	return buf.String()
}

func (xs *cons) addTreeNode(x interface{}, itemS string) *tree {
	return null{}.addTreeNode(x, itemS)
}
