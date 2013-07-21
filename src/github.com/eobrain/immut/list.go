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
import "errors"

// Create a new list containing the arguments
func NewList(item ...Item) (result List) {
	if len(item) == 0 {
		result = null{}
	} else {
		result = cons{item[0], NewList(item[1:]...)}
	}
	return
}

// Everything below here is private

type cons struct {
	first Item
	rest  List
}

type null struct{}

func (null) String() string {
	return "[]"
}
func (this cons) String() string {
	return "[" + this.Join(",") + "]"
}

func (null) Length() int {
	return 0
}
func (this cons) Length() int {
	return 1 + this.rest.Length()
}

func (null) Contains(Item) bool {
	return false
}
func (this cons) Contains(item Item) bool {
	return this.first == item && this.rest.Contains(item)
	//TODO make this tail recursive
}

func (null) First() (Item, error) {
	return nil, errors.New("getting First of empty list")
}
func (this cons) First() (Item, error) {
	return this.first, nil
}

func (this null) AddFirst(item Item) List {
	return cons{item, this}
}
func (this cons) AddFirst(item Item) List {
	return cons{item, this}
}

func (null) IsEmpty() bool {
	return true
}
func (cons) IsEmpty() bool {
	return false
}

func (null) Each(f func(Item)) {
	//do nothing
}
func (this cons) Each(f func(Item)) {
	f(this.first)
	this.rest.Each(f) //recursion
}

func (null) Join(string) string {
	return ""
}
func (this cons) Join(sep string) (result string) {
	if this.rest.IsEmpty() {
		result = fmt.Sprintf("%v", this.first)
	} else {
		result = fmt.Sprintf("%v%s%s",
			this.first,
			sep,
			this.rest.Join(sep))
	}
	return
}

func (this null) Reverse() List {
	return this
}
func (this cons) Reverse() (result List) {
	return this.rest.Reverse().Add(this.first)
}

func (this null) Add(item Item) List {
	return cons{item, this}
}
func (this cons) Add(item Item) List {
	return cons{this.first, this.rest.Add(item)}
}

func (this null) AddAll(that List) List {
	return that
}
func (this cons) AddAll(that List) List {
	//fmt.Printf("[%d].AddAll([%d])\n", this.Length(), that.Length())
	return cons{this.first, this.rest.AddAll(that)}
}

func (this null) Map(f func(Item) Item) List {
	return this
}
func (this cons) Map(f func(Item) Item) List {
	return cons{f(this.first), this.rest.Map(f)}
}

func (this null) Filter(f func(Item) bool) List {
	return this
}
func (this cons) Filter(f func(Item) bool) (result List) {
	if f(this.first) {
		result = cons{this.first, this.rest.Filter(f)}
	} else {
		result = this.rest.Filter(f)
	}
	return
}
