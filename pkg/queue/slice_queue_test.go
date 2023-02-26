package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceQueue(t *testing.T) {
	q := NewSliceQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Pop()

	assert.Equal(t, []int{2, 3}, q.values())
}
