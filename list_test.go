package golist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicFunctionality(t *testing.T) {
	nilElt := (*Element[int])(nil)
	l := MakeList(1, 2, 3)

	elt := l.Front()
	assert.Equal(t, 1, elt.Value())
	elt = elt.Next()
	assert.Equal(t, 2, elt.Value())
	elt = elt.Prev()
	assert.Equal(t, 1, elt.Value())
	elt = l.Back()
	assert.Equal(t, 3, elt.Value())

	l.PushFront(0)
	l.PushBack(4)
	elt = l.Front()
	assert.Equal(t, 0, elt.Value())
	elt = l.Back()
	assert.Equal(t, 4, elt.Value())

	val := l.Remove(elt)
	assert.Equal(t, 4, val)
	elt = l.Back()
	assert.Equal(t, 3, elt.Value())

	l2 := l.Init()
	assert.Equal(t, 0, l.Len())
	assert.Equal(t, l, l2)

	l = MakeList(1, 2)
	l2 = MakeList(3, 4)
	l.PushBackList(l2)
	assert.Equal(t, []int{1, 2, 3, 4}, l.ToSlice())
	l.PushFrontList(l2)
	assert.Equal(t, []int{3, 4, 1, 2, 3, 4}, l.ToSlice())

	l = New[int]()
	assert.Equal(t, 0, l.Len())
	assert.Equal(t, nilElt, l.Front())
	assert.Equal(t, nilElt, l.Back())
	l.PushBack(1)
	assert.Equal(t, nilElt, l.Front().Next())
	assert.Equal(t, nilElt, l.Front().Prev())

	l = MakeList(1, 2, 3)
	l.InsertBefore(8, l.Back())
	assert.Equal(t, []int{1, 2, 8, 3}, l.ToSlice())
	l.InsertAfter(9, l.Front())
	assert.Equal(t, []int{1, 9, 2, 8, 3}, l.ToSlice())

	l = MakeList(1, 2, 3)
	l.MoveBefore(l.Front(), l.Back())
	assert.Equal(t, []int{2, 1, 3}, l.ToSlice())

	l = MakeList(1, 2, 3)
	l.MoveAfter(l.Back(), l.Front())
	assert.Equal(t, []int{1, 3, 2}, l.ToSlice())

	l = MakeList(1, 2, 3)
	l.MoveToFront(l.Back())
	assert.Equal(t, []int{3, 1, 2}, l.ToSlice())

	l = MakeList(1, 2, 3)
	l.MoveToBack(l.Front())
	assert.Equal(t, []int{2, 3, 1}, l.ToSlice())
}

func TestLoopThroughList(t *testing.T) {
	l := MakeList(1, 2, 3)
	results := make([]int, 0)

	for l.Len() > 0 {
		results = append(results, l.PopFront())
	}

	assert.Equal(t, []int{1, 2, 3}, results)
}

func TestLoopThroughListBackward(t *testing.T) {
	l := MakeList(1, 2, 3)
	results := make([]int, 0)

	for l.Len() > 0 {
		results = append(results, l.PopBack())
	}

	assert.Equal(t, []int{3, 2, 1}, results)
}
