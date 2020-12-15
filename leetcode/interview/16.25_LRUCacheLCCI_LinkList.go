package interview

type LinkNode struct {
	key int
	value int
	pre,next *LinkNode
}

type LRUCache struct {
	m map[int]*LinkNode
	capacity int
	head, tail *LinkNode
}


func Constructor(capacity int) LRUCache {
	head := LinkNode{-1, -1, nil, nil}
	tail := LinkNode{-1, -1, nil, nil}
	head.next = &tail
	tail.pre = &head
	return LRUCache{make(map[int]*LinkNode), capacity, &head, &tail}
}

func (this *LRUCache) addToHead(node *LinkNode){
	node.next = this.head.next
	node.next.pre = node
	this.head.next = node
	node.pre = this.head
}

func (this *LRUCache) remove(node *LinkNode)  {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) moveToHead(node *LinkNode)  {
	this.remove(node)
	this.addToHead(node)
}

/**
  如果能从缓存中获取到相应的数据，则将该数据插入到头部
  如果获取不到，则返回-1
 */
func (this *LRUCache) Get(key int) int {
	m := this.m
	n,ok := m[key]
	if ok == true {
		this.moveToHead(n)
		return n.value
	}
	return -1
}


func (this *LRUCache) Put(key int, value int) {
	m := this.m
	if node,ok := m[key];ok {
		node.value = value
		m[key] = node
		this.moveToHead(node)
	}else {
		linkNode := LinkNode{key, value, nil, nil}
		if len(m) >= this.capacity {
			delete(m,this.tail.pre.key)
			this.remove(this.tail.pre)
		}
		m[key] = &linkNode
		this.addToHead(&linkNode)
	}

}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */