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

// Note, no attempt to keep this binary tree balanced

// Create a new ordered set containing the arguments. O(n*log(n))
func Set(item ...Item) Seq {
	if len(item) == 0 {
		return null{}
	}
	first := item[0]
	return (&tree{first, s(first), null{}, null{}}).buildTreeFrom(
		item[1:])
}

// Everything below here is private

// Recursively build a binary tree. O(n*log(n))
func (xs *tree) buildTreeFrom(remaining []Item) *tree {
	if len(remaining) == 0 {
		return xs
	}
	x := remaining[0]
	return xs.addTreeNode(x, s(x)).buildTreeFrom(remaining[1:])
}

type tree struct {
	value  Item
	valueS string //hack: use string compare for ordering
	left   Seq
	right  Seq
}

func s(x Item) string {
	return fmt.Sprintf("%v", x)
}

func (xs *tree) Len() int {
	return 1 + xs.left.Len() + xs.right.Len()
}

func (xs *tree) Contains(x Item) bool {
	itemS := s(x) //inefficiently re-creating on every recursion
	return x == xs.value ||
		itemS < xs.valueS && xs.left.Contains(x) ||
		xs.right.Contains(x)
}

func (xs *tree) Front() (Item, error) {
	if xs.left.IsEmpty() {
		return xs.value, nil
	}
	return xs.left.Front()
}

// O(n^2 * log(n))
func (xs *tree) Rest() (Seq, error) {
	if xs.left.IsEmpty() {
		return xs.right, nil
	}
	// Perhaps not most efficient
	leftRest, _ := xs.left.Rest() // guaranteed not empty
	return leftRest.AddFront(xs.value).AddAll(xs.right), nil
}

func (xs *tree) IsEmpty() bool {
	//log.Printf("%v.IsEmpty()", xs)
	return false
}

func (xs *tree) Each(f func(Item)) {
	xs.left.Each(f)
	f(xs.value)
	xs.right.Each(f)
}

func (xs *tree) Join(sep string, buf *bytes.Buffer) {
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

//func (xs *tree) Join(sep string) string {
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
func (xs *tree) addTreeNode(x Item, itemS string) *tree {
	if x == xs.value {
		//set semantics -- cannnot have more than one of any value
		return xs
	}
	//hack: use string compare for ordering
	if itemS < xs.valueS {
		//put on left
		return &tree{xs.value,
			xs.valueS,
			xs.left.addTreeNode(x, itemS),
			xs.right}
	}
	//put on right
	return &tree{xs.value,
		xs.valueS,
		xs.left,
		xs.right.addTreeNode(x, itemS)}
}

// Cannot reverse a sorted set, so just return the set itself
func (xs *tree) Reverse() Seq {
	return xs
}

// O(log n)
func (xs *tree) AddFront(x Item) Seq {
	//log.Printf("%v.Add(%v)\n", xs, x)
	return xs.addTreeNode(x, s(x))
}

func (xs *tree) AddBack(x Item) Seq {
	return xs.AddFront(x) // same
}

// O(n*log(n))
func (xs *tree) AddAll(that Seq) Seq {
	//fmt.Printf("[%d].AddAll([%d])\n", xs.Len(), that.Len())
	first, err := that.Front()
	if err != nil {
		//that is empty
		return xs
	}
	rest, _ := that.Rest() //error guaranteed to be non null TODO: add tests for Rest
	return xs.AddFront(first).AddAll(rest)
	//TODO, avoid xs creating very unbalanced trees
}

func (xs *tree) Forall(f func(Item) bool) bool {
	return f(xs.value) && xs.left.Forall(f) && xs.right.Forall(f)
}

func (xs *tree) Map(f func(Item) Item) Seq {
	mappedValue := f(xs.value)
	mappedValueS := s(mappedValue)
	return &tree{
		mappedValue,
		mappedValueS,
		xs.left.Map(f),
		xs.right.Map(f)}

}

func (xs *tree) Filter(f func(Item) bool) Seq {
	if xs.Forall(f) {
		return xs
	}
	if f(xs.value) {
		// root is included
		return &tree{
			xs.value,
			xs.valueS,
			xs.left.Filter(f),
			xs.right.Filter(f)}
	}
	// exclude root
	if xs.left.IsEmpty() {
		if xs.right.IsEmpty() {
			return null{}
		}
		return xs.right.Filter(f)
	}
	if xs.right.IsEmpty() {
		return xs.left.Filter(f)
	}
	//tricky case: root is filtered out but left and right are not null
	return xs.left.Filter(f).AddAll(xs.right.Filter(f))
}
func (xs *tree) String() string {
	var buf bytes.Buffer
	buf.WriteString("{")
	xs.Join(",", &buf)
	buf.WriteString("}")
	return buf.String()
}

//func (xs *tree) String() string {
//	return fmt.Sprintf("(%v %v %v)", xs.left, xs.value, xs.right)
//}
