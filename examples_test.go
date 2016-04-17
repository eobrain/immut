package immut_test

import (
	"fmt"
	"github.com/eobrain/immut"
	"github.com/eobrain/immut/list"
	"github.com/eobrain/immut/ordered"
	"github.com/eobrain/immut/unordered"
	"github.com/eobrain/immut/vector"
	"os"
)

func ExampleIsEmpty() {
	seqs := []immut.Seq{
		list.New(),
		list.New(1, 2, 3),
		vector.New(),
		vector.New(1, 2, 3),
		ordered.New(),
		ordered.New(1, 2, 3),
		unordered.New(),
		unordered.New(1, 2, 3),
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
	// true
	// false
}

func ExampleLen() {
	seqs := []immut.Seq{
		list.New(),
		list.New(1, 2, 3),
		vector.New(),
		vector.New(1, 2, 3),
		vector.Repeat(1000000, "foo"),
		ordered.New(),
		ordered.New(1, 2, 3),
		list.New(make([]interface{}, 999)...),
		unordered.New(),
		unordered.New(1, 2, 3),
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
	// 0
	// 3
	// 1000000
}

func ExampleFront() {
	seqs := []immut.Seq{
		list.New(3, 2, 1),
		vector.New(3, 2, 1),
		ordered.New(3, 2, 1),
		unordered.New(999),
	}

	for _, xs := range seqs {
		fmt.Println(xs.Front())
	}

	// Output:
	// 3
	// 3
	// 1
	// 999
}

func ExampleGet() {
	seqs := []immut.Seq{
		vector.New(4),
		vector.New(4, 3, 2, 1),
		list.New(4),
		list.New(4, 3, 2, 1),
		ordered.New(4),
		ordered.New(4, 3, 2, 1),
		unordered.New(999),
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
	// false
}

func ExampleRest() {
	seqs := []immut.Seq{
		list.New(3, 2, 1),
		vector.New(3, 2, 1),
		ordered.New(3, 2, 1),
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

func ExampleFront_vector() {
	stooges := vector.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(stooges.Front())
	// Output:
	// Larry
}

func ExampleBack_vector() {
	stooges := vector.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(stooges.Back())
	// Curly
}

func ExampleGet_vector() {
	stooges := vector.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(stooges.Get(2))
	// Output:
	// Moe true
}

func ExampleFront_ordered() {
	stooges := ordered.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(stooges.Front()) // get first alphabetically
	// Output:
	// Curly
}

func ExampleBack_ordered() {
	stooges := ordered.New("Larry", "Shemp", "Moe", "Curly")
	fmt.Println(stooges.Back()) // get last alphabetically
	// Shemp
}

func ExampleGet_ordered() {
	stooges := ordered.New("Larry", "Shemp", "Moe", "Curly")
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

func ExampleRemove_vector() {
	empty := vector.New()
	ints := vector.New(1, 2, 3)
	strings := vector.New("one", "two", "three", "four")
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

func ExampleRemove_ordered() {
	empty := ordered.New()
	ints := ordered.New(1, 2, 3)
	strings := ordered.New("one", "two", "three", "four")
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

func ExampleRemove_unordered() {
	empty := unordered.New()
	one := unordered.New("one")
	both := unordered.New("one", "two")
	fmt.Println(empty.Remove(33))
	fmt.Println(one.Remove("nope"))
	fmt.Println(one.Remove("one"))
	fmt.Println(both.Remove("one"))
	// Output:
	// {}
	// {one}
	// {}
	// {two}
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

func ExampleAddAll_vector() {
	ints := vector.New(1, 2, 3)
	strings := vector.New("one", "two", "three", "four")
	fmt.Println(ints.AddAll(strings))
	fmt.Println(strings.AddAll(ints))
	// Output:
	// [1,2,3,one,two,three,four]
	// [one,two,three,four,1,2,3]
}

func ExampleAddAll_ordered() {
	ints := ordered.New(1, 2, 3)
	strings := ordered.New("one", "two", "three", "four")
	fmt.Println(ints.AddAll(strings))
	fmt.Println(strings.AddAll(ints))

	fmt.Println(ints.AddAll(strings))
	fmt.Println(strings.AddAll(ints))
	fmt.Println(
		ordered.New("a", "b", "c", "e", "d", "f", "g", "h").AddAll(ordered.New("X")))
	fmt.Println(
		ordered.New("X").AddAll(ordered.New("a", "b", "c", "d", "e", "g", "f", "h")))
	fmt.Println(
		ordered.New("a", "b", "c", "d", "e", "f", "g", "h").AddAll(ordered.New("X", "Y")))
	fmt.Println(
		ordered.New("X", "Y").AddAll(ordered.New("a", "b", "c", "d", "e", "f", "g", "h")))
	fmt.Println(
		ordered.New("a", "b", "c", "d", "e", "f", "g", "h").AddAll(ordered.New("X", "Y")))
	fmt.Println(
		ordered.New("X", "Y").AddAll(ordered.New("a", "b", "c", "d", "e", "f", "g", "h")))
	fmt.Println(
		ordered.New("a", "b", "c", "d", "e", "f", "g", "h").AddAll(ordered.New("X", "Y", "Z")))
	fmt.Println(
		ordered.New("X", "Y", "Z").AddAll(ordered.New("a", "b", "c", "d", "e", "f", "g", "h")))
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
	vector := vector.New("one", "two", "three", "four")
	list := list.New("one", "two", "three", "four")
	ordered := ordered.New("one", "two", "three", "four")
	unordered := unordered.New("one", "two", "three", "four")
	fmt.Println(list.AddFront("iiiii"))
	fmt.Println(vector.AddFront("iiiii"))
	fmt.Println(ordered.AddFront("iiiii"))
	fmt.Println(unordered.AddFront("iiiii").Len())
	// Output:
	// [iiiii,one,two,three,four]
	// [iiiii,one,two,three,four]
	// {four,iiiii,one,three,two}
	// 5
}

func ExampleAddFront_ordered() {
	ints := ordered.New(1, 2, 3)

	fmt.Println(ordered.New(1).AddFront(2))
	fmt.Println(ordered.New(2).AddFront(1))
	fmt.Println(ordered.New("aaa").AddFront("bbb"))
	fmt.Println(ordered.New("bbb").AddFront("aaa"))
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
	vector := vector.New("one", "two", "three", "four")
	list := list.New("one", "two", "three", "four")
	ordered := ordered.New("one", "two", "three", "four")
	unordered := unordered.New("one", "two", "three", "four")
	fmt.Println(list.AddBack("iiiii"))
	fmt.Println(vector.AddBack("iiiii"))
	fmt.Println(ordered.AddBack("iiiii"))
	fmt.Println(unordered.AddBack("iiiii").Len())
	// Output:
	// [one,two,three,four,iiiii]
	// [one,two,three,four,iiiii]
	// {four,iiiii,one,three,two}
	// 5
}

func ExampleDo_square() {
	vector := vector.New(2, 30, 40)
	list := list.New(2, 30, 40)
	ordered := ordered.New(2, 30, 40)

	printSquare := func(item interface{}) {
		i := item.(int)
		fmt.Println(i * i)
	}

	vector.Do(printSquare)
	list.Do(printSquare)
	ordered.Do(printSquare)

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

func ExampleDo_sum() {
	seqs := []immut.Seq{
		vector.New(2, 30, 40),
		list.New(2, 30, 40),
		unordered.New(2, 30, 40),
		ordered.New(2, 30, 40),
	}
	for _, seq := range seqs {
		total := 0
		seq.Do(func(x interface{}) {
			total += x.(int)
		})
		fmt.Println(total)
	}

	// Output:
	// 72
	// 72
	// 72
	// 72
}

func ExampleDoBackwards() {
	vector := vector.New(2, 30, 40)
	list := list.New(2, 30, 40)
	ordered := ordered.New(2, 30, 40)

	printSquare := func(item interface{}) {
		i := item.(int)
		fmt.Println(i * i)
	}

	vector.DoBackwards(printSquare)
	list.DoBackwards(printSquare)
	ordered.DoBackwards(printSquare)

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

func ExampleDoBackwards_sum() {
	seqs := []immut.Seq{
		vector.New(2, 30, 40),
		list.New(2, 30, 40),
		unordered.New(2, 30, 40),
		ordered.New(2, 30, 40),
	}
	for _, seq := range seqs {
		total := 0
		seq.DoBackwards(func(x interface{}) {
			total += x.(int)
		})
		fmt.Println(total)
	}

	// Output:
	// 72
	// 72
	// 72
	// 72
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

func ExampleJoin_vector() {
	strings := vector.New("one", "two", "three", "four")
	strings.Join("|", os.Stdout)
	fmt.Println()

	ints := vector.New(1, 2, 3)
	ints.Join(" <--> ", os.Stdout)
	fmt.Println()

	// Output:
	// one|two|three|four
	// 1 <--> 2 <--> 3
}

func ExampleJoin_ordered() {
	strings := ordered.New("one", "two", "three", "four")
	strings.Join("|", os.Stdout)
	fmt.Println()

	ints := ordered.New(1, 2, 3)
	ints.Join(" <--> ", os.Stdout)
	fmt.Println()

	// Output:
	// four|one|three|two
	// 1 <--> 2 <--> 3
}

func Example_sort() {
	fmt.Println(ordered.New(333, 111, 222))
	fmt.Println(ordered.New(3, 11, 222))
	fmt.Println(ordered.New(4, 900, 1600))

	// Output
	// {111,222,333}
	// {11,222,3}
	// {1600,4,900}
}

func ExampleMap_integers() {
	vector := vector.New(2, 30, 40)
	list := list.New(2, 30, 40)
	ordered := ordered.New(2, 30, 40)

	square := func(item interface{}) interface{} {
		i := item.(int)
		return i * i
	}

	fmt.Println(vector.Map(square))
	fmt.Println(list.Map(square))
	fmt.Println(ordered.Map(square)) // sort alphabetically, not numerically

	// Output:
	// [4,900,1600]
	// [4,900,1600]
	// {1600,4,900}
}

func ExampleMap_strings() {
	vector := vector.New("BBB", "AAA", "CCC")
	list := list.New("BBB", "AAA", "CCC")
	ordered := ordered.New("BBB", "AAA", "CCC")
	unordered := unordered.New("BBB", "AAA", "CCC")

	constant := func(item interface{}) interface{} { return "foo" }

	fmt.Println(vector.Map(constant))
	fmt.Println(list.Map(constant))
	fmt.Println(ordered.Map(constant))   // set semantics: just one element
	fmt.Println(unordered.Map(constant)) // set semantics: just one element

	// Output:
	// [foo,foo,foo]
	// [foo,foo,foo]
	// {foo}
	// {foo}
}

func ExampleFilter_integers() {
	vector := vector.New(2, 30, 40)
	list := list.New(2, 30, 40)
	ordered := ordered.New(2, 30, 40)

	endsWithZero := func(item interface{}) bool {
		i := item.(int)
		return i%10 == 0
	}

	fmt.Println(vector.Filter(endsWithZero))
	fmt.Println(list.Filter(endsWithZero))
	fmt.Println(ordered.Filter(endsWithZero))

	// Output:
	// [30,40]
	// [30,40]
	// {30,40}
}

func ExampleFilter_strings() {
	vector := vector.New("BBB", "AAA", "CCCCC")
	list := list.New("BBB", "AAA", "CCCCC")
	ordered := ordered.New("BBB", "AAA", "CCCCC")

	isTriple := func(item interface{}) bool { return len(item.(string)) == 3 }

	fmt.Println(vector.Filter(isTriple))
	fmt.Println(list.Filter(isTriple))
	fmt.Println(ordered.Filter(isTriple))

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

func Example_vector() {
	// Ported from http://java.ociweb.com/mark/clojure/article.html#Collections

	count := vector.New(19, "yellow", true).Len()

	reverse := vector.New(2, 4, 7).Reverse()

	mapped := vector.New(2, 4, 7).Map(func(x interface{}) interface{} {
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

func Example_ordered() {
	// Ported from http://java.ociweb.com/mark/clojure/article.html#Collections

	count := ordered.New(19, "yellow", true).Len()

	reverse := ordered.New(2, 4, 7).Reverse()

	mapped := ordered.New(2, 4, 7).Map(func(x interface{}) interface{} {
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
