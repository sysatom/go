package container

type LRUCache struct {
	m     map[int]*node
	cache DoublyLinkedList
	cap   int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{cap: capacity, m: make(map[int]*node), cache: *NewDoublyLinkedList()}
}

func (c *LRUCache) get(key int) int {
	if _, ok := c.m[key]; !ok {
		return -1
	}
	val := c.m[key].Val
	c.put(key, val)
	return val
}

func (c *LRUCache) put(key, val int) {
	x := &node{
		Key: key,
		Val: val,
	}
	if _, ok := c.m[key]; ok {
		c.cache.Remove(*c.m[key])
		c.cache.AddFirst(x)
		c.m[key] = x
	} else {
		if c.cap == c.cache.Size() {
			last := c.cache.RemoveLast()
			delete(c.m, last.Key)
		}
		c.cache.AddFirst(x)
		c.m[key] = x
	}
}
