package set

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

// Note, no attempt to keep this binary tree balanced

// Create a new ordered set containing the arguments. O(n*log(n))
func New(item ...interface{}) immut.Seq { return newTreeNode(item...) }

// A Seq implemented as a binary tree, containing at least one value
type Tree struct {
	value  interface{}
	valueS string //hack: use string compare for ordering
	left   immut.Seq
	right  immut.Seq
}

// An empty Seq
type Empty struct{}

// Everything below here is private

func newTreeNode(item ...interface{}) treeNode {
	if len(item) == 0 {
		return Empty{}
	}
	first := item[0]
	return (&Tree{first, s(first), Empty{}, Empty{}}).buildTreeFrom(
		item[1:])
}

// Both Tree and Empty implement this
type treeNode interface {
	immut.Seq
	addTreeNode(x interface{}, itemS string) *Tree
}

// Recursively build a binary tree. O(n*log(n))
func (xs *Tree) buildTreeFrom(remaining []interface{}) *Tree {
	if len(remaining) == 0 {
		return xs
	}
	x := remaining[0]
	return xs.addTreeNode(x, s(x)).buildTreeFrom(remaining[1:])
}

func s(x interface{}) string {
	return fmt.Sprintf("%v", x)
}

// O(log n)
func (xs *Tree) Len() int {
	return 1 + xs.left.Len() + xs.right.Len()
}
func (Empty) Len() int { return 0 }

// O(log n)
func (xs *Tree) Contains(x interface{}) bool {
	itemS := s(x) //inefficiently re-creating on every recursion
	return x == xs.value ||
		itemS < xs.valueS && xs.left.Contains(x) ||
		xs.right.Contains(x)
}
func (Empty) Contains(interface{}) bool { return false }

// O(log n)
func (xs *Tree) Front() (interface{}, error) {
	if xs.left.IsEmpty() {
		return xs.value, nil
	}
	return xs.left.Front()
}
func (Empty) Front() (interface{}, error) {
	return nil, fmt.Errorf("getting Front of empty seq")
}

// O(n^2 * log(n))
func (xs *Tree) Rest() (immut.Seq, error) {
	if xs.left.IsEmpty() {
		return xs.right, nil
	}
	// Perhaps not most efficient
	leftRest, _ := xs.left.Rest() // guaranteed not empty
	return leftRest.AddFront(xs.value).AddAll(xs.right), nil
}
func (Empty) Rest() (immut.Seq, error) {
	return nil, fmt.Errorf("getting Rest of empty seq")
}

// O(1)
func (xs *Tree) IsEmpty() bool { return false }
func (Empty) IsEmpty() bool    { return true }

// O(n)
func (xs *Tree) Each(f func(interface{})) {
	xs.left.Each(f)
	f(xs.value)
	xs.right.Each(f)
}
func (Empty) Each(f func(interface{})) {}

// O(n)
func (xs *Tree) Join(sep string, buf *bytes.Buffer) {
	if !xs.left.IsEmpty() {
		xs.left.Join(sep, buf)
		buf.WriteString(sep)
	}
	buf.WriteString(xs.valueS)
	if !xs.right.IsEmpty() {
		buf.WriteString(sep)
		xs.right.Join(sep, buf)
	}
}
func (Empty) Join(string, *bytes.Buffer) {}

//func (xs *Tree) Join(sep string) string {
//	var buf bytes.Buffer
//	xs.join(sep, &buf)
//	return buffer.String()

/*
	//TODO: make more efficient http://stackoverflow.com/a/1766304/978525
	if xs.left.IsEmpty() {
		if xs.right.IsEmpty() {
			return xs.valueS
		}
		return xs.valueS + sep + xs.right.Join(sep)
	}
	if xs.right.IsEmpty() {
		return xs.left.Join(sep) + sep + xs.valueS
	}
	return xs.left.Join(sep) + sep +
		xs.valueS + sep +
		xs.right.Join(sep)
*/
//}

