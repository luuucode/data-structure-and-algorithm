package array

import "testing"

func TestPointers(t *testing.T)  {
	Pointers()
}
func TestStructs(t *testing.T) {
	//Structs()
	//PointerStructs()
	InitStructs()
}

func TestArrays(t *testing.T) {
	Arrays()
	Slice()
}

func TestAdd(t *testing.T) {
	primes := []int{2, 3, 4, 5, 6}
	add := Add(primes, 4, 18)
	t.Log(add)
}

func TestLRUArrays(t *testing.T)  {
	obj := Constructor(3)
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