GO language immutable structure-sharing collection classes
==========================================================

PACKAGE
-----

package immut

    import "github.com/eobrain/immut"

    Immutable Structure-Sharing Types

TYPES
-----

type Item interface{}

    An item in the seq

type Seq interface {

    //O(n) return number of elements
    Len() int
    //O(n) or O(log(n)) whether item is in seq
    Contains(Item) bool
    //O(1) or O(log(n)) return first item, or an error if seq is empty
    Front() (Item, error)
    //O(1) or O(???) return new list with all except the first item
    //or an error if seq is empty
    Rest() (Seq, error)
    //O(1) is this the empty seq
    IsEmpty() bool
    //O(n) Apply the function to each item in the seq
    Each(func(Item))
    //O(???) Return a concatentaion of the string representations of the items separated by sep
    Join(sep string) string
    //O(n) or O(???) return a new seq with the item added on to the end
    Add(Item) Seq
    //return a new seq that is a concatenation of this seq with the given one
    AddAll(Seq) Seq

    //whether function is true for all items, or if there are no items
    Forall(func(Item) bool) bool

    //return a new seq where each item is the result of running the function on the corresponding item of this seq
    Map(func(Item) Item) Seq
    //return a new seq with a subset of the items for which the function is true
    Filter(func(Item) bool) Seq
    // contains filtered or unexported methods
}

    An immutable sequence of Items Where multiple O(...) given, first is for
    list, second is for tree set

func List(item ...Item) Seq

    Create a new list containing the arguments

func Remove(xs Seq, x Item) Seq

    Return sequence resulting from removing the item, or the sequence itself
    if item not contained in it

func Set(item ...Item) Seq

    Create a new set containing the arguments


