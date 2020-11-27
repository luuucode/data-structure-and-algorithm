package sort

import (
	"data-structure-and-algorithm/sort/insert"
	"data-structure-and-algorithm/sort/swap"
	"testing"
)


func TestInsertionSort(t *testing.T)  {
	sortArr := []int{ 5,3,6,1,8}
	insert.InsertionSort(sortArr,0,len(sortArr) -1)
	t.Log(sortArr)
}


func TestQuickSort(t *testing.T) {
	sortArr := []int{ 5,3,6,1,8}
	swap.QuickSort(sortArr,0, len(sortArr)-1)
	t.Log(sortArr)
}

func TestSwap(t *testing.T) {
	sortArr := []int{ 5,3,6,1,8}
	swap.Swap(0, len(sortArr) - 1,sortArr)
	t.Log(sortArr)
}

func TestPartition(t *testing.T) {
	sortArr := []int{ 5,3,6,1,8}
	i := swap.Partition(sortArr,0,len(sortArr)-1,2)
	t.Log(i)
	t.Log(sortArr)
}

func TestFloat(t *testing.T) {
	sortArr := []int{ 5,3,6,1,8}
	swap.Float(sortArr, len(sortArr));
	t.Log(sortArr)
}

func TestBubbleSort(t *testing.T) {
	sortArr := []int{ 5,3,6,1,8}
	swap.BubbleSort(sortArr, len(sortArr));
	t.Log(sortArr)
}
