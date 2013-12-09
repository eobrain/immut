//Immutable structure-sharing types
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

// Note no attempt to keep this binary tree balanced

//recursively buyild a binary tree
func (this tree) buildTreeFrom(remaining []Item) tree {
	if len(remaining) == 0 {
		return this
	}
	item := remaining[0]
	return this.addTreeNode(item, s(item)).buildTreeFrom(remaining[1:])
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

func s(item Item) string {
	return fmt.Sprintf("%v", item)
}

func (this tree) Length() int {
	return 1 + this.left.Length() + this.right.Length()
}

func (this tree) Contains(item Item) bool {
	itemS := s(item) //inefficiently re-creating on every recursion
	return item == this.value ||
		itemS < this.valueS && this.left.Contains(item) ||
		this.right.Contains(item)
}

func (this tree) First() (Item, error) {
	if this.left.IsEmpty() {
		return this.value, nil
	}
	return this.left.First()
}

func (this tree) Rest() (Seq, error) {
	//log.Printf("%v.Rest()\n", this)
	if this.left.IsEmpty() {
		//log.Printf("%v returning right %v", this, this.right)
		return this.right, nil
	}
	// Perhaps not most efficient
	leftRest, _ := this.left.Rest() // guaranteed not empty
	return leftRest.Add(this.value).AddAll(this.right), nil
}

func (this tree) IsEmpty() bool {
	//log.Printf("%v.IsEmpty()", this)
	return false
}

func (this tree) Each(f func(Item)) {
	this.left.Each(f)
	f(this.value)
	this.right.Each(f)
}
func (this tree) Join(sep string) string {
	//TODO: make more eficient http://stackoverflow.com/a/1766304/978525
	if this.left.IsEmpty() {
		if this.right.IsEmpty() {
			return this.valueS
		}
		return this.valueS + sep + this.right.Join(sep)
	}
	if this.right.IsEmpty() {
		return this.left.Join(sep) + sep + this.valueS
	}
	return this.left.Join(sep) + sep +
		this.valueS + sep +
		this.right.Join(sep)

}

func (this tree) addTreeNode(item Item, itemS string) tree {
	if item == this.value {
		//set semantics -- cannnot have more than one of any value
		return this
	}
	//hack: use string compare for ordering
	if itemS < this.valueS {
		//put on left
		return tree{this.value,
			this.valueS,
			this.left.addTreeNode(item, itemS),
			this.right}
	}
	//put on right
	return tree{this.value,
		this.valueS,
		this.left,
		this.right.addTreeNode(item, itemS)}
}

func (this tree) Add(item Item) Seq {
	return this.addTreeNode(item, s(item))
}

func (this tree) AddAll(that Seq) Seq {
	//fmt.Printf("[%d].AddAll([%d])\n", this.Length(), that.Length())
	first, err := that.First()
	if err != nil {
		//that is empty
		return this
	}
	rest, _ := that.Rest() //error guaranteed to be non null TODO: add tests for Rest
	return this.Add(first).AddAll(rest)
	//TODO, avoid this creating very unbalanced trees
}

func (this tree) Forall(f func(Item) bool) bool {
	return f(this.value) && this.left.Forall(f) && this.right.Forall(f)
}

func (this tree) Map(f func(Item) Item) Seq {
	mappedValue := f(this.value)
	mappedValueS := s(mappedValue)
	return tree{
		mappedValue,
		mappedValueS,
		this.left.Map(f),
		this.right.Map(f)}

}

func (this tree) Filter(f func(Item) bool) Seq {
	if this.Forall(f) {
		return this
	}
	if f(this.value) {
		return tree{
			this.value,
			this.valueS,
			this.left.Filter(f),
			this.right.Filter(f)}
	}
	if this.left.IsEmpty() {
		if this.right.IsEmpty() {
			return null{}
		}
		return this.right.Filter(f)
	}
	if this.right.IsEmpty() {
		return this.left.Filter(f)
	}
	//tricky case: root is filtered out but left and right are not null
	return this.left.Filter(f).AddAll(this.right.Filter(f))
}

func (this tree) String() string {
	return "[" + this.Join(",") + "]"
}

//func (this tree) String() string {
//	return fmt.Sprintf("(%v %v %v)", this.left, this.value, this.right)
//}
