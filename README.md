GO language immutable structure-sharing collection classes
==========================================================

    PACKAGE
    
    package immut
        import "github.com/eobrain/immut"
    
        Immutable structure-sharing types
    
    
        Immutable structure-sharing types
    
        Immutable structure-sharing types
    
    TYPES
    
    type Item interface{}
        An item in the list
    
    type List interface {
        //O(n) return number of elements
        Length() int
        //O(n) whether item is in list
        Contains(Item) bool
        //O(1) return first item, or an error if list is empty
        First() (Item, error)
        //O(1) return a new list with the item prepended
        AddFirst(Item) List
        //O(1) is this the empty list
        IsEmpty() bool
        //Apply the function to each item in the list
        Each(f func(Item))
        //Return a concatentaion of the string representations of the items separated by sep
        Join(sep string) string
        //O(n) return a new list with the item added on to the end
        Add(Item) List
        //return a new list that is a concatenation of this list with the given one
        AddAll(List) List
        //return a new list that is the reverse of this one
        Reverse() List
        //return a new list where each item is the result of running the function on the corresponding item of this list
        Map(func(Item) Item) List
        //return a new list with a subset of the items for which the function is true
        Filter(func(Item) bool) List
    }
        An immutable singly-list list with structure sharing
    
    func NewList(item ...Item) (result List)
        Create a new list containing the arguments
    
    
