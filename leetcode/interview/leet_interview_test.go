package interview

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	obj := Constructor(3)
	obj.Put(10,1)
	obj.Put(11,2)
	obj.Put(12,3)
	obj.Put(12,4)
	obj.Put(12,5)
	for _, node := range obj.m {
		t.Log(node.key,node.value)
	}
}
func TestLRUArrays(t *testing.T)  {
	obj := ArrayConstructor(3)
	obj.Put(10,1)
	obj.Put(11,2)
	obj.Put(12,3)
	obj.Put(12,4)
	obj.Put(12,5)
	obj.Put(13,1)
	obj.Get(12)
	for i, node := range obj.arr {
		t.Log(i,node.key,node.value)
	}
}