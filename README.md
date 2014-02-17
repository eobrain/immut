````go

PACKAGE DOCUMENTATION

package immut
    import "github.com/eobrain/immut"

    The immut package contains immutable structure-sharing collections in
    the style of Scala or Clojure.


````

FUNCTIONS
````go


func Join(xs Seq, sep string) string


````

TYPES
````go


type Item interface{}
    An Item is an element in a Seq.


func Back(xs Seq) (Item, error)


func Nth(xs Seq, n uint) (Item, error)


func Second(xs Seq) (Item, error)



type Seq interface {

    // Len is the number of elements. O(n) or O(log n)
    Len() int

    // Contains is whether the Item is in the Seq. O(n) or O(log n)
    Contains(Item) bool

    // Front returns the first item. O(1) or O(log n)
    Front() (Item, error)

    // Rest returns new list with all except the first item. O(1) or  O(n^2 * log(n))
    Rest() (Seq, error)

    // IsEmpty is whether this is the empty seq. O(1)
    IsEmpty() bool

    // Each Apply the function to each item in the seq. O(n)
    Each(func(Item))

    // Join writes a concatentaion of the string representations
    // of the items separated by sep into the Writer. O(n)
    Join(string, *bytes.Buffer)

    // AddFront returns a new seq with the item added on to the beginning. O(1) or O(log n)
    AddFront(Item) Seq

    //O(n) or O(1) return a new seq with the item added on to the end
    AddBack(Item) Seq

    //return a new seq that is a concatenation of this seq with the given one
    AddAll(Seq) Seq

    //return a new seq that is the reverse of this one
    Reverse() Seq

    //whether function is true for all items, or if there are no items
    Forall(func(Item) bool) bool

    //return a new seq where each item is the result of running
    //the function on the corresponding item of this seq
    Map(func(Item) Item) Seq

    //return a new seq with a subset of the items for which the
    //function is true
    Filter(func(Item) bool) Seq
    // contains filtered or unexported methods
}
    A Seq is an immutable sequence of Items. Where multiple O(...) given,
    first is for list, second is for set (average case, assuming it is a
    balanced tree),


func List(item ...Item) Seq
    Create a new list containing the arguments


func Remove(xs Seq, x Item) Seq
    Return sequence resulting from removing the item, or the sequence itself
    if item not contained in it


func Set(item ...Item) Seq
    Create a new ordered set containing the arguments. O(n*log(n))




````

