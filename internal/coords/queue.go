package coords

type node struct {
	coords Coords
	next   *node
}

type Queue struct {
	head *node
	tail *node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(c Coords) {
	newNode := &node{
		coords: c,
	}
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
		return
	}

	q.tail.next = newNode
	q.tail = newNode
}

func (q *Queue) Pop() (c Coords) {
	c = q.head.coords
	q.head = q.head.next
	return c
}

func (q *Queue) PeekAll() (c []Coords) {
	for n := q.head; n != nil; n = n.next {
		c = append(c, n.coords)
	}
	return c
}

func (q *Queue) PeekLast() Coords {
	return q.tail.coords
}
