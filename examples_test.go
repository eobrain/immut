package immut_test

import (
	"fmt"
	"github.com/eobrain/immut"
	"github.com/eobrain/immut/list"
	"github.com/eobrain/immut/set"
	"github.com/eobrain/immut/slice"
)

func ExampleFront_list() {
	stooges := list.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(stooges.Front())
	// Output:
	// Larry <nil>
}

func ExampleSecond_list() {
	stooges := list.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(immut.Second(stooges))
	// Shemp <nil>
}

func ExampleBack_list() {
	stooges := list.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(immut.Back(stooges))
	// Curly <nil>
}

func ExampleNth_list() {
	stooges := list.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(immut.Nth(stooges, 2))
	// Output:
	// Moe <nil>
}

func ExampleFront_slice() {
	stooges := slice.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(stooges.Front())
	// Output:
	// Larry <nil>
}

func ExampleSecond_slice() {
	stooges := slice.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(immut.Second(stooges))
	// Shemp <nil>
}

func ExampleBack_slice() {
	stooges := slice.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(immut.Back(stooges))
	// Curly <nil>
}

func ExampleNth_slice() {
	stooges := slice.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(immut.Nth(stooges, 2))
	// Output:
	// Moe <nil>
}

func ExampleFront_set() {
	stooges := set.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(stooges.Front()) // get first alphabetically
	// Output:
	// Curly <nil>
}

func ExampleSecond_set() {
	stooges := set.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(immut.Second(stooges)) // get second alphabetically
	// Larry <nil>
}

func ExampleBack_set() {
	stooges := set.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(immut.Back(stooges)) // get last alphabetically
	// Shemp <nil>
}

func ExampleNth_set() {
	stooges := set.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(immut.Nth(stooges, 2)) // get third alphabetically
	// Output:
	// Moe <nil>
}

func ExampleRemove_list() {
	empty := list.New()
	ints := list.New(1, 2, 3)
	strings := list.New("one", "two", "three", "four")
	fmt.Println(immut.Remove(empty, 33))
	fmt.Println(immut.Remove(ints, 33))
	fmt.Println(immut.Remove(ints, "foo"))
	fmt.Println(immut.Remove(ints, 1))
	fmt.Println(immut.Remove(ints, 2))
	fmt.Println(immut.Remove(ints, 3))
	fmt.Println(immut.Remove(strings, "one"))
	fmt.Println(immut.Remove(strings, "two"))
	fmt.Println(immut.Remove(strings, "four"))
	// Output:
	// []
	// [1,2,3]
	// [1,2,3]
	// [2,3]
	// [1,3]
	// [1,2]
	// [two,three,four]
	// [one,three,four]
	// [one,two,three]
}

func ExampleRemove_slice() {
	empty := slice.New()
	ints := slice.New(1, 2, 3)
	strings := slice.New("one", "two", "three", "four")
	fmt.Println(immut.Remove(empty, 33))
	fmt.Println(immut.Remove(ints, 33))
	fmt.Println(immut.Remove(ints, "foo"))
	fmt.Println(immut.Remove(ints, 1))
	fmt.Println(immut.Remove(ints, 2))
	fmt.Println(immut.Remove(ints, 3))
	fmt.Println(immut.Remove(strings, "one"))
	fmt.Println(immut.Remove(strings, "two"))
	fmt.Println(immut.Remove(strings, "four"))
	// Output:
	// []
	// [1,2,3]
	// [1,2,3]
	// [2,3]
	// [1,3]
	// [1,2]
	// [two,three,four]
	// [one,three,four]
	// [one,two,three]
}

func ExampleRemove_set() {
	empty := set.New()
	ints := set.New(1, 2, 3)
	strings := set.New("one", "two", "three", "four")
	fmt.Println(immut.Remove(empty, 33))
	fmt.Println(immut.Remove(ints, 33))
	fmt.Println(immut.Remove(ints, "foo"))
	fmt.Println(immut.Remove(ints, 1))
	fmt.Println(immut.Remove(ints, 2))
	fmt.Println(immut.Remove(ints, 3))
	fmt.Println(immut.Remove(strings, "one"))
	fmt.Println(immut.Remove(strings, "two"))
	fmt.Println(immut.Remove(strings, "four"))
	// Output:
	// {}
	// {1,2,3}
	// {1,2,3}
	// {2,3}
	// {1,3}
	// {1,2}
	// {four,three,two}
	// {four,one,three}
	// {one,three,two}
}
