package golist

//
// import "container/list"
//
// type List[T any] interface {
// 	Back() Element[T]
// 	Front() Element[T]
// 	Init() List[T]
// 	InsertAfter(v any, mark Element[T]) Element[T]
// 	InsertBefore(v any, mark Element[T]) Element[T]
// 	Len() int
// 	MoveAfter(e, mark Element[T])
// 	MoveBefore(e, mark Element[T])
// 	MoveToBack(e Element[T])
// 	MoveToFront(e Element[T])
// 	PushBack(v any) Element[T]
// 	PushBackList(other List[T])
// 	PushFront(v any) Element[T]
// 	PushFrontList(other List[T])
// 	PopFront() T
// 	PopBack() T
// 	Remove(e Element[T]) any
// }
//
// type Element[T any] interface {
// 	Next() Element[T]
// 	Prev() Element[T]
// }
//
// func MakeList[T any](values ...T) List[T] {
// 	resultList := intList[T]{lst: list.New()}
// 	for _, val := range values {
// 		resultList.lst.PushBack(val)
// 	}
// 	return &resultList
// }
//
// type intList[T any] struct {
// 	lst *list.List
// }
//
// type intElement[T any] struct {
// 	elt *list.Element
// }
//
// func (e *intElement[T]) Next() Element[T] {
// 	return &intElement[T]{e.elt.Next()}
// }
//
// func (e *intElement[T]) Prev() Element[T] {
// 	return &intElement[T]{e.elt.Prev()}
// }
//
// func (l *intList[T]) Back() Element[T] {
// 	return &intElement[T]{l.lst.Back()}
// }
//
// func (l *intList[T]) Front() Element[T] {
// 	return &intElement[T]{l.lst.Back()}
// }
//
// func (l *intList[T]) Init() List[T] {
// 	l.lst.Init()
// 	return l
// }
//
// func (l *intList[T]) InsertAfter(v any, mark Element[T]) Element[T] {
// 	return &intElement[T]{l.lst.InsertAfter(v, mark.(*intElement[T]).elt)}
// }
//
// func (l *intList[T]) InsertBefore(v any, mark Element[T]) Element[T] {
// 	return &intElement[T]{l.lst.InsertBefore(v, mark.(*intElement[T]).elt)}
// }
//
// func (l *intList[T]) Len() int {
// 	return l.lst.Len()
// }
//
// func (l *intList[T]) MoveAfter(e, mark Element[T]) {
// 	l.lst.MoveAfter(e.(*intElement[T]).elt, mark.(*intElement[T]).elt)
// }
//
// func (l *intList[T]) MoveBefore(e, mark Element[T]) {
// 	l.lst.MoveBefore(e.(*intElement[T]).elt, mark.(*intElement[T]).elt)
// }
//
// func (l *intList[T]) MoveToBack(e Element[T]) {
// 	l.lst.MoveToBack(e.(*intElement[T]).elt)
// }
//
// func (l *intList[T]) MoveToFront(e Element[T]) {
// 	l.lst.MoveToFront(e.(*intElement[T]).elt)
// }
//
// func (l *intList[T]) PushBack(v any) Element[T] {
// 	return &intElement[T]{l.lst.PushBack(v)}
// }
//
// func (l *intList[T]) PushBackList(other List[T]) {
// 	l.lst.PushBackList(other.(*intList[T]).lst)
// }
//
// func (l *intList[T]) PushFront(v any) Element[T] {
// 	return &intElement[T]{l.lst.PushFront(v)}
// }
//
// func (l *intList[T]) PushFrontList(other List[T]) {
// 	l.lst.PushFrontList(other.(*intList[T]).lst)
// }
//
// func (l *intList[T]) Remove(e Element[T]) any {
// 	return l.lst.Remove(e.(*intElement[T]).elt).(T)
// }
//
// func (i *intList[T]) PopFront() T {
// 	return i.lst.Remove(i.lst.Front()).(T)
// }
//
// func (i *intList[T]) PopBack() T {
// 	return i.lst.Remove(i.lst.Back()).(T)
// }
//
// func New[T any]() List[T] {
// 	return &intList[T]{list.New()}
// }
