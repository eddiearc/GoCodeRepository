package lru

type Cache struct {
	m        map[int]*node
	head     *node
	tail     *node
	capacity int
}
type node struct {
	prev *node
	next *node
	key  int
	val  int
}

func Constructor(capacity int) Cache {
	head := Node(-1, -1)
	tail := Node(-1, -1)
	head.next = tail
	tail.prev = head
	return Cache{
		m:        map[int]*node{},
		capacity: capacity,
		head:     head,
		tail:     tail,
	}
}

func Node(key, val int) *node {
	return &node{
		key: key,
		val: val,
	}
}

func (n *node) removeFromList() {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (n *node) insertAtHead(h *node) {
	h.next.prev = n
	n.next = h.next
	h.next = n
	n.prev = h
}

func (this *Cache) Get(key int) int {
	n := this.m[key]
	if n == nil {
		return -1
	}
	n.removeFromList()
	n.insertAtHead(this.head)
	return n.val
}

func (this *Cache) removeLast() {
	n := this.tail.prev
	delete(this.m, n.key)
	n.removeFromList()
}

func (this *Cache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	n := this.m[key]
	if n != nil {
		n.val = value
		n.removeFromList()
		n.insertAtHead(this.head)
		return
	}

	n = Node(key, value)
	this.m[key] = n
	n.insertAtHead(this.head)

	if len(this.m) > this.capacity {
		this.removeLast()
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
