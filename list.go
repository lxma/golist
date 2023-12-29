package golist

import "container/list"

func MakeList[T any](values ...T) *List[T] {
	resultList := List[T]{lst: list.New()}
	for _, val := range values {
		resultList.lst.PushBack(val)
	}
	return &resultList
}

type List[T any] struct {
	lst *list.List
}

type Element[T any] struct {
	elt *list.Element
}

func (e *Element[T]) Next() *Element[T] {
	return &Element[T]{e.elt.Next()}
}

func (e *Element[T]) Prev() *Element[T] {
	return &Element[T]{e.elt.Prev()}
}

func (e *Element[T]) Value() T {
	return e.elt.Value.(T)
}

func (l *List[T]) Back() *Element[T] {
	return &Element[T]{l.lst.Back()}
}

func (l *List[T]) Front() *Element[T] {
	return &Element[T]{l.lst.Back()}
}

func (l *List[T]) Init() *List[T] {
	l.lst.Init()
	return l
}

func (l *List[T]) InsertAfter(v any, mark *Element[T]) *Element[T] {
	return &Element[T]{l.lst.InsertAfter(v, mark.elt)}
}

func (l *List[T]) InsertBefore(v any, mark *Element[T]) *Element[T] {
	return &Element[T]{l.lst.InsertBefore(v, mark.elt)}
}

func (l *List[T]) Len() int {
	return l.lst.Len()
}

func (l *List[T]) MoveAfter(e, mark *Element[T]) {
	l.lst.MoveAfter(e.elt, mark.elt)
}

func (l *List[T]) MoveBefore(e, mark *Element[T]) {
	l.lst.MoveBefore(e.elt, mark.elt)
}

func (l *List[T]) MoveToBack(e *Element[T]) {
	l.lst.MoveToBack(e.elt)
}

func (l *List[T]) MoveToFront(e *Element[T]) {
	l.lst.MoveToFront(e.elt)
}

func (l *List[T]) PushBack(v any) *Element[T] {
	return &Element[T]{l.lst.PushBack(v)}
}

func (l *List[T]) PushBackList(other *List[T]) {
	l.lst.PushBackList(other.lst)
}

func (l *List[T]) PushFront(v any) *Element[T] {
	return &Element[T]{l.lst.PushFront(v)}
}

func (l *List[T]) PushFrontList(other *List[T]) {
	l.lst.PushFrontList(other.lst)
}

func (l *List[T]) Remove(e *Element[T]) any {
	return l.lst.Remove(e.elt).(T)
}

func (i *List[T]) PopFront() T {
	return i.lst.Remove(i.lst.Front()).(T)
}

func (i *List[T]) PopBack() T {
	return i.lst.Remove(i.lst.Back()).(T)
}

func New[T any]() *List[T] {
	return &List[T]{list.New()}
}
