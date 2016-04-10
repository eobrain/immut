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
	"errors"
)

type null struct{}

func (null) String() string {
	return "<nil>"
}

func (null) Len() int {
	return 0
}

func (null) Contains(interface{}) bool {
	return false
}

func (null) Front() (interface{}, error) {
	return nil, errors.New("getting Front of empty seq")
}

func (null) Rest() (Seq, error) {
	//log.Println("null.Rest() -- ERROR: calling Rest() of empty seq")
	return nil, errors.New("getting Rest of empty seq")
}

func (null) IsEmpty() bool {
	//log.Println("null.IsEmpty()")
	return true
}

func (null) Each(f func(interface{})) {
	//do nothing
}

func (null) Join(string, *bytes.Buffer) {
	// do nothing
}

func (this null) Reverse() Seq {
	return this
}

func (this null) AddFront(item interface{}) Seq {
	return &cons{item, this}
}

func (this null) AddBack(item interface{}) Seq {
	return &cons{item, this}
}

func (this null) AddAll(that Seq) Seq {
	return that
}

func (this null) Forall(f func(interface{}) bool) bool {
	return true
}

func (this null) Map(f func(interface{}) interface{}) Seq {
	return this
}

func (this null) Filter(f func(interface{}) bool) Seq {
	return this
}

func (this null) addTreeNode(item interface{}, itemS string) *tree {
	return &tree{item, itemS, null{}, null{}}
}
