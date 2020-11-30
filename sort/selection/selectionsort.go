package selection

import (
	"data-structure-and-algorithm/sort/swap"
)

func SelectionSort(arr[] int, len int)  {
	for i := 0; i < len; i++ {
		origin := arr[i]
		index :=  i
		for j := i + 1; j < len -1 ; j++ {
			origin,index = swap.Min(index,j,arr)
		}
		if origin < arr[i] && i != index {
			swap.Swap(i,index,arr)
		}
	}
}