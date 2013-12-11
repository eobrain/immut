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
	"fmt"
	//"log"
)

// Note, no attempt to keep this binary tree balanced

//recursively build a binary tree
func (xs tree) buildTreeFrom(remaining []Item) tree {
	if len(remaining) == 0 {
		return xs
	}
	x := remaining[0]
	return xs.addTreeNode(x, s(x)).buildTreeFrom(remaining[1:])
}

// Create a new set containing the arguments
func Set(item ...Item) Seq {
	if len(item) == 0 {
		return null{}
	}
	first := item[0]
	return tree{first, s(first), null{}, null{}}.buildTreeFrom(
		item[1:])
}

// Everything below here is private

type tree struct {
	value  Item
	valueS string //hack: use string compare for ordering
	left   Seq
	right  Seq
}

func s(x Item) string {
	return fmt.Sprintf("%v", x)
}

func (xs tree) Len() int {
	return 1 + xs.left.Len() + xs.right.Len()
}

func (xs tree) Contains(x Item) bool {
	itemS := s(x) //inefficiently re-creating on every recursion
	return x == xs.value ||
		itemS < xs.valueS && xs.left.Contains(x) ||
		xs.right.Contains(x)
}

func (xs tree) Front() (Item, error) {
	if xs.left.IsEmpty() {
		return xs.value, nil
	}
	return xs.left.Front()
}

func (xs tree) Rest() (Seq, error) {
	//log.Printf("%v.Rest()\n", xs)
	if xs.left.IsEmpty() {
		//log.Printf("%v returning right %v", xs, xs.right)
		return xs.right, nil
	}
	// Perhaps not most efficient
	leftRest, _ := xs.left.Rest() // guaranteed not empty
	return leftRest.Add(xs.value).AddAll(xs.right), nil
}

func (xs tree) IsEmpty() bool {
	//log.Printf("%v.IsEmpty()", xs)
	return false
}

func (xs tree) Each(f func(Item)) {
	xs.left.Each(f)
	f(xs.value)
	xs.right.Each(f)
}
func (xs tree) Join(sep string) string {
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

}

func (xs tree) addTreeNode(x Item, itemS string) tree {
	if x == xs.value {
		//set semantics -- cannnot have more than one of any value
		return xs
	}
	//hack: use string compare for ordering
	if itemS < xs.valueS {
		//put on left
		return tree{xs.value,
			xs.valueS,
			xs.left.addTreeNode(x, itemS),
			xs.right}
	}
	//put on right
	return tree{xs.value,
		xs.valueS,
		xs.left,
		xs.right.addTreeNode(x, itemS)}
}

func (xs tree) Add(x Item) Seq {
	//log.Printf("%v.Add(%v)\n", xs, x)
	return xs.addTreeNode(x, s(x))
}

func (xs tree) AddAll(that Seq) Seq {
	//fmt.Printf("[%d].AddAll([%d])\n", xs.Len(), that.Len())
	first, err := that.Front()
	if err != nil {
		//that is empty
		return xs
	}
	rest, _ := that.Rest() //error guaranteed to be non null TODO: add tests for Rest
	return xs.Add(first).AddAll(rest)
	//TODO, avoid xs creating very unbalanced trees
}

func (xs tree) Forall(f func(Item) bool) bool {
	return f(xs.value) && xs.left.Forall(f) && xs.right.Forall(f)
}

func (xs tree) Map(f func(Item) Item) Seq {
	mappedValue := f(xs.value)
	mappedValueS := s(mappedValue)
	return tree{
		mappedValue,
		mappedValueS,
		xs.left.Map(f),
		xs.right.Map(f)}

}

func (xs tree) Filter(f func(Item) bool) Seq {
	if xs.Forall(f) {
		return xs
	}
	if f(xs.value) {
		// root is included
		return tree{
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
func (xs tree) String() string {
	return "[" + xs.Join(",") + "]"
}

//func (xs tree) String() string {
//	return fmt.Sprintf("(%v %v %v)", xs.left, xs.value, xs.right)
//}
