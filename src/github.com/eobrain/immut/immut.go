package immut

import "fmt"
import "errors"

type Item interface{}

type List interface{

	Length() int

	First() (Item, error)

	IsEmpty() bool
	Each(f func(Item))
	Join(sep string) string

	Add(Item) List
	Reverse() List

	Map(func(Item) Item) List
	Filter(func(Item) bool) List

}


type Cons struct {
	first Item
	rest List
}

type Nil struct{}


func (Nil) Length() int {
	return 0
}
func (this Cons) Length() int {
	return 1 + this.rest.Length()
}


func (Nil) First() (Item, error) {
	return nil, errors.New("getting First of empty list")
}
func (this Cons) First() (Item, error) {
	return this.first, nil
}


func (Nil) IsEmpty() bool{
	return true
}
func (Cons) IsEmpty() bool{
	return false
}


func (Nil) Each(f func(Item)) {
	//do nothing
}
func (this Cons) Each(f func(Item)) {
	f(this.first)
	this.rest.Each(f) //recursion
}


func (Nil) Join(string) string{
	return ""
}
func (this Cons) Join(sep string) (result string){
	if this.rest.IsEmpty() {
		result = fmt.Sprintf("%v", this.first)
	}else{
		result = fmt.Sprintf("%v%s%s",
			this.first,
			sep,
			this.rest.Join(sep))
	}
	return
}


func (this Nil) Reverse() List{
	return this
}
func (this Cons) Reverse() (result List){
	return this.rest.Reverse().Add(this.first)
}


func (this Nil) Add(item Item) List {
	return Cons{item, this}
}
func (this Cons) Add(item Item) List {
	return Cons{this.first, this.rest.Add(item)}
}


func (this Nil) Map(f func(Item) Item) List{
	return this
}
func (this Cons) Map(f func(Item) Item) List{
	return Cons{f(this.first), this.rest.Map(f)}
}


func (this Nil) Filter(f func(Item) bool) List{
	return this
}
func (this Cons) Filter(f func(Item) bool) (result List){
	if f(this.first){
		result = Cons{this.first, this.rest.Filter(f)}
	}else{
		result = this.rest.Filter(f)
	}
	return
}


func NewList(item ...Item) (result List) {
	if len(item)==0 {
		result = Nil{}
	} else {
		result = Cons{item[0], NewList(item[1:]...)}
	}
	return
}

