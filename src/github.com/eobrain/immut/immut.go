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

// An item in the list
type Item interface{}

// An immutable singly-list list with structure sharing
type List interface {
	//O(n) return number of elements
	Length() int
	//O(n) whether item is in list
	Contains(Item) bool
	//O(1) return first item, or an error if list is empty
	First() (Item, error)
	//O(1) return a new list with the item prepended
	AddFirst(Item) List
	//O(1) is this the empty list
	IsEmpty() bool
	//Apply the function to each item in the list
	Each(f func(Item))
	//Return a concatentaion of the string representations of the items separated by sep
	Join(sep string) string
	//O(n) return a new list with the item added on to the end
	Add(Item) List
	//return a new list that is a concatenation of this list with the given one
	AddAll(List) List
	//return a new list that is the reverse of this one
	Reverse() List
	//return a new list where each item is the result of running the function on the corresponding item of this list
	Map(func(Item) Item) List
	//return a new list with a subset of the items for which the function is true
	Filter(func(Item) bool) List
}
