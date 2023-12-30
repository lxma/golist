package golist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Iter(t *testing.T) {
	lst := New[int]()
	result := make([]int, 0)
	for value := range lst.Iter() {
		result = append(result, value)
	}
	assert.Equal(t, []int{}, result)

	lst = MakeList(1, 2, 3)
	result = make([]int, 0)
	for value := range lst.Iter() {
		result = append(result, value)
	}
	assert.Equal(t, []int{1, 2, 3}, result)
}
