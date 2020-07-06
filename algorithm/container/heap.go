package container

type MaxPriorityQueue struct {
	heap []int
	n    int
}

func NewMaxPriorityQueue() *MaxPriorityQueue {
	return &MaxPriorityQueue{heap: []int{-1}}
}

func (t *MaxPriorityQueue) Max() int {
	if len(t.heap) > 1 {
		return t.heap[1]
	}
	return -1
}

func (t *MaxPriorityQueue) Insert(val int) {
	t.n += 1
	if len(t.heap) > t.n {
		t.heap[t.n] = val
	} else {
		t.heap = append(t.heap, val)
	}

	t.swim(t.n)
}

func (t *MaxPriorityQueue) DelMax() int {
	max := t.Max()
	t.exch(1, t.n)
	t.heap[t.n] = -1 // NULL
	t.n -= 1
	t.sink(1)
	return max
}

func (t *MaxPriorityQueue) swim(k int) {
	for k > 1 && t.less(t.parent(k), k) {
		t.exch(t.parent(k), k)
		k = t.parent(k)
	}
}

func (t *MaxPriorityQueue) sink(k int) {
	for t.left(k) <= t.n {
		older := t.left(k)
		if t.right(k) <= t.n && t.less(older, t.right(k)) {
			older = t.right(k)
		}
		if t.less(older, k) {
			break
		}
		t.exch(k, older)
		k = older
	}
}

func (t *MaxPriorityQueue) exch(i, j int) {
	t.heap[i], t.heap[j] = t.heap[j], t.heap[i]
}

func (t *MaxPriorityQueue) less(i, j int) bool {
	return t.heap[i] < t.heap[j]
}

func (t *MaxPriorityQueue) parent(root int) int {
	return root / 2
}

func (t *MaxPriorityQueue) left(root int) int {
	return root * 2
}

func (t *MaxPriorityQueue) right(root int) int {
	return root*2 + 1
}


type MinPriorityQueue struct {
	heap []int
	n    int
}

func NewMinPriorityQueue() *MinPriorityQueue {
	return &MinPriorityQueue{heap: []int{-1}}
}

func (t *MinPriorityQueue) Min() int {
	if len(t.heap) > 1 {
		return t.heap[1]
	}
	return -1
}

func (t *MinPriorityQueue) Insert(val int) {
	t.n += 1
	if len(t.heap) > t.n {
		t.heap[t.n] = val
	} else {
		t.heap = append(t.heap, val)
	}

	t.swim(t.n)
}

func (t *MinPriorityQueue) DelMin() int {
	max := t.Min()
	t.exch(1, t.n)
	t.heap[t.n] = -1 // NULL
	t.n -= 1
	t.sink(1)
	return max
}

func (t *MinPriorityQueue) swim(k int) {
	for k > 1 && t.more(t.parent(k), k) {
		t.exch(t.parent(k), k)
		k = t.parent(k)
	}
}

func (t *MinPriorityQueue) sink(k int) {
	for t.left(k) <= t.n {
		older := t.left(k)
		if t.right(k) <= t.n && t.more(older, t.right(k)) {
			older = t.right(k)
		}
		if t.more(older, k) {
			break
		}
		t.exch(k, older)
		k = older
	}
}

func (t *MinPriorityQueue) exch(i, j int) {
	t.heap[i], t.heap[j] = t.heap[j], t.heap[i]
}

func (t *MinPriorityQueue) more(i, j int) bool {
	return t.heap[i] > t.heap[j]
}

func (t *MinPriorityQueue) parent(root int) int {
	return root / 2
}

func (t *MinPriorityQueue) left(root int) int {
	return root * 2
}

func (t *MinPriorityQueue) right(root int) int {
	return root*2 + 1
}