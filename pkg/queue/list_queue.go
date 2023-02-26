package queue

type node struct {
	v    int
	next *node
}

type Queue struct {
	head *node
	tail *node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(v int) {
	newNode := &node{
		v: v,
	}
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
		return
	}

	q.tail.next = newNode
	q.tail = newNode
}

func (q *Queue) Pop() (v int) {
	v = q.head.v
	q.head = q.head.next
	return v
}

func (q *Queue) values() (values []int) {
	for n := q.head; n != nil; n = n.next {
		values = append(values, n.v)
	}
	return values
}
