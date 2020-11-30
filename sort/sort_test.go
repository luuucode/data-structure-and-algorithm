package sort

import (
	"data-structure-and-algorithm/sort/insert"
	"data-structure-and-algorithm/sort/merge"
	"data-structure-and-algorithm/sort/selection"
	"data-structure-and-algorithm/sort/swap"
	"testing"
)

var sortArr = []int{-1,-2,5,3,6,1,8}

func TestMergeSortTopDown(t *testing.T)  {
	sortArr = merge.MergeSortTopDown(sortArr)
	t.Log(sortArr)
}

func TestSelectionSort(t *testing.T) {
	selection.SelectionSort(sortArr,len(sortArr))
	t.Log(sortArr)
}

func TestInsertionSort(t *testing.T)  {
	insert.InsertionSort(sortArr,0,len(sortArr) -1)
	t.Log(sortArr)
}


func TestQuickSort(t *testing.T) {
	swap.QuickSort(sortArr,0, len(sortArr)-1)
	t.Log(sortArr)
}

func TestSwap(t *testing.T) {
	swap.Swap(0, len(sortArr) - 1,sortArr)
	t.Log(sortArr)
}

func TestPartition(t *testing.T) {
	i := swap.Partition(sortArr,0,len(sortArr)-1,2)
	t.Log(i)
	t.Log(sortArr)
}

func TestFloat(t *testing.T) {
	swap.Float(sortArr, len(sortArr));
	t.Log(sortArr)
}

func TestBubbleSort(t *testing.T) {
	swap.BubbleSort(sortArr, len(sortArr));
	t.Log(sortArr)
}
