package test

import "github.com/eobrain/immut"
import "testing"
import "fmt"

var empty = immut.NewList()
var ints = immut.NewList(1, 2, 3)
var strings = immut.NewList("one", "two", "three", "four")

func TestIsEmpty(t *testing.T) {
	if !empty.IsEmpty() {
		t.FailNow()
	}
	if ints.IsEmpty() {
		t.FailNow()
	}
}

func TestLength(t *testing.T) {
	if empty.Length() != 0 {
		t.FailNow()
	}
	if ints.Length() != 3 {
		t.FailNow()
	}
}

func TestFirst(t *testing.T) {
	firstStrings, errStrings := strings.First()
	if errStrings != nil {
		t.FailNow()
	}
	if firstStrings != "one" {
		t.FailNow()
	}

	firstInts, errInts := ints.First()
	if errInts != nil {
		t.FailNow()
	}
	if firstInts != 1 {
		t.FailNow()
	}

	firstEmpty, errEmpty := empty.First()
	if errEmpty == nil {
		t.FailNow()
	}
	if firstEmpty != nil {
		t.FailNow()
	}
}

func BenchmarkNilIsEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		empty.IsEmpty()
	}
}
func BenchmarkIntsIsEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ints.IsEmpty()
	}
}

func ExampleJoin() {
	fmt.Println(strings.Join("|"))
	fmt.Println(ints.Join(" <--> "))
	// Output:
	// one|two|three|four
	// 1 <--> 2 <--> 3
}
