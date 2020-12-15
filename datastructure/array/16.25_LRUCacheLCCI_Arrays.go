package array

type LRUCache struct {
	arr[] *Node
	capacity int
}

type Node struct {
	key int
	value int
}

func (this *LRUCache)exist(key int) (int,bool) {
	for i, n := range this.arr {
		if n.key == key {
			return i,true
		}
	}
	return -1,false
}

func (this *LRUCache) remove(i int)  {
	this.arr = append(this.arr[:i], this.arr[i+1:]...)
}

func (this *LRUCache) addToHead(node *Node)  {
	arr := this.arr
	this.arr = append(make([]*Node, 0), node)
	for _, n := range arr {
		this.arr = append(this.arr,n)
	}
}

func (this *LRUCache)moveToHead(i int)  {
	node := this.arr[i]
	this.remove(i)
	this.addToHead(node)
	
}

func Constructor(capacity int) LRUCache {
	return LRUCache{make([]*Node,0),capacity}
}


func (this *LRUCache) Get(key int) int {
	if i,ok := this.exist(key);ok {
		node := this.arr[i]
		this.moveToHead(i)
		return node.value
	}
	return -1

}


func (this *LRUCache) Put(key int, value int)  {
	node := Node{key, value}
	if i,ok := this.exist(key); ok {
		this.arr[i].value = value
	}else {
		if len(this.arr) >= this.capacity {
			this.remove(len(this.arr) - 1)
		}
		this.addToHead(&node)
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */