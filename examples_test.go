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
		list.New(3, 2, 1),
		slice.New(3, 2, 1),
		set.New(3, 2, 1),
	}

	for _, xs := range seqs {
		fmt.Println(xs.Front())
	}

	// Output:
	// 3
	// 3
	// 1
}

func ExampleGet() {
	seqs := []immut.Seq{
		slice.New(4),
		slice.New(4, 3, 2, 1),
		list.New(4),
		list.New(4, 3, 2, 1),
		set.New(4),
		set.New(4, 3, 2, 1),
	}

	for _, xs := range seqs {
		fmt.Println(xs.Get(2))
	}

	// Output
	// false
	// 2 true
	// false
	// 2 true
	// false
	// 3 true
}

func ExampleRest() {
	seqs := []immut.Seq{
		list.New(3, 2, 1),
		slice.New(3, 2, 1),
		set.New(3, 2, 1),
	}

	for _, xs := range seqs {
		fmt.Println(xs.Rest())
	}

	// Output:
	// [2,1]
	// [2,1]
	// {2,3}
}

func ExampleFront_list() {
	stooges := list.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(stooges.Front())
	// Output:
	// Larry
}

func ExampleBack_list() {
	stooges := list.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(stooges.Back())
	// Curly
}

func ExampleGet_list() {
	stooges := list.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(stooges.Get(2))
	// Output:
	// Moe true
}

func ExampleFront_slice() {
	stooges := slice.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(stooges.Front())
	// Output:
	// Larry
}

func ExampleBack_slice() {
	stooges := slice.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(stooges.Back())
	// Curly
}

func ExampleGet_slice() {
	stooges := slice.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(stooges.Get(2))
	// Output:
	// Moe true
}

func ExampleFront_set() {
	stooges := set.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(stooges.Front()) // get first alphabetically
	// Output:
	// Curly
}

func ExampleBack_set() {
	stooges := set.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(stooges.Back()) // get last alphabetically
	// Shemp
}

func ExampleGet_set() {
	stooges := set.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(stooges.Get(2)) // get third alphabetically
	// Output:
	// Moe true
}

func ExampleRemove_list() {
	empty := list.New()
	ints := list.New(1, 2, 3)
	strings := list.New("one", "two", "three", "four")
	fmt.Println(empty.Remove(33))
	fmt.Println(ints.Remove(33))
	fmt.Println(ints.Remove("foo"))
	fmt.Println(ints.Remove(1))
	fmt.Println(ints.Remove(2))
	fmt.Println(ints.Remove(3))
	fmt.Println(strings.Remove("one"))
	fmt.Println(strings.Remove("two"))
	fmt.Println(strings.Remove("three"))
	fmt.Println(strings.Remove("four"))
	// Output:
	// []
	// [1,2,3]
	// [1,2,3]
	// [2,3]
	// [1,3]
	// [1,2]
	// [two,three,four]
	// [one,three,four]
	// [one,two,four]
	// [one,two,three]
}

func ExampleRemove_slice() {
	empty := slice.New()
	ints := slice.New(1, 2, 3)
	strings := slice.New("one", "two", "three", "four")
	fmt.Println(empty.Remove(33))
	fmt.Println(ints.Remove(33))
	fmt.Println(ints.Remove("foo"))
	fmt.Println(ints.Remove(1))
	fmt.Println(ints.Remove(2))
	fmt.Println(ints.Remove(3))
	fmt.Println(strings.Remove("one"))
	fmt.Println(strings.Remove("two"))
	fmt.Println(strings.Remove("four"))
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
	fmt.Println(empty.Remove(33))
	fmt.Println(ints.Remove(33))
	fmt.Println(ints.Remove("foo"))
	fmt.Println(ints.Remove(1))
	fmt.Println(ints.Remove(2))
	fmt.Println(ints.Remove(3))
	fmt.Println(strings.Remove("one"))
	fmt.Println(strings.Remove("two"))
	fmt.Println(strings.Remove("three"))
	fmt.Println(strings.Remove("four"))
	// Output:
	// {}
	// {1,2,3}
	// {1,2,3}
	// {2,3}
	// {1,3}
	// {1,2}
	// {four,three,two}
	// {four,one,three}
	// {four,one,two}
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

	fmt.Println(ints.AddAll(strings))
	fmt.Println(strings.AddAll(ints))
	fmt.Println(
		set.New("a", "b", "c", "e", "d", "f", "g", "h").AddAll(set.New("X")))
	fmt.Println(
		set.New("X").AddAll(set.New("a", "b", "c", "d", "e", "g", "f", "h")))
	fmt.Println(
		set.New("a", "b", "c", "d", "e", "f", "g", "h").AddAll(set.New("X", "Y")))
	fmt.Println(
		set.New("X", "Y").AddAll(set.New("a", "b", "c", "d", "e", "f", "g", "h")))
	fmt.Println(
		set.New("a", "b", "c", "d", "e", "f", "g", "h").AddAll(set.New("X", "Y")))
	fmt.Println(
		set.New("X", "Y").AddAll(set.New("a", "b", "c", "d", "e", "f", "g", "h")))
	fmt.Println(
		set.New("a", "b", "c", "d", "e", "f", "g", "h").AddAll(set.New("X", "Y", "Z")))
	fmt.Println(
		set.New("X", "Y", "Z").AddAll(set.New("a", "b", "c", "d", "e", "f", "g", "h")))
	// Output:
	// {1,2,3,four,one,three,two}
	// {1,2,3,four,one,three,two}
	// {1,2,3,four,one,three,two}
	// {1,2,3,four,one,three,two}
	// {X,a,b,c,d,e,f,g,h}
	// {X,a,b,c,d,e,f,g,h}
	// {X,Y,a,b,c,d,e,f,g,h}
	// {X,Y,a,b,c,d,e,f,g,h}
	// {X,Y,a,b,c,d,e,f,g,h}
	// {X,Y,a,b,c,d,e,f,g,h}
	// {X,Y,Z,a,b,c,d,e,f,g,h}
	// {X,Y,Z,a,b,c,d,e,f,g,h}
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

func ExampleAddFront_set() {
	ints := set.New(1, 2, 3)

	fmt.Println(set.New(1).AddFront(2))
	fmt.Println(set.New(2).AddFront(1))
	fmt.Println(set.New("aaa").AddFront("bbb"))
	fmt.Println(set.New("bbb").AddFront("aaa"))
	fmt.Println(ints.AddFront(1))
	fmt.Println(ints.AddFront(2))
	fmt.Println(ints.AddFront(3))
	fmt.Println(ints.AddFront(0))
	fmt.Println(ints.AddFront(4))
	// Output:
	// {1,2}
	// {1,2}
	// {aaa,bbb}
	// {aaa,bbb}
	// {1,2,3}
	// {1,2,3}
	// {1,2,3}
	// {0,1,2,3}
	// {1,2,3,4}
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

func ExampleDoBackwards() {
	slice := slice.New(2, 30, 40)
	list := list.New(2, 30, 40)
	set := set.New(2, 30, 40)

	printSquare := func(item interface{}) {
		i := item.(int)
		fmt.Println(i * i)
	}

	slice.DoBackwards(printSquare)
	list.DoBackwards(printSquare)
	set.DoBackwards(printSquare)

	// Output:
	// 1600
	// 900
	// 4
	// 1600
	// 900
	// 4
	// 1600
	// 900
	// 4
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
	fmt.Println(set.Map(constant)) // set semantics: just one element

	// Output:
	// [foo,foo,foo]
	// [foo,foo,foo]
	// {foo}
}

func ExampleFilter_integers() {
	slice := slice.New(2, 30, 40)
	list := list.New(2, 30, 40)
	set := set.New(2, 30, 40)

	endsWithZero := func(item interface{}) bool {
		i := item.(int)
		return i%10 == 0
	}

	fmt.Println(slice.Filter(endsWithZero))
	fmt.Println(list.Filter(endsWithZero))
	fmt.Println(set.Filter(endsWithZero))

	// Output:
	// [30,40]
	// [30,40]
	// {30,40}
}

func ExampleFilter_strings() {
	slice := slice.New("BBB", "AAA", "CCCCC")
	list := list.New("BBB", "AAA", "CCCCC")
	set := set.New("BBB", "AAA", "CCCCC")

	isTriple := func(item interface{}) bool { return len(item.(string)) == 3 }

	fmt.Println(slice.Filter(isTriple))
	fmt.Println(list.Filter(isTriple))
	fmt.Println(set.Filter(isTriple))

	// Output:
	// [BBB,AAA]
	// [BBB,AAA]
	// {AAA,BBB}
}

func Example_list() {
	// Ported from http://java.ociweb.com/mark/clojure/article.html#Collections

	count := list.New(19, "yellow", true).Len()

	reverse := list.New(2, 4, 7).Reverse()

	mapped := list.New(2, 4, 7).Map(func(x interface{}) interface{} {
		return x.(int) + 3
	})

	fmt.Println(count)
	fmt.Println(reverse)
	fmt.Println(mapped)

	// Output:
	// 3
	// [7,4,2]
	// [5,7,10]
}

func Example_slice() {
	// Ported from http://java.ociweb.com/mark/clojure/article.html#Collections

	count := slice.New(19, "yellow", true).Len()

	reverse := slice.New(2, 4, 7).Reverse()

	mapped := slice.New(2, 4, 7).Map(func(x interface{}) interface{} {
		return x.(int) + 3
	})

	fmt.Println(count)
	fmt.Println(reverse)
	fmt.Println(mapped)

	// Output:
	// 3
	// [7,4,2]
	// [5,7,10]
}

func Example_set() {
	// Ported from http://java.ociweb.com/mark/clojure/article.html#Collections

	count := set.New(19, "yellow", true).Len()

	reverse := set.New(2, 4, 7).Reverse()

	mapped := set.New(2, 4, 7).Map(func(x interface{}) interface{} {
		return x.(int) + 3
	})

	fmt.Println(count)
	fmt.Println(reverse)
	fmt.Println(mapped)

	// Output:
	// 3
	// {2,4,7}
	// {10,5,7}
}
