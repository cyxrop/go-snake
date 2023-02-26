package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(2)
	q.Pop()
	q.Push(3)
	q.Pop()
	q.Push(4)

	assert.Equal(t, []int{3, 4}, q.values())
}
