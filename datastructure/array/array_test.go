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