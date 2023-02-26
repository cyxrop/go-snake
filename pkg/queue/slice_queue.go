package queue

type SliceQueue struct {
	queue []int
}

func NewSliceQueue() *SliceQueue {
	return &SliceQueue{}
}

func (q *SliceQueue) Pop() (v int) {
	v, q.queue = q.queue[0], q.queue[1:]
	return v
}

func (q *SliceQueue) Push(v int) {
	q.queue = append(q.queue, v)
}

func (q *SliceQueue) values() []int {
	return q.queue
}
