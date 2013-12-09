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

// An item in the seq
type Item interface{}

// An immutable sequence of Items
// Where multiple O(...) given, first is for list, second is for tree set
type Seq interface {
	//O(n) return number of elements
	Length() int
	//O(n) or O(log(n)) whether item is in seq
	Contains(Item) bool
	//O(1) or O(log(n)) return first item, or an error if seq is empty
	First() (Item, error)
	//O(1) or O(???) return new list with all except the first item
	//or an error if seq is empty
	Rest() (Seq, error)
	//O(1) is this the empty seq
	IsEmpty() bool
	//O(n) Apply the function to each item in the seq
	Each(f func(Item))
	//O(???) Return a concatentaion of the string representations of the items separated by sep
	Join(sep string) string
	//O(n) or O(???) return a new seq with the item added on to the end
	Add(Item) Seq
	//return a new seq that is a concatenation of this seq with the given one
	AddAll(Seq) Seq
	//return a new seq that is the reverse of this one 
	//TODO move out of Seq into List
	//Reverse() Seq

	//whether function is true for all items, or if there are no items
	Forall(func(Item) bool) bool

	//return a new seq where each item is the result of running the function on the corresponding item of this seq
	Map(func(Item) Item) Seq
	//return a new seq with a subset of the items for which the function is true
	Filter(func(Item) bool) Seq

	addTreeNode(Item, string) tree
}
