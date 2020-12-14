package interview

import "testing"

func TestLRUCache(t *testing.T) {
	obj := Constructor(2)
	obj.Put(10,1)
	obj.Put(11,2)
	obj.Put(12,3)
	for _, node := range obj.m {
		t.Log(node.key,node.value)
	}
}