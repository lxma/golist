package golist

// Iter which returns a channel that delivers the elements of a list to
// permit iteration.
//
// Note: the list must not be modified during this iteration.
//
// Example:
//
//	lst := MakeList(1,2,3)
//	for value := range lst.Iter() {
//	   fmt.Print(value)
//	}
func (l *List[T]) Iter() chan T {
	c := make(chan T)
	go func() {
		elt := l.lst.Front()
		for elt != nil {
			c <- elt.Value.(T)
			elt = elt.Next()
		}
		close(c)
	}()
	return c
}
