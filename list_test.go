package golist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoopThroughList(t *testing.T) {
	l := MakeList([]int{1, 2, 3}...)
	results := make([]int, 0)

	for l.Len() > 0 {
		results = append(results, l.PopFront())
	}

	assert.Equal(t, []int{1, 2, 3}, results)
}

func TestLoopThroughListBackward(t *testing.T) {
	l := MakeList([]int{1, 2, 3}...)
	results := make([]int, 0)

	for l.Len() > 0 {
		results = append(results, l.PopBack())
	}

	assert.Equal(t, []int{3, 2, 1}, results)
}
