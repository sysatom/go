package container

type Queue struct {
	size  int
	items []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(element int) {
	q.size++
	q.items = append(q.items, element)
}

func (q *Queue) Dequeue() int {
	if q.size == 0 {
		return -1
	}
	q.size--
	first := q.items[0]
	q.items = q.items[1:]
	return first
}

func (q *Queue) Size() int {
	return q.size
}
