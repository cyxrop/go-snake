package queue

import "testing"

const profileNum = 1_000_000

func TestProfileSliceQueue(t *testing.T) {
	q := NewSliceQueue()
	q.Push(-1)
	for i := 0; i < profileNum; i++ {
		q.Push(i)
		q.Pop()
	}
}

func TestProfileQueue(t *testing.T) {
	q := NewQueue()
	q.Push(-1)
	for i := 0; i < profileNum; i++ {
		q.Push(i)
		q.Pop()
	}
}
