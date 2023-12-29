package golist

import "container/list"

// List is a typed wrapper around list.List
type List[T any] struct {
	lst *list.List
}

// Element is a typed wrapper around list.Element
type Element[T any] struct {
	elt *list.Element
}

// Next retrieves the next element from a list.
func (e *Element[T]) Next() *Element[T] {
	nextElt := e.elt.Next()
	if nextElt == nil {
		return nil
	} else {
		return &Element[T]{nextElt}
	}
}

// Prev retrievs the previous element from a list
func (e *Element[T]) Prev() *Element[T] {
	prevElt := e.elt.Prev()
	if prevElt == nil {
		return nil
	} else {
		return &Element[T]{prevElt}
	}
}

// Value retrieves the value from a list (in a typed manner)
func (e *Element[T]) Value() T {
	return e.elt.Value.(T)
}

// MakeList creates a list from a slice.
func MakeList[T any](values ...T) *List[T] {
	resultList := List[T]{lst: list.New()}
	for _, val := range values {
		resultList.lst.PushBack(val)
	}
	return &resultList
}

// Back returns the last element from a list
func (l *List[T]) Back() *Element[T] {
	backElt := l.lst.Back()
	if backElt == nil {
		return nil
	} else {
		return &Element[T]{backElt}
	}
}

// Front returns the first element from a list
func (l *List[T]) Front() *Element[T] {
	frontElt := l.lst.Front()
	if frontElt == nil {
		return nil
	} else {
		return &Element[T]{frontElt}
	}
}

// Init erases & initializes the current list
func (l *List[T]) Init() *List[T] {
	l.lst.Init()
	return l
}

// InsertAfter inserts a new element with value v to a list after the element mark
func (l *List[T]) InsertAfter(v T, mark *Element[T]) *Element[T] {
	return &Element[T]{l.lst.InsertAfter(v, mark.elt)}
}

// InsertAfter inserts a new element with value v to a list before the element mark
func (l *List[T]) InsertBefore(v T, mark *Element[T]) *Element[T] {
	return &Element[T]{l.lst.InsertBefore(v, mark.elt)}
}

// Len returns the length of a list
func (l *List[T]) Len() int {
	return l.lst.Len()
}

// MoveAfter moves a given element e after another element mark. e and mark must not be nil.
func (l *List[T]) MoveAfter(e, mark *Element[T]) {
	l.lst.MoveAfter(e.elt, mark.elt)
}

// MoveBefore moves a given element e before another element mark. e and mark must not be nil.
func (l *List[T]) MoveBefore(e, mark *Element[T]) {
	l.lst.MoveBefore(e.elt, mark.elt)
}

// MoveToBack moves e to the back of the list. If e is not an element of the list, nothing is changed
func (l *List[T]) MoveToBack(e *Element[T]) {
	if e != nil {
		l.lst.MoveToBack(e.elt)
	}
}

// MoveToBack moves e to the front of the list.
// If e is not an element of the list, nothing is changed
func (l *List[T]) MoveToFront(e *Element[T]) {
	if e != nil {
		l.lst.MoveToFront(e.elt)
	}
}

// PushBack creates a new element with value v at the end of the list and returns it.
func (l *List[T]) PushBack(v T) *Element[T] {
	return &Element[T]{l.lst.PushBack(v)}
}

// PushBackList appends values of another list to a given list. The other
// list is not changed.
func (l *List[T]) PushBackList(other *List[T]) {
	l.lst.PushBackList(other.lst)
}

// PushFront creates a new element with value v at the start of the list and returns it.
func (l *List[T]) PushFront(v T) *Element[T] {
	return &Element[T]{l.lst.PushFront(v)}
}

// PushFrontList puts values of another list to the front of a given list. The other
// list is not changed. The other list must not be nil.
func (l *List[T]) PushFrontList(other *List[T]) {
	l.lst.PushFrontList(other.lst)
}

// Remove removes an element from a list. The element must not be nil.
func (l *List[T]) Remove(e *Element[T]) T {
	return l.lst.Remove(e.elt).(T)
}

// PopFront Removes the first element of a list and returns its value.
// The list must not be empty.
func (l *List[T]) PopFront() T {
	return l.lst.Remove(l.lst.Front()).(T)
}

// PopBack Removes the last element of a list and returns its value.
// The list must not be empty.
func (l *List[T]) PopBack() T {
	return l.lst.Remove(l.lst.Back()).(T)
}

// New Creates a new list
func New[T any]() *List[T] {
	return &List[T]{list.New()}
}

// ToSlice returns a slice of the values of the elements of a list.
func (l *List[T]) ToSlice() []T {
	result := make([]T, l.Len())
	i := 0
	elt := l.Front()
	for elt != nil {
		result[i] = elt.Value()
		i++
		elt = elt.Next()
	}
	return result
}
