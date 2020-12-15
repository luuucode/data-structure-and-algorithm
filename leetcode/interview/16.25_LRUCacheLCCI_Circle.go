package interview

/**
//TODO 后续来写
因为是最简单的LRU，没有散列表，也没有新老生代，用数组来模拟LRU的思想我们可以借鉴MySQL的redo log环，
用两个指针来表示当前数组的可用空间，分别是write pos来标记当前写入的位置，check point来标记清除数据的检查点。
假设数组的长度为n，初始write pos和check point都是0，数组为空。当写入的时候write pos不断递增，
这里有一个特殊操作，我们要让write pos和check point对n取模以实现环状的效果，
当write pos追上check point的时候（第一次表现为write pos % n = 0，check point此时也为0），
check point就向前推进某个step，这里是1，当然也可以设置更大的值，比如说 0.1 * n，在这个环上，
write pos 到 check point的区间就表征了可用空间，check point到 write pos就代表写入的数据。
*/
type LRUCacheCircle struct {

}


func CirCleConstructor(capacity int) LRUCache {
	return LRUCache{}
}


func (this *LRUCacheCircle) Get(key int) int {
	return -1
}


func (this *LRUCacheCircle) Put(key int, value int)  {

}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
