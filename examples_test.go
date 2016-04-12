package immut_test

import (
	"fmt"
	"github.com/eobrain/immut"
	"github.com/eobrain/immut/list"
	"github.com/eobrain/immut/set"
	"github.com/eobrain/immut/slice"
	"os"
)

func ExampleIsEmpty() {
	seqs := []immut.Seq{
		list.New(),
		list.New(1, 2, 3),
		slice.New(),
		slice.New(1, 2, 3),
		set.New(),
		set.New(1, 2, 3),
	}
	for _, xs := range seqs {
		fmt.Println(xs.IsEmpty())
	}
	// Output:
	// true
	// false
	// true
	// false
	// true
	// false
}

func ExampleLen() {
	seqs := []immut.Seq{
		list.New(),
		list.New(1, 2, 3),
		slice.New(),
		slice.New(1, 2, 3),
		slice.Repeat(1000000, "foo"),
		set.New(),
		set.New(1, 2, 3),
		list.New(make([]interface{}, 999)...),
		list.Repeat(1000000, "foo"),
	}
	for _, xs := range seqs {
		fmt.Println(xs.Len())
	}
	// Output:
	// 0
	// 3
	// 0
	// 3
	// 1000000
	// 0
	// 3
	// 999
	// 1000000
}

func ExampleFront() {

	seqs := []immut.Seq{
		list.New(),
		list.New(3, 2, 1),
		slice.New(),
		slice.New(3, 2, 1),
		set.New(),
		set.New(3, 2, 1),
	}

	for _, xs := range seqs {
		fmt.Println(xs.Front())
	}

	// Output:
	// <nil> getting Front of empty seq
	// 3 <nil>
	// <nil> getting Front of empty seq
	// 3 <nil>
	// <nil> getting Front of empty seq
	// 1 <nil>
}

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

func ExampleAddAll_list() {
	ints := list.New(1, 2, 3)
	strings := list.New("one", "two", "three", "four")
	fmt.Println(ints.AddAll(strings))
	fmt.Println(strings.AddAll(ints))
	// Output:
	// [1,2,3,one,two,three,four]
	// [one,two,three,four,1,2,3]
}

func ExampleAddAll_slice() {
	ints := slice.New(1, 2, 3)
	strings := slice.New("one", "two", "three", "four")
	fmt.Println(ints.AddAll(strings))
	fmt.Println(strings.AddAll(ints))
	// Output:
	// [1,2,3,one,two,three,four]
	// [one,two,three,four,1,2,3]
}

func ExampleAddAll_set() {
	ints := set.New(1, 2, 3)
	strings := set.New("one", "two", "three", "four")
	fmt.Println(ints.AddAll(strings))
	fmt.Println(strings.AddAll(ints))
	// Output:
	// {1,2,3,four,one,three,two}
	// {1,2,3,four,one,three,two}
}

func ExampleAddFront() {
	slice := slice.New("one", "two", "three", "four")
	list := list.New("one", "two", "three", "four")
	set := set.New("one", "two", "three", "four")
	fmt.Println(list.AddFront("iiiii"))
	fmt.Println(slice.AddFront("iiiii"))
	fmt.Println(set.AddFront("iiiii"))
	// Output:
	// [iiiii,one,two,three,four]
	// [iiiii,one,two,three,four]
	// {four,iiiii,one,three,two}
}
func ExampleAddBack() {
	slice := slice.New("one", "two", "three", "four")
	list := list.New("one", "two", "three", "four")
	set := set.New("one", "two", "three", "four")
	fmt.Println(list.AddBack("iiiii"))
	fmt.Println(slice.AddBack("iiiii"))
	fmt.Println(set.AddBack("iiiii"))
	// Output:
	// [one,two,three,four,iiiii]
	// [one,two,three,four,iiiii]
	// {four,iiiii,one,three,two}
}

func ExampleDo() {
	slice := slice.New(2, 30, 40)
	list := list.New(2, 30, 40)
	set := set.New(2, 30, 40)

	printSquare := func(item interface{}) {
		i := item.(int)
		fmt.Println(i * i)
	}

	slice.Do(printSquare)
	list.Do(printSquare)
	set.Do(printSquare)

	// Output:
	// 4
	// 900
	// 1600
	// 4
	// 900
	// 1600
	// 4
	// 900
	// 1600
}

func ExampleJoin_list() {
	strings := list.New("one", "two", "three", "four")
	strings.Join("|", os.Stdout)
	fmt.Println()

	ints := list.New(1, 2, 3)
	ints.Join(" <--> ", os.Stdout)
	fmt.Println()

	// Output:
	// one|two|three|four
	// 1 <--> 2 <--> 3
}

func ExampleJoin_slice() {
	strings := slice.New("one", "two", "three", "four")
	strings.Join("|", os.Stdout)
	fmt.Println()

	ints := slice.New(1, 2, 3)
	ints.Join(" <--> ", os.Stdout)
	fmt.Println()

	// Output:
	// one|two|three|four
	// 1 <--> 2 <--> 3
}

func ExampleJoin_set() {
	strings := set.New("one", "two", "three", "four")
	strings.Join("|", os.Stdout)
	fmt.Println()

	ints := set.New(1, 2, 3)
	ints.Join(" <--> ", os.Stdout)
	fmt.Println()

	// Output:
	// four|one|three|two
	// 1 <--> 2 <--> 3
}

func Example_sort() {
	fmt.Println(set.New(333, 111, 222))
	fmt.Println(set.New(3, 11, 222))
	fmt.Println(set.New(4, 900, 1600))

	// Output
	// {111,222,333}
	// {11,222,3}
	// {1600,4,900}
}

func ExampleMap_integers() {
	slice := slice.New(2, 30, 40)
	list := list.New(2, 30, 40)
	set := set.New(2, 30, 40)

	square := func(item interface{}) interface{} {
		i := item.(int)
		return i * i
	}

	fmt.Println(slice.Map(square))
	fmt.Println(list.Map(square))

	// TODO(eob) Fix this. set.Map results in unsorted set
	fmt.Println(set.Map(square)) // sort alphabetically, not numerically

	// Output:
	// [4,900,1600]
	// [4,900,1600]
	// {1600,4,900}
}
func ExampleMap_strings() {
	slice := slice.New("BBB", "AAA", "CCC")
	list := list.New("BBB", "AAA", "CCC")
	set := set.New("BBB", "AAA", "CCC")

	constant := func(item interface{}) interface{} { return "foo" }

	fmt.Println(slice.Map(constant))
	fmt.Println(list.Map(constant))

	// TODO(eob) Fix this. set.Map results repeates in set
	fmt.Println(set.Map(constant)) // set semantics: just one element

	// Output:
	// [foo,foo,foo]
	// [foo,foo,foo]
	// {foo}
}
