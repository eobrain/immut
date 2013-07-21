package immut

import "fmt"

type Item interface{}

type List interface{

	IsEmpty() bool
	Each(f func(Item))
	Join(sep string) string

	Append(Item) List
	Reverse() List

	Map(func(Item) Item) List
	Filter(func(Item) bool) List

	//Pop() Item
	//Max( /*iterator ?? */ ) Item

}



type Slice []Item

type Cons struct {
	first Item
	rest List
}

type Nil struct{}


func (Nil) IsEmpty() bool{
	return true
}
func (Cons) IsEmpty() bool{
	return false
}
func (this Slice) IsEmpty() bool{
	return len(this)==0
}


func (Nil) Each(f func(Item)) {
	//do nothing
}
func (this Cons) Each(f func(Item)) {
	f(this.first)
	this.rest.Each(f) //recursion
}
func (this Slice) Each(f func(Item)) {
	for _, item := range this {
		f(item)
	}
}


func (Nil) Join(string) string{
	return ""
}
func (this Cons) Join(sep string) (result string){
	//fmt.Printf("%v.Join(%s,)\n", this, sep)
	//return ToString(this.first) + sep + this.rest.Join(sep, ToString)
	if this.rest.IsEmpty() {
		result = fmt.Sprintf("%v", this.first)
	}else{
		result = fmt.Sprintf("%v%s%s", this.first, sep, this.rest.Join(sep))
	}
	return
}
func (this Slice) Join(sep string) (result string){
	switch len(this) {
	case 0:
		result = ""
	case 1:
		result = fmt.Sprintf("%v", this[0])
	default:
		result = fmt.Sprintf("%v%s%s", this[0], sep, this[1:].Join(sep))
	}
	return
}


func (this Nil) Reverse() List{
	return this
}
func (this Cons) Reverse() (result List){
	return this.rest.Reverse().Append(this.first)
}
func (this Slice) Reverse() List{
	m := len(this)
	result := make([]Item, m)
	for i,item := range this {
		result[m-i] = item
	}
	return Slice(result)
}


func (this Nil) Append(item Item) List {
	return Cons{item, this}
}
func (this Cons) Append(item Item) List {
	return Cons{this.first, this.rest.Append(item)}
}
func (this Slice) Append(item Item) List {
	m := len(this)
	result := make([]Item, m+1)
	copy(result,this)
	result[m] = item
	return Slice(result)
}


func (this Nil) Map(f func(Item) Item) List{
	return this
}
func (this Cons) Map(f func(Item) Item) List{
	return Cons{f(this.first), this.rest.Map(f)}
}
func (this Slice) Map(f func(Item) Item) List{
	m := len(this)
	result := make([]Item, m)
	for i,item := range this {
		result[i] = f(item)
	}
	return Slice(result)
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
func (this Slice) Filter(f func(Item) bool) (result List){
	if this.IsEmpty() {
		result = this
	}else if f(this[0]){
		result = append(this[0:1], this[1:].Filter(f))
	}else{
		result = this[1:].Filter(f)
	}
	return
}

func NewConsList(item ...Item) (result List) {
	if len(item)==0 {
		result = Nil{}
	} else {
		result = Cons{item[0], NewConsList(item[1:]...)}
	}
	return
}

func NewSlice(item ...Item) (result List) {
	return Slice(item)
}
