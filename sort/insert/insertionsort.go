package insert

import "data-structure-and-algorithm/sort/swap"

func InsertionSort(arr[] int, left int, right int)  {
	for i := left; i <= right; i++ {
		for j := i; j > left ; j-- {
			if arr[j] < arr[j-1] {
				swap.Swap(j,j-1,arr)
			}
		}
	}
}