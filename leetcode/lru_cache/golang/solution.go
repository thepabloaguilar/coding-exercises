package golang

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

type LRUCache struct {
	capacity int
	cache    map[int]*node
	head     *node
	tail     *node
}

func Constructor(capacity int) LRUCache {
	head, tail := &node{}, &node{}
	head.next, tail.prev = tail, head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*node),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.remove(n)
	this.insert(n)

	return n.value
}

func (this *LRUCache) Put(key int, value int) {
	n, ok := this.cache[key]
	if ok {
		this.remove(n)
	}

	n = &node{key: key, value: value}
	this.insert(n)
	this.cache[key] = n

	if len(this.cache) > this.capacity {
		lruNode := this.head.next
		this.remove(lruNode)
		delete(this.cache, lruNode.key)
	}
}

func (this *LRUCache) remove(n *node) {
	n.prev.next, n.next.prev = n.next, n.prev
}

func (this *LRUCache) insert(n *node) {
	prevNode, nextNode := this.tail.prev, this.tail
	prevNode.next, nextNode.prev = n, n
	n.prev, n.next = prevNode, nextNode
}
