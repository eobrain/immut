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

import "fmt"

// Create a new list containing the arguments
func List(item ...Item) Seq {
	if len(item) == 0 {
		return null{}
	}
	return cons{item[0], List(item[1:]...)}
}

// Everything below here is private

type cons struct {
	first Item
	rest  Seq
}

func (this cons) Length() int {
	return 1 + this.rest.Length()
}

func (this cons) Contains(item Item) bool {
	return this.first == item || this.rest.Contains(item)
	//TODO make this tail recursive
}

func (this cons) First() (Item, error) {
	return this.first, nil
}

func (this cons) Rest() (Seq, error) {
	return this.rest, nil
}

func (cons) IsEmpty() bool {
	return false
}

func (this cons) Each(f func(Item)) {
	f(this.first)
	this.rest.Each(f) //recursion
}
func (this cons) Join(sep string) string {
	if this.rest.IsEmpty() {
		return fmt.Sprintf("%v", this.first)
	}
	return fmt.Sprintf("%v%s%s",
		this.first,
		sep,
		this.rest.Join(sep))
}

//func (this cons) Reverse() Seq {
//	return this.rest.Reverse().Add(this.first)
//}

// Add to beginning
func (this cons) Add(item Item) Seq {
	return cons{item, this}
}

/*func (this cons) AddLast(item Item) Seq {
	return cons{this.first, this.rest.Add(item)}
}*/

func (this cons) AddAll(that Seq) Seq {
	//fmt.Printf("[%d].AddAll([%d])\n", this.Length(), that.Length())
	return cons{this.first, this.rest.AddAll(that)}
}

func (this cons) Forall(f func(Item) bool) bool {
	return f(this.first) && this.rest.Forall(f)
}

func (this cons) Map(f func(Item) Item) Seq {
	return cons{f(this.first), this.rest.Map(f)}
}

func (this cons) Filter(f func(Item) bool) Seq {
	if f(this.first) {
		return cons{this.first, this.rest.Filter(f)}
	}
	return this.rest.Filter(f)
}

func (this cons) String() string {
	return "[" + this.Join(",") + "]"
}

func (this cons) addTreeNode(item Item, itemS string) tree {
	return null{}.addTreeNode(item, itemS)
}
