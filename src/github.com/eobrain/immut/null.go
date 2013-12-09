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
	"errors"
	//"log"
)

type null struct{}

func (null) String() string {
	return "[]"
}

func (null) Length() int {
	return 0
}

func (null) Contains(Item) bool {
	return false
}

func (null) First() (Item, error) {
	return nil, errors.New("getting First of empty seq")
}

func (null) Rest() (Seq, error) {
	//log.Println("null.Rest() -- ERROR: calling Rest() of empty seq")
	return nil, errors.New("getting Rest of empty seq")
}

func (null) IsEmpty() bool {
	//log.Println("null.IsEmpty()")
	return true
}

func (null) Each(f func(Item)) {
	//do nothing
}

func (null) Join(string) string {
	return ""
}

//func (this null) Reverse() Seq {
//	return this
//}

func (this null) Add(item Item) Seq {
	return cons{item, this}
}

func (this null) AddAll(that Seq) Seq {
	return that
}

func (this null) Forall(f func(Item) bool) bool {
	return true
}

func (this null) Map(f func(Item) Item) Seq {
	return this
}

func (this null) Filter(f func(Item) bool) Seq {
	return this
}

func (this null) addTreeNode(item Item, itemS string) tree {
	return tree{item, itemS, null{}, null{}}
}
