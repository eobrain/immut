package test

import "github.com/eobrain/immut"
import "testing"
import "fmt"

var empty = immut.NewConsList()
var short = immut.NewConsList(1,2,3)


func TestIsEmpty(t *testing.T) {
	if !empty.IsEmpty() {
		t.FailNow()
	}
	if short.IsEmpty() {
		t.FailNow()
	}
}


func BenchmarkNilIsEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		empty.IsEmpty()
	}
}
func BenchmarkShortIsEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		short.IsEmpty()
	}
}


func ExampleJoin() {
	three := immut.NewConsList("one","two","three")
	fmt.Println(three.Join("|"))
	fmt.Println(short.Join(" <--> "))
	// Output:
	// one|two|three
	// 1 <--> 2 <--> 3
}