/// O(log n)
func (xs *Tree) addTreeNode(x interface{}, itemS string) *Tree {
	if x == xs.value {
		//set semantics -- cannnot have more than one of any value
		return xs
	}
	//hack: use string compare for ordering
	if itemS < xs.valueS {
		//put on left
		return &Tree{xs.value,
			xs.valueS,
			asTreeNode(xs.left).addTreeNode(x, itemS),
			xs.right}
	}
	//put on right
	return &Tree{xs.value,
		xs.valueS,
		xs.left,
		asTreeNode(xs.right).addTreeNode(x, itemS)}
}
func (Empty) addTreeNode(item interface{}, itemS string) *Tree {
	return &Tree{item, itemS, Empty{}, Empty{}}
}

func asTreeNode(xs immut.Seq) treeNode {
	// Avoid converting if not necessary
	switch xs := xs.(type) {
	case *Tree:
		return xs
	case Empty:
		return xs
	default:
		// Possibly expensive convert slice, and then to tree
		return newTreeNode(xs.Items()...)
	}
}

// Cannot reverse a sorted set, so just return the set itself
func (xs *Tree) Reverse() immut.Seq {
	return xs
}
func (n Empty) Reverse() immut.Seq { return n }

// O(log n)
func (xs *Tree) AddFront(x interface{}) immut.Seq {
	return xs.addTreeNode(x, s(x))
}
func (Empty) AddFront(item interface{}) immut.Seq { return New(item) }

// O(log n)
func (xs *Tree) AddBack(x interface{}) immut.Seq {
	return xs.AddFront(x) // same
}
func (n Empty) AddBack(item interface{}) immut.Seq { return New(item) }

// O(n*log(n))
func (xs *Tree) AddAll(that immut.Seq) immut.Seq {
	//fmt.Printf("[%d].AddAll([%d])\n", xs.Len(), that.Len())
	first, err := that.Front()
	if err != nil {
		//that is empty
		return xs
	}
	rest, _ := that.Rest() //error guaranteed to be non nil TODO: add tests for Rest
	return xs.AddFront(first).AddAll(rest)
	//TODO, avoid xs creating very unbalanced trees
}
func (n Empty) AddAll(other immut.Seq) immut.Seq { return other }

func (xs *Tree) Forall(f func(interface{}) bool) bool {
	return f(xs.value) && xs.left.Forall(f) && xs.right.Forall(f)
}
func (Empty) Forall(f func(interface{}) bool) bool { return true }

func (xs *Tree) Map(f func(interface{}) interface{}) immut.Seq {
	mappedValue := f(xs.value)
	mappedValueS := s(mappedValue)
	return &Tree{
		mappedValue,
		mappedValueS,
		xs.left.Map(f),
		xs.right.Map(f)}
}
func (n Empty) Map(f func(interface{}) interface{}) immut.Seq { return n }

func (xs *Tree) Filter(f func(interface{}) bool) immut.Seq {
	if xs.Forall(f) {
		return xs
	}
	if f(xs.value) {
		// root is included
		return &Tree{
			xs.value,
			xs.valueS,
			xs.left.Filter(f),
			xs.right.Filter(f)}
	}
	// exclude root
	if xs.left.IsEmpty() {
		if xs.right.IsEmpty() {
			return Empty{}
		}
		return xs.right.Filter(f)
	}
	if xs.right.IsEmpty() {
		return xs.left.Filter(f)
	}
	//tricky case: root is filtered out but left and right are not empty
	return xs.left.Filter(f).AddAll(xs.right.Filter(f))
}
func (n Empty) Filter(f func(interface{}) bool) immut.Seq { return n }

func (xs *Tree) String() string {
	var buf bytes.Buffer
	buf.WriteString("{")
	xs.Join(",", &buf)
	buf.WriteString("}")
	return buf.String()
}
func (Empty) String() string { return "<nil>" }

func (xs *Tree) Items() (ys []interface{}) {
	ys = make([]interface{}, xs.Len())
	i := 0
	xs.Each(func(x interface{}) {
		ys[i] = x
		i++
	})
	return
}
func (Empty) Items() []interface{} { return []interface{}{} }

//func (xs *Tree) String() string {
//	return fmt.Sprintf("(%v %v %v)", xs.left, xs.value, xs.right)
//}
