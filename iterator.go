package golist

import "container/list"

// Iterator sets ground for a Java-Style iterator for lists.
// permit iteration.
//
// Example:
//
//	iter := MakeList(1,2,3).Iterator()
//	for iter.HasNext() {
//	   fmt.Print(iter.Next())
//	}
type Iterator[T any] interface {
	// ForEachRemaining calls a given function for each of the remaining elements
	ForEachRemaining(func(T))

	// HasNext returns true if there is an element to read from using Next
	HasNext() bool

	// Next returns the current element and proceeds the iterator to the next element
	Next() T

	// Remove removes the last returned element from the underlying list. Example:
	//
	//    iter := MakeList(1, 2, 3).Iterator()
	//    for iter.HasNext() {
	//        if iter.Next() == 2 {
	//            iter.Remove()
	//        }
	//    }
	//
	// leaves [1,3]
	Remove()
}

type listIterator[T any] struct {
	element *list.Element
	list    *list.List
}

// Iterator creates a Java style iterator for the given list.
// It should be used like:
//
//	iter := MakeList(1,2,3).Iterator()
//	for iter.HasNext() {
//	   fmt.Print(iter.Next())
//	}
func (l *List[T]) Iterator() Iterator[T] {
	front := l.lst.Front()
	return &listIterator[T]{front, l.lst}
}

func (iter *listIterator[T]) ForEachRemaining(f func(T)) {
	for iter.element != nil {
		f(iter.element.Value.(T))
		iter.element = iter.element.Next()
	}
}

func (iter *listIterator[T]) HasNext() bool {
	return iter.element != nil
}

func (iter *listIterator[T]) Next() T {
	if iter.element == nil {
		panic("List has no next element")
	}
	val := iter.element.Value.(T)
	iter.element = iter.element.Next()
	return val
}

func (iter *listIterator[T]) Remove() {
	if iter.element != nil {
		iter.list.Remove(iter.element.Prev())
	} else if iter.list.Len() > 0 {
		iter.list.Remove(iter.list.Back())
	}
}
