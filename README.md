# golist
Typed lists for Go

This is an attempt for typed lists in Go. Basically, this is about how (in my opionon)
lists should have been designed in the first place. My expectation is that sooner or later,
typed lists will be part of the Go standard library. Until then this is a substitute, as I
find working with untyped lists unnecessarily cumbersome and error prone.

The implementation merely wraps (untyped) Go standard lists. This library uses the same
methodology and method names as found there and follows the signatures of those methods
as far as they make sense (in the context of typed lists).

Some methods were added to make life a little easier (for typed lists):
```go
func (l *List[T]) PopFront() T
func (l *List[T]) PopBack() T 
```
to get the first (rsp. last) value of a list and remove the corresponding element. Both methods will panic
in case the List is empty.
```go
func MakeList[T any](values ...T) *List[T]
func (l *List[T]) ToSlice() []T 
```
to create a list from a slice (or from single values) and to turn a list into a slice.

The method `Value()` was added to `Element` to actually retrieve a value from an element.

# Iterator
In addition, a Java style iterator is provided (also type-safe). It can be used like:
```go
iter := MakeList(1, 2, 3, 4).Iterator()
for iter.HasNext() {
    fmt.Printf("%d\n", iter.Next())
}
```
Java style iterators are preferred over `Iter()` returning a channel as we concern the latter
to be prone to create stale goroutines.