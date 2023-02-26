package queue

import "testing"

func BenchmarkSliceQueue(b *testing.B) {
	q := NewSliceQueue()
	q.Push(-1)
	for i := 0; i < b.N; i++ {
		q.Push(i)
		q.Pop()
	}
}

func BenchmarkQueue(b *testing.B) {
	q := NewQueue()
	q.Push(-1)
	for i := 0; i < b.N; i++ {
		q.Push(i)
		q.Pop()
	}
}
