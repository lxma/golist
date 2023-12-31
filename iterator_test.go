package golist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Iter(t *testing.T) {
	iter := New[int]().Iterator()
	result := make([]int, 0)
	for iter.HasNext() {
		result = append(result, iter.Next())
	}
	assert.Equal(t, []int{}, result)

	iter = MakeList(1, 2, 3).Iterator()
	result = make([]int, 0)
	for iter.HasNext() {
		result = append(result, iter.Next())
	}
	assert.Equal(t, []int{1, 2, 3}, result)
}

func TestListIterator_Remove(t *testing.T) {
	lst := MakeList(1, 2, 3, 4)
	iter := lst.Iterator()
	for iter.HasNext() {
		if iter.Next() == 2 {
			iter.Remove()
		}
	}
	assert.Equal(t, []int{1, 3, 4}, lst.ToSlice())

	iter = lst.Iterator()
	for iter.HasNext() {
		if iter.Next() == 4 {
			iter.Remove()
		}
	}
	assert.Equal(t, []int{1, 3}, lst.ToSlice())
}

func TestListIterator_ForEachRemaining(t *testing.T) {
	iter := MakeList(1, 2, 3, 4).Iterator()
	slc := make([]int, 0)
	iter.Next()
	iter.ForEachRemaining(func(n int) {
		slc = append(slc, n)
	})
	assert.Equal(t, []int{2, 3, 4}, slc)

}
